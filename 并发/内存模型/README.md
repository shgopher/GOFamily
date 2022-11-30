# 内存模型
go官方介绍go内存模型的时候说：探究在什么条件下，goroutine 在读取一个变量的值的时，能够看到其它 goroutine 对这个变量进行的写的结果。

我们为什么需要内存模型？由于cpu指令重排，以及多级的内存cache的存在，比如go语言存在的多级内存模型，不同的cpu架构，例如x86，arm 等等，而且编译器的优化也会对于指令进行重排，所以编程语言需要一个内存规范，即为：内存模型。

介绍一些基本的操作：
> 除了基本的读和写之外都是同步操作，包括以下列举的，并且更多以这些基础操作衍生出的操作。 

- 读
  - 基础的读
    - 对超过机器word大小的值读，可以看作是拆成word大小的几个读的无序进行
  - 原子包的读
  - 使用sync.mux 的读
  - 使用channel的收这种形式的读
- 写
  - 基础的写
    - 变量的初始化就是一个写操作
    - 对超过机器word大小的值写，可以看作是拆成word大小的几个写的无序进行
  - 原子包的写
  - sync.mux 的写
  - 使用channel的发
- atomic的 compare and swap 操作，兼顾了写和读


我们讨论go内存模型的时候，主要聚焦下面四种细节的**差异性**：

- 操作的种类：
  - 普通的读或者写
  - 同步操作，比如利用sync包进行同步，利用channel在不同的goroutine中间进行同步，利用atomic包进行同步等。
- 这个操作者在程序中的位置
- 被操作者在内存位置，或者是被访问的变量的位置
- 被操作者值的类型


内存模型主要用**两个目的**：

- 给程序员一个最根本的法则 --- 你只要遵循这个法则，那么在进行多 goroutine 数据访问的时候，做串行控制（使用同步操作）一定会成功。

- 给编译器优化开发者一个法则，只要遵循这个法则，就能不出错的做出编译器级别的优化。

实际上，我们在内存模型中主要探究的**具体表现**有两方面：

- **探究可见性**
- **happens- before**：x 操作一定在 y 操作之前执行完毕，这里说的并不是时间，而是“执行完毕”这个结论，我们可以放心的在y中使用x的执行结果。主要突出的是这个完成性。

这里有**三条定律**：
- 在一个goroutine中，程序执行的顺序一定是符合代码写的顺序的。但是不同的goroutine中这个定理失效。

- 在**同步操作**的时候，从a goroutine看到b goroutine的执行顺序，必须符合同步操作的顺序，这里举个例子，举例：b的同步操作是什么顺序，那么a读取的就是什么顺序。比如 b中有三个Channel，他们执行的顺序是 先 a 发 b 收，b 发，c 收，那么在a中读取的顺序跟这个同步的操作顺序是一致的，**如果不是同步操作就可能不一致。**

- 对于非同步操作的普通读和写，如果b操作了a，那么必须成立
  - b发生在a之前

我们可以使用 `go build -race` 来发现数据竞争。

接下来我们对于内存模型的具体表现进行更具体的介绍。
## 重排和可见性

由于指令的重排，代码**不一定**会按照你写的顺序执行。

这里给出一个案例：有两个 goroutine 分别是g1，和g2，g1读某个变量的数据，g2写这个变量的数据，当g1读取到g2写的新数据之后，你也不一定能读取到在 goroutine g2 中排列在这个写数据操作之前的操作：

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
这里有两点无法确认，第一：主 goroutine 无法**确认** r 的状态，这里指的就是 **可观测性**，因为这里是非同步操作，所以主 goroutine的读无法观测到 `go set（）` 这里 `r = d` 这里的写，注意是无法完全确认，通常来说这段代码是可以运行的，我说的是通常，第二 无法保证 `d.data` 一定会运行，所以有可能输出的是空字符串，这里说的就是因为编译优化导致的指令重排。
## happens-before

go 不直接提供cpu 屏障，来保证编译器或者cpu保证顺序性，而是使用不同架构的内存屏障指令来实现统一的并发原语。

上文说到，在一个goroutine中，happens- before 其实就等于代码书写的顺序，这一点是严格成立的。

### init函数
go语言的初始化是在单一的goroutine中执行的，也就是main goroutine ，**p 包导入了q包，那么q的init函数一定 happens  before p的任何初始化代码**。

这里有点像树，导入的过程，以及初始化的过程，会形成一个多叉树，从最底层的树开始进行初始化，逐步的忘上层走，先进后出，栈一样的感觉，并且相同的包只会导入一次，进而也只会初始化一次，比如q包被第三层导入了，在第五层也导入了，那么这个包在从底层往上走的过程中，只会在第五层先全局变量后init函数的初始化一次。这导入的包全部初始化一遍之后，才开始进行main包的mian函数，然后进而去运行接下来的逻辑。
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
初始化的顺序是这样的，首先初始化a，因为a = c + b ，那么系统开始初始化 c 和 b ，c 和b 等于一个函数f，那么就开始初始化这个函数，函数里面是d 那么开始初始化d，所以 b = 3+1  c = 3+1+1 a = 9 ，这个时候d已经++ 两次了，所以最终的初始化以后 a = 9 b = 4 c = 5 d = 5

包级别的变量按照代码顺序初始化，同一个包的众多文件会按照文件名的排列顺序进行初始化。
### goroutine
启动goroutine的go语言执行，一定happens - before 此goroutine内的代码执行。退出可就没有任何happens- before 的理论了，所以退出我们通常要部署同步语句。

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
从单个goroutine的happens- before的关系来看， a = “hi” 一定 happens- before go f（） ，从新goroutine的开辟来说，go f（） 一定happens- before fmt.Println 所以这个代码中，hi一定能被输出。

### channel
### Mutex/RWMutex
### WaitGroup
### Once
### atomic 
### 其它延伸的并发原语
## 总结




## 参考资料
- https://go.dev/ref/mem
- https://time.geekbang.org/column/article/307469

