# go编程模式

导读：
- 面向接口编程
- 错误处理编程模式
- 函数式编程
- “控制代码”独立模式
- Map-Reduce
- 修饰器
- pipeline
- k8s visitor模式


## 面向接口编程
在go语言中，面向接口编程是一个重要的编程模式。Java，python等编程语言中，推崇的是“面向对象”的编程方式，特征是拥有类这种抽象接口，以及事例对象。当然他们也有接口，理论上java python也可以进行面向接口编程，但是在这些编程语言中oop的优先级比面向接口更高一些。

下面这段代码可以看出，面向接口编程，把接口，以及接口设立的抽象方法，当作抽象的主体，使用隐藏式的接口实现，从而实现了编码上的解耦合，下面这段代码中，control实现的就是一个抽象方法，因为它使用了抽象的接口方法，通过这个抽象方法control，将control中的流程与struct具体的实现通过接口方法，进行了解除耦合。

在go的标准库中，这种实现的方法很多，也非常常见，比较著名的案例比如 `io.Read` 和 `ioutil.ReadAll`

```go
package main

func main() {
	g := new(Girl)
	b := new(Boy)
	Control(g)
	Control(b)
}

//接口
type Speaker interface {
	GetName(id int) string
}
type Boy struct {
}

type Girl struct {
}

// boy实现接口
func (b *Boy) GetName(id int) string {
	return ""
}

// girl实现接口
func (b *Girl) GetName(id int) string {
	return ""
}

// 处理函数
func Control(s Speaker) bool {
	return true
}

```
### 面向接口编程中的"接口实现验证"
我们要确定一个对象已经实现了，并且是实现了接口的全部方法，这个时候我们需要进行一个验证
```go
type Speaker interface{
    a()
    b()
    }
// 假设 Gril仅仅实现了a()
type Girl struct

var _ Speaker = (*Girl)(nil)
```
这个时候，编译期间一定会报错，因为无法将类型是 *Girl的nil转为 Speaker类型的nil 
> 注意一下，nil是有类型的哦。
## 错误处理模式
go推崇要处理各种错误，但是有些场景是这些错误都是一类错误，没有必要每一个都处理，那么我们这里就是要处理这种情况。

error check hell:

```go
	if err := binary.Read(); err != nil {
		return nil, err
	}

	if err := binary.Read(); err != nil {
		return nil, err
	}

	if err := binary.Read(); err != nil {
		return nil, err
	}

	if err := binary.Read(); err != nil {
		return nil, err
	}
```
我们可以定义一个stuct，给一个方法，里面放进去if err != nil 就可以简略这个过程

```go
type A stuct {
	//
	err error
}
func (a *A)Read(){
	if a.err == nil {
		a.err = //
	}
}
func main(){
	a = new(A)
	a.Read()
	a.Read()
	a.Read()
	a.Read()
	if a.err != nil {
		//
	}
}
```
也就是这种场景可以用这种方法进行优化，实际上大多数场景都需要我们老老实实的去处理各种err，错误处理的越仔细，发现问题的时候定位错误就会越容易。

## 函数式编程模式
我们先看一个例子

```go

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

func NewServerTimeout()*Server {
}

func NewServer() *Server{
}

func NewXXX() *Server{

}

```
这段代码的意义就是new一个server，并且不同的场景，比如说timeout了呀，正常状态啊，是不同的函数，要new好多个，那么这个时候我们就可以使用函数式编程。

```go
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}
// 要用的函数类型
type Option func(*Server)

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func Stl(stl *tls.Config) Option {
	return func(s *Server) {
		s.TLS = stl
	}
}
func NewServer(addr string, port int, opions ...Option) *Server {
	// 给定默认值
	ser := &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "xx",
		Timeout:  time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	// 二次赋值
	for _, option := range opions {
		option(ser)
	}
	return ser
}

```
这种模式非常的直观，而且new函数只需要一个，并且除了端口和地址，我们需要强制给定，其他的都有默认值，只有给具体的值时，才会进行二次的赋值，可以说这种方法非常的直观，简洁，并且高可扩展。
## “控制代码”独立模式
这种编程模式的核心就是讲逻辑代码和控制代码分离，逻辑代码去使用控制代码去做事，控制代码相当于构造器一样的没有实际意义的辅助函数，这种写法拥有高度的可扩展性。
```go
type IntCout struct {
	value map[int]bool
}

func (i *IntCout) Add(value int) {
	i.value[value]=true
}
func(i *IntCout)Delete(value int){
	delete(i.value,value)
}


```
当我们想增加功能的时候：

```go
type AnotherIntCout struct {
	intCount
	something
}
//override
func(a *AnotherIntCout)Add(){}
func(a *AnotherIntCout)Del(){}
func(a *AnotherIntCout)AnotherMethod(){}
```
这种办法其实就是控制代码侵入了逻辑代码，我们要做的事情就是，讲控制代码单拎出来，然后让他实现功能，逻辑代码嵌套控制代码，因为控制代码基本上是很稳定的，毕竟功能较少，而且不轻易改动，所以代码可以改成这样

首先先讲控制代码单拎出来
```go
type Something struct{}
func(*Something)AnotherMethod(){}
```
然后这个时候，让逻辑代码去继承这个控制代码

