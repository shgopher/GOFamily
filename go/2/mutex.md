# 同步原语和锁

本章会含有以下内容：
- sync.Mutex
- sync.RWMutex
- sync.WaitGroup
- sync.Once
- sync.Cond
- sync.Pool
- golang/sync/errgroup.Group
- golang/sync/semaphore.Weighted
- golang/sync/singleftlight.Group
- github.com/marusama/cyclicbarrier

导读：
mutex是互斥锁，也就是最基础的锁，RWMutex是读写锁，读写分离，写是强互斥，保持串行，读也有锁，但是降低锁的颗粒度，可以并发执行，但是读还是要让位于写，遇到写读就会强制等待，等写完成，总的来说还是提高了读的性能，waitgroup是让单goroutine等待多goroutine执行完毕，再去做某事，cond的应用场景就是等待以及通知，once就是只允许执行一次，pool就是线程池原理，这里传入的是临时变量，errgroup是将单任务变成多子任务的功能，semaphore是信号量，所谓信号量就是通过信号的增加和减少来控制某个操作的行为，可以理解为多个人用一个网吧会员，singleflight叫做合并请求，意思是多个一模一样的请求合并为一个，cyclicbarrier叫做循环栅栏，先说栅栏意思就是多个goroutine达到某个条件后然后再去执行某件事，循环的意思就是它可以循环使用，并且循环的适合使用简单，不会Panic

## sync.Mutex

这是go语言最基础的锁，也是很多原语都要内置的锁，比如chan中就使用到了这个锁，通常来说我们对锁的优化方法就是降低锁的颗粒度，举个简单的例子，现在1000个人公用一把锁，后面我们1000个人用10把锁，每100个人用一把，这就是降低了锁的颗粒度。

其中操作的时候可以这么做
```go
l := new(sync.Mutex)
// 加锁
l.Lock()
// 解锁
l.Unlock()
```

我们看一下sync.Mutex的底层数据结构：
```go
type Mutex struct {
	state int32
	sema  uint32
}
```
可以看出，state+sema，也就是状态+信号量一共64位，8个字节，就表示了整个go语言的互斥锁.

state按照位运算，后三位表示三种模式，即：mutexlock，mutexwwoken，mutexstarving，分别的意思是锁定，醒来，和饥饿，然后剩下的位都用来表示当前等待的goroutine数量。

sync.Mutex 分为正常模式和饥饿模式。正常模式下，按照先入先得的模式来获取锁，但是当一个被唤醒的goroutine跟一个新建的同时去获得锁的时候，就会发现，刚被唤醒的竞争不过这个新建的，体现就是这个goroutine超过了1ms都获取不到锁，这个时候防止它被饿死，锁立马切换成饥饿模式，率先让队头的goroutine获得锁，把这个新建的goroutine放到队尾部，然后，当获得锁的goroutine在队尾，以及goroutine获取锁的时间小于1ms的时候，那么这个锁就会重新切换成正常模式。饥饿模式可以防止某个goroutine一直无法获取锁，而被饿死的情况发生，引入了公平机制。

加锁和解锁过程，主要是控制的mutexlock字段，当锁的位置是0的时候，那么就使用atom包的原子操作，将值改成1，当互斥锁的状态不是0的时候，就会把这个goroutine改成自旋的方式，继续去等待这个锁，当然了这个进入自旋是有条件的，比如必须是多核心cpu，比如自旋的次数不能大于4，比如必须存在至少一个正在运行的p，且队列是空。处理完自旋以后，系统就会更新mutexwoken mutexstarving这几个字段的状态。如果这个时候没有获得锁，那么系统就会通过信号量的方式，不允许两个goroutine获取到锁，并且将此goroutine改成休眠状态，然后等待被唤醒。在正常模式下，这段代码会设置唤醒和饥饿标记、重置迭代次数并重新执行获取锁的循环；在饥饿模式下，当前 goroutine 会获得互斥锁，如果等待队列中只存在当前 goroutine，那么互斥锁就会从饥饿模式退回正常模式。

解锁的过程就是通过原子操作将mutexlock字段从1改为0，如果不等于0，那么就会尝试慢速解锁。首先解锁会查看是否已经解锁，如果已经解锁再调用unlock就会发生Panic，正模式下，如果发现没有等待者，或者mutexlock这几个字段不都为0，那么就会直接退出，不需要唤醒其它goroutine，如果发现有其它goroutine，那么就会唤醒其它goroutine，然后移交控制权。饥饿模式下，这个锁会直接移交给下一个等待的goroutine，并且这个时候饥饿模式不会被取消掉。

