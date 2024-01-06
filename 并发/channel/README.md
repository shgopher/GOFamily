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

## select
## 定时器
## 数据交流
## 消息传递
## 信号通知
## 任务编排
## 锁
## channel 的实现原理
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