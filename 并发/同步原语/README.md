<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-12-29 00:23:49
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
  key int32 // 锁是否被持有的标志，1 被持有，0 没有被持有
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
在这个演变中，go 的互斥锁调度会优先将锁权分给新创建的 goroutine
```go
const (
  // 1
  mutexLock = 1 << iota
  // 2
  mutexWoken
  // 2
  mutexWaiterShift = iota
)
type Mutex struct {
  state int32
  sema uint32
}

```
state 的内容变成了三个：
- 第一个字段：mutexWaiters 阻塞等待的数量
- 第二个字段：mutexWoken 唤醒标记
- 第三个字段：mutexLocked 持有锁的标记

我们先前知道，如果想要获取锁的 goroutine 没有机会获取到锁，就会进行休眠，但是在锁释放唤醒之后，它并不能像先前一样直接获取到锁，还是要和正在请求锁 goroutine 进行竞争。这会给后来请求锁的 goroutine 一个机会，也让 CPU 中正在执行的 goroutine 有更多的机会获取到锁，在一定程度上提高了程序的性能。

在这次演变中，go 的调度器会将新的 goroutine 给赋予锁，因为新的 goroutine 就是正在 cpu 执行片段中执行的 goroutine，这个时候将锁给他们无疑是效率最高的。

#### 互斥锁演变的四个阶段三：多给机会，优先新 goroutine 以及被唤醒的 goroutine

如果新来的 goroutine 或者是被唤醒的 goroutine 首次获取不
到锁，它们就会通过**自旋** (spin，通过循环不断尝试) 的方式，尝试检查锁是否被释放。在尝试**一定**的自旋次数后，再执行原来的逻辑。

#### 互斥锁演变的四个阶段四：在第三个阶段的基础上引入饥饿模式
为什么会加入饥饿模式呢？就是因为如果都优先给新的 goroutine，那么等待队列中的 goroutine 就永远不会被执行，所以引入了饥饿模式，优先执行等待中的 goroutine
然后新的 gorontine 就被添加到了等待队列中的队尾，这个时期称之为饥饿模式，因为新的 goroutine 基本上都要在 cpu 的时间片段中执行，所以在这个模式下，执行效率反而是底下的，因为正在执行的 goroutine 被强行放到队尾了。

当等待的队首的 goroutine 等待时间超过 1ms 就会进入这个模式

当队首的 goroutine 等待时间小于 1ms 或者已经执行到队尾了，那么这个模式就会从饥饿模式改为正常的模式
### sync.Mutex 易错的几种场景
#### Lock/Unlock 不是成对出现
因为 go 语言中，互斥锁是无法获取 goroutine 的信息的，所以存在 a 锁 b 解的情况，即：a goroutine 上了锁，b goroutine 给解开了。

如果你不是为了实现锁，是为了任务编排，那么可以这么做。

如果是为了锁，请不要这么做，因为这么做的后果就是这将不能形成锁这个概念

或者说当你使用了 lock 的时候忘记 unlock 了，那么最终都会导致系统走向失败
#### Copy 已经使用的 Mutex
go 语言的 mutex 使用 state 字段去表示锁的含义，所以当你 copy 一个锁的时候，实际上已经 copy 了这个锁的状态，这将导致错误的结果

go 语言的同步原语众多，使用的底层都是 mutex (包括 channel)，所以说，不仅仅是 mutex 不能使用 copy，其它的同步原语都不能。
#### 重入
所谓重入，就是多次上锁，注意这里是拥有锁的这个线程去请求这把锁

go 语言不支持重入，系统会 panic，这种重入锁无法实现也跟 go 语言的互斥锁没有记录使用它的 goroutine 有关系

那么如果 go 语言也实现一个重入锁，核心就是将持有锁的 goroutine 的 id 记录下来
#### 死锁
当多线程的情况下，多个线程陷入了争抢资源的情况，当他们都陷入了停滞状态，或者阻塞状态的时候，就会发生死锁，deaadlock

一般来讲，当你发现系统多个线程都堵死的时候，就会发生死锁情况了，但是通常发生死锁是发生在满足这四个情况下

