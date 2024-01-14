# channel
channel 是 csp 并发模型中的重要组成部分，它完成的使命是 goroutine 之间的通信。也就是 csp 中的 c。

channel 是 go 语言语言级的线程安全的通信组件，它是 go 语言的一个内置类型，也因此 go 语言才被称之为在设计之初就考虑了并发的语言。

go 语言之父 rob pike 曾说到 “不要使用共享内存的方式去通信，而是通过通信的方式去共享内存” 这句话就是形容的 channel，这句话的意思是说，我们正在执行的程序不应该通过共享内存的方式去完成这个并发操作，各个线程应该各自运行，互不干涉，使用通信的方式去交流去分享内存信息即可。
## 基础用法
### 创建 channel
channel 是一个引用类型，所以创建一个 channel 必须划定底层内存给 channel，所以需要使用 make，不使用 make 只用来声明的 channel，其初始值就是 nil，这一点跟其它引用类型，例如接口，切片，map 是一致的。

对于 nil 的 channel，发送和接收数据总是会阻塞的，对于 nil 的 channel，关闭它直接会 panic `panic: close of nil channel`

```go
// 创建一个无缓存channel
ch := make(chan int)
// 创建一个有缓存的channel
ch := make(chan int, 10)
```
一个无缓存的 channel，只有接收了数据，才能发送下一个数据往通道中，一个有缓存的数据，比如这里是 10，那么可以往这个 channel 中发送 10 个数据，并且不需要接收，只有发送第十一个数据的时候才会发送阻塞。
> 关于 channel 的多线程规则，请参考[内存模型](../内存模型)这一章节

### channel 分为三种类型
- 只能发送的 channel `chan<-`
- 只能接收的 channel `<-chan`
- 可发可收的 channel `chan`

channel 里面理论上可以存在任意的类型数据，当然也包括 channel 自己，那么当 channel 中出现了很多符号时该如何区分呢？
```go
chan<- chan int 
```
这里有个规则，符号总是跟左侧的 chan 在一起表示含义，比如这里，就是一个只能发送的 channel，里面保存的是一个 int 类型的 channel

我们知道，channel 可以拥有三种不同的类型，能发能收的 channel 是满足前两者的，也就是说，当你声明的是一个只能接收的 channel，你可以传入一个能收能发的 channel，但当你声明的是一个能收也能发的 channel 时，就不能传递进去一个只能发或者只能收的 channel
```go
package main

func main() {
	c := make(chan int, 10)
	age(c)
}

func age(c <-chan int) {}
```
```go
// ❌ 
// cannot use c (variable of type <-chan int) as chan int value in argument to age

package main

func main() {
	c := make(<-chan int, 10)
	age(c)
}

func age(c chan int) {}
```
所以，通常我们实际使用当中都是这样的场景，某个函数中的参数是一个只读或者只发的一个 channel 类型，我们传递进去的是一个能收能发的 channel。

因为我们在函数的参数中这么设置是为了避免使用中出现 bug，比如某个函数它就只能收数据，所以给定它的是一个只能收的 channel 非常合理，但是为什么我们创建时要创建一个能收也能发的 channel 呢？因为这个 channel 必定要在其它的 goroutine 中被发送数据，这个其它 goroutine 参数可能被设置为只能发送。

总结一下，设置只能发或者只能收的 channel 本质上只是为了限制函数的能力是为了避免不必要的 bug，但是 channel 的本质我们还是要清楚的：在不同的 goroutine 之间传递信息

同时，除了上面这个禁忌之外，**只能收**的 channel 不能 close，一旦 close 就会 panic，不过只能发的 channel 倒是可以 close 这个操作
```go
//❌
// invalid operation: cannot close receive-only channel c (variable of type <-chan int)
func main() {
	c := make(<-chan int)
	close(c)
}
//❌ 
//invalid operation: cannot close receive-only channel c (variable of type <-chan int)
func age(c <-chan int) {
	close(c)
}

// ✅
func main() {
	c := make(chan<- int)
	close(c)
}
//✅
func age(c chan<- int) {
	close(c)
}
```
### 发送数据
```go
ch := make(chan int)
ch <- 12
```
### 接收数据
```go
func age(ch chan int){
  // 接收但舍弃数值
  // <- ch
  
  // 接收并赋值
  c := <- ch
  fmt.Println(c)
}
```
### 关闭 channel
```go
ch := make(chan int)
close(ch)
```
### 判断从 channel 中取出的值是否有效
```go
v,ok := <-ch

if ok {
  // 是有效值
}else{
  //是零值
}
```
ok 的值表示的是 “接收是否会成功”，而不是 channel 是否已关闭

