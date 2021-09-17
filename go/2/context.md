# context包的使用
我们先看一下context最核心的 context interface的定义：

```go
type Context interface {
    // 返回done的时间。
    Deadline() (deadline time.Time, ok bool)

    // done返回一个chan，也就是说，当取消运行的时候，这个done就会返回这么个chan。其实就是个信号
    Done() <-chan struct{}

    // 如果done没有正常的关闭，那么就会返回错误。它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
    Err() error

    // value会返回一个key - value
    Value(key interface{}) interface{}
}
```

在context包中有以下几个函数返回的type实现了context接口：

- context.TODO()
- context.Background()
- context.WithDeadline()
- context.WithValue()
- context.WithCancle()
- context.WithTimeout()

这些函数返回的类型都实现了context接口。当然你也可以自己实现以下。毕竟这些都是可外漏的方法。

不过，像todo和background返回的都是没有任何意义的context，也可以把他们叫做root 上下文，也就是上下文的最顶端。

context，中文名就是上下文，它的主要目的就是作为树枝把goroutine串成一棵树，然后统一调度，也就是说，让这些个goroutine统一取消就一起取消，所以它才被翻译为上下文，因为真的是控制上下文的作用。

下面这个例子就是一个context使用的案例

```go
func main() {
	ctx,cancel := context.WithCancel(context.TODO())
	ctx = context.WithValue(ctx,"0","0")
	ctx = context.WithValue(ctx,"1","1")
	go get1(ctx)
	go get2(ctx)
	cancel()
	time.Sleep(time.Second)
}

func get1(ctx context.Context) {
L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("该退出了get1",ctx.Value("0"))
			break L
		}
	}
	go get2(ctx)

}
func get2(ctx context.Context) {
L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出 get2",ctx.Value("1"))
			break L
		}
	}

}
// output : 
//退出 get2 1
//该退出了get1 0
//退出 get2 1

```
使用context比并发原语或者扩展语句更加的简单。主要是操作的时候减少了deadlock的机会，不过也会引入狗皮膏药一样的ctx，不过总体来看context包确实能解决好了控制goroutine树这个问题。另外不要使用withvalue来轻易传递值，因为它寻址的时候是递归的往父树上去寻找，很慢的。