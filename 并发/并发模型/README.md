<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-02-04 15:55:30
 * @FilePath: /GOFamily/并发/并发模型/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# go 并发模型
## 并发和并行的关系
并发是问题域的一种概念，它强调处理多个同时 (或者**近似同时**) 发生的事件。

并行是方法域的一种概念，将问题分解为多个部分，同时并行执行来加速解决问题。

并发可以不是同时进行的，但是并行强调的就是必须同时进行。

> rob pike：并发不是并行，并发关乎结构，并行关乎执行

并发但不并行：一位老师，在听一个学生朗读的时候，她可以暂停学生的朗读，然后回答学生的问题，再次开始学生的朗读，虽然她一次只能干一件事 (所以不满足并行)，但可以处理近乎同时发生的***多个事件***

并行但不是并发：让全班同学制作贺卡，全班同学每个学生制作五枚，全班同学同时开始做 (满足并行)，但不是并发，因为只有一个事件 (不满足并发提出的同时处理多个事件)

并发和并行：两位老师，一个老师提问，一个老师解决学生的问题，这就是满足了并行和并发：同时 (这里是近似同时) 处理多个事件：提问和解决问题 (满足了并发)，两个老师同时开始这满足了物理层面的同时进行 (满足了并行)

可以看出来，并发强调多个事件，并行强调物理层面的同时执行

获得真正的并行，必须在具有多个物理处理器的计算机上运行程序，针对单个处理器的计算机，只能实现并发执行

go 在执行并发任务时，如果所在物理机器为多核，那么并行的数量就等于 `runtime.GOMAXPROCS` 的值，如果机器为单核，那么只能执行并发操作了

在使用中，***并发且并行***是最常见的行为，首先，我们通常不会只执行一个事件，并且所用机器不太可能是单核

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

goroutine 是一种轻量的可以被大量创建的用户态线程。
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
### 异步函数决定权交给函数调用方
我们看一个场景：

我们要读取一个目录下的路径，首先我们可以这么写函数：
```go
func ListDirectory(dir string)([]string,error)
```
这是一个同步函数，我们传入的是目录的地址，返回的是值和错误，只不过我们需要阻塞的等待所有的路径全部扫描完成才能返回

如果我们不想让业务阻塞到这里，可以改造成异步函数：
```go
func ListDirectoryAsync(dir string)chan string{
	go func(){
		//
	}()
	return c
}
```
将数据传递给 channel，只需要不断的去读取 channel 就可以变成非阻塞的业务。

不过这里我们发现还是会有一些问题，比如，如果我读取到了想要的数据想结束这个函数，该如何操作呢？读取过程中我如何分辨是读取完成了 close 掉了这个 channel 还是出现了错误 close 这个 channel 呢，所以这个函数还是需要改造

我们可以将这个函数设置成一个同步函数，让调用者来决定是否异步的启动新 goroutine 去调用这个函数，这给了程序更大的灵活性

***将异步执行函数的异步权交给调用方***是更好的设计思想，

因为如果这个函数内部启动了一个 goroutine，但是它并没有提供给你详细的退出机制，那么非常容易出现 goroutine 的泄漏问题

```go
func ListDirectory(dir string, fn func(path string, info os.FileInfo, err error)) {
	info, err := os.Lstat(dir)
	if err != nil {
		err = fn(root, nil, err)
	} else {
		err = walk(root, info, fn)
	}
	if err == SkipDir || err == SkipAll {
		return nil
	}
	return err
}

// 同步调用
func retrieveData(root string) (value []string, err error) {
	// 使用一个切片来存储结果
	var result []string

	// 调用ListDirectory，这里不再使用goroutine
	err = ListDirectory(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果文件不是普通文件，直接返回nil
		if !info.Mode().IsRegular() {
			return nil
		}
		// 将路径添加到结果切片中
		result = append(result, path)
		return nil
	})

	// 如果没有错误，返回结果切片
	if err == nil {
		return result, nil
	}
	// 如果有错误，返回错误
	return nil, err
}

// 异步调用调用：

func retrieveData(root string) (value chan string, err chan error) {
	err = make(chan error, 1)
	value = make(chan string)
	go func() {
		defer close(value)
		// 调用时再决定是同步还是异步
		err <- ListDirectory(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// if the file is noe regular, it mean the file is done,you should return
			if !info.Mode().IsRegular() {
				return nil
			}
			value <- path
			return nil
		})
	}()
	return
}

```

可以看上文，同步操作也可以，我们也可以使用 `go func` 的方式异步执行它，因为要传入一个函数，所以如果我们使用了异步，就在函数中使用 channel 来传递结果，如果我们是同步，那么就不再使用 channel，使用一个切片即可


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

go goroutine 执行完毕就会直接退出，

***程序的退出跟主 goroutine 有关，只要主 goroutine 不退出程序就会正常执行下去，反之，主 goroutine 如果退出了，其它的 goroutine 即便没有执行完毕，整个程序还是会结束。***

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