四种会发生错误的操作一定要避免

- 解锁已经解锁的锁，会Panic
- 复制这个锁，因为底层来看都是有状态的，复制以后就不是初始值了，肯定会出错
- mutex不支持重入，意思就是某个持有的goroutine可劲再调用锁，go不支持，因为go的锁不记录goroutine的信息。

几种可以提高锁的性能的方法

1. 降低锁的颗粒度
2. 引入排它锁，就是某个goroutine持有了锁，其他锁就不再争夺了。

## sync.RWMutex

读写锁，这是细颗粒度的互斥锁，写锁跟互斥锁一样，但是读锁的颗粒度就会降低，所以说不是说读就不加锁，而且锁的颗粒度降低。所以说读锁在读的时候会并发执行，来加快读的速度，读锁如果发现这个变量被写锁占用，那么就会等待，也就是说读会让位写，让写优先。读的时候并发执行来代替串行自然就提高了速度。
## sync.WaitGroup

- sync.WaitGroup 必须在 sync.WaitGroup.Wait 方法返回之后才能被重新使用；
- sync.WaitGroup.Done 只是对 sync.WaitGroup.Add 方法的简单封装，我们可以向 sync.WaitGroup.Add 方法传入任意负数（需要保证计数器非负）快速将计数器归零以唤醒等待的 Goroutine；
- 可以同时有多个 Goroutine 等待当前 sync.WaitGroup 计数器的归零，这些 Goroutine 会被同时唤醒；

## sync.Once

可以保证某段代码只执行一次。sync.Once.Do 方法中传入的函数只会被执行一次，哪怕函数中发生了 panic；两次调用 sync.Once.Do 方法传入不同的函数只会执行第一次调传入的函数；

## sync.cond
Go 标准库提供 Cond 原语的目的是，为等待 / 通知场景下的并发问题提供支持。Cond 通常应用于等待某个条件的一组 goroutine，等条件变为 true 的时候，其中一个 goroutine 或者所有的 goroutine 都会被唤醒执行。
使用的机会较少。一般我们都是使用channel来代替这个东西

## sync.Pool

go实现的临时线程池，可以往池里加东西，然后取东西，但是无法清楚的获取到底是取得到哪个东西

## errgroup.Group
处理一组子任务，就该使用errgroup等这种控制分组的操作，它适合于将一个通用的父任务拆成几个小任务并发执行的场景。

它有三个方法
- WithContext
    传入一个context，并返回一个group
- go
    传入子任务
- wait
    等待子任务的完成

```go
package main
import (
    "errors"
    "fmt"
    "time"
    "golang.org/x/sync/errgroup"
)
func main() {
    var g errgroup.Group
    // 启动第一个子任务,它执行成功
    g.Go(func() error {
        time.Sleep(5 * time.Second)
        fmt.Println("exec #1")
        return nil
    })
    // 启动第二个子任务，它执行失败
    g.Go(func() error {
        time.Sleep(10 * time.Second)
        fmt.Println("exec #2")
        return errors.New("failed to exec #2")
    })
    // 启动第三个子任务，它执行成功
    g.Go(func() error {
        time.Sleep(15 * time.Second)
        fmt.Println("exec #3")
        return nil
    })
    // 等待三个任务都完成
    if err := g.Wait(); err == nil {
        fmt.Println("Successfully exec all")
    } else {
        fmt.Println("failed:", err)
    }
}
```



## semaphore.Weighted
信号量顾名思义，就是使用某个变量比如 0 1的状态不同来去表示不同的状态，这就是信号量。
在信号量里有两个操作p和v，p就是减少信号量，v就是增加信号量。

golang.org/x/sync/semaphore 中是以下内容：

```go

type Weighted struct{}
// 返沪一个 weighted
func NewWeighted
// p  操作
func Acquire
// v 操作
func Release
// 尝试获取 n 个资源，但是它不会阻塞，要么成功获取 n 个资源，返回 true，要么一个也不获取，返回 false。
func TryAcquire

```
所以说信号量用非常容易的话来说就是，假如你在网吧办了一张会员，你和你的朋友都能用，你们每用一次就是增加来一次使用次数，假设一共能用10次，越接近十次就越少机会继续使用，任何充钱每次充钱就把这个信号减去1，代表你们可以用的次数又多了。简单的信号量就是这么容易，使用某个信号充当某个作用。