当 channel 已关闭且数据已取完，再接收时 ok 会为 false，表示 “接收不会成功”

不过值得注意的是，即使 channel 已经关闭，如果之前 channel 中还有数据未被消费，那么数据也能被正常的读取：
```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	// 向channel发送3个数据
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
    fmt.Println("已经关闭了channel，并且没有取值")
	}()
  time.Sleep(time.Second*2)
	// channel中还有未消费的数据,可以正常读取
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
	// channel数据已取完,此时会立即得到零值
	fmt.Println(<-ch) // 0
	// 再次关闭channel会panic
	close(ch)
}
```
### range channel
```go
ch := make(chan int, 3)
for value := range ch{
    fmt.Println(value)
}
```
值得注意的是，一个 channel 如果没有被关闭，那么 range 操作将会一直阻塞，所以通常我们都会关闭这个 channel，好让程序继续执行
```go
go func() {
		wg1.Wait()
		close(b)
	}()

for i := 0; i < 32; i++ {
		go func(i int) {
			defer wg2.Done()
			for i := range b {
        //
			}
		}(i)
	}  
```
注意看，这个关闭的操作其实是另起了一个 goroutine，因为你为了 waitgroup 的完全执行完毕，所以使用新的 goroutine 去关闭 channel 的操作是比较常见的做法，你可以参考这个[项目](https://github.com/shgopher/collie)

### 其它操作：len cap

```go
ch := make(chan int, 10)

// 获取channel中缓存的中还未被取走的元素数量
fmt.Println(len(ch))

// 获取channel的容量
fmt.Println(cap(ch))
```
### 不同状态 channel 的总结
|状态|发送数据|接收数据|关闭chan|
|:---:|:---:|:---:|:---:|
|正常|正常|正常|正常|
|nil|阻塞|阻塞|panic|
|closed|panic|正常值+零值|panic|

对于正常的 channel，如果它没有缓存，那么读写都会发生阻塞，除非另一方准备好，对于有 buffer 的 channel，当它还有缓存的时候，随意发送数据，这个时候收可以不准备好，当他满了，其实跟没有 buffer 的是一样的，这里就不过多讨论了，可以去查看[内存模型](../内存模型/README.md)那一章

## select
select 是 go 语言提供的，供 channel 去操作的一个组件，它的基本使用方法如下
```go
func main() {
  var ch = make(chan int,10)
  var ch1 = make(chan int,10)
  var ch2 = make(chan int,10)
  for i:=0;i< 10;i++ {
    select {
      case <-ch:
        fmt.Println("ch被关闭了")
      case ch1<- i:
        fmt.Println("ch1 发送成功")
      case v,ok := <- ch2:
        if ok {
          fmt.Println("ch2 接收到值",v)
        }
      default:
        fmt.Println("ch没有被关闭")
    }
  }
}
```
首先，select 的 case 中只能存放 channel 的收和发，以及一个 default 分支，各个分支如果在相同时间满足了条件是会随机去走分支的，不存在先后顺序，在 select 中也是可以去做判断的，判断 channel 是否存在正常值

select 本身不具有循环性质，所以通常被配合 for 循环使用

select 在没有任何 case 的时候会陷入阻塞，我们如果希望有一个阻塞存在，使用 select 是极好的：
```go
func main(){
//...
select {}
}


```
### 对于 select for 和 time.Sleep 的阻塞机制的理解
***1。使用 for 循环不会造成 cpu 的执行吗？还是说 cpu 陷入了休眠状态，time.Sleep 呢？***

使用 `for{}` 循环会导致 CPU 高速空转。

for 循环本质上是一个忙等待/空转，CPU 会一直执行循环体内容，只不过这里的循环体为空，所以 CPU 会一直空转浪费资源。

相比而言，time.Sleep() 会**让出 CPU 时间片**，把 goroutine 暂停一段时间，在这段时间 CPU 可以去执行其他任务。当 sleep 时间结束，goroutine 会重新得到时间片继续执行。

所以从资源利用效率来说，time.Sleep() 明显优于 for 循环的空转方式。

***2。那么 select {} 呢是高速空转还是让出资源呢？***

当 select 里所有的 case 都不 ready 时，它会释放 CPU 时间片，使当前 goroutine 进入阻塞状态，这就避免了空转。

***3。select 的不同 case，在等待 case ready 的时候，select 是靠调度器去看 case 是否 ready 还是不停的轮询呢？***

select 内部实现了更精密的监听逻辑：

- 为每个 case 的 channel 设置监听器
- 当前 goroutine 阻塞时释放 CPU 执行权
- 监听器异步监控 channel 状态
- 有 channel ready 则唤醒 goroutine

所以 select 不是通过忙碌的轮询来判断 channel ready，而是通过异步监听的方式，只在必要的时刻唤醒 goroutine。

这种方式可以极大地减少 CPU 的占用，效率也更高。

综上所述，如果你想设置一个阻塞时，使用一个没有 case 的 select 是比一个空 for 循环更好的方法，如果你想设置一个有时间的阻塞时，使用 time.Sleep 无疑是更好的选择，不过这里需要注意的是 time.Sleep 的精度并不高，特别高精度的阻塞不要使用这个函数，select 设计非常精妙，在等待 case 的时候并不会大量耗费 cpu 的执行时间，而是让出 cpu 的执行片段，设置监听，异步的去获取状态，所以 select 的性能是非常不错的

## 定时器
在 go 语言中，go 提供了定时器给用户，基本的使用方法如下：
```go
func main(){
  for {
    select {
      // 在一秒后给case发送信息
    case <-time.After(time.Second):
      return
    // 心跳信号，每间隔一秒发送一次信息
    case <-time.Tick(time.Second):
    }
  }
}
```
通常来说，这是为了超时而去设置的跳出机制

## 使用反射执行未知数量的 channel
当你不确定需要多少个 channel 去处理时，你只能选择在运行时去创建未知数量的 channel，这个时候就要使用反射了。

我们使用 `reflect.Select(cases []reflect.SelectCase)(chosen int,recv reflect.Value,recvok bool)` 去实现 select

```go
package main

import (
	"fmt"
	"reflect"
)

// 使用反射创建的select
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	var cases = createSelectCase(ch1, ch2)

	for i := 0; i < 10; i++ {
		// chosen：channel 切片的索引
		// recv：收到的值
		// ok：是否是通道上发送的值（而不是因为通道关闭而接收到的零值）
		// Select最多支持 65536 个channel
		chosen, recv, ok := reflect.Select(cases)

		if recv.IsValid() { // 收
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // 发
			fmt.Println("send:", cases[chosen].Dir)
		}
	}
}

func createSelectCase(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建 收 channel
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,  // 发 or 收
			Chan: reflect.ValueOf(ch), // 具体 Channel
		})
	}
	// 创建 发 channel
	for i, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,  // 发 or 收
			Chan: reflect.ValueOf(ch), // 具体 channel
			Send: reflect.ValueOf(i),  // 发送的值
		})
	}
	return cases
}


```
## 数据交流---生产者/消费者模式
著名的 worker pool，如果使用 channel 去实现的话，就是一个标准的生产者消费者模式

为了组成这个结构，需要一个存储任务的数据结构，那么这里肯定是使用一个 channel

基本原理就是，一边往 channel 中发送数据，一边从 channel 中取数据，然后使用固定数量的 goroutine 去消费 channel 中的数据，刚好形成一个完整的生产者消费者模式，这就是复用 goroutine 的模式。

具体的额外操作还有控制 worker 数量，任务放入 woker 池，以及从 woker 池取出任务这个操作

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cal := context.WithTimeout(context.TODO(), time.Second*2)
	defer cal()
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				select {
				case c := <-ch:
					fmt.Println(i, ":", c)
					time.Sleep(time.Second)
				case <-ctx.Done():
					fmt.Println(i, "退出")
					return
				}
			}
		}(i)
	}
	// 往 channel 中 发送数据
	go func() {
		for {
			ch <- 89
		}

	}()

	time.Sleep(time.Second * 20)
}