```go
type IntCout struct {
	value map[int]bool
	Something
}
// 重写这个逻辑代码
func(*IntCount)Add(){}
```

使用这种方法，即便是逻辑代码再怎么改，控制代码丝毫不变，就跟你家灯随便换，但是开关不需要怎么动，而且可以更好的扩展更多的代码。

## Map-Reduce
我将map-filter-reduce模式称之为做菜理论，map的作用是将菜洗干净，filter的作用是将洗好的菜中，老的不新鲜的菜取出来扔掉，reduce的作用是将这些菜拌一拌变成一道佳肴。

![](https://gitee.com/shgopher/img/raw/master/%E5%A4%A7%E7%9B%98%E9%B8%A1.png)

- map: 怎么进怎么出。
- filter: 怎么进怎么出，只是数量少了。
- reduce: 多个进，一个出，要成品了。

### Map:

map的意义就是数据预处理。前面的切片中的数据调用一个map函数，然后处理一下。

```go
func main() {
	fmt.Println(MapStrToStr([]string{"A", "B"}, func(str string) string {
		return str
	}))

	fmt.Println(MapStrToInt([]string{"1", "2"}, func(str string) int {
		i, _ := strconv.ParseInt(str, 10, 0)
		return int(i)
	}))
}
// 
func MapStrToStr(str []string, fn func(str string) string) []string {
	var ma []string
	for _, value := range str {
		ma = append(ma, fn(value))
	}
	return ma
}
func MapStrToInt(str []string, fn func(str string) int) []int {
	var ma []int
	for _, value := range str {
		ma = append(ma, fn(value))
	}
	return ma
}

```
### Reduce:
reduce的意义就跟你将切好的菜，融会贯通给它融合了做成一盘菜。所以说你看进入了一个str的slice，只出来了一个sum
```go
func Reduce(str []string,fn func(string)int)int{
	sum := 0
	for _,v := range str{
		sum+= fn(v)
	}
	return sum
}
```
所以说通常来说 map进去什么样子，出来还是那个基本造型，比如进去是一个切片出来还是个切片，但是reduce就是进去很多东西但是出来不一样了，例如这个例子，进去了很多slice，出来了一个东西sum。
### Filter:
filter就是摘菜，通过if fn方法，将可以使用的再形成一个新的slice输出。
```go
func Filter(str []string,fn func(string)bool)[]string{
	ma := []string{}
	for _,v := range str{
		if fn(v) {
			ma = append(ma,v)
		}
	}
	return ma
}
```

### interface{}泛型：
我们使用map函数为例

```go
func Map(data interface{}, fn interface{}) []interface{} {
	dataR := reflect.ValueOf(data)
	fnR := reflect.ValueOf(fn)
	result := make([]interface{}, dataR.Len())
	for i:= 0;i < dataR.Len();i++ {
		result[i]= fnR.Call([]reflect.Value{dataR.Index(i)})[i].Interface()
	}

	return result
}
```
调用的时候可以使用

```go
Map([]string{"1"}, func(i string)string {
	return i + i
})
```
or

```go
Map([]int{1}, func(i int)int {
	return i*i
})
```
### go1.18+ 泛型：
> go verison 1.18+

在使用了泛型后我们的代码就可以更改为下面这种表达方式：

```go
func Map[T any](data []T, fn func(T)T) []T {
	var ma []T
	for _, value := range data {
		ma = append(ma, fn(value))
	}
	return ma
}
```

调用的时候就可以这样做：

```go
	Map([]string{"1"}, func(i string)string {
		return i + i
	})

	Map([]int{1}, func(i int)int {
		return i + i
	})
```

不得不承认啊，泛型真香😂
## 修饰器
这种模式其实就是函数式编程的一种，它的主要思想就是传入一个函数，然后返回的还是一个函数，我们将传入的这个函数进行二次修饰，然后再返回，进而调用使用：

```go
func decorator(fn func(s string)string)func(string)string {
	return func(s string) string {
		return fn(s) + "。。。"
	}
}
```

调用：

```go
	hello  := func(s string)string {
		return s
	}

	defn := decorator(hello)
	defn("你好")
	// print: 你好。。。
```
## pipeline

```go
func A(des ...fn()){
	for _,v := range des {
		v()
	}
}

// 调用的时候
A(fn1,fn2,fn3,fn4)

```
这就是属于基本的pipeline模式了。
## k8s visitor模式

k8s的visitor模式的意义就是将sturct数据结构和算法，解耦。


```go
package main

func main() {
	p := new(Peo)
	p.year = 10
	p.name = "a"
	p.Did(Run)
	// 这个时候，即便我们改变了p的值，那么这个算法 Run 也不会有任何的关联，他们俩完全解除耦合了
	p.year = 100
	p.Did(Run)
}

type Visitor func(Do)
type Do interface {
	Did(Visitor)
}

// 数据结构
type Peo struct {
	name string
	year int
}

func (p *Peo) Did(v Visitor) {
	v(p)
}


// 算法
func Run(do Do) {
	// 这里就是通过接口对象，来进行一系列的操作，真实的数据结构和这里的算法完全解除耦合
	fmt.Println(do)
}

```
