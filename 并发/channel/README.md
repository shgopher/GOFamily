<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-03-02 15:18:19
 * @FilePath: /GOFamily/并发/channel/README.md
 * @Description: 
 * 
 * Copyright (c) 2024 by shgopher, All Rights Reserved. 
-->
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
值得注意的是，一个 channel 如果没有被关闭，那么 range 操作将会一直阻塞，所以通常我们都会关闭这个 channel，好让程序继续执行，所以当我们控制 channel 的时候，尽可能的还是让发送方去控制 channel 的关闭，不要在接收方去控制。
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
### 对于 select，for，time.Sleep 的阻塞机制的理解
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
### cron 规则定时任务框架 robfig/cron

robfig/cron 是一个 Go 语言编写的 Cron 计划任务管理器。

这个项目的主要功能和特点包括：

- 实现了 cron 规范，可以基于 cron 表达式来调度任务
- 支持在代码中以声明式方式定义定时任务
- 支持在配置文件中定义定时任务
- 支持任务链，一个任务可以触发其他任务
- 支持任务锁，防止并发执行
- 支持失败重试
- 有默认的日志记录器，也可以自定义日志记录器
- 轻量级，没有外部依赖

robfig/cron 因为其简单、轻量和高效而被许多 Go 项目用来实现定时任务。它可以用于项目中不同粒度的调度需求，如每分钟执行、每小时执行等。

下面解释一下 cron 机制的含义：

cron 表达式调度是 cron 作业调度程序的核心机制。

cron 表达式是一个字符串，由 5 或 6 个用空格分隔的字段组成，每个字段表示一个时间单位。Cron 会根据表达式中设置的值，在特定的时间点触发相应的任务。

cron 表达式的字段与意义如下：

```go
# 文件格式說明
# ┌──分钟（0 - 59）
# │ ┌──小时（0 - 23）
# │ │ ┌──日（1 - 31）
# │ │ │ ┌─月（1 - 12）
# │ │ │ │ ┌─星期（0 - 6，表示从周日到周六）
# │ │ │ │ │ ┌─年份（ * 表示全部年份，2024 表示只在2024年）
# │ │ │ │ │ │
# * * * * * *

# 例如
# 0 0 12 * * *
# 每天 12 点触发
```

每个字段可以使用特殊字符表示多种时间设置：

- `*`：表示匹配该字段的任意值
- `,`：表示分隔多个值，如在分钟字段使用 “5,15,25” 就表示5分、15分、25分都触发
- `-`：表示一个范围，如在小时字段使用 “9-17” 表示 9 点到 17 点之间每个整点都触发
- `/`：表示一个间隔，如在分钟字段使用 “0/15” 表示每15分钟触发一次，从0分开始
- `?`：表示不指定的值，会匹配该字段的任意值。

一些例子：

- `0 0 12 * * ?`：每天 12 点触发
- `0 15 10 ? * *`：每天 10:15 触发
- `0 0/5 14 * * ?`：每天 14:00 到 14:55 每5分钟触发一次
- `0 0/5 14,18 * * ?`：每天 14:00 到 14:55 和 18:00 到 18:55 每5分钟触发一次

通过这种表达式，cron 可以非常灵活地设置触发时间表，从而实现各种调度需求。

下面给出这个项目的 demo：

```go
c := cron.New() 

// 每小时的第30分钟执行
c.AddFunc("30 * * * *", func() { fmt.Println("每小时的第30分钟") })

// 在每天早上3点到6点,晚上8点到11点的范围内,每小时执行一次  
c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println("在每天早上3点到6点,晚上8点到11点的范围内,每小时执行一次") }) 

// 每天早上东京时间4点30分执行一次
c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("每天早上东京时间4点30分执行一次") })

// 从现在起每小时执行一次
c.AddFunc("@hourly", func() { fmt.Println("从现在起每小时执行一次") })

// 从现在起每隔1小时30分执行一次  
c.AddFunc("@every 1h30m", func() { fmt.Println("从现在起每隔1小时30分执行一次") })

c.Start()

// 被添加的函数会在自己的goroutine中异步执行
... 

// 也可以向已经启动的Cron添加新的函数
c.AddFunc("@daily", func() { fmt.Println("每天执行") })

// 检查cron任务下次和上次的执行时间
inspect(c.Entries()) 

// 停止调度器(已经启动的任务不会被停止)
c.Stop()
```

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
### shutdown 和信号传递结合优雅退出
Shutdown 方法通常与信号处理结合使用，以便在接收到系统退出信号 (如 SIGTERM 或 SIGINT) 时优雅地关闭服务器。这样可以确保所有正在进行的操作都能完成，而不是突然中断，从而提供更好的用户体验