```
这就是基本的生产者消费者模式，以及 channel 去数据交流的最基础的原理，下面我们来实现一个真正的可用的 worker pool

一个真正可用的 worker pool 不仅需要一个任务 channel，还需要存储任务 channel 的 woker pool，而这个 woker pool 是一个 chan chan 类型

```go
type worker struct {
   wokerPool chan *worker
   jobChannel chan Job
}
```
之所以需要一个 chan chan 原因也很简单，如果只有一个 chan，那么 10 几个固定数目的 goroutine 将会互斥的读取一个 channel 的数据，如果是 chan chan 则互不打扰各自读取自己的 channel 即可。

在这个项目中，我们会将读取的 channel 和处理的 channel 分开，读写分离，进行解耦，
读取数据的 channel 只需要一个即可，处理的 channel 可以有多个

```go
func(w *worker)run(){
  go func(){
   for {
    // 将 单个channel（可以简单的这么认为为一个 channel） 
    // 放入 wokerPool 中
    // 如果 这个 单个woker（这里代表单个channel）不再使用的话
    w.workerPool <- w
    select{
      case job := <- w.jobChannel: // 从单个channel中读取数据并处理
        job()
        return        
    }
     }
  }()
}


```
接下来我们需要一个任务分发的函数，也就是从读取的单个 channel 中读取数据，然后发送给多个处理任务的 channel

```go
type dispatcher struct {
   workerPool chan *worker
   OneChannel chan Job
}
func(dispatcher *dispatcher)run(){
  for {
    select {
      case job := <- dispatcher.OneChannel: // 从单个channel中读取数据
          // 从 wokerPool 中获取一个worker
          work :=<- disatcher.workerPool
          // 给这个worker发送任务
          work.jobChannel <- job
    }

  }
}

