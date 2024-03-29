# 内存模型
go 官方介绍 go 内存模型的时候说：探究在什么条件下，A goroutine 在读取一个变量的值的时，能够看到其它 goroutine 对这个变量进行的写的结果。

我们为什么需要内存模型？

CPU 指令可能会被重排序，并且存在多级缓存，例如 Go 语言中的多级内存模型。不同的 CPU 架构 (比如 x86 和 ARM) 也会对指令顺序产生影响，编译器优化也可能改变指令顺序。

因此，编程语言需要一个内存规范，也就是内存模型，来规定程序中的内存访问和同步的语义，保证在不同的硬件和编译器下，程序的执行结果保持一致

## 介绍一些 go 语言中的基本操作
> 除了基本的读和写之外都是同步操作，包括以下列举的，并且更多以这些基础操作衍生出的操作。

- 读
  - 基础的读
    - 对超过机器 word 大小的值读，可以看作是拆成 word 大小的几个读的无序进行
  - 原子包的读
  - 使用 sync.Mux 的读
  - 使用 channel 形式的读
- 写
  - 基础的写
    - 变量的初始化就是一个写操作
    - 对超过机器 word 大小的值写，可以看作是拆成 word 大小的几个写的无序进行
  - 原子包的写
  - sync.mux 的写
  - 向 channel 发送数据
- atomic 的 compare and swap 操作，兼顾了写和读


我们讨论 go 内存模型的时候，主要聚焦下面四种细节的**差异性**：

- 操作的种类：
  - 普通的读或者写
  - **同步操作**，比如利用 sync 包进行同步，利用 channel 在不同的 goroutine 中间进行同步，利用 atomic 包进行同步等。
- 这个操作者在程序中的位置
- 被操作者在内存位置，或者是被访问的变量的位置
- 被操作者值的类型


## 内存模型存存在的意义

- 给程序员一个最根本的法则 --- 你只要遵循这个法则，那么在进行多 goroutine 数据访问的时候，做串行控制 (使用同步操作) 一定会成功。

- 给编译器优化开发者一个法则，只要遵循这个法则，就能不出错的做出编译器级别的优化。

### 具体表现
- **探究可见性**
- **happens- before**：x 操作一定在 y 操作之前执行完毕，这里说的并不是时间，而是 “执行完毕” 这个结论，我们可以放心的在 y 中使用 x 的执行结果。主要突出的是这个完成性。

### 三条定律
- 在一个 goroutine 中，程序执行的顺序一定是符合代码写的顺序的。但是不同的 goroutine 中这个定理失效。

- 在**同步操作**的时候，从 a goroutine 看到 b goroutine 的执行顺序，必须符合同步操作的顺序，这里举个例子，举例：b 的同步操作是什么顺序，那么 a 读取的就是什么顺序。比如 b 中有三个 Channel，他们执行的顺序是先 a 发 b 收，b 发，c 收，那么在 a 中读取的顺序跟这个同步的操作顺序是一致的，**如果不是同步操作就可能不一致。**

- 对于非同步操作的普通读和写，如果 b 操作了 a，那么必须成立
  - b 发生在 a 之前

我们可以使用 `go build -race` 来发现数据竞争。

接下来我们对于内存模型的具体表现进行更具体的介绍。
## 重排和可见性

由于指令的重排，代码**不一定**会按照你写的顺序执行。

这里给出一个案例：有两个 goroutine 分别是 g1，和 g2，g1 读某个变量的数据，g2 写这个变量的数据，当 g1 读取到 g2 写的新数据之后，你也不一定能读取到在 goroutine g2 中排列在这个写数据操作之前的操作：

```go
package main

import (
	"fmt"
)

var (
	a    = 1
	b    = 2
	done bool
)

func main() {

	go func() {
		a = 3
		b = 4
		done = true
	}()
	for !done {
	}
	fmt.Println(b)
	fmt.Println(a) // 可能观察到的就是输出的 4 3，但是无法保证必须是4 3 ，能观察到 != 保证
}

```

