# 通道 channel
我们使用make来初始化这个chan，参数是容量，拥有容量就是可以拥有buffered，没有容量就是必须收了发的数据，才能传递下一个。
## 通道的基本使用
```go
// 初始化：
ch := make(chan Type,Cap)
// 例如：
ch := make(chan int,10)
// 写入数据
ch <- 1
// 或者可以使用select写入数据
select {
    case ch <- 1:
}
// 读取数据
<- ch // 这种方式即是读取数据，然后舍弃数据
f(<-ch) // 这种方式是读取数据，传入参数
a := <-ch // 读取数据，并且赋值给a

//当然读取也可以使用select

select{
    case <-ch :
}
```
另外chan中存在着几种内置函数，close关闭这个chan，len返回队列中的数据，cap是队列的容量，make后面的是cap不是len。make chan中后面只有容量着一个参数。

chan拥有下面三种形式：

- chan 可发可收的channel
- chan<- 只能发的chan
- <- chan 只能收的chan

注意一点，selec中的case中，发chan和收chan是平等的，都可以存在在case中，以及chan也是可以是chan存储的对象的，比如 `chan<- <-chan` 这个意思就是只发chan中存储了一个只收chan，<- 尽量跟左侧的chan结合，比如说 `chan<- chan` 意思就是只发chan中存储的是收发chan，如果想改变以下意思，变成chan中装只收chan可以加小括号来进行约束`chan (<-chan)` 

清空一个chan是有快捷方式的：

```go
for range ch {

}
```
## 通道的实现原理
chan在go语言中的具体实现来自runtime.hchan,具体的数据结构如下：
```go
type hchan struct {
	qcount   uint           //  循环队列中的元素个数，调用len，返回的就是这个变量。
	dataqsiz uint           // 循环队列的大小，就是cap返回时盗用的变量。
	buf      unsafe.Pointer // 循环队列的指针
	elemsize uint16 // chan中数据的大小
	closed   uint32 // 是否关闭通道
	elemtype *_type // chan中元素的类型
	sendx    uint   // send在buffered中的索引，意思就是发送数据的时候，这个队列中的index所在的位置
	recvx    uint   // receive 在buffered中的索引，同样的，取数据的时候，index所在这个循环队列中的位置。
	recvq    waitq  // recv的等待队列，如果接收者，因为没办法接收了，那么它就会被加入到这个队列中进行等待
	sendq    waitq  // send中的等待队列，如果不能发送了，那么就会加入到这个队列中等待。
	lock mutex // **互斥锁**
}
```

初始化chan的底层实现：

```go
func makechan(t *chantype, size int) *hchan {
	elem := t.elem

	// 对于size进行检查
	if elem.size >= 1<<16 {
		throw("makechan: invalid channel element type")
	}
	if hchanSize%maxAlign != 0 || elem.align > maxAlign {
		throw("makechan: bad alignment")
	}


	mem, overflow := math.MulUintptr(elem.size, uintptr(size))
	if overflow || mem > maxAlloc-hchanSize || size < 0 {
		panic(plainError("makechan: size out of range"))
	}

	var c *hchan
	switch {
	case mem == 0:
		// 这里是指的是make后面没有指定cap的情况。
		c = (*hchan)(mallocgc(hchanSize, nil, true))
		// 这种情况下，就不必创建buffered了。
		c.buf = c.raceaddr()
	case elem.ptrdata == 0:
		// 如果chan中存储的不是指针类型，那么就给hcan数据结构分配空间。
		c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
        // 并且给buf分配空间
		c.buf = add(unsafe.Pointer(c), hchanSize)
	default:

		元素存在指针，单独分配buf
		c = new(hchan)
		c.buf = mallocgc(mem, elem, true)
	}
    // 记录元素的大小类型和容量
	c.elemsize = uint16(elem.size)
	c.elemtype = elem
	c.dataqsiz = uint(size)
	lockInit(&c.lock, lockRankHchan)

	if debugChan {
		print("makechan: chan=", c, "; elemsize=", elem.size, "; dataqsiz=", size, "\n")
	}
	return c
}
```

这段代码对于这三种方式给出了不同的分配内存的方式。
- 有buffered
- 无buffered
    - 值类型
    - 指针类型



发送：

