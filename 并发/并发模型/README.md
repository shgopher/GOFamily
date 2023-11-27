<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-11-27 16:14:01
 * @FilePath: /GOFamily/并发/并发模型/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# go 并发模型
## 并发和并行的关系
并发是问题域的一种概念，它强调处理多个同时 (或者近似同时) 发生的事件。

并行是方法域的一种概念，将问题分解为多个部分，同时并行执行来加速解决问题。

> rob pike：并发不是并行，并发关乎结构，并行关乎执行

并发：一位老师，在听一个学生朗读的时候，她可以暂停学生的朗读，然后回答学生的问题，再次开始学生的朗读，虽然她一次只能干一件事，但是这也能看作处理多个近乎同时发生的事件

并发和并行：两位老师，一个老师提问，一个老师解决学生的问题，这就是满足了并行和并发

并行但不是并发：让全班同学制作贺卡，全班同学每个学生制作五枚，全班同学同时开始做，只能是并行，但不是并发，因为只有一个事件。

## 多线程并发模型
使用共享内存的方式去完成并发就是多线程并发模型，它的核心就是使用锁的方法，让某个线程单独拥有某块内存，其他线程只能访问该内存，从而实现了并发。

go 语言中的锁就是 sync.Mutex，这也是 go 语言实现多线程并发的核心，一共有：
- sync.Mutex 互斥锁，可以同时对一个资源进行读写操作，但是只能有一个线程可以对该资源进行写操作。
- sync.RWMutex 读写锁，可以同时对一个资源进行读操作，但是只能有一个写操作
- sync.Cond 条件变量，可以让一个线程等待另一个线程满足某个条件
- sync.Once 单例模式，保证某个资源只被初始化一次
- sync.Pool 资源池，可以让一个资源在内存中被复用，避免了重复创建资源的开销
- sync.Map 线程安全的 map，可以让多个线程安全的对 map 进行读写操作
- sync.Pool 资源池，可以让一个资源在内存中被复用，避免了重复创建资源的开销
- sync.WaitGroup 等待组，可以让多个线程等待，直到某个线程完成某个任务
- golang.org/x/sync/errgroup 为处理公共任务的子任务的 goroutine 组提供同步、错误传播和上下文取消
- golang.org/x/sync/semaphore 提供了一个加权信号量实现。
- golang.org/x/sync/singleflight 提供了重复函数调用抑制机制，中文叫栅栏机制
- golang.org/x/sync/syncmap 提供了一个并发映射实现
## csp
go 语言推荐的并发模型使用的就是 csp 模型，csp 的核心思想就是讲各个任务等同于进程，进程顺序执行互不牵连，进程可以收发信息，使用通道的方式进行信息的通信。

所以如果使用 channel 的方式进行通信就是使用的 csp 并发模型

CSP 模型中的进程通信原语包括：

- 发送消息：一个进程可以通过发送消息到另一个进程来与之进行通信。
- 接收消息：一个进程可以通过接收消息来获取另一个进程发送的消息。
- 原子发送-接收：一个进程可以通过原子发送-接收操作来发送消息并等待接收消息，这相当于发送和接收两个操作的组合。

这些进程通信的都是通过内置的 channel 对象去实现的。
## 了解 goroutine
这里我们现简单的了解一些基本的使用 goroutne 的方法，后面的 channel 篇和并发原语 context atomic 定时器会进行更加详细的介绍。

我们知道 go 使用了用户线程也就是 goroutine 去替代了传统的线程，所以在 go 语言中我们能操作的线程就是 goroutine，我们无法去触及真实的线程，线程和 goroutine 之间的关系是 go 语言运行时的调度器去调度的。
### 创建 goroutine
使用 go 关键字加上函数去创建一个 goroutine，当然后面跟方法也可以。

```go
func age() {
	// 注意后面跟的是一个函数的运行，这跟 defer 一致
	c := make(chan int, 1)
	go func() {
		time.Sleep(10000)
		c <- 12
	}()
	d := <-c
	fmt.Println(d)
}
```
我们可以看到，这里使用了 csp 的并发模型，下面我们看一下使用传统的共享内存的并发模式

```go
func age() {
	// 注意后面跟的是一个函数的运行，这跟 defer 一致
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(i)
		}(i)
	}
	time.Sleep(200)
}
```


### goroutine 退出
goroutine 使用代价很低，因为它并不是操作系统的线程，创建成本非常低，go 推荐可以多多使用 goroutine

goroutine 退出有两种方式：
- 主动退出：使用 return 或者 panic 关键字退出
- 非主动退出：使用 sync.WaitGroup，context 等方法，当所有的 goroutine 都退出后，等待组会自动退出

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
    defer wg.Done()
	}()
    wg.Wait()
}
``` 
go goroutine 执行完毕就会直接退出，程序的执行跟主 goroutine 有关，只要主 goroutine 不退出程序就会正常执行下去，反之，主 goroutine 如果退出了，其它的 goroutine 可能没有执行完毕，但是整个程序还是结束了。

```go
func main() {
	go age()
}
func age() {
	time.Sleep(10000)
	fmt.Println("hi there")
}
```
这个程序将无法保证能输出 hi there

要想让 age 输出正常的值，必须保证主 goroutine 不能退出，比如使用 sync.WaitGroup，比如直接让主 goroutine 休眠

```go
func main() {
	go age()
  time.Sleep(20000)
}
func age() {
	time.Sleep(10000)
	fmt.Println("hi there")
}
```
下面让我们正式学习 go 并发语言，channel 等相关知识！
## 参考资料
- https://mp.weixin.qq.com/s/TvHY2i1FX1zS_WHdCvK-wA
- https://book.douban.com/subject/26337939/ 
- https://book.douban.com/subject/35720728/ 315 页 - 317 页
- 极客时间《go 进阶训练营》