- 互斥：资源具有排他性，只能有一个 goroutine 访问
- **持有和等待**：goroutine 持有资源，并还在请求其它资源
- 不可剥夺：资源只有被它持有的 goroutine 释放
- **环路等待**：发生了环路等待事件，下面这个案例可以解释这个理论

举一个案例
```go
package main

import(
  "sync"
  "fmt"
  "time"
)

func main(){
  var wg sync.WaitGroup
  var mu1 sync.Mutex
  var mu2 sync.Mutex
  wg.Add(2)
  go func(){
    defer wg.Done()
    mu1.Lock()
    defer mu1.Unlock()
    time.Sleep(1000)
    mu2.Lock()
    defer mu2.Unlock()
  }()
  go func(){
    defer wg.Done()
    mu2.Lock()
    defer mu2.Unlock()
    time.Sleep(1000)
    mu1.Lock()
    defer mu1.Unlock()
  }()
  wg.Wait()
}
```
在这个案例中，mu1 和 mu2 代表两个资源，两个 goroutine 在争夺这两个资源，下面我们盘点一下上文说的四个理论知识：

- 互斥性：资源只能被一个 goroutine 持有
- 持有和等待，一个 goroutine 获取了一把锁，还想获取第二把
- 不可剥夺，持有锁的 goroutine 释放锁后，其他 goroutine 不能再获取该锁
- 环路等待，两个 goroutine 陷入了环路这个概念总，第一个先持有 mu1，第二个 goroutine 先持有 mu2，他们又分别要获取另一个锁，所以陷入了环路等待中

![h](./ringwaiter.png)

所以这个案例中，发生了死锁
```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc0000a4020?)
	/usr/local/go-faketime/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0x0?)
	/usr/local/go-faketime/src/sync/waitgroup.go:116 +0x48
main.main()
	/tmp/sandbox1712910389/prog.go:31 +0x125

goroutine 17 [sync.Mutex.Lock]:
sync.runtime_SemacquireMutex(0x1?, 0x58?, 0x459218?)
	/usr/local/go-faketime/src/runtime/sema.go:77 +0x25
sync.(*Mutex).lockSlow(0xc0000a2028)
	/usr/local/go-faketime/src/sync/mutex.go:171 +0x15d
sync.(*Mutex).Lock(...)
	/usr/local/go-faketime/src/sync/mutex.go:90
main.main.func1()
	/tmp/sandbox1712910389/prog.go:20 +0xd1
created by main.main in goroutine 1
	/tmp/sandbox1712910389/prog.go:15 +0xb9

goroutine 18 [sync.Mutex.Lock]:
sync.runtime_SemacquireMutex(0x1?, 0x58?, 0x459218?)
	/usr/local/go-faketime/src/runtime/sema.go:77 +0x25
sync.(*Mutex).lockSlow(0xc0000a2020)
	/usr/local/go-faketime/src/sync/mutex.go:171 +0x15d
sync.(*Mutex).Lock(...)
	/usr/local/go-faketime/src/sync/mutex.go:90
main.main.func2()
	/tmp/sandbox1712910389/prog.go:28 +0xd1
created by main.main in goroutine 1
	/tmp/sandbox1712910389/prog.go:23 +0x119

```
### sync.Mutex 扩展
这里主要讲解如何对 sync.Mutex 进行功能扩展

基本的原理就是**将 sync.Mutex 嵌入到一个新的接口体中，然后重载 Lock 和 Unlock 的方法进行改造**

