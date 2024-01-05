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
### channel 分为三种类型
- 只能发送的 channel `chan<-`
- 只能接收的 channel `<-chan`
- 可发可收的 channel `chan`

channel 里面理论上可以存在任意的类型数据，当然也包括 channel 自己，那么当 channel 中出现了很多符号时该如何区分呢？
```go
chan<- chan int 
```
这里有个规则，符号总是跟左侧的 chan 在一起表示含义，比如这里，就是一个只能发送的 channel，里面保存的是一个 int 类型的 channel

我们知道，channel 可以拥有三种不同的类型，能发能收的 channel 是满足前两者的，也就是说，当你声明的是一个只能接收的 channel，你完全可以传入一个能收能发的 channel，但是当你声明的是一个能收也能发的 channel 时，就不能传递进去一个只能发或者只能收的 channel
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



## channel 的实现原理
## select
## 定时器
## 基本操作
## 数据交流
## 消息传递
## 信号通知
## 任务编排
## 锁
## 注意事项
## issues
### 有无 buffer 的 channel 区别
### channel close 后，read write close 的区别
### channle 底层是什么
### channle 和运行时调度如何交互
## 参考资料
- https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
- https://github.com/fortytw2/leaktest
- https://cloud.tencent.com/developer/article/1921580