```go
func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
    // 第一部分先做判断，如果chan没有初始化（所以才是nil），
	if c == nil {
		if !block {
			return false
		}
        // 那么就无法继续往下走，gopark的意义就是挂起这个goroutine
		gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}

	if debugChan {
		print("chansend: chan=", c, "\n")
	}

	if raceenabled {
		racereadpc(c.raceaddr(), callerpc, funcPC(chansend))
	}

	// 这部分的意思是，如果队列中满了，但是不想阻塞，那么就会直接返回false
	if !block && c.closed == 0 && full(c) {
		return false
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

// 这部分的意思就是当一个已经关闭的chan，继续往里面send数据，那么就会Panic
	lock(&c.lock) // 添加 互斥锁
// 这里的直接意义就是，close关闭了，然后程序走到这一步了（因为是send函数，所以意思就是我关闭了你还send）
	if c.closed != 0 {
		unlock(&c.lock) // 在这种情况下的解锁。
		panic(plainError("send on closed channel"))
	}
// 这一部分的意思是从接收者的等待队列中，拿出来一个接收者，然后把发送的数据交给它。很明显啊，这个地方结合下面的buf，我们可以得出一个结论，也就是说，不论是否拥有buf，都会是优先去调用接收函数。
	if sg := c.recvq.dequeue(); sg != nil {
		// 可以看到 这个if中也解锁了。
		send(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true
	}
// 这种情况就是说buf还没满，需要把数据放入buf，然后返回即可。
	if c.qcount < c.dataqsiz {
		qp := chanbuf(c, c.sendx)
		if raceenabled {
			racenotify(c, c.sendx, nil)
		}
		typedmemmove(c.elemtype, qp, ep)
		c.sendx++
		if c.sendx == c.dataqsiz {
			c.sendx = 0
		}
		c.qcount++
        // 解锁了
		unlock(&c.lock)
		return true
	}
// 如果buf满了，就返回false
	if !block {
		unlock(&c.lock)
		return false
	}
    
	....
}
```

接收：

```go
//如果没有第二个参数的时候就直接调用chanrev
func chanrecv1(c *hchan, elem unsafe.Pointer) {
	chanrecv(c, elem, true)
}

//如果拥有第二个参数的时候，那么就按照这种方式调用 v,ok := <- chan 的形式
func chanrecv2(c *hchan, elem unsafe.Pointer) (received bool) {
	_, received = chanrecv(c, elem, true)
	return
}
```
接下来我们探究一下 这个chanrecv函数(片段，有删节)：

```go
func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool){
    ...
    // 这个的意思是说，如果这个chan没有初始化，那么接收的时候调用者就会被阻塞。
    if c == nil {
		if !block {
			return
		}
        // 这个函数的意思就是阻塞挂起
		gopark(nil, nil, waitReasonChanReceiveNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}
    ...

// 这部分的意思是说，首先先加锁
    lock(&c.lock)
//如果chan被关闭了，然后队列中没有元素了，那么返回true和false并且解锁
	if c.closed != 0 && c.qcount == 0 {
		if raceenabled {
			raceacquire(c.raceaddr())
		}
		unlock(&c.lock)
		if ep != nil {
			typedmemclr(c.elemtype, ep)
		}
		return true, false
	}
    // 这部分的意思是 sendq中有等待者（意思就是等着收的接收者队列中有等待者）那么如果buf中有数据，直接从buf中读数据，如果没有，那么直接将这个数据给从队列中取出的这个接收者。
if sg := c.sendq.dequeue(); sg != nil {
		
		recv(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true, true
	}
    //没有等待的发送者，buf中还存在数据，取出一个数据给reciver
    if c.qcount > 0 {
		// Receive directly from queue
		qp := chanbuf(c, c.recvx)
		if raceenabled {
			racenotify(c, c.recvx, nil)
		}
		if ep != nil {
			typedmemmove(c.elemtype, ep, qp)
		}
		typedmemclr(c.elemtype, qp)
		c.recvx++
		if c.recvx == c.dataqsiz {
			c.recvx = 0
		}
		c.qcount--
		unlock(&c.lock)
		return true, true
	}

	if !block {
		unlock(&c.lock)
		return false, false
	}

    ...

    // 下面的处理内容就是 buf中没有数据，reciver阻塞，直到它从sender中取出来了数据。

```

close:

