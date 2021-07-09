# net pool  网络轮询器
io模型分为以下几种
- 阻塞式

    当系统调用read或者write的时候，系统从用户态切换为内核态，内核会检查文件描述符的状态是否可读，如果可读就读取数据，并将数据复制给用户，并切换为用户态，反之，内核将一直阻塞，直到文件描述符可读。
- 非阻塞式
    
    非阻塞式，系统会立刻回返回文件描述符的返回值，不会再阻塞等待数据，但是返回值可能是error或者是可读转台，如果是error就证明并没有准备好数据，这个时候需要系统不断的轮询，当可读以后，系统就可以读取系统缓冲区的数据，并将数据复制给用户，非阻塞式，可以在系统轮询的间隙去做另外的工作，提高了系统的利用率。
- 信号驱动
- 异步模型
- io多路复用
## 设计原理
不同的操作系统底层的模型是不一样的，都要实现以下方法：

```go
func netpollinit()
func netpollopen(fd uintptr, pd *pollDesc) int32
func netpoll(delta int64) gList
func netpollBreak()
func netpollIsPollDescriptor(fd uintptr) bool
```
### go语言的网络io模型
## 数据结构
## 多路复用
## io.rader/writer unix 文件哲学
## net/http 基于goroutine的http服务器