下面这种情况，a goroutine 无法观察到 b goroutine 是否完成了写操作，有可能造成 a goroutine 一直被卡的现象：

```go
package main

import (
	"fmt"
)

type Result struct {
	data string
}

var (
	r *Result
)

func set() {
	d := new(Result)
	d.data = "hi there" // 这里无法确认一定会运行
	r = d // 这里无法确认观察到
}

func main() {
	go set()
	for r == nil {
	}
	fmt.Println(r.data)
}

```
>there is no guarantee that the write to done will ever be observed by main，**since there are no synchronization events between the two threads**。The loop in main is not guaranteed to finish。“

这一句是 go blog 中对于这代代码的描述，重点是说，这俩线程中，没有存在同步原语，所以无法做到从 main goroutine 去观测到 r 的状态，也就是说，在 goroutine 中只要存在了同步原语，那么就可以把两个 goroutine 当作正常的语句来看。

这里有两点无法确认，第一：主 goroutine 无法**确认** r 的状态，这里指的就是**可观测性**，因为这里是非同步操作，所以主 goroutine 的读无法观测到 `go set（）` 这里 `r = d` 这里的写，注意是无法完全确认，通常来说这段代码是可以运行的，我说的是通常，第二无法保证 `d.data` 一定会运行，所以有可能输出的是空字符串，这里说的就是因为编译优化导致的指令重排。
## 同步操作的 happens-before

go 不直接提供 cpu 屏障，来保证编译器或者 cpu 保证顺序性，而是使用不同架构的内存屏障指令来实现统一的并发原语。

### 单个 goroutine
上文说到，**在一个 goroutine 中，happens- before 其实就等于代码书写的顺序，这一点是严格成立的**。

### init 函数
go 语言的初始化是在单一的 goroutine 中执行的，也就是 main goroutine。

**p 包导入了 q 包，那么 q 的 init 函数一定 happens before p 的任何初始化代码**。

这里有点像树，导入的过程，以及初始化的过程，会形成一个多叉树，从最底层的树开始进行初始化，逐步的忘上层走，先进后出，栈一样的感觉，并且相同的包只会导入一次，进而也只会初始化一次，比如 q 包被第三层导入了，在第五层也导入了，那么这个包在从底层往上走的过程中，只会在第五层先全局变量后 init 函数的初始化一次。这导入的包全部初始化一遍之后，才开始进行 main 包的 mian 函数，然后进而去运行接下来的逻辑。
```go
var(
  a = c + b

  b = f()
  
  c = f()

  d = 3 
)

func f() int{
  d ++ 
  return d 
}
```
初始化的顺序是这样的，首先初始化 a，因为 a = c + b，那么系统开始初始化 c 和 b，c 和 b 等于一个函数 f，那么就开始初始化这个函数，函数里面是 d 那么开始初始化 d，所以 b = 3+1 c = 3+1+1 a = 9，这个时候 d 已经++ 两次了，所以最终的初始化以后 a = 9 b = 4 c = 5 d = 5

包级别的变量按照代码顺序初始化，同一个包的众多文件会按照文件名的排列顺序进行初始化。
### goroutine
启动 goroutine 的 go 语言执行，一定 happens - before 此 goroutine 内的代码执行。退出可就没有任何 happens- before 的理论了，所以退出我们通常要部署同步语句。

```go
var a string
func f(){
  fmt.Println(a)
}
func main(){
  a = "hi"
  go f()
}
```
从单个 goroutine 的 happens- before 的关系来看，a = “hi” 一定 happens- before go f ()，从新 goroutine 的开辟来说，go f () 一定 happens- before fmt.Println 所以这个代码中，hi 一定能被输出。

### channel
go 语言有一句经典名言，不要使用共享内存来通信 (sync.Mutex) 而应该采用通信的方式来共享内存，这个后者说的就是 channel，channel 是同步操作的首选。

- 往 channel 中发数据，happens - before 从这个 channel 接收数据**完成**。注意这里说的是接收完成，**没说**发数据 happens- before 这个 channel 接受数据开始的时候。