```go
func closechan(c *hchan) {
	if c == nil { // chan是nil的时候panic，所以不能close一个没有初始化的chan，结局就是下面这一句
		panic(plainError("close of nil channel"))
	}

	lock(&c.lock)
	if c.closed != 0 { 
        // 如果chan的close已经被标记为0了（意思就是这个chan已经被关闭了）也是panic
		unlock(&c.lock)
		panic(plainError("close of closed channel"))
	}

...

	c.closed = 1

	var glist gList

	// 释放所有的reader
	for {
		sg := c.recvq.dequeue()
		...
		gp := sg.g
		...
		glist.push(gp)
	}

	// 释放所有的writer，他们会panic
	for {
		sg := c.sendq.dequeue()
		...
		gp := sg.g
		...
		glist.push(gp)
	}
	unlock(&c.lock)

	...

```

## 使用chan最容易犯的错误
- panic
- 协程泄漏

### panic
下面这些，上面源码分析中都有提示

- close chan等于nil的chan
- send已经close的chan
- close已经close的chan
### 协程泄漏
举一个简单的例子

```go
func Test(){
    for i:= 0;i < 100;i++ {
        go func(){
               select {

               } 
        }()
    }
    
}
```
这种情况下，这些个协程就永远被阻塞了，也就是永远不会被gc掉，也就是所谓的goroutine泄漏，下面给出一个稍微生产环境一点的例子:

```go
func A(timeout time.Duration)int{
    ch := make(chan int)
    go func(){
        // 模拟耗时处理
        time.Sleep(time.Second + timeout)
        ch <-1
        fmt.Println("test ")
    }()
    select {
        case v := <-ch :
            return v
        case <- time.After(timeout):

            return -1
    }
}
```
这种情况下 ch绝对不会被接收，因为select中走 time.After了，然后上面那个协程，ch就用于啊不会发送1，因为当select运行的时候，收已经准备好，这个时候按道理是可以接收发传来的data的，然后发并没有准备好，所以收的case只能阻塞了，那么当time.After搞定以后，就直接走这个case了，那么结局就是收关闭了，上面的goroutine就会一直阻塞到ch <- 1这个地方，原因就是内存模型：无buf的chan只有收准备好，才能发，收不可能准备好了，所以发的操作就挂起阻塞了。

**解决这个问题也很简单**，就是给这个buf设置1，这个时候发就可以开始了，当然没有收它的reader，这个chan很快就会被gc掉。
```go

func A(timeout time.Duration)int{
    // 设置一个buf
    ch := make(chan int,1)
    go func(){
        // 模拟耗时处理
        time.Sleep(time.Second + timeout)
        // 这里将不再阻塞
        ch <-1
        fmt.Println("test ")
    }()
    select {
        case v := <-ch :
            return v
        case <- time.After(timeout):

            return -1
    }
}
```
## 什么时候用chan，什么时候用并发原语

- 共享资源的并发，使用并发原语

- 复杂的任务编排和消息的传递使用chan

- 消息通知机制使用chan

- 等待所有任务的完成使用并发原语（waitgroup），当然了如果是可循环的那种，直接使用 cyclicbarrier，【也就是说符合并发原语典型场景的用并发原语就对了，比如合并请求用singleleftlight，分组的使用errgroup】

- 需要跟select结合的用chan

- 需要和超时结合的，使用chan和context

## 值得注意的事情
一个通道中如果含有数据，这个时候数据并没有被接收读取，然后它被close了，那么这个值仍然可以被读取，然后再读才是零值，当然了，被close掉的chan是不能再往里面send数据了

||nil|空|满|不空也不满|closeed|
|:---:|:---:|:---:|:---:|:---:|:---:|
|接收者|阻塞|阻塞|正常读取|正常读取|读取队列中未读取的值，然后再读取到零值|
|发送者|阻塞|正常发送|阻塞|正常发送|panic|
|close内置函数|panic|正常关闭|正常关闭，数据可以正常被读取完|正常关闭，数据可以被正常读取完|panic|

# 通道的使用场景
chan通常会被当作一个线程安全的消息队列和buffer来使用。底层实现是用array实现的循环队列。

go中的chan包括含有buffer的队列和没有buffer的队列，如果 chan 中还有数据，那么，从这个 chan 接收数据的时候就不会阻塞，如果 chan 还未满（“满”指达到其容量），给它发送数据也不会阻塞，否则就会阻塞。unbuffered chan 只有读写都准备好之后才不会阻塞，这也是很多使用 unbuffered chan 时的常见 Bug。
## 使用反射来处理多chan或者是编译期间才能确定的chan数量

