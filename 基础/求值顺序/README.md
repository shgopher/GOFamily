# 求值顺序
> 本文讲述的是单个goroutine的求值顺序，多goroutine时可以参考一下[这篇文章](../../并发/内存模型/)

在包级别，初始化的依赖关系决定了变量声明中各个初始化表达式的求值顺序，当没有依赖关系的时候，在计算表达式，赋值，返回语句时，函数的调用，方法的返回，chanel操作，都按照从左到右的顺序进行计算：

```go
y[f()], ok = g(h(), i()+x[j()], <-c), k()
```
因为没有依赖关系，所以执行顺序是 f() h() i() j() <-c g() k() 

如果存在依赖关系，那么这是一个例子：

```go
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// functions u and v are independent of all other variables and functions

```
按照从左到右，先a 的f（） 但是 需要c ，所以先找c ，c需要 sqr，sqr需要 u 和v ，所以顺序是这样的，u  sqr v  f v g 。 

函数的调用，和通信按照上面说的顺序进行调用但是仍然存在不可排序的情况，下面看一些无法确定的顺序：
```go
a := 1
f := func() int { a++; return a } 
x := []int{a, f()}           // x 可能是 1 2 或者是 2 2 
m := map[int]int{a: 1, a: 2}  //  可能是 {2:1} 或者是 {2:2}
n := map[int]int{a: f()} // 可能是 {2:3} {3:3}
```




## 参考资料
- https://go.dev/ref/spec#Order_of_evaluation
- https://book.douban.com/subject/35720728/ 132页 - 142页