# go语言并发模型 --- csp 通信顺序进程

csp在go中的实现就是goroutine之间顺序执行，chan在之间交流通信。

在cspgo的实现里，有着非常流行的一句话，不要通过共享内存来通信，而是应该通过通信来共享内存。通过通信就是通过chan。

通常来说，不同的goroutine之间是没有任何的关联的，然后一个g通过chan往另一个g传递了信号，这个时候他们就有了关联

```go
func main(){
    ch := make(chan int)
    // 1 g
    go func(){
        ch <- 1
        close(ch)
    }()
    for v := range ch {
        fmt.Println(v)
    }
}
```
这段代码中，1g就往主g中发送了一个信息1