这个方法允许你安全地停止服务器，而不会中断正在处理的请求

Shutdown 方法的工作原理如下：

- 关闭监听器：首先，它会关闭所有打开的监听器，这意味着服务器将不再接受新的连接请求。

- 处理现有连接：然后，它会等待所有现有的连接 (如 HTTP 请求) 完成处理。这包括那些正在处理中的请求，以及那些已经完成但尚未关闭的连接。

- 上下文管理：Shutdown 方法接受一个 context.Context 参数，这个上下文可以用来设置超时。如果上下文在服务器关闭完成之前被取消 (例如，因为超时)，Shutdown 方法会返回一个错误。

- 不处理长连接：需要注意的是，Shutdown 方法不会尝试关闭那些长时间保持连接的连接，如 WebSockets。对于这类连接，你需要在服务器上注册一个关闭通知函数，以便在需要时手动通知这些连接关闭。

- 服务器不可重用：一旦调用了 Shutdown，服务器就不能再被重用了。尝试再次调用 Serve、ListenAndServe 或 ListenAndServeTLS 方法将返回 ErrServerClosed 错误。

下面我们看一个案例：

有两个 listenAndServe 服务，其中一个是 debug 服务一个是正式服务
 ```go
func main(){
	done := make(chan error,2)
	stop := make(chan struct{})

	// 使用方去控制服务的关闭
	// 使用方控制并发的操作
	go func(){
		done <- serveDubug(stop)
	}()
	go func(){
		done <- serveApp(stop)
	}()

	var stopped bool

	for i := range cap(done){
		if err := <- done; err != nil{
			fmt.Println(err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}

// 省略
//func serveDubug(){serve()}
//func serveApp(){serve()}

func serve(addr string,handler http.Handler,stop <-chan struct{})error {
	s := http.Server{
		Addr: addr,
		Handler: handler,
	}
	// 信号传递和关闭服务结合
	// 这里不能使用 os.Exit()
	// os.exit 会立刻全部停止，不优雅，比较暴力
	// log.Fetal 这个函数通常不要使用，因为它的底层调用的就是os.Exit()
	// 如果要是使用这个 log.Fetal 函数，只用在程序的起始位置，比如main的开头，或者init函数中
	go func(){
		<- stop
		s.Shutdown(contex.Background())
	}()
	return s.ListenAndServe()
}
 ```
