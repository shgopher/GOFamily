<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-12-25 23:04:39
 * @FilePath: /GOFamily/并发/同步原语/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 同步原语
传统同步原语是 go 提供的相对底层的同步机制，它更加灵活，但是同时也更加复杂，如果可能的话，我们应该尽量使用 csp 的并发模型，使用 channel 去代替传统同步原语。

本单元讲的传统同步原语，channel 和 contex 也属于同步原语，他们属于 csp 的并发模型，所以他们会单独为列。

下面讲一下这么多同步原语 (也叫并发原语) 的基本功能：

- `sync.Mutex, sync.RWMutex` 共享资源，避免数据竞争 (data race)
- `sync.Waitgroup, channel` 任务编排，各个 goroutine 所代表的任务的前后顺序关系
- `channel` 传递消息

这是基本的划分，当然这个划分还不严谨，但是你只需要知道，这些属于最常见的同步原语，以及他们最常见的功能。
## sync.Mutex
下面介绍的众多并发原语，甚至下一章的 channel，都使用了这个核心内容，它就是 `sync.Mutex`，它是所有同步原语包含 channel 的底层核心。

Mutex 也就是互斥锁，它主要是为了解决多线程下数据的竞争问题，所以互斥锁是同步原语的最底层最核心的组件

让我们看几个场景

1. 计数器，多个线程去更新一个计数器的时候，如果不加锁，就会出现数据的错误，本来你只加上了 1，正当你读的时候你发现别的线程也加上了 1，所以导致读取的错误

2. 秒杀服务，这也是一个常见的场景，如果不加锁，就会导致超卖的情况出现

3. 往一个 buffer 中传入数据，如果不加锁，数据的顺序就会发生乱序
### 临界区
临界区的概念是指在多线程的时候，临界区域的内容会被不同的线程去获取和释放，这个区域会发生数据的争夺问题。

这个区域因为会发生争夺，所以会被保护起来，可以这么说，临界区是**多线程中整体数据中的共享部分**

所以临界区是要保护的区域，**一次只能让一个线程去使用**

所以 sync.Mutex 就是用来保护临界区的，它可以保证临界区的互斥

共享资源通常是某个变量，例如临界区是变量 count，临界区操作是 count++，只要在临界区前面获取
锁，在离开临界区的时候释放锁，就能解决 data race 的问题。

### sync.Mutex 基础操作
我们现看一下它的基础使用功能：
```go
func age(){
  var mu sync.Mutex
  for i := 0; i <10; i++ {
    go func(){
      mu.Lock()
      defer mu.Unlock()
    }()
  }
}

// 跨goroutine 加解锁
func age1() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	go func() {
		defer mu.Unlock()
		time.Sleep(1000)
		fmt.Println("hi there")
	}()
	mu.Lock()
}
```
sync.Mutex 即为互斥锁，规则是：

- 锁的加锁和解锁可以跨 goroutine 使用，比如 a goroutine 加锁，在 b goroutine 将锁解开。
- 只有现解锁才能继续上锁，happens-before 就是：***n 次解锁一定 happens-before n+1 次加锁***

### sync.Locker 接口
```go
type Locker interface{
  Lock()
  Unlock()
}
```
Mutex 就实现了这个 Locker 接口

Locker 定义了锁的基本方法，加锁 + 解锁

### 什么是 data race 的本质
我们常说 data race 的情况，是多线程同时对于某块内存进行数据的变更，那么问题来了，这个地方谈到的同时，是真的物理层面的同时还是近似同时？

关于 data race 中的 “同时” 通常是**指逻辑上的同时或近似同时，而不是物理层面严格的同一时刻。**

主要原因有以下几点：

- 现代 CPU 中，同一时刻真正执行指令的只有一个核心。不同核心之间以及同一核心的不同执行周期之间，不存在物理层面严格的同步。
- 即使在单核 CPU 上，由于指令流水线、内存缓存、分支预测等机制，实际执行顺序也可能与代码顺序不一致，很难定义物理层面严格的同步。
- 不同线程之间进行切换的时间间隔非常小 (几十到几百纳秒)，对程序逻辑而言可以视为同时进行。
- 要构成 data race，不同线程对同一地址的访问之间间隔不能太长，必须在一个指令/操作的启始和完成之间，所以也符合逻辑上的近似同时。

