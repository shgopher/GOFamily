# go语言运行时的调度器
进程是操作系统资源共享的基本单位，线程是操作系统调度的最基本单位，go决定使用协程作为基本的执行单元，常见的有多进程，多线程，多进程的上下文切换耗时太长，多线程虽然有所改善，但是仍然比较长，go选协程可以大幅度降低资源在各个执行单元中切换的时间

go语言的调度模型在各个时期是有不同的模型的，那么我们先了解一下它的演变过程。

- 单线程调度
    
    在这个版本里，runtime代码还是用c写的，这个单线程调度就是同时仅仅支持一个g来执行，每次先上锁，g执行，执行完释放锁再去执行下一个g
- 多线程调度

    这个时候程序内就存在多个同时运行的线程，所以引入了多线程调度器，并且引入了 GOMAXPROCS 这个可以控制最大并发的设置。这个阶段容易产生全局锁颗粒度大，可执行线程之间经常交换g造成时间上的浪费等问题
- 任务窃取调度

    正式引入G:M:P 模型，工作窃取的意思就是目前这个p下面没有正在运行的g，那么就从其它的p中窃取


    下面是p这个调度器的结构体，可以看出，它拥有了两个g参数，一个是队列，一个是g变量，其中还包括了一个m，所以它的调度方式就是将队列中的队头拉到m中去执行，每一个m对应一个p，每一个m对应了cpu中的一个线程（比如4核8线程），这个线程可以成为一个内核了，最起码操作系统是这么认为的。
    ```c
    struct P {
	Lock;

	uint32	status;
	P*	link;
	uint32	tick;
	M*	m;
	MCache*	mcache;

	G**	runq;
	int32	runqhead;
	int32	runqtail;
	int32	runqsize;

	G*	gfree;
	int32	gfreecnt;
    };

    ```
    此版本的的**p没有可以停止g执行的权利**，意思是，只有让g主动让出cpu的执行才可以，p没有那个能力去干预。
- 抢占式调度

    在这个版本中，p拥有了抢占能力，**可以主动去干预g的执行问题**。go会在**stop the world**的时候发现运行时间超过`10ms`的g，让它让出cpu执行。

    目前的抢占只发生在**垃圾回收阶段**。
- 非均匀内存访问调度器

    目前是为了增加数据的局部性减少数据的全局性，因为全局就意味着锁的颗粒度不够小，各种资源统一调度会有资源利用率低等问题。
## G:M:P 模型   
### G:M:P的数据结构
|G|M|P|
|:---:|:---:|:---:|
|对应了go内部的最基本的可执行单元|直接对应操作系统的线程|对应了本地调度器| 

G:
```go
type g struct {
    // 表示这个g的栈内存的范围
	stack       stack
    // 这个字段用于调度器的抢占式调度
	stackguard0 uintptr
    // 每一个g中都会保存panic和defer的调用链表
    _panic       *_panic 
	_defer       *_defer 
    // 对应了调用它的m
    m              *m
    // 跟调度器有关
	sched          gobuf
    // g的状态
	atomicstatus   uint32
    // goroutine id，对于开发者不可见
	goid           int64
}
// 这几个字段，决定了g在系统的调度中使用，其中栈指针和计数器准确的记录了程序执行的位置，方便下次在cpu中执行继续某个位置
type gobuf struct {
    // 栈指针
	sp   uintptr
    // 程序计数器
	pc   uintptr
    // 调用了runtime.gobuf的g
	g    guintptr
    // 系统的返回值
	ret  sys.Uintreg
}
```
g的执行状态可以大概分为三种，等待中，可执行未执行，正在执行

M:

m可以创建10000个，但是只有GOMAXPROCS个才能被激发活跃运行，其余的都会被回收。系统默认的GOMAXPROCS等于内核数，比如四核八线程的cpu就会创建8个。
```go
type m struct {
    // g0持有调度栈
	g0   *g
    //正在m中执行的g
	curg *g
}
```
P:

p上拥有一个本地队列的g，g会优先放置到本地的队列中，然后全局还拥有一个g的队列，这个存在全局的g就是当队列中超过256个的时候就会把本地的一半g移到全局q中，当本地没有g的时候就从全局的g中抓取数据，或者m也会从其它m旗下p的队列中抓取一半的g放置到自己旗下p的队列里。

```go
type p struct {
    // 表示反向持有的系统线程
	m           muintptr
    // 表示p上面挂载的队列头部
	runqhead uint32
    // 尾部
	runqtail uint32
    // 队列，注意看队列是用一个256的数组表示的，说明创建的g数量是优先的，不能超过256 * p的个数，如果是4核8线程就是不能超过 256 * 8 = 2048个
	runq     [256]guintptr
	runnext guintptr
	...
}
```
## runtime.gosched()基于阻塞的协程调度

Gosched 允许其他goroutine运行,它不会挂起当前的goroutine,因此会自动恢复执行。

意思就是这个函数，让这个g让出了cpu执行权利，让位给其它的g，然后某个时间点再回来执行这个g。

```go
func main(){
    for i:= 0;i < 10;i++ {
        go func(i int){
                fmt.Println(i)
        }(i)
    }
    fmt.Println("over")
}
```
这种情况下，就会输出over，上面那是个g的i就可能不会执行。

我们使用runtime.Gosched()，那么上面的g一定会执行。（不一定全部执行，但是一定有一部分执行）
```go
func main(){
    for i:= 0;i < 10;i++ {
        go func(i int){
                fmt.Println(i)
        }(i)
    }
    // 这里就可以暂时让主g让出cpu执行片段
    runtime.Gosched()
    fmt.Println("over")
}
```
底层的逻辑就是，将这个g从m中取消，放到本地q的队尾