#### 改造一个重入锁
```go
type RecursiveMutex struct {
  sync.Mutex
  owner int64 // goroutine id
  recursion int64 //当前goroutine重入次数
}
```
lock 操作
```go
func(m *RecursiveMutex)Lock(){
  // 第三方包，目的是获取正在获取锁的lock 操作的 runtime id
  // github.com/petermattis/goid
  gid := goid.Get()

  // 判断当前goroutine是否是要重入的goroutine
  if atomic.LoadInt64(&m.owner) == gid {
    // 重入指标+1
    atomic.AddInt64(&m.recursion, 1)
    return
  }
  m.Mutex.Lock()
  // 获取得到锁的goroutine的id
  atomic.StoreInt64(&m.owner, gid)
  atomic.StoreInt64(&m.recursion, 1)

}
```
unlock 操作
```go
func(m *RecursiveMutex) Unlock(){
  gid := goid.Get()

  // 非持有锁的goroutine去释放锁直接panic
  if atomic.LoadInt64(&m.owner) == gid {
    panic("wrong",m.owner,gid)
  }
  // 重入指标-1
  atomic.StoreInt64(&m.recursion,&m.recursion -1)

  if atomic.LoadInt64(&m.recursion) != 0 { // 持有的goroutine还没有全部unlock
    return 
  }
  // 将 重入指标设置为-1
  atomic.StoreInt64(&m.recursion,-1)
  m.Mutex.Unlock()
}
```
## sync.RWMutex
上文我们提到了互斥锁，互斥锁是真的锁，不论是读还是写都只能有一个 goroutine 去操作，所以说互斥锁才是真的锁，那么我们这里讲的读写锁是什么含义呢？

读写锁：允许多个 goroutine 同时去读一个数据，但是这个时候是不允许写的；只允许一个 goroutine 去写一个数据，并且不允许其它 goroutine 去读这个数据。

所以，读写锁跟互斥锁相比，把读的权限给增大了，但是写的权限不变。

当我们遇到一个读多写少的场景，那么使用读写锁的效率要比互斥锁的效率高的多。

读写锁拥有以下几个方法：

- Lock/Unlock：写操作时加的锁
- Rlock/RUnlock：读操作时加的锁
- Rlocker：为读操作返回一个 Locker 接口的对象

RWMutex 跟 mutex 一样，初始值都是未加锁的状态，当然他们都是有状态的结构体，所以也不能复制锁，因为初始值就是未加锁，所以直接声明即可。

总结一下读写锁的几个注意事项
- 上文提到了，不可复制
- lock 和 unlock 要成对出现；Rlock 和 RUnlock 也要成对出现

下面我们看一下，使用读写锁造成的死锁问题

```go
// 这个写法出问题的原因是：读锁还没结束就开启了写锁
func main() {
	var mu sync.RWMutex

	mu.RLock()
	mu.Lock()
	mu.Unlock()
	mu.RUnlock()

}
```
出现死锁的原因就是出现了环形等待，读锁等待写锁解锁，写锁等待读锁解锁

## sync.WaitGroup
它的作用是任务的编排

waitgroup 一共有三个方法：

- Add(delta int)：增加 delta 个任务
- Done()：任务完成，减少一个任务
- Wait()：阻塞等待，直到所有任务完成

```go
func main() {
  var wg sync.WaitGroup
  wg.Add(2)
  go func(){
    defer wg.Done()
    time.Sleep(1000)
  }()
  go get(&wg)
  wg.Wait()
}

func get(wg *sync.WaitGroup) {
  defer wg.Done()
  time.Sleep(1000)
}
```
### 常见错误
- add 的时候数值传入负值
- done 的次数超过 add 中的次数
- 在 add 之前调用了 wait
- 当前一个 waitgroup 还没有完结就开始重用了 waitgroup，也就是说，必须等到上一轮的 wait 执行完毕了才能开启新的一轮

最后这个错误我们看一个案例：

```go
func main(){
  var wg sync.WaitGroup
  wg.Add(1)
  go func(){
    time.Sleep(1000)
    wg.Done()
    wg.Add(1) // AA
  }()
  wg.Wait()
}
```
在这个案例中，AA 处的 add 行为有可能发生在主 goroutine 之后，那么相当于一轮未结束又开启了新的一轮，就会发生 panic 行为

当然了，我不是说 add 只能调用一次，但是 add 虽然能调用多次，但是不能发生在 wait 之后

正确的多次调用 add 的案例：

```go
func main(){
  var wg sync.WaitGroup

  for i:=0;i<10;i++{
    wg.Add(1)
    go func(){
      defer wg.Done()
      time.Sleep(1000)
    }()
  }
  wg.Wait()
}
```
注意，我们这里每次循环都调用了一次 add，但是 add 的调用始终发生在 wait 之前，这还是属于同一轮的多次 add 调用，这符合 waigroup 的规定

## singleflight

## cyclicBarrier

## errgroup