所以说如果按照这种逻辑，使用channel来实现一个信号量也很容易，初始化就是给定一个chan多少容量，每用一次就往chan中添加几个次数，然后每次充钱就从chan中取回来多少信号。

使用chan有个缺点，就是一次只能实现取得一个资源但是golang.org的这种就可以一次获取n个资源，这也是golang.org实现这种方式的信号量的原因把。

## singleflight.Group
SingleFlight 的作用是将并发请求合并成一个请求，以减少对下层服务的压力，在处理多个 goroutine 同时调用同一个函数的时候，只让一个 goroutine 去调用这个函数，等到这个 goroutine 返回结果的时候，再把结果返回给这几个同时调用的 goroutine，这样可以减少并发调用的数量。在面对**秒杀**等大并发请求的场景，而且这些请求都是读请求时，你就可以把这些请求合并为一个请求，这样，你就可以将后端服务的压力从 n 降到 1

在高并发秒杀服务中，这种合并请求简直就是无敌的存在啊～
## cyclicbarrier
CyclicBarrier 是一个可重用的栅栏并发原语，用来控制一组请求同时执行的数据结构。它常常应用于重复进行一组 goroutine 同时执行的场景中。举个例子，有一种场景，需要在多个goroutine同时达到某个条件时，然后让这些goroutine共同发生某种操作，并且这种等待然后操作是循环的，这个时候就可以使用循环栅栏，这个waitgroup也可以，不过因为waigroup重复使用要又注意事项，就是在wait结束后才能add所以容易出现bug，这个时候使用cyclicbarrier更好。

WaitGroup 更适合用在“一个 goroutine 等待一组 goroutine 到达同一个执行点”的场景中，或者是不需要重用的场景中。CyclicBarrier 更适合用在“固定数量的 goroutine 等待同一个执行点”的场景中，而且在放行 goroutine 之后，CyclicBarrier 可以重复利用

接下来我们来看一个使用了循环栅栏和信号量的一段代码，这段代码的意义是生成水分子。

```go
// github.com/marusama/cyclicbarrier
// golang/x/sync/semaphore
package main

import (
	"context"
	"fmt"
	cyc "github.com/marusama/cyclicbarrier"
	sema "golang.org/x/sync/semaphore"
	"math/rand"
	"sort"
	"sync"
	"time"
)
func main(){
	CreatH20()
}
type H2O struct {
	H *sema.Weighted
	O *sema.Weighted
	b cyc.CyclicBarrier
}
func NewH2O()*H2O {
	return &H2O{
        // 这里配置好配方，一o二h，共计3个才能下一轮
		H: sema.NewWeighted(2),
		O: sema.NewWeighted(1),
		b: cyc.New(3),
	}
}
func (h *H2O)CreatH(r func()){
	h.H.Acquire(context.Background(),1)
	r()
    // 这里的wait的意思就是如果达不到三个，那么就 一直wait，这里的信号量的意义就跟我说的网吧原理一样，一共有三个硬币，h能完两次最多，o能玩一次，每次使用的h是1，每次使用的o也是1.
	h.b.Await(context.Background())
	h.H.Release(1)

}
func (o *H2O)Creat0(r func()){
	o.O.Acquire(context.Background(),1)
	r()
	o.b.Await(context.Background())
	o.O.Release(1)
}

func CreatH20(){
	var N = 100
	wg := new(sync.WaitGroup)
	h2o := NewH2O()
	ch := make(chan string,N*3)
	r1 := func() {
		ch <- "H"
	}
	r2 := func() {
		ch <- "o"
	}
	wg.Add(N*3)
	for i:= 0;i <N*2;i++ {
		go func() {
			defer wg.Done()
            // 假设创造h不需要时间等待
			h2o.CreatH(r1)
		}()
	}

	for i:= 0;i <N;i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.Creat0(r2)
		}()
	}
	wg.Wait()
    // 这里 chan初始化给定的是容量，但是len给不了，len是值的队列中实际存在的数据。所以初始化的时候其实len=0，容量是300
	if len(ch) != 3*N {
		fmt.Println("error,n is wrong",len(ch))
	}
	t := make([]string,3)
	for i:= 0;i< N;i++{ // 这里不要使用range，因为需要关闭chan才能跳出，使用select也可以。
		t[0]=  <- ch
		t[1] = <- ch
		t[2] = <- ch
		sort.Strings(t)
		water := t[0] + t[1]+t[2]
		if water != "HHo" {
			fmt.Println("error",water)
		}

	}
	fmt.Println("yes ,im ok ")
}
```