```go
var c = make(chan int ,10)
var a string 

func f() {
  //这里之所以a = hi 发生在 c <-0 之前，就是因为这个goroutine和main goroutine 存在同步原语，
  //只要存在，那么代码就会按照程序员写的顺序执行，并不会进行重排。
  a = "hi" 
  c <-0  // 这里因为有了同步操作，所以 a = “hi” 在main goroutine看 也是 a= “hi” happens- before c <- 0
}
func main(){
  go f()
  <-c // 这里只要收到数据了，那么 c <- 0 肯定已经运行了，并且发送完毕了。
  print(a) 
}
```

- close 一个 channel 一定 happens- before 从关闭的 channel 中读取一个零值。我们都知道可以从一个 closed 的一个 channel 中读取零值是可以的，这里说明的是读零值这个过程一定在关闭这个操作之后。

- 一个 unbuffered 的 channle，读的准备好一定 happens-before 发准备好，意思是说，只有收数据准备好了，才能发，否则就会一直卡在发那个地方。

- 一个容量大于 0 为 m 的 channel，第 n 个接收，一定 happens- before 第 n+m 的发送，意思是说，如果容量满了，必须先拿出来一个才能往里面再塞进去一个。比如：n=1 m=2，第一个接收一定 happens- before 第三个发送，因为容量一共是 2，要想往里面塞进去第三个，**必须拿出来一个并且拿出来这个操作完毕了**，才能发送第三个。

### Mutex/RWMutex

- 第 n 次的 unlock 一定 happends- before 第 n+1 次 lock 方法的返回
意思是说，只有先解锁才能再次加锁。
- 读写锁虽然有两把锁，但是不能同时使用，必须等待一个锁解锁后，另一个锁才能上锁，比如读锁解锁后才能再次上读锁，或者才能上写锁。

互斥锁可以由 a goroutine 上锁，b goroutine 解锁

```go
var mu sync.Mutex
var s string

func f(){
  s = "hi"
  mu.Unlock() // 这里拥有了一个同步原语，所以 s = hi 一定发生在 mu.unlock() 之前。（相当于在一段锁的内部）
}
func main() {
  mu.Lock()
  go f()
  mu.Lock() // 这里的lock 一定发生在  另一个goroutine 中f 的 mu.Unlock() 之后。
  print(s)
}
```

### WaitGroup
- wait 方法等到计数值归零才能返回 (运行完毕)
### Once
- 对于 once.Do(f) f 一定会在 do 方法返回之前执行。
```go
var s string
var once sync.Once

func f(){
  time.Sleep(time.Second)
  s = "hi"
}
func main(){
  once.Do(f)
  print(s) // 这里 f 的执行完毕 一定在 Do()返回之前,所以一定能输出 s = hi
}
```
### atomic
按照 atomic load / store 顺序来保证 happens - before 不过 go 官方并没有严格定义，直到目前为止并没有严格定义。

```go

func main() {
  var a, b int32 = 0, 0

  go func() {
    atomic.StoreInt32(&a, 1)
    atomic.StoreInt32(&b, 1)
  }()

  for atomic.LoadInt32(&b) == 0{
    runtime.Gosched()
  }
    fmt.Println(atomic.LoadInt32(&a))
}
```
## issues

`使用channel实现一个互斥锁`

我们利用 channel 发送数据 happens- before 收到数据完毕这个特性来实现互斥锁。
```go
type Locker struct {
	ch  chan struct{}
}
// 初始状态是已经有一个
func NewLocker()*Locker{
	ch := make(chan struct{},1)
	ch <- struct{}{}
	return &Locker{
		ch: ch,
	}
}
// lock 从缓存是1编程0，从chan中取出来数据
func (l *Locker) Lock()  {
	<- l.ch
}
// unlock 将缓存从0变成1，向ch中放入数据。
func (l *Locker) Unlock() {
	l.ch <- struct{}{}
}
```

一个小小的经验：channel 拥有 buffer 比没有 buffer 常用很多。

完结。

## 参考资料
- https://go.dev/ref/mem
- https://time.geekbang.org/column/article/307469