我们要使用反射的能力来处理这个需求：`reflect.Select`

```go
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
```
这个函数的意义就是传入一个selectcase类型的slice，然后随机的去执行case，chosen就是被选择的参数在slice中的index，recv返回的是接收值，recvok是返回的是否可以拥有可以执行的参数，如果有就是true，否则就是false

下面我们实现一个可以动态收发的一个操作
```go

package main

import (
	"fmt"
	"reflect"

)

func main(){
	a,b := make(chan int,10),make(chan int,10)
	s := creatSelectCase(a,b)
	for i:= 0;i<10;i++ {
		c,v,ok:= reflect.Select(s)
		// 收
		if v.IsValid() {
			fmt.Println(s[c].Dir,v,ok) 
		//发
		}else {
			fmt.Println(s[c].Dir,ok)
		}
	}

}
// 创建reflect.SelectCase
func creatSelectCase(ch ...chan int)[]reflect.SelectCase{
	var SelectCase []reflect.SelectCase
	// send
	for i,v := range ch {
		SelectCase = append(SelectCase,reflect.SelectCase{
			Dir: reflect.SelectSend,
			Chan: reflect.ValueOf(v),
			Send: reflect.ValueOf(i),

		})
	}
	// recive

	for _,v := range ch {
		SelectCase = append(SelectCase,reflect.SelectCase{
			Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(v),
		})
	}
	return SelectCase
}



```


## chan拥有以下五种使用场景：

- 数据交流
- 数据传递
- 信号通知
- 任务编排
- 锁

接下来我们一一分析

## 数据交流，channel基于生产者消费者模型的无锁队列
数据交流只得是多个chan进行交流，数据传递是一个传递给另一个。其中生产者消费者模型就是典型的数据交流。我们可以实现单一的生产者单一的消费者模型，也可以实现多生产者和多消费者的模型。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

var (
	buf = make(chan int, 1000)
	msg1 = make(chan struct{})
)

func main() {
	ctx,cal := context.WithTimeout(context.Background(),time.Second<<2)
	defer cal()
	go Writer(1,ctx)
	go Writer(2,ctx)
	go Reader(ctx)
	go Reader(ctx)
	<- msg1
}

// 生产者
func Writer(i int,ctx context.Context) {
	L:
	for {
		select {
		case buf <- i:
			fmt.Println("生产者发")
			time.Sleep(time.Second >>2)
		case <- ctx.Done():
			fmt.Println("生产者要退出了")
			msg1 <- struct{}{}
			break L
		}
	}

}

// 消费者
func Reader(ctx context.Context) {
	L:
	for {
		select {
		case v := <- buf:
			fmt.Println("消费者要收")
			fmt.Println(v)
		case <- ctx.Done():
			fmt.Println("时间到")
				break L
		}
	}

}

```

## 数据传递，简单的理解就是从一个g传递到另一个g，击鼓传花
击鼓传花是一个很典型的数据传递的例子，在各个g中，使用通道来传递信息，就是这个场景的解释。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 对 通道进行初始化
	cS := make([]chan struct{}, 4)
	for i := 0; i < 4; i++ {
		cS[i] = make(chan struct{})
	}
	// 初始化g
	for i := 0; i < 4; i++ {
		i := i
		go Gu(cS[i], cS[(i+1)%4], i+1)
	}
	cS[0] <- struct{}{}
	select {}
}

//Gu 信息传递的数据处理中心，这个场景就是经典的信息传递，从一个g传递到另一个g
func Gu(c1, c2 chan struct{}, n int) {
	for {
		select {
		case <-c1:
			fmt.Println(n)
			time.Sleep(time.Second >>1)
			c2 <- struct{}{}
		}
	}

}

```

## 信号通知，使用chan传递信号，最典型的就是传递退出信号，跟context结合的时候
```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 这里模拟一个执行的业务
	go func() {
		time.Sleep(time.Second*10)
	}()

	term := make(chan os.Signal,1)
	// 这个函数是监听的系统的输入信号，其中底层代码中有一个新开辟的g，用来监听
	signal.Notify(term,syscall.SIGINT,syscall.SIGTERM)
	// 这个地方用来阻塞使用，当监听到了系统给予的信号以后，这里开始接收数据并抛弃，然后继续向下执行
	<- term
	fmt.Println("\n退出")
}

```
使用上面这种方式进行优雅的退出。