## sync.Once
once 用来执行仅发生一次的动作，常用与单例模式，对象初始化的行为，并且经常在 init 函数中使用

sync.Once 仅仅暴漏了一个 do 方法，而且多次调用 do，仅有第一次的无返回值的 f 函数可以执行，即便 f 不同：

```go
var once sync.Once

func init() {
  once.Do(func() {
    // 仅执行一次
  })
  // 这次不会执行
  once.Do(func() {
    
  })
}
```

## 讨论线程安全的 map 在多线程中的使用
go 语言中的 map 并不是并发安全的，一个 map 如果不加锁的去处理数据的时候就会出现 panic 的情况，比如：
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for i := 0; i < 100000; i++ {
			m[1] = 1 //设置key
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			fmt.Println(m[2]) //访问这个map
		}
	}()
	time.Sleep(1000000)
}
```
这种写法就会发生 panic，原因是 go 语言不支持并发读写 map，必须加锁

其实我们如果分析这段代码，并没有说对同一个 key 值进行读写，也没有涉及到扩容的问题，但是仍然会 panic，go 在操作 map 时会进行 data race 的检测，只要检测有，就会直接 panic
### 直接加锁
我们可以人为的加锁，这样就可以避免 data race 的行为
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for i := 0; i < 100000; i++ {
			mu.Lock()
			m[1] = 1 //设置key
			mu.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			mu.Lock()
			fmt.Println(m[2]) //访问这个map
			mu.Unlock()
		}
	}()
	time.Sleep(1000000)
}
```
如果数据量比较低的话，这么做毫无问题，如果数据量较大，或者每次操作都比较耗时，读写公用一锁就比较浪费了。

那么可以使用读写锁吗？当然可以啦，我们使用读写锁可以更优秀的去解决这个问题


```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for i := 0; i < 100000; i++ {
			mu.Lock()
			m[1] = 1 //设置key
			mu.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			mu.RLock()
			fmt.Println(m[2]) //访问这个map
			mu.RUnlock()
		}
	}()
	time.Sleep(1000000)
}
```
RWMutex 在同时有读写需求时，会优先获取写锁，读锁需要等待

如果当前有读锁，则后续的写锁请求会被阻塞，但读锁可以继续获取，

如果当前有写锁，则后续的读锁和写锁请求都会被阻塞。

所以，如果读多写少，使用读写锁是非常方便的，假如读和写都异常的高，那么读和写其实是不能同时进行的，如果读贼多，写就可能被阻塞等待了。

### 细颗粒度并发安全 Map
我们都知道，锁对于性能的影响是特别大的，尤其是线程非常多的时候，那么多线程公用一个锁，各种等待，能不影响效率吗，那么怎么做能提高效率呢？

降低锁的颗粒度就可以提高效率，换言之就是本来 1000 个线程公用一个，现在，我们把数据分为 10 份，100 个线程用一个锁，性能就能大范围的提高

我们可以使用 https://github.com/orcaman/concurrent-map 这个分片儿锁去替代互斥锁

分片儿锁的基本原理就是将一个大的 map 的内容，变成 10 个或者是更多个 map 的内容，我们可以这么做：

本身需要一个 map 的数据结构，我们改成一个 slice，slice 含有 10 个 map 的数据结构
我们还需要一个定位分片的算法，基本上都是使用一个哈希算法先定位分片，然后后续就跟一般的互斥锁一致了。
用法如下：

```go
// Create a new map.
	m := cmap.New[string]()

	// Sets item within map, sets "bar" under key "foo"
	m.Set("foo", "bar")

	// Retrieve item from map.
	bar, ok := m.Get("foo")

	// Removes item under key "foo"
	m.Remove("foo")
```
### sync.Map
这是 go 官方提供的一个线程安全的 map，先说使用场景，只写一次，大量读的场景。

sync.Map 跟分片锁不同，分片锁是直接降低颗粒度，sync.Map 它的基本原理是读写分离，用空间换时间。通过一个只读的数据结构来提高读取速度。

当读少写多的时候，它的效率甚至比直接使用互斥锁还低，总之如果不是写少读多的场景，千万不要用，这个包的使用率挺低的。

