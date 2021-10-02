## 策略
结构模式的策略模式，这个设计模式的意义依旧是解耦（可以发现很多设计模式出现的意义就是为了解除耦合）它具体的意义是：对象和对象要做的策略进行解除耦合，举个例子：对象 a ，和它要做的策略【加法，减法，乘法，除法】进行解耦和。

所以策略模式可以分为三步：
1. 设置策略：抽象化策略接口 + 实现接口的具体策略
2. 设置执行策略的对象：一个内置抽象策略接口的struct 
3. 将执行策略的对象和策略们集合：使用set和do两个方法，其中set的目的是将具体执行的策略和对象的策略match在一起，do的目的就是执行方法。

这样做的好处就是，我们可以尽情的添加新的策略，不用改变旧的体制。

我们来看一下代码：

```go
func main() {
stu := new(Student)

// 执行策略 add
stu.setStra(&add{})
resutAdd := stu.doStra(1,3)
fmt.Println(resutAdd)
// 执行策略 minus
stu.setStra(&minus{})
resutMinus := stu.doStra(1,2)
fmt.Println(resutMinus)
// 接下来我们看一下，是如何解除耦合的，我们改变一下，我们的策略执行者，你会发现下面的执行不会有任何的影响。
}

// 策略1的抽象接口,给策略定义一个基本的定义
type StrategyStudentOne interface {
	do(int, int) int
}

type add struct {
}
type minus struct {
}

func (ad *add) do(a, b int) int {
	return a + b
}

func (m *minus) do(a, b int) int {
	return a - b
}

// 定义执行策略的用户,将使用到的策略放入进去即可
type Student struct {
	strategy StrategyStudentOne
}

// 设置策略
func (s *Student) setStra(strategy StrategyStudentOne) {
	s.strategy = strategy

}

// 执行策略

func (s *Student) doStra(a, b int) int {
	return s.strategy.do(a, b)
}

```