下面再看一个例子，这个例子更透明更简单：

```go
func main() {
	c1 := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		// 这里就是执行完毕本职任务的时候，给等待的g发送停止的信号
		c1 <- true
	}()
	// 这个chan阻塞等被唤醒
	<- c1
	// 这里可以设置清理工作
	 doClean()
	fmt.Println("执行正式完毕，可以退出了")
}

```
真实场景下，我们做的比如有退出的清理工作，比如说清除缓存这类的，那么清理的过程中，如果时间过长，用户就会受不了，这个时候我们可以设置超时时间。

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	// 正常的业务代码
	closing := make(chan struct{})
	closed := make(chan struct{})
	go func() {
		for {
			select {
				case <-closing:
					return
			default:
				// 正常的业务计算模拟
				time.Sleep(time.Second << 3)
			}
		}
	}()
	termChan := make(chan os.Signal,1)
	signal.Notify(termChan,syscall.SIGINT,syscall.SIGTERM)
	<- termChan
	close(closing)
	// 这里就是结束以后的处理过程了
	go func() {
		time.Sleep(time.Second)
		close(closed)
	}()
	select {
	case <- closed:
		fmt.Println("被手动关闭")
	case <- time.After(time.Second):
		fmt.Println("超时自动跳出")
	}
	fmt.Println("正式退出")
}
```
## 实现锁，锁不只有互斥锁，还有例如排外锁等

### 使用chan来实现互斥锁
```go
/Mutex 使用chan实现的互斥锁
type Mutex struct {
	mu chan struct{}
}

// NewMutex 返回一个带有锁的互斥锁
// 其中chan中含有一个数据，视为含有一个未使用的锁
func NewMutex() *Mutex {
	m := &Mutex{make(chan struct{}, 1)}
	m.mu <- struct{}{}
	return m
}

//Lock 得到锁
func (m *Mutex) Lock() {
	<-m.mu
}

// Unlock 释放锁
func (m *Mutex) Unlock() {
	select {
	case m.mu <- struct{}{}:
	default:
		panic("不应该给已经解锁的锁，再次解锁")
	}
}

// IsLock 是否某个元素已经上锁
func (m *Mutex) IsLock() bool {
	return len(m.mu) == 0
}

```

### 使用chan来实现排外锁trylock

```go
// TryLock 排外锁,true表示可以上锁，并且给予锁，false表示已经上锁，死心吧。
func (m *Mutex) Trylock() bool {
	select {
	case <-m.mu:
		return true
	default:
		return false
	}
}
```
### 使用chan实现超时功能

```go
// TimeOutLock 在设置时间后，如果还没得到锁，将不再获取锁，直接退出
func (m *Mutex) TimeOutLock(d time.Duration)bool {
	t := time.NewTimer(d)
	select {
		case <- m.mu:
			t.Stop()
			return true
		case <-t.C:
			return false
	}
}

```
## 任务编排，如果按照某种逻辑，使用chan来使得这些g按照我们设想的逻辑顺序去执行

任务编排通常来说有下面这几种模式
- 流水线模式
- Or-Done
- 扇入
- 扇出
- Stream
- Map-Reduce
### chan实现sync.WaitGroup功能

```go

package main

import "fmt"

func main() {
	g := NewWaitGroup(10)
	g.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer g.Done()
			fmt.Println(i)
		}(i)
	}
	g.Wait()
}

//WaitGroup 使用chan实现waitgroup
type WaitGroup struct {
	c chan struct{}
}

func NewWaitGroup(n int) *WaitGroup {
	return &WaitGroup{
		make(chan struct{}, n),
	}
}
func (w *WaitGroup) Add(n int) {
	if n > cap(w.c) {
		panic("add中的数据不可以大于new中的数据")
	}
	for i := 0; i < n; i++ {
		w.c <- struct{}{}
	}
}
func (w *WaitGroup) Done() {
	<-w.c
}
func (w *WaitGroup) Wait() {
	for len(w.c) > 0 {
	}
}