## sync.Pool
sync.Pool 是一个对象池，如果我们有一些重复使用的，并且需要频繁的申请和释放的临时对象，那么可以用这个对象池来提高性能。不过这个池子里的对象有可能会被垃圾回收，所以非常重要的数据不能使用这种方法

我来描述一种场景，我们有数据需要被 goroutine 去处理，但是谁处理都行，不 care 是哪位，那么我们就可以创建很多的 goroutine，然后放入到 goroutine 池中 (就跟外包一样。。。)

sync.Pool 有两个注意事项，首先，它线程安全，其次，不能复制 sync.Pool，如果你复制一个 sync.Pool，实际上得到的只是一个指针的拷贝，并不会复制本地池子，所以多个拷贝的 sync.Pool 指针指向的是同一个本地池子，达不到复用的目的。应该定义一个全局的 sync.Pool 实例，不同的 goroutine 都使用这个实例，才能达到对象复用和减少内存分配的目的

pool 包拥有三个方法

- New，pool 的 new 方法后面跟的是创建的内容，一旦 pool 里面空了，就会调用这个 new 方法，当然你也可以不指定这个参数，那么创建的就是 nil 了
- Get 从 pool 中获取一个对象
- Set 将对象送回 pool 中
 
下面举一个演示的例子：
```go
var buffer = sync.Pool{
  New: func() any {
    return new(bytes.Buffer)
  },
}

func GetBuffer()*bytes.Buffer {
  return buffer.Get().(*bytes.Buffer)
}
func PutBuffer(buf *bytes.Buffer){
  // Reset会将缓冲区重置为空，但它会保留底层存储以供将来写入使用
  buf.Reset()
  buffers.Put(buf)
}
```
这段代码是会有内存泄露的风险的，原因也很简单，buf.Reset() 并不会删除底层 slice 的容量，它会保存底层数据结构中的容量，所以 gc 的时候这个 buffer 就非常有可能不会被回收，造成内存泄露

改正方法也很简单，底层数据太大了直接丢弃 gc 就 ok 了，小的放到池子里
```go
func PutBuffer(buf *bytes.Buffer){
  // Reset会将缓冲区重置为空，但它会保留底层存储以供将来写入使用
  buf.Reset()
  if buf.Cap() > maxSize{
    return 
  }
  buffers.Put(buf)
}
```
### pool 内存浪费
当我们需求的内存比池子中的内存小很多的时候，就会造成内存浪费，解决方法就是多造几个池子，比如小池子，中池子，大池子，使用这种方案合理的使用内存

```go
var(
  readerPool   sync.Pool
  reader2kPool sync.Pool
  reader4kPool sync.Pool
)
```
推荐一个三方库，特点是能更加高效的发挥系统性能节约内存以及避免内存泄露问题等，他可以动态的去调节池子的 dafault size 和 max size

https://github.com/valyala/bytebufferpool

连接池也是一个很常见的需求，但是通常不会使用 pool，因为 pool 会被 gc，所以并不靠谱，通常我们使用 map 或者 slice 去实现一个需要真的长时间稳定连接的连接池比如：

- http client 池
- tcp 连接池，推荐三方库：https://github.com/fatih/pool
- 数据库连接池
- memcached client 连接池
### worker pool
worker pool 其实就是 goroutine 的池子，虽然 go 语言的 goroutine 非常的轻量化，但是如果几十万上百万的 goroutine 被创建还是会出问题的

因为 goroutine 是一种要求长期存在的资源，所以一般不使用 pool，而是使用 channel 去作为池子，毕竟 channel 自带线程安全

Worker pool 的目的是为了重用 goroutine。但它重用的不是 goroutine 本身，而是通过 goroutine 执行任务的能力。

如果直接在 channel 中传递 goroutine，那么每个 goroutine 只能执行唯一的一个任务，执行完就退出了。这并没有实现重用。

而 worker pool 的设计是：

创建固定数量的 goroutine 作为 workers。
workers 在 loop 中持续从任务 channel 接收任务并执行。
向任务 channel 不断发送不同的任务。
这样，每个 worker(goroutine) 就可以执行多个任务，实现重用。

所以 channel 中的内容不是 goroutine，而是任务信息，用于指导 goroutine 执行什么任务。goroutine 从 channel 收任务，利用自身的执行能力反复执行不同的任务。