### 信号传递之 or-do 模式
channel 信号传递之 or-do 模式跟 go 并发原语的 [singleflight](../同步原语/README.md#singleflight) 比较类似

区别是，channel 的 or-do 是多个信号传递，只有一个 channel 获取信号之后就可以停止其它 channel 了，并不需要将这个信号传递给其它的 channel，singleflight 是需要分享情报的。

下面有两种实现的方式，递归和反射，递归比较适合递归树较低的场景，反射就适合了递归树更高的场景，因为如果递归过深的话很容易出现内存溢出的情况，如果你认为你项目中的 goroutine 也就几十那么递归就可以了，毕竟反射可是很浪费执行时间的，如果 goroutine 成千上万，选用反射即可。

使用递归的方法：
```go

func Ordo(chs ...<-chan any) <-chan any {
	// 递归退出条件
	switch len(chs) {
	case 0:
		return nil
	case 1:
		return chs[0]
	}
	or := make(chan any)
	go func() {
		defer close(or)
		switch len(chs) {
		case 2:
			select {
			case <-chs[0]:
			case <-chs[1]:
			}
		default:
			mi := len(chs) / 2
			select {
			case <-Ordo(chs[:mi]...):
			case <-Ordo(chs[mi:]...):
			}
		}
	}()
	return or
}
```

使用反射的方式：
```go

func Ordo1(chs ...<-chan any) <-chan any {
	// 递归退出条件
	switch len(chs) {
	case 0:
		return nil
	case 1:
		return chs[0]
	}
	or := make(chan any)
	
  go func() {
		defer close(or)
		var cases []reflect.SelectCase
		for _, v := range chs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(v),
			})

			reflect.Select(cases)
		}
	}()
	return or
}
```
开始测试：

```go
package main

import "time"

func main() {
  // output: 7
  start := time.Now()
	<-Ordo(sig(time.Second*7), sig(time.Second*10), sig(time.Second*13))
	fmt.Println(time.Since(start))
}

func sig(t time.Duration) <-chan any {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		time.Sleep(t)
	}()
	return ch
}

```

测试证明这两种方法均可实现 or-do 信号通知模式，不过就我看来，反射的这种方法，完全无脑，也不用考虑递归，也不必去考虑递归导致的内存溢出，从实现逻辑上，非常的清晰，只需要组成一个 select 随便 case 即可，非常清晰，我首推这种实现方法

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
### fan-in
核心思想就是 “多个进，单个出”，这是数字电路的一个概念，场景是，多个 channel 发送数据给一个处理函数，此函数返回一个 channel 去承载众多信息，我们只需要监听最后输出的 channel 就可以达到监听众多信息的目的。

基本样貌如下：
```go
func fanin(chs ...<-chan any) <-chan any {
}
```
从这个基本构型也能看出来，我们要在函数中处理多个 channel，这跟上文 channel 传递信号的 or-do 模式类似，所以免不了我们要使用反射或者递归的方式去处理函数，首先我们使用递归的方式去实现这个功能

递归版 `fan-in`：
```go
func fanin(chs ...chan any) chan any {
	switch len(chs) {
	case 0:
		c := make(chan any)
		close(c)
		return c
	case 1:
		return chs[0]
	case 2:
		return merge(chs[0], chs[1])
	default:
		m := len(chs) / 2
		return merge(fanin(chs[:m]...), fanin(chs[m:]...))
	}
}

func merge(ch1, ch2 chan any) chan any {
	c := make(chan any)
	go func() {
		defer close(c)
		for ch1 != nil || ch2 != nil {
			select {
			case v1, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				c <- v1
			case v2, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				c <- v2
			}
		}

	}()
	return c
}
```
反射版 `fan-int`：
```go
func fanintflect(chs ...chan any) chan any {
	out := make(chan any)
	go func() {
		defer close(out)

		var cases []reflect.SelectCase

		for _, vl := range chs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(vl),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			// 从 cases 取数据，一旦取完，cases 会取消已经取完的 case
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			// out 发送数据
			out <- v.Interface()
		}
	}()
	return out
}
```
反射版本这里有个小问题，从 cases 中取出数据，往 out 中添加，这个操作只能一个 case 一个 case 的取，但是你如果观察上文的递归实现模式，你会发现它开启了很多 goroutine 所以是并发的去取数据往 out 中送。
### fan-out
fan out 跟 fan in 模式是相反的，fan out 指的是拥有一个输入源，多个输出源，实际上这是一种广播机制，读取一个数据，广播给众多接受者。

下面看一下代码：

```go
func fanout(value chan any, out []chan any, async bool) {
	go func() {
		defer func() {
			for _, v := range out {
				close(v)
			}
		}()
		// 对一个nil的通道进行 for range 遍历会导致阻塞(block)。
		for v := range value {
			for _, vi := range out {
				vi := vi // if go version is lower then 1.22
				if async {
					go func() {
						vi <- v
					}()
				} else {
					vi <- v
				}
			}
		}
	}()
}
```

实际上这段代码有 goroutine 泄露的风险，如果 out 一直不被消费，value 的数据越来越多，那么这么多等待发送信息的 goroutine 就会发生泄露问题

解决方案有下面几种：

- 使用信号量去控制最多的 goroutine 数量
- 给每一个 goroutine 设置超时时间
- 控制发送的 value channel 的发送频率
- 使用 worker pool 去限制并发数量

那么让我们用代码去实现这些改进的方案：

方案一：使用信号量控制
```go
var (
  maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers))
)
// 在 value传递结束之后，记得close value
func fanout(value chan any, out []chan any, async bool) {
	go func() {
		defer func() {
			for _, v := range out {
				close(v)
			}
		}()
		// 对一个nil的通道进行 for range 遍历会导致阻塞(block)。
		for v := range value {
			for _, vi := range out {
				vi := vi // if go version is lower then 1.22
				if async {
          ctx := context.Background()
					if err := sema.Acquire(ctx, 1); err != nil {
						break
					}
					go func() {
						defer sema.Release(1)
						vi <- v
					}()
				} else {
					vi <- v
				}
			}
		}
	}()
}
```

方案二：使用超时时间
```go
func fanout(value chan any, out []chan any, async bool) {
	go func() {
		defer func() {
			for _, v := range out {
				close(v)
			}
		}()
		// 对一个nil的通道进行 for range 遍历会导致阻塞(block)。
		for v := range value {
			for _, vi := range out {
				vi := vi // if go version is lower then 1.22
				if async {
					go func() {
            select{
              case vi <- v:
              case time.After(time.Second):
              return
            }
						
					}()
				} else {
					vi <- v
				}
			}
		}
	}()
}
```

方案三：控制发送 value 的发送频率
```go
go func() {
		for {
			select {

			case <-time.After(time.Second * 3):
				fmt.Println("关闭了吗")
				close(value)
				return
			case value <- "%":
				time.Sleep(time.Second >> 5)
				fmt.Println("发送了吗")

			}
		}
	}()
```

方案四：使用 worker 池复用 goroutine 去控制并发的 goroutine 数量
```go
// github.com/shgopher/grpool
func fanout(value chan any, out []chan any, async bool) {
	pool := grpool.NewPool(100, 50)
	go func() {
		defer pool.Release()
		defer func() {
			for _, v := range out {
				close(v)
			}
		}()
		// 对一个nil的通道进行 for range 遍历会导致阻塞(block)。
		for v := range value {
			for _, vi := range out {
				if async {
					pool.JobQueue <- func() {
						vi <- v
					}
				} else {
					vi <- v
				}
			}
		}
	}()
}

```

按照我的工作经验，使用工作池和信号量控制 goroutine 数量的方法最为常用，他们都是保证最多同时存在 n 个 goroutine，这样就避免了 goroutine 泄漏问题
### map-reduce
我们在函数那个[章节](https://github.com/shgopher/GOFamily/blob/aca91397692f99cd744fabcaca2055da3ec92c6f/%E5%9F%BA%E7%A1%80/%E5%87%BD%E6%95%B0%E6%96%B9%E6%B3%95/1.md?plain=1#L11)介绍过 map-reduce 模式，不过并没有牵涉到并发，这里我们介绍一下并发版的 map-reduce 模式
map
```go
func mapChan(in <-chan any,fn func(any)any)<-chan {
  out := make(chan any)
  if in == nil {
    close(out)
    return out
  }

  go func() {
    defer close(out)
    for v := range in {
      out <- fn(v)
    }
  }()
  return out
}
```
reduce
```go
func reduceChan(in <-chan,fn func(r,v any)any)any {
  if in == nil {
    return nil
  }

  out <- in

  for v := range in {
    out = fn(out,v)
  }    
  return out
}
```
### pipeline-流水线模式
pipeline 模式的核心思想就是顺序，单线模式，数据从头到尾，顺序执行，一个阶段的输出是下一阶段的输入。

```go
var a A
a.get().then().Download()
```
这就是 pipeline 的一个简单掩饰，它的底层可能是这样实现的

```go
type A struct{}
func (a *A)get()*A{return a}
func (a *A)then()*A{return a}
func (a *A)Download()*A{return a}
```

那么在并发的场景里，如何使 goroutine 实现 pipeline 模式呢？

这显然跟一般的流水线实现方法不同。

我们将 channel 比作一个 token，所谓令牌，只要我们控制获取令牌的顺序，那么就可以控制持有这些令牌的 goroutine 顺序。

下面我们使用 4 个 goroutine 依次打印 “a b c d”，300 次，我们使用流水线的方式实现：

```go
package main

import (
	"fmt"
)

type Token struct{}

var (
	GNumber int
	Number  int
	sign    = make(chan struct{})
)

func world(id int) {
	switch id {
	case 0:
		fmt.Println("a")
	case 1:
		fmt.Println("b")
	case 2:
		fmt.Println("c")
  case 3:
    fmt.Println("d")
  }
}
func worker(id int, in chan Token, out chan Token, number int,fn func(int)) {
	for number > 0 {
		token := <-in
		fn(id)
    
    // 这段代码是为了保证最后一个 goroutine 执行完毕退出 
		if id == (GNumber-1) && number == 1 {
			close(sign)
			break
		}

		out <- token
		number--
	}
}
func RunPipeline(GNumber int,Number int,chs[]chan Token,fn func(int)) {
  
	for i := range GNumber {
    // 核心代码: 
    // 传入的in chan 和 out chan 刚好前后顺序
    // 这样 token 才能前后传递
		go worker(i, chs[i], chs[(i+1)%GNumber], Number,fn)
	}
	chs[0] <- struct{}{}
	<-sign

}
func main() {
  GNumber = 4
	Number = 300
  chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
  RunPipeline(GNumber,Number,chs,world)
}
```
 
### stream-流模式
将 channel 当做流式管道，提供多种功能，比如筛选元素，跳过元素等

将数据转化为流，其中 channel 就是流，众多数据这里是 `...any`

创建流
```go
func stream(done chan struct{}, values ...any) chan any {
	c := make(chan any)
	go func() {
		defer close(c)
		for _, v := range values {
			select {
			case <-done:
				return
			case s <- v:
			}
		}
	}()
	return c
} 
```
根据流，我们可以有以下的处理
- takeN：只取流中的前 n 个数据
- takeFn：筛选流中的数据，只保留满足条件的数据
- takeWhile：只取前面满足条件的数据，一旦不满足，就不再取了
- skipN：跳过流中的前 n 个数据
- skipFn：跳过满足条件的所有数据
- skipWhile：跳过前面满足条件的数据，一旦不满足条件了，当前这个元素和以后的元素都会输出

#### takeN
```go
func TakeN(done chan struct{},in chan any,num int)chan any{
	c := make(chan any)
	go func(){
		defer close(c)
		for i := 0; i < num; i++ {
			select {
				case <-done:
					return 
				case v,ok := <- in:
					if ok  {
						c <-v
					}					
			}
		}
	}()
	return c
}
```
#### takeFn
```go
func TakeFn(done chan struct{},in chan any,fn func(any)bool)chan any{
	c := make(chan any)
	go func(){
		defer close(c)
		for v := range in {
			if fn(v){
			select {
				case <-done:
					return
				case c <- v:					
			  }
			}
		}
	}()
	return c
}
```

#### takeWhile
```go
func takeWhile(done chan struct{}, in chan any, fn func(any) bool) chan any {
    c := make(chan any)
    go func() {
        defer close(c)
        for v := range in {
            if !fn(v) {
                return
            }
            select {
            case <-done:
                return
            case c <- v:
            }
        }
    }()
    return c 
}
```
#### skipN
```go
func skipN(done chan struct{}, in chan any, n int) chan any {
    c := make(chan any)
    go func() {
        defer close(c)
        i := 0
        for v := range in {
            if i < n {
                i++
                continue
            }
            select {
            case <-done:
                return
            case c <- v:
            }
        }
    }()
    return c
}
```
#### skipFn
```go
func skipFn(done chan struct{}, in chan any, fn func(any) bool) chan any {
    c := make(chan any)
    go func() {
        defer close(c)
        for v := range in {
            if !fn(v) {
                select {
                case <-done:
                    return
                case c <- v:
                }
            }
        }
    }()
    return c
}
```
#### skipWhile
```go

func skipWhile(done chan struct{}, in chan any, fn func(any) bool) chan any {
    c := make(chan any)
    go func() {
        defer close(c)
        skip := true
        for v := range in {
            if skip && fn(v) {
                continue
            }
            skip = false
            select {
            case <-done:
                return
            case c <- v:
            }
        }
    }()
    return c
}
```
#### 测试 stream
```go
func TestStream(t *testing.T) {

  done := make(chan struct{})
  defer close(done)

  in := stream(done, 1, 2, 3, 4, 5)

  // takeN
  out := takeN(done, in, 3)
  want := []int{1, 2, 3}
	
	// reflect.DeepEqual 更适用于判断复杂类型的变量,
	// 像数组、切片、map、结构体等是否相等,
	// 简单类型可以用==运算符判断。
  if got := drain(out); !reflect.DeepEqual(got, want) {
    t.Errorf("takeN = %v, want %v", got, want)
  }

  // takeFn
  fn := func(v any) bool {
    n := v.(int)
    return n%2 == 0
  }
  out = takeFn(done, in, fn)
  want = []int{2, 4}
  if got := drain(out); !reflect.DeepEqual(got, want) {
    t.Errorf("takeFn = %v, want %v", got, want) 
  }

  // takeWhile
  out = takeWhile(done, in, func(v any) bool {
    return v.(int) < 4
  })
  want = []int{1, 2, 3}
  if got := drain(out); !reflect.DeepEqual(got, want) {
    t.Errorf("takeWhile = %v, want %v", got, want)
  }

  // skipN 
  out = skipN(done, in, 2)
  want = []int{3, 4, 5}
  if got := drain(out); !reflect.DeepEqual(got, want) {
    t.Errorf("skipN = %v, want %v", got, want)
  }

  // skipFn
  out = skipFn(done, in, func(v any) bool {
    return v.(int)%2 == 0
  })
  want = []int{1, 3, 5}
  if got := drain(out); !reflect.DeepEqual(got, want) {
    t.Errorf("skipFn = %v, want %v", got, want)
  }

  // skipWhile
  out = skipWhile(done, in, func(v any) bool {
    return v.(int) < 3
  })
  want = []int{3, 4, 5}
  if got := drain(out); !reflect.DeepEqual(got, want) {
    t.Errorf("skipWhile = %v, want %v", got, want) 
  }
}

func drain(in <-chan any) []any {
  out := make([]any, 0)
  for v := range in {
    out = append(out, v)
  }
  return out
}
```
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

这里插一句，main goroutine 只要退出，其它 goroutine 不管有没有执行完毕也会退出，所以如果这种代码在 main 函数中出现，那么是不会发生 goroutine 泄露问题的，因为：

***main 函数结束以后，其它 goroutine 自动结束***

## channle 和运行时调度如何交互
channel 和 Go 运行时调度器的交互方式如下：

1. 当 goroutine 向 channel 发送数据时，会调用 runtime.chansend 函数。该函数会判断 channel 是否已满，如果满了则让 goroutine 进入阻塞等待。

2. 当 goroutine 从 channel 接收数据时，会调用 runtime.chanrecv 函数。该函数会判断 channel 是否为空，如果为空则让 goroutine 进入阻塞等待。

3. 运行时调度器通过维护等待队列来管理被阻塞的 goroutine。当 channel 可读写时，会从等待队列中弹出 goroutine 继续运行。

4. 运行时调度器以非常高频率运行，会在 goroutine 之间快速切换。这给用户的感觉就是 goroutine 并发运行。

5. channel 的发送接收动作必须匹配，否则程序会死锁。调度器试图重排发送接收顺序来避免死锁。

总结来说，channel 的阻塞是建立在运行时调度器的基础上。调度器管理被阻塞的 goroutine 队列，在条件允许时恢复 goroutine 继续运行。channel 的使用保证了数据的同步传递。

### 如果去掉 channel，goroutine 就可以真实的并行了吗

如果没有 channel 的话，goroutine 可以真正并行执行，不会被阻塞。

channel 的发送接收机制，会在对端未准备好时造成 goroutine 阻塞。这就会引起 goroutine 之间的同步关系。

而如果没有 channel：

1. goroutine 之间不存在同步依赖，可以任意交叉执行。

2. 程序需要自己维护竞争和同步逻辑，较难处理好并发问题。

3. goroutine 无法通过 channel 通信，必须使用共享变量，增加竞争风险。

4. 程序整体并发效果可能更好，但复杂度和错误风险都更高。

5. 丧失了 channel 带来的可组合性、表达能力等优势。

所以 channel 的引入让编程变复杂，但收获是可靠性和可维护性。需要根据实际情况权衡 channel 的引入带来的收益。

如果非要最大程度发挥并行，确实可以考虑去掉 channel，但成功的并发程序往往不是依靠纯并行来实现的。

这需要根据具体情形来权衡 channel 的引入带来的收益和复杂度。
### channel 底层实际上是互斥锁，阻塞机制也符合互斥锁对于 goroutine 的阻塞规则吗

channel 的发送接收在底层会使用互斥锁来实现同步。所以它继承了互斥锁一些内在的特性：

1. 阻塞的 goroutine 会形成 FIFO (先进先出) 的等待队列，这确保了公平性。

2. runtime 会优先唤醒等待时间最长的 goroutine，防止饥饿。

3. 发送和接收有时会采用短暂的自旋等待锁，避免过早休眠线程。

4. 发送方和接收方争用同一个锁，这会引起优先级反转等问题。

但是

channel 的实现要比直接使用锁更复杂，它内部做了很多优化：

1. 非阻塞发送和接收会被优先执行，避免不必要的阻塞。

2. 尽量避免线程休眠，减少锁的争用。

3. 多核心会并行匹配发送和接收等。

4. 可调节的缓冲区大小可根据场景优化吞吐量。

综上，channel 继承了互斥锁的一些内在特性，但它的实现要复杂很多，做了大量优化来减少锁争用的影响。

所以使用 channel 的信号量模式去实现的互斥锁并不适合，因为它比 go 本身的互斥锁更加复杂
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

channel 在 Go 语言中的底层实现主要涉及以下几个方面：

1. 数据结构：channel 在内部维护一个消息队列，用于存储缓冲区的数据 (有缓存的话)，同时有发送方、接收方的指针。

2. 同步机制：使用互斥锁 (Mutex) 和条件变量 (Cond) 实现同步，保证并发安全。

3. 调度机制：通过运行时调度器管理 goroutine 的阻塞和唤醒，维护等待队列。

4. 内存管理：与运行时内存管理紧密结合，缓冲区数据的内存分配与释放。

5. 性能优化：负载均衡、本地队列等方式提高性能。

具体而言，channel 的数据结构 hchan 中包含以下字段：

- buf：循环队列，维护缓冲区
- sendx/recvx：发送和接收指针
- lock：互斥锁
- sendsema/recvsema：发送和接收的条件变量

发送接收时需要获取互斥锁，阻塞时等待条件变量。调度器 basis 上维护等待队列。

### 编程题，使用三个 goroutine 打印 abc 100 次
上文提到的 channel 任务编排中的 pipeline 流水线模式完美解决这个问题。
```go
package main

import (
	"fmt"
)

type Token struct{}

var (
	GNumber int
	Number  int
	sign    = make(chan struct{})
)

func world(id int) {
	switch id {
	case 0:
		fmt.Println("a")
	case 1:
		fmt.Println("b")
	case 2:
		fmt.Println("c")
  }
}
func worker(id int, in chan Token, out chan Token, number int,fn func(int)) {
	for number > 0 {
		token := <-in
		fn(id)
		if id == (GNumber-1) && number == 1 {
			close(sign)
			break
		} 
		out <- token
		number--
	}
}
func RunPipeline(GNumber int,Number int,chs[]chan Token,fn func(int)) {
  
	for i := range GNumber {
		go worker(i, chs[i], chs[(i+1)%GNumber], Number,fn)
	}
	chs[0] <- struct{}{}
	<-sign

}
func main() {
  GNumber = 3
	Number = 100
  chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
  RunPipeline(GNumber,Number,chs,world)
}
```
包括之前的水生成工厂那道题，也可以使用 pieline 模式来解决，替换掉之前使用的 waitgroup 和信号量，以及循环栅栏
## 参考资料
- https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
- https://github.com/fortytw2/leaktest
- https://cloud.tencent.com/developer/article/1921580