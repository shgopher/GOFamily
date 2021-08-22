# go编程模式
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

## 错误处理
## function options
## 委托和反转控制
## Map-Reduce
## Go Generation
## 修饰器
## pipeline
## k8s visitor模式