所以，data race 中的 “同时” 就是指逻辑上近似同时，或者无法确定准确执行顺序的情况，而不是物理层面严格同一时刻。这种近似同时从程序角度就是可能造成冲突，需要进行同步处理。
### 检测 data race 的工具
并发问题不是一定能肉眼看出来的，如果只是基础的，容易看出来的也就罢了，但是很多隐藏的 data race 问题必须使用专业的工具去鉴别，go 语言提供了 `-race` 功能，在编译，测试，run 的时候，会自动检测到 data race 问题，并且给出详细的错误信息。

```bash
 go run -race main.go
```

我们看一个例子
```go
func main() {
	value := 0
	for i := 0; i < 10000; i++ {
		 go func(){
			value ++
		 }()
	}
	time.Sleep(1000)
	fmt.Println(value)
}
```
本能的你会以为能输出 10000，但是结果确实 9000 多，而且还不一定，这是为啥呢？

因为你以为 ++ 操作是原子操作，其实并不是。

++ 操作分为三个步骤
- 获取 value 值
- 值+1
- 将新值写回

这其实是三个步骤，10000 个线程，假如同时有 10 个去读了这个 value，在他们看来都是初始值是 0，然后他们+1，然后写回去了结果 value 是 1，相当于 10 个 goroutine 都去写，本来应该是 10，但是赋值回去都变成了 1

这个时候，如果你使用 run -race 就能检测出来

```bash
go1 go run -race main.go 
==================
WARNING: DATA RACE
Read at 0x00c00010a018 by goroutine 8:
  main.main.func1()
      /Users/shgopher/Desktop/1/go1/main.go:23 +0x2c

Previous write at 0x00c00010a018 by goroutine 6:
  main.main.func1()
      /Users/shgopher/Desktop/1/go1/main.go:23 +0x3c

Goroutine 8 (running) created at:
  main.main()
      /Users/shgopher/Desktop/1/go1/main.go:22 +0x48

Goroutine 6 (finished) created at:
  main.main()
      /Users/shgopher/Desktop/1/go1/main.go:22 +0x48
==================
9957
Found 1 data race(s)
exit status 66
```
多线程多某个区域的内存进行同时 (或者近似同时) 操作，这就是数据竞争

使用这个内置工具有个很大的缺陷，就是只有在数据真的执行中发生了数据竞争才能被发现，并且，使用-race 会影响编译的程序执行速度

如果我们使用 `go tool compile -race -S x.go` 的时候就发现使用-race 之后，编译的时候 go 编译器往代码中加入了很多运行时监控代码，**这些运行时的监控代码**会影响性能
```go
	0x001c 00028 (a.go:3)	PCDATA	$1, $0
	0x001c 00028 (a.go:3)	CALL	runtime.racefuncenter(SB)
	0x0020 00032 (a.go:4)	MOVD	$type:int(SB), R0
	0x0028 00040 (a.go:4)	CALL	runtime.newobject(SB)
```
`runtime.xx` 的代码就是添加的运行时监控
> 小知识，使用 go tool compile 的时候不要加入第三方库，标准库也不行，编译工具只能编译本文件，跟 go build 那种能查找依赖树的方式不同

### 将互斥锁嵌入到结构体中
```go
type MyAge struct {
  mu sync.Mutex
  value int
}
```
或者直接嵌入
```go
type MyAge struct {
  sync.Mutex
  value int
}
```
```go
func main() {
	var age MyAge
	for i := 0; i < 10000; i++ {
		go func(){
			age.Lock()
			defer age.Unlock()
			age.value ++
		}()
	}
	time.Sleep(1000)
	fmt.Println(age.value)
}
```
### sync.Mutex 互斥锁的实现原理
go 语言互斥锁的实现非常简单，只有这一个结构体就是核心：
```go
type Mutex struct {
  state int32
  sema uint32
}
```
> 在阅读下面互斥锁的几个阶段之前，建议先读一下 G:M:P 模型