这才是 worker pool 设计的核心思想 - 通过固定数量的 goroutine 反复执行不同的任务，以重用 goroutine 实现高效调度。

所以，worker pool 的目的是重用 goroutine 的执行能力，而不是重用 goroutine 本身。这是通过 channel+goroutine 的组合实现的，channel 用于传递任务信息，goroutine 负责执行任务

让我们简单的实现一个方案：
```go
// 任务类型
type Task struct {
  f func() // 任务函数
}

// 执行任务
func (t *Task) Execute() {
  t.f() 
}

// 工作goroutine
type Worker struct {
  TaskChan chan *Task
}

func (w *Worker) Start() {
  for {
    task := <-w.TaskChan
    task.Execute() 
  }
}

// goroutine池
type Pool struct {
  TaskChan chan *Task
  Workers []*Worker
} 

// 创建goroutine池
func NewPool(numWorkers int) *Pool {
  
  taskChan := make(chan *Task)

  // 初始化workers
  var workers []*Worker
  for i := 0; i < numWorkers; i++ {
    worker := &Worker{
      TaskChan: taskChan,
    }
    workers = append(workers, worker)
    go worker.Start()
  }

  return &Pool{
    TaskChan: taskChan,
    Workers: workers,
  }
}

// 分发任务
func (p *Pool) Schedule(task *Task) {
  p.TaskChan <- task
}
```
使用方式
```go
func main(){
// 定义任务
task1 := &Task{ 
  f: func() {
    // do something
  }}

pool := NewPool(10) 

// 分发任务
pool.Schedule(task1)
}
```

提供一些个优秀的 worker pool 方案

- gammazero/workerpool
- ivpusic/grpool
- dpaks/goworkers
- https://github.com/alitto/pond


## semaphore
信号量 (英语：semaphore) 又称为信号标，是一个同步对象，用于保持在 0 至指定最大值之间的一个计数值。

在系统中，给予每一个进程一个信号量，代表每个进程目前的状态，未得到控制权的进程会在特定地方被强迫停下来，等待可以继续进行的信号到来

根据信号量的不同可以分为计数信号量和二进制信号量，前者使用一个整数作为信号量，后者使用一个二进制 0 1 作为信号量

信号量拥有两个操作：
- p 操作会减少信号量的数值
- v 操作会增加信号量的数值

其中二进制信号量是特殊的信号量，它就是互斥锁的功能

go 语言在 x/sync 中提供了一个 weighted 的包，它就是提供的信号量的功能

- Acquire p 操作，减少信号量的数值，表示获取了资源
- Release v 操作，增加信号量的数值，表示释放了资源
- TryAcquire 尝试获取资源，如果获取成功，则返回 true，否则返回 false
它类似于 trylock 锁，也就是失败直接返回 false，并不会阻塞

让我们使用信号量来实现一个 worker pool
```go
package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"
)

var (
  // 最大的worker数量
	maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers)) //信号量
	task       = make([]int, maxWorkers*4)                // 任务数，是worker的四倍
)

func main() {
	ctx := context.Background()
	for i := range task {
		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}
		// 启动worker goroutine
		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
			task[i] = i + 1
		}(i)
	}
	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("获取所有的worker失败: %v", err)
	}
	fmt.Println(task)
}
```
### 使用信号量时的注意事项
- 请求了资源，忘记了释放
- 释放了从未请求的资源
- 长时间持有一个资源但是不使用它
- 不持有一个资源，但是直接使用了它
### 使用 channel 去实现一个信号量
使用一个缓存为 n 的 channel 去实现一个信号量

```go
package main

import "sync"

// Semaphore 数据结构，并且还实现了Locker接口
type semaphore struct {
	ch chan struct{}
}

// 创建一个新的信号量
func NewSemaphore(capacity int) sync.Locker {
	if capacity <= 0 {
		capacity = 1 // 容量为1就变成了一个互斥锁
	}
	return &semaphore{ch: make(chan struct{}, capacity)}
}

// 请求一个资源
func (s *semaphore) Lock() {
	s.ch <- struct{}{}
}

// 释放资源
func (s *semaphore) Unlock() {
	<-s.ch
}

```
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