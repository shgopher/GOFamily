# 计时器

时间对于系统来说尤其重要，比如分布式系统中，相对的时间非常重要，不过目前最佳的分布式系统也是存在误差的，go计时器的目的是为了获取相对的时间。标准库中还提供了定时器、休眠等接口能够我们在 Go 语言程序中更好地处理过期和超时等问题。

go语言的计时器发展经历了三个阶段分别是：
- 全局四叉堆
- 分片四叉堆
- 单独管理四叉堆，网络轮询器触发

## 全局四叉堆

```go
var timers struct {
	lock         mutex
	gp           *g
	created      bool
	sleeping     bool
	rescheduling bool
	sleepUntil   int64
	waitnote     note
    // 这个字段存储的就是四叉堆
	t            []*timer
}
```
其中这个字段中的t，后面跟的[]*timer就是表示的四叉堆，这个期间，全局维护一个timers，使用一个lock进行互斥加锁，当然了，明显颗粒度较大。

在go的运行时期间，有两种情况会唤醒计时器，1四叉堆中的计时器到期了，2是四叉堆中加入了更早触发的计时器
## 分片四叉堆
由于全局的计时器，因为互斥锁的颗粒度实在是太大，所以系统改成了64片的互斥锁，这个64基本上是根据内核核心数设置的，如果p(G:M:P 模型中的p）的个数大于了64，那么就将节点改成桶，往桶里叠加即可。不过这种行为就会造成cpu和线程的上下文切换时间增加。切换造成的时间浪费又成为了新的麻烦。

```go
const timersLen = 64

// 这是一个含有64个timerbucket的全局切片。
var timers [timersLen]struct {
	timersBucket
}

type timersBucket struct {
	lock         mutex
	gp           *g
	created      bool
	sleeping     bool
	rescheduling bool
	sleepUntil   int64
	waitnote     note
	t            []*timer
}
```

这种情况下可以将锁的颗粒度从1改为64
## 单独管理四叉堆

最新的四叉堆，取消了计时桶，所有的四叉堆都存放在处理器p（P:M:G中的p，属于运行时的上下文调度处理器）中，

```go
type p struct {
    // 锁
	timersLock mutex
    // 四叉堆
	timers []*timer
// 存放处理器中的四叉堆的数量
	numTimers     uint32
    // adjust状态的计时器的数量
	adjustTimers  uint32
    // deleted状态的计时器的数量
	deletedTimers uint32
}
```
目前计时器都交由处理器的网络轮询器（netpool）和调度器（P:M:G）触发。

## 计时器的数据结构

内部的计时器数据结构下面所示，这种结构存放在各自处理器的四叉堆中，所以下面的数据结构其实就是四叉堆中存放的元素的数据结构。

```go
type timer struct {
	pp puintptr

	when     int64
	period   int64
	f        func(interface{}, uintptr)
	arg      interface{}
	seq      uintptr
	nextwhen int64
	status   uint32
}
```

暴漏出来的计时器：

```go
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
```

time.Timer 计时器必须通过 time.NewTimer、time.AfterFunc 或者 time.After 函数创建。 当计时器失效时，订阅计时器 Channel 的 Goroutine 会收到计时器失效的时间