#### 互斥锁演变的四个阶段一：简单实现
```go
// 最初的代码
type Mutex struct {
  state int32 // 锁是否被持有的标志，1 被持有，0 没有被持有
  sema uint32 // 锁的具体状态，此信号量具有高级语意，用来控制goroutine的状态
}
// compare and swap 操作：val 和 old 进行对比，如果相同，使用new去替代 val的值
func cas(val *int32, old,new int32)bool{}
// val 数据添加一个 delta 值，返回新的值
func xadd(val *int32, delta int32)(new int32){
  for {
    v := *val
    if cas(val,v,v+delta) {
      return v+delta
    }
  }
  panic("unreached")
}

func (m *mutex)Lock() {
  if xadd(&m.key,1) == 1 { // 标识+1 ，如果等于1 获取到锁
    return 
  }
  // 这里就是说，当key >1 的时候，我们通知goroutine休眠等待锁
  // 只有key 等于 1 才能算获取锁
  semacquire(&m.sema)
}
func (m *mutex)Unlock() {
  if xadd(&m.key,-1) == 0 { // 标识-1 ，如果等于0 表示没有其它等待者
    return
  }
  // 这个函数是汇编语言写的，上面那个 semacquire 也是
  semrelease(&m.sema) // 唤醒等待锁的其它的goroutine中的一个
}
``` 

可以看到，这种简单的实现，完全没有考虑任何的情况，仅仅是简单的加锁和解锁，不考虑 goroutne 的任何情况，就是随机的，随机的获取锁然后解锁。

注意 go 语言的锁可以 a 加 b 解，所以一定要谁加锁就谁解锁，不然就无法构成互斥锁这个概念了。
#### 互斥锁演变的四个阶段二：优先新 goroutine

#### 互斥锁演变的四个阶段三：优先新 goroutine 以及被唤醒的 goroutine

#### 互斥锁演变的四个阶段四：在第三个阶段的基础上引入饥饿模式

## sync.RWMutex
## sync.Locker
## sync.WaitGroup
## sync.Cond
## sync.Once
## 讨论 map 在多线程中的场景
## Pool
## errgroup
## semaphore
## singleflight
## syncmap
## issues
### 问题一：有互斥锁就一定有临界区吗？
互斥锁的存在不等于必须存在临界区。

构成一个合理的临界区，需要满足：

- 有真正需要互斥访问的共享资源 (比如共享内存，变量等)
- 通过加锁，在访问该资源前后形成互斥的代码区域
- 确保同一时间只有一个线程/goroutine 可以进入该互斥区域

所以互斥锁只是手段之一，用于保证临界区互斥性的需要。

如果没有需要保护的共享资源，或者互斥逻辑不严密，那么使用再多的锁也不等于形成了临界区
### 问题二：如果 Mutex 已经被一个 goroutine 获取了锁，其它等待中的 goroutine 们只能一直等待。那么，等这个锁释放后，等待中的 goroutine 中哪一个会优先获取 Mutex 呢？

上文中的锁的饥饿模式和正常模式可以解释这个问题。

如果是正常的模式下，就是按照正常队列 FIFO 的顺序去获取锁，除非这个时候有新的 goroutine 生成，那么这个 goroutine 会优先获取锁。

但是如果一个队头的 goroutine 等待时间过长超过了 1ms，那么它就会将互斥锁的模式变成饥饿模式，自动获取锁
### 问题三：Mutext 的底层中，使用 state 和 sema 来表示锁的状态，sema 是信号量，为什么在有信号量表示锁的状态之后还需要一个 state 表示锁是否上锁呢？

- state 作为一个 boolean 变量，可以非常简单直观地表示锁的基本状态。
- sema 是一个整数计数器，可以灵活地表示多种状态，如等待队列长度等。
- 将两者分开，可以清晰地分离基本锁状态和高级同步语义，符合分而治之的设计思想。
- sema 可直接 reused 现成的信号量实现代码，state 又足够轻量不需要复杂机制。
- 将两者组合可以充分发挥各自的优势，实现一个功能完备但设计简单的 mutex。
- 如果全部只依赖 sema 来表示所有状态，实现可能会更复杂，语义也不够清晰。
## 参考资料
- https://mp.weixin.qq.com/s/iPpWd8vjyaN2sJFwxzN9Bg
- https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/
- https://time.geekbang.org/column/intro/100061801
- https://colobu.com/2018/12/18/dive-into-sync-mutex/
- 《go 语言精进之路》