```
### 流水线模式
上文中的数据传递，击鼓传花模式就是流水线模式的一种表现形式。流水线模式，顾名思义，就是跟工厂中的流水线类似，一个工位搞定以后再往下一个工作走，属于单线程模式。上文中的击鼓传花数据传递，那个数据在不同的g之间流转一个接着一个，单线程工作，是典型的流水线模式。
### Or-Done
这是上文中信号通知的延伸，上文中信号是当某一个任务完成的时候，我们给予了信号通知的方式，这里我们是多任务，当多任务中的任意一个完成的时候，就发起通知。

现在有一个场景，我们需要发送一个信息到分布式的多服务器节点，我们要求，只要有一个节点服务器做出了应答，我们就可以返回这个信号

我们先看一下使用传统的递归分治的算法：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	<- Or(
		w(1*time.Second),
		w(2*time.Second),
		w(3*time.Second),
		w(4*time.Second),
		w(5*time.Second),
		w(6*time.Second),
		w(7*time.Second),
		w(8*time.Second),
		w(9*time.Second),
		w(10*time.Second),
		w(11*time.Second),
		w(12*time.Second),
		)
	fmt.Println("测试结束",time.Since(n))
}

//Or 这是Or-done模式,本方法采用分治的方法
func Or(c ...chan interface{})chan interface{}{
	// 特殊情况，有零个或者一个的情况
	// 必须拥有特殊值的处理，不然无法跳出递归
	switch len(c) {
	case 0:
		return nil
	case 1:
		return c[0]
	}
	t := make(chan interface{})
	go func() {
		defer close(t)
		// 只有两个
		switch len(c) {
		// 必须拥有特殊值的处理，不然无法跳出递归
		case 2:
			select {
				case <- c[0]:
				case  <- c[1]:
			}
		default:
			mi := len(c)/2
			select {
			// 这里开始分治
				case <- Or(c[:mi]...):
				case <- Or(c[mi:]...):
			}
		}
	}()
	// 返回分治的结果
	return t
}
//w 测试的工作文件
func w(d time.Duration) chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(d)
	}()
	return c
}
```

当节点过多的时候，我们就要很深的递归，这个时候很容易发生内存泄漏，所以我们在海量的g的时候可以使用反射这种方式来实现or-done这种模式。

```go
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	t1 := time.Now()

	<-Or(
		w(1*time.Second),
		w(2*time.Second),
		w(3*time.Second),
		w(4*time.Second),
		w(5*time.Second),
		w(6*time.Second),
	)
	fmt.Println(time.Since(t1))

}

//Or 实现的功能：输入众多chan interface{}的时候，我们只需要一个执行完毕就可以让这个or返回chan interface{}值
// 由于本场景goroutine众多，所以此处采用reflect包，反射的方式来写。
func Or(c ...chan interface{}) chan interface{} {
	switch len(c) {
	case 0:
		return nil
	case 1:
		return c[0]
	}
	// 下面就是大于等于2的时候，使用反射做出来的操作：
	done := make(chan interface{})
	go func() {
		defer close(done)
		var f []reflect.SelectCase
		for _, v := range c {
			f = append(f, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(v),
			})
		}
		reflect.Select(f)
	}()
	return done
}

//w 此处模拟了工作状态
func w(d time.Duration) chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		time.Sleep(d)
	}()
	return done
}

```
### 扇入
在chan中的定义就是多个输入端，一个输出端口

使用反射的方法实现扇入
```go
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	fi := FanInReflect(
		work(1*time.Second),
		work(2*time.Second),
		work(3*time.Second),
		work(4*time.Second),
		work(5*time.Second),
		work(6*time.Second),
		work(7*time.Second),
		work(8*time.Second),
		work(9*time.Second),
		work(10*time.Second),
		work(11*time.Second),
		work(12*time.Second),
	)
	for v := range fi {
		fmt.Println(v)
	}

}

//FanInReflect
func FanInReflect(channels ...chan interface{}) chan interface{} {
	out := make(chan interface{})
	go func() {
		// 结尾关闭out
		defer close(out)
		// 构造一个收的反射select
		var rcs []reflect.SelectCase
		for _, v := range channels {
			rcs = append(rcs, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(v),
			},
			)
		}
		// 这是将已经关闭的chan给筛出去。
		for len(rcs) > 0 {
			i, r, ok := reflect.Select(rcs)
			if !ok {
				fmt.Println("关闭一个chan，剩余",len(rcs))
				// 证明这个chan是关闭的，这个c的内容就不用发送出去了，所以是continue
				rcs = append(rcs[:i], rcs[i+1:]...)
				continue
			}
			out <- r
		}

	}()
	return out
}

// work 这里就是作为测试使用
func work(d time.Duration) chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		done <- d.String()
		time.Sleep(d)
	}()
	return done
}

```
使用分治的方法实现扇入
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	fi := FanInRec(
		work(1*time.Second),
		work(2*time.Second),
		work(3*time.Second),
		work(4*time.Second),
		work(5*time.Second),
		work(6*time.Second),
		work(7*time.Second),
		work(8*time.Second),
		work(9*time.Second),
		work(10*time.Second),
		work(11*time.Second),
		work(12*time.Second),
	)
	for v := range fi {
		fmt.Println(v)
	}

}

//FanInRec 使用归并的思想进行处理，可以类比算法中的多路归并。
func FanInRec(channels ...chan interface{}) chan interface{} {
	switch len(channels) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return channels[0]
	case 2:
		return mergeTwoChannel(channels[0],channels[1])
	default:
		m := len(channels)/2
		return mergeTwoChannel(
			FanInRec(channels[:m]...),
			FanInRec(channels[m:]...),
		)
	}
}

// mergeTwoChannel
func mergeTwoChannel(a, b chan interface{}) chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		for a!= nil || b != nil {
			select {
			case v,ok :=  <- a:
				if !ok {
					a = nil
					continue
				}
				done <- v
			case v,ok :=  <- b:
				if !ok {
					b = nil
					continue
				}
				done <- v
			}
		}
	}()
	return done
}

// work 这里就是作为测试使用
func work(d time.Duration) chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		done <- d.String()
		time.Sleep(d)
	}()
	return done
}

```
### 扇出
在chan中的定义就是一个输入端口，多个输出端口

我们经常使用这种模式来完成设计模式中的观察者模式，即使用一个目标去同时通知多个目标，这一个通知者就是所谓的观察者。

我们可以设想这么一个场景，有一群的节点，然后他们拥有一个配置中心，需要让配置中心充当观察者，只要发生了改变，观察者去通知各个节点，下面我们将使用
两种方法去实现这么一种扇出 (fanOut)的模式
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan interface{})
	go func() {
		defer close(c1)
		for i := 0; i < 12; i++ {
			c1 <- i
		}
	}()
	cs := make([]chan interface{}, 6)
	for k := range cs {
		cs[k] = make(chan interface{})
	}
	FanOut(c1, cs, true)
	wg := new(sync.WaitGroup)
	for k := range cs {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for {
				select {
				case v := <- cs[k]:
					fmt.Println(v)
				case <- time.After(time.Second>>2):
					return
				}
			}
		}(k)
	}
	wg.Wait()
}

//FanOut 扇出模式
// victor 为观察者，也就是它一个人的数据要传递给out这个slice中的各个文件，async表示是否异步传递
func FanOut(victor chan interface{}, out []chan interface{}, async bool) {
	go func() {
		//defer func() {
		//	for _ ,v := range out {
		//		close(v)
		//	}
		//}()
		for v := range victor {
			for i:= 0;i < len(out);i++ {
				i:= i
				if async {
					go func(v interface{}) {
						out[i] <- v
					}(v)
				} else {
					out[i] <- v
				}
			}
		}
	}()
}

```
### Stream
流式管道用法，提供跳过几个元素，或者只取某几个元素这种功能,将数据从数组转化为流, 注意这个从切片改为流跟上文中的扇入颇为相似，
只不过这里是切片中的数据改为流，上文中是切片中的各个chan改为一个chan,这是本质的区别。

```go
//TobeStream 将数据从数组转化为流，注意这个从切片改为流跟上文中的扇入颇为相似，
//只不过这里是切片中的数据改为流，上文中是切片中的各个chan改为一个chan
// 这是本质的区别。
func TobeStream(ctx context.Context, values ...interface{}) chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		for i := 0; i < len(values); i++ {
			select {
			case done <- values[i]:
			case <-ctx.Done():
			}
		}
	}()
	return done
}
```

下面有几种可以实现的方法：
- takeN 只取前n个
	```go

	//TakeN 只取前n个
	func TakeN(ctx context.Context, stream chan interface{}, n int) chan interface{} {

			if n > len(stream) {

			panic("n 不应该大于len(stream)")
			
			}

		done := make(chan interface{})
		go func() {
			defer close(done)
			for i := 0; i < n; i++ {
				select {
				case done <- <-stream:
				case <-ctx.Done():
					return
				}
			}
		}()
		return done
	}

	```
