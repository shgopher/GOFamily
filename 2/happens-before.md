# go语言的内存模型 --- happens-before
此处讨论的内存模型跟内存分配和内存gc一点关系都没有，是属于并发的概念。它描述的是在并发环境中，读取相同变量时，变量的可见性，比如某个g在什么时候可以读取到另一个g修改过的变量。

## 指令重排和可见行的问题

go编译器为了实现优化，有的时候你写的代码的顺序不一定按照你写的顺序执行，

```go
package main

var a, b int

func main() {
	go b1()
	a1()

}

func a1() {
	print(a)
	print(b)
}
func b1() {
	a = 1
	b = 2
}

```

这里 ，print的b的值不一定是2，有可能是0，因为a1（）不一定能看得到b1中的a b的先后关系。因为跨goroutine了。

go内存模型中有个非常重要的概念 happens-before，如果严格按照这个概念去执行，那么程序代码在并发的多goroutine中对于变量的读取，便不会出现上述的问题。

## happens- before

happens-before的定义就是xx一定发生在xx之前，并且其它干扰项不会同一时间出现进行干扰，也就是说happens-before发生的时候就是单指这两者，不能有第三方去干扰。

在一个goroutine内部，程序的执行和代码的顺序严格保持一致，即使指令重排了也不会改变这个原则。这不过另一个goroutine看不到这种规则，因为这个规则只限于单个的goroutine内部。那么对于外部的goroutine我们要记住happens-before就可以避免这些问题。

go内存模型通过 happens-before和happens-after定义了三个状态，即：发生在我之前，发生在我之后，还有就是同时发生。

go的导包过程中的happens-before的执行过程是这样的，比如main包里面有init和main以及var和const，main导入的包是a1，a1同样拥有 const var init a1又有a2导入，a2同样有这么一套东西，那么执行的顺序就是 a2的cosnt var init（如果不止一个，一个包内可以有很多个，但是一个文件只能有一个，按照文件名称顺序执行init）然后再来就是a1的const var init 再来是 main的cosnt var init然后最后是main函数。

启动goroutine的go语句这个执行的过程，一定是happens-before 这个goroutine内部的执行的。不过值得注意的是，goroutine退出的时候并没有任何的happens-before保证。

channel中的happens-before：

- 往这个chan中发送数据一定happens-before 从这个chan中取出这个数据，即第n个发一定比第n个收早。
- 关闭一个chan一定happens-before从这个关闭的通道住读取它的一个零值。
- 对于一个没有buffered的chan，从这个chan中读取数据这个操作一定happens-before往这个chan中发送数据这个操作。
- 对于一个容量为m（>0）的的chan，第n个recive，一定是happens-before n+m个send。意思就是这些缓存可以在没有收的情况下即使是充满了也不会去看send有没有准备好。所以收的这个happens-bere就是跟上面的一样，before  n个发，只不过这里有缓存所以就是n+m个发。

mutex和rwmutex中的happens-before：
- 第n次的unlocok一定happens-before 第n+1次的lock
- 对于读写锁，只有释放了写锁，才能去调用读锁
- 对于读写锁，lock必须等待读锁释放了才能获取
waitgroup中的happens-before：
- wait方法只有等到计数值归零了才会返回
once中的happens-before：
- sync.Once(f)，f函数一定是在once的返回值返回之前执行，意思是说，这个once执行的时候，f函数一定是最早执行的那个，但是不一定能执行完啊，只是说它一定最早执行。

atomic中的happens-before：

目前 atomic 无法保证。
