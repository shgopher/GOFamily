# go编程模式
导读：
- 面向接口编程
- 错误处理编程模式
- 函数式编程
- 


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
## 委托和反转控制
## Map-Reduce
## Go Generation
## 修饰器
## pipeline
## k8s visitor模式