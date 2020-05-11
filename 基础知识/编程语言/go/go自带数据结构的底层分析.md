# go自带数据结构的底层分析
- chan
- slice
- map
- struct
- iota
- string
- interface

## chan

底层
```go
type hchan struct {
    qcount   uint           // 当前队列中剩余元素个数
    dataqsiz uint           // 环形队列长度，即可以存放的元素个数
    buf      unsafe.Pointer // 环形队列指针
    elemsize uint16         // 每个元素的大小
    closed   uint32            // 标识关闭状态
    elemtype *_type         // 元素类型
    sendx    uint           // 队列下标，指示元素写入时存放到队列中的位置
    recvx    uint           // 队列下标，指示元素从队列的该位置读出
    recvq    waitq          // 等待读消息的goroutine队列
    sendq    waitq          // 等待写消息的goroutine队列
    lock mutex              // 互斥锁，chan不允许并发读写
}
```
其中储存的单位是一个指针指向的array，这个array的功能就是一个环形的队列 也就是说是index =  index % length(array) 得出来的结果。
而且最下面的lock是一个互斥锁，证明了一个chan每次只能在一个goruntine中使用。所以使用chan通信当然不用加锁了，因为go底层已经加上了。

其中lock mutext 中的mutex

互斥锁。 在无竞争的情况下，
和自旋锁一样快（只有一些用户级别的说明），
但是在争用路径上，它们在内核中休眠。
零位的互斥锁被解锁（无需初始化每个锁）。
初始化对于静态锁排名很有帮助，但不是必需的。
```go
type mutex struct {
	// Empty struct if lock ranking is disabled, otherwise includes the lock rank
	lockRankStruct
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	key uintptr
}
```

此锁对比sync.Mutex的锁

```go
// A Mutex is a mutual exclusion lock.
// The zero value for a Mutex is an unlocked mutex.
//
// A Mutex must not be copied after first use.
type Mutex struct {
	state int32
	sema  uint32
}
```
很明显两者是不同的，前者的颗粒度更少，占用的资源更少，他们都是使用数字进行记录状态，然后通过atomic的原子操作进行加锁的。
这个原子操作其实就是 相加或者是比较这种东西，祝不过是原子性的，原子性的东西是一瞬间只能一个执行。

## slice
```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```
- 每个切片都指向一个底层数组
- 每个切片都保存了当前切片的长度、底层数组可用容量
- 使用len()计算切片长度时间复杂度为O(1)，不需要遍历切片
- 使用cap()计算切片容量时间复杂度为O(1)，不需要遍历切片
- 通过函数传递切片时，不会拷贝整个切片，因为切片本身只是个结构体而已
- 使用append()向切片追加元素时有可能触发扩容，扩容后将会生成新的切片

### map
```go
type hmap struct {
    count     int // 当前保存的元素个数

    B         uint8  // 指示bucket数组的大小

    buckets    unsafe.Pointer // bucket数组指针，数组的大小为2^B

}

```
### iota

它的数据是根据索引的下标来决定的。简而言之就是根据的是index
所以这块的const中不管你写了几个iota 它表示的都是索引值，所以值是不会变化的。

```go
const(
  a = iota,b = iota-1// 0,-1
  c ,d // 1,1
)
```

### string
```go
type stringStruct struct {
    str unsafe.Pointer
    len int
}
```
需要注意的是 []byte 和string的转化是需要内存的复制的，所以是会消耗额外的内存空间的。（编译器也会优化，比如某些条件下并不会需要内存的复制）