func RunDispatcher(workerPool chan *workerPool){
   // 启动固定数目的goroutine去消费任务，并且每一个goroutine拥有一个独立的channel
   for i:= 0;i<cap(workerPool);i++ {
      worker := newWorker(workerPool)
      worker.run()
   }
  // 启动分发器
   go  dispatcher.run()
}
```

大致的运行规律就是这些，其它的完整功能请查看这里：https://github.com/shgopher/grpool

## 传递信号/通知
当使用 channel 去传递信号的用法

例如使用一个 channel 去充当信号量，当 channel 没有被关闭的时候，那么就会一直阻塞，一旦 closed，那么就能读取到数据，自然就完成了信号的传递，这种用法非常常见，比如：

```go
func age(){
  ch := make(chan struct{})
  go func(){
    //...
    close(ch)
  }()
  <- ch
}
```
这里使用的 channel 就是充当了一个传递信号的功能。一旦信号传递过来，整个函数就可以继续运行下去了。

或者我们也可以传递一个空的结构体而不是直接 close(ch)：

```go
func age(){
  ch := make(chan struct{})
  go func(){
    //...
    ch <- struct{}{}
  }()
  <- ch
}
```
### 收发同时进行的信号
你也可能见过这种表现方式，当 channel 充当信号的时候，发和收同时进行，看一个例子：

```go
func main() {
	go func() {
		for {
			select {
			case <-w.stop: // B
				w.stop <- struct{}{} // C
				return
			}
		}
	}()

	// 另一段代码
	w.stop <- struct{}{} // A
	<-w.stop             // D
}
```
先说结论，这种用法提供了两个意思，提供信号+判断是否运行完毕

A 代码等 B 准备好之后，发送了信号；
B 代码接受到了信号，这算完成了第一个含义，提供信号

C 代码等 D 代码准备完毕，发送了信号，D 接受完毕，这个表示 A 通知 B 干的事儿，圆满完成，这里是表示判断是否执行完毕
### 使用计时器去解决长等待问题

一般来说，我们都是希望程序执行完毕之后自己主动退出，比如：
```go
func main() {
  ch := make(chan struct{})
  go func(){
    time.sleep(time.Second)
    close(ch)
  }()
  <- ch
}
```
很合理，一般执行完毕之后，就会通知主 goroutine 优雅退出，但是如果执行的时间过长，用户很有可能就会失去信息，这个时候就需要我们再设置一个超时时间更加优雅的退出

```go
func main() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 100)
		close(ch)
	}()
	select {
	case <-ch:
	case <-time.After(time.Second):
		fmt.Println("处理超时")
	}
}
```
## 锁
我们不仅可以使用 sync.Mutext 去实现互斥锁，也可以使用 channel 去做锁，锁本质上来说，其实就是一种信号量，标准的 pv 操作，p 减少数据，获取到锁，v 增加数据释放掉锁，锁是一种二进制信号量

刚好，channel 的收发就符合信号量的指征

```go
type lock struct {
   ch chan struct{}
}
func(l *lock)lock(){
   <- l.ch 
}
func(l *lock)unlock(){
   l.ch <- struct{}{}
}
func (l *lock) tryLock() bool {
	select {
	case <-l.ch:
		return true
	default:
	}
	return false
}