- takeFn 筛选只满足fn的数据
	```go
	//TakeFn 筛选只满足fn的数据
	func TakeFn(ctx context.Context, stream chan interface{}, fn func() bool) chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			for len(stream) > 0 {
				if fn() {
					select {
					case done <- <-stream:
					case <-ctx.Done():
						return
					}
				}else {
					<- stream
				}

			}
		}()
		return done
	}

	```
- takeFnWhile 从满足条件开始取，一旦遇到不满足的直接走人
	```go
	//TakeFnWhile 从满足条件开始取，一旦遇到不满足的直接走人
	func TakeFnWhile(ctx context.Context, stream chan interface{}, fn func() bool) chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			key := 0
			for len(stream) > 0 {
				if fn() {
					key = 1
					select {
					case done <- <-stream:
					case <-ctx.Done():
						return
					}
				} else if !fn() && key == 1 {
					break
				}else {
					<- stream
				}
			}
		}()
		return done
	}

	```
- skipN 跳过前n个
	```go
	//SkipN 跳过前n个
	func SkipN(ctx context.Context, stream chan interface{}, n int) chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			i := 0
			for len(stream) > 0 {
				i++
				if i > n {
					select {
					case done <- <-stream:
					case <-ctx.Done():
						return
					}
				}else {
					<- stream
				}
			}
		}()
		return done
	}

	```
- skipFn 跳过满足条件的
	```go
	//SkipFn 跳过满足条件的
	func SkipFn(ctx context.Context, stream chan interface{}, fn func() bool) chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			for len(stream) > 0 {
				if !fn() {
					select {
					case done <- <-stream:
					case <-ctx.Done():
						return
					}
				}else {
					<- stream
				}
			}
		}()
		return done
	}

	```
- skipFnWhile 跳过满足条件的，直到不满足，然后就不跳走了。
	```go
	//SkipFnWhile 跳过满足条件的，直到不满足，然后就不跳走了
	func SkipFnWhile(ctx context.Context, stream chan interface{}, fn func() bool) chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			key := 0
			for len(stream) > 0 {
				if !fn() || fn() && key == 1 {
					key = 1
					select {
					case done <- <-stream:
					case <-ctx.Done():
						return
					}
				} else if fn() && key == 0{
					continue
				}
			}
		}()
		return done
	}

	```


### Map-Reduce
map-reduce可以简单的看作两个步骤，第一，将数据进行映射，也就是分类，第二将数据进行处理。这个意思就是将那些分类好的数据，按照某个次序，进行排列后放入一个统一的结果中。

map:

```go
//MapChan 将流中的数据映射到不同的处理单元
func MapChan(in chan interface{},fn func(v interface{})interface{})chan interface{}{
	done := make(chan interface{})
	// 如果in是nil的化，直接向外传递一个已经close掉的信号即可。
	if in == nil {
		close(done)
		return done
	}
	go func() {
		defer close(done)
		for v := range in {
			// 这里就是fn是map函数，使用这个函数将v传入，看看返回值。
			done <- fn(v)
		}
	}()
return done
}

```

reduce:

```go
//ReduceData
func ReduceData(in chan interface{}, fn func(v1, v2 interface{}) interface{}) interface{} {
	if in == nil {
		return nil
	}
	v := <-in
	for vi := range in {
		v = fn(v, vi)
	}
	return v
}

```

下面我们来演示一下场景：

```go
func main() {
mapfn1 := func(v interface{}) interface{}{
		if v.(int) %2 == 0{
			return v
		}
		return nil
	}
	reduceFn := func(v1, v2 interface{}) interface{}{
		return v1.(int) *2 + v2.(int) *3
	}

	data := []interface{} {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	ctx,cal := context.WithCancel(context.Background())
	defer cal()
	in := streamData(ctx,data...)
	m1 := MapChan(in,mapfn1)
	fmt.Println(ReduceData(m1,reduceFn))

}
// 将数组合并为流
func streamData(ctx context.Context,data ...interface{})chan  interface{}{
	done := make(chan interface{})
	go func() {
		defer close(done)
		for _,v := range data{
			select {
			case <- ctx.Done():
			case done <- v :
			}
		}
	}()
	return done
}
```
> map-reduce的进阶版，在[go编程模式](../3/codeSpecification.md#Map-Reduce)这一个章节