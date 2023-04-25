# go常见设计模式

常见的设计模式有创建型，结构型，行为型，这里只讲go项目中常见的设计模式。

创建型：特点是关注创建对象的方式；结构型：特点是关注对象的结构，各种组合的模式；行为型：关注对象之间的通信。

创建型：
- 单例模式
- 工厂三种模式

结构型：
- 策略模式
- 模版模式

行为型：
- 代理模式
- 选项模式

本文档在这里只会涉及到这些内容。因为在go的项目中，这几个才是最常用的，更多关于设计模式的内容，在下面的408篇中会更加详细的阐述。

## 单例模式
单例模式的意义就是，只执行一次，在go语言中，我们使用懒汉模式，或者饿汉模式，都是不优雅的，使用`sync.Once` 才是正确的选择

> sync.Once 属于加锁的懒汉模式 

```go
type singleton struct {}
var app *singleton
var once sync.Once
// 使用CreateApp once.Do只能执行一次
func CreateApp()*singleton{
	once.Do(func() {
		app = &singleton{}
	})
	return app
}
```

## 简单工厂
简单工厂的目的也是为了创建，我们可以看一下一段代码
```go
type People struct {
	name string
	year int
}

func (p *People) Post() {
	fmt.Println(p.name, p.year)
}

func NewPeople(name string, year int) *People {
	return &People{
		name: name,
		year: year,
	}
}
```
使用这种方式，newpeople可以接受参数，然后返回一个结构体。
## 抽象工厂
抽象工厂和简单工厂的区别是它不返回结构体，返回一个接口。
```go
type People interface {
	Do(req *http.Request) (*http.Response, error)
}
type people struct {
	name string
	year int
}

func (p *people) Do(req *http.Request)(*http.Response, error) {
	rec := httptest.NewRecorder()
	return rec.Result(),nil
}

func Newpeople(name string, year int) People {
	return &people{
		name: name,
		year: year,
	}
}
```
可以看到 new方法返回的是一个接口类型，当然了return的是一个实现了接口的structure，只不过这里有一个隐含的类型转换。

这样做的好处是什么呢，比如现在有另一个new，newStudent ,它当然也定义了一个Do

## 工厂方法
工厂方法的宗旨就是解耦，使用子函数来代替自身，这样就可以解除耦合，举个简单的例子：

```go
func main() {
	// 16岁的分成一组
	student16 := NewStudent(16)
	student16("red")
	student16("green")
	student16("blue")
	// 17岁的分成一组
	student17 := NewStudent(17)
	student17("red")
	student17("green")
	student17("blue")
	// 18岁的分成一组
	student18 := NewStudent(18)
	student18("red")
	student18("green")
	student18("blue")
}

type Student struct {
	name string
	year int
}

func NewStudent(year int) func(name string) Student {
	return func(name string) Student {
		return Student{
			name: name,
			year:year,
		}
	}
}


```
我们可以清晰的看出，我们将年龄作为一个root，然后name是child ，那么就是将 year和name解耦了，我们使用子函数去创建name这child选项。

所以可以看到上文中，我们创建了16岁的红蓝绿，17岁的红蓝绿，和18岁的红蓝绿 我们将year和name解耦了。

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

## 模版
模版模式的核心是，公有方法 和子方法的解耦。
也就是说，一个接口中的众多方法，分为公有方法和子方法，公有方法是事先定义好的，强制子方法去实现各自的方法。

```go
func main(){
	a1 := new(A1)	
	Do(a1)
	// 开始第二个部分的执行。
	a2 := new(A2)	
	Do(a2)
}

type Cook interface {
	first()
	second()
	last()
}

type A struct {
}

func (A) first()  {
	fmt.Println("first")
}

type A1 struct {
	A
}

func (a *A1) second()  {
	fmt.Println("second1")
}
func (a *A1) last()  {
	fmt.Println("last1")
}
type A2 struct {
	A
}
func (a *A2) second()  {
	fmt.Println("second2")
}
func (a *A2) last()  {
	fmt.Println("last2")
}

func Do(cook Cook){
	cook.first()
	cook.second()
	cook.last()
}
```
## 代理
```go
type Student struct {
	name string
}

type ProxyStudent struct {
	student *ProxyStudent
}
```
这就是代理模式。
## 选项