func (l *lock) tryLockeTimeout(timeout time.Duration) bool {
	select {
	case <-l.ch:
		return true
	case <-time.After(timeout):
	}
  return false
	
}
func NewLock() *lock{
  l:= &lock{
    ch: make(chan struct{},1),
  }
  l.ch <- struct{}{}
  return l
}
```
在初始阶段，给予一个拥有一个缓存的 channel 一个数据，获取锁就是将这个数据取出来，释放锁就是将这个数据放回 channel，这样就实现了信号量以及互斥锁

那么让我们试一下我们自己实现的互斥锁：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	l := NewLock()
	value := 0
	for i := 0; i < 120; i++ {

		go func() {
			l.lock()
			value++
			l.unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(value)

}
// 120
```
## 任务编排
### or-do 模式
### fan-in
### fan-out
### map-reduce
### pipeline-流水线模式
### stream-流模式
### pipeline 流水线模式和 stream 流模式的对比
流水线模式 (Pipeline Pattern) 和流模式 (Stream Pattern) 都是将任务分解成多个阶段来处理，但两者还是有一些区别：

- 数据流动方向不同
  - 流水线模式是一条主线，数据从头流到尾，每个阶段处理完输出给下一个阶段。
  - 流模式是多条分支，数据可以从多个来源流入，经过处理流向多个去处。
- 执行方式不同
  - 流水线模式强调阶段间**串行执行**。**一个阶段的输出是下一阶段的输入**。
  - 流模式可以**并行执行**，**不同阶段可以同时操作不同数据**。
- 连接方式不同
  - 流水线模式中，每个阶段靠固定管道连接，顺序不能改变。
  - 流模式中，流可以更灵活自由地连接。
- 扩展性不同
  - 流水线模式扩展一个阶段可能影响整个流水线。
  - 流模式扩展一个阶段对其他阶段影响很小。

总结：
- 流水线模式注重将任务划分多个线性执行的固定阶段。
- 流模式注重构建灵活的流处理流程，流可以并行运动。

所以两者侧重点不同，在使用上也有区别。
## channel 注意事项
发生下面事项一定会触发 panic：

- 向已经关闭的 channel 中发送数据
- 关闭一个 nil channel
- 关闭一个已经被关闭的 channel

### goroutine 泄露
```go
func age(){
  ch := make(chan int)
  go func(){
    time.Sleep(time.Second*10)
    ch <- 12
    }()
    select{
      case <-ch:
        fmt.Println("ch 被关闭了")
      case <-time.After(time.Second):
    }
}
```
当 time.After(time.Second) 执行完毕以后，上面那个 goroutine 因为无法接收数据，所以就会一直阻塞在发送数据那个地方，所以这个代码中，goroutine 就会泄露了。

解决之道就是将容量设置为 1 即可：

```go
func age(){
  ch := make(chan int,1)
  go func(){
    time.Sleep(time.Second*10)
    ch <- 12
    }()
    select{
      case <-ch:
        fmt.Println("ch 被关闭了")
      case <-time.After(time.Second):
    }
}
```
当设置为 1 的时候，即使没有接受者了，发送这个地方的代码也能执行完毕，所以这个 goroutine 是不会泄露了。

这里插一句，main goroutine 只要退出，其它 goroutine 不管有没有执行完毕也会退出，所以如果这种代码在 main 函数中出现，那么是不会发生 goroutine 泄露问题的，因为 main 函数结束以后，其它 goroutine 自动结束
## channel 的实现原理
### 创建
### send
### recive
### close
## issues
### channel 是并发银弹吗？

在 go 语言中，绝大多数情况下都是是使用 channel 更有优势，比如上文提到的那几种场景，例如数据的传递以及任务编排，需要跟 select 或者 context 结合，那么这都是 channel 的适用场景

不过如果是对于共享资源的并发访问，使用传统的互斥锁更有优势一些

如果只是线程安全的对于某个变量的数据变更，使用原子包显然是更加合适的选择
### 有无 buffer 的 channel 区别
go 语言中经常会出现一个 bug，就是死锁，很多都很没有设置 channel 缓存有关，有的时候给 channel 设置一个缓存往往可以规避很多的 panic 风险

### channel close 后，read write close 的区别
- read：正常值+零值
- close：panic
- write：panic
### channle 底层是什么
一个内部有锁的循环队列
### channle 和运行时调度如何交互

### 编程题，使用三个 goroutine 打印 abc 100 次
上文提到的 channel 任务编排中的 pipeline 流水线模式完美解决这个问题。
## 参考资料
- https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
- https://github.com/fortytw2/leaktest
- https://cloud.tencent.com/developer/article/1921580