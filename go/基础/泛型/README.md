# 泛型
**导读：**
- 约束
- 使用方法
- 实现原理 
- 跟其它语言的泛型进行对比

> 泛型需满足 `go1.18+`
## 约束
go使用interface作为约束，约束的意思是约束了这个泛型都具有哪些实际类型。
```go
type st interface{
  int | string
}
```
这里 st约束拥有int和string，请注意这里的st是约束，不是泛型类型

go内置了很多约束，比如说 any 和 comparable ，意思是任何类型和可以比较的类型。以后应该会有一个新的内置约束的包叫做`package constraints` 例如any comparable ，Ordered 等等约束都会内置到标准库中

约束不仅仅可以单独写出来，还可以内置于函数内部。

```go
func Age[T int| string,B float64| string](i T,j B){}
```
这种形式下，T 和 B 的约束就是仅限此函数使用 

下面我们看一种形式，这种情况下约束的不仅仅是string和int，而是包含了底层是他们的所有数据，比如说 `type DD int` 也符合这个约束，请记住只能使用在底层类型上，如果使用`~DD`是不被允许的

```go
type st interface{
	~string | ~int
}
```
于此同时，约束也不仅仅是基础类型，约束的内容是方法也是可以的

```go

func ConcatTo[S Stringer, P Plusser](s []S, p []P) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = p[i].Plus(v.String())
	}
	return r
}

type Plusser interface {
	Plus(string) string
}
type Stringer interface {
	String() string
}

```
所有说在引入泛型之后，go的interface的功能其实是扩充了，也可以将泛型中的约束就称之为接口也没问题。

我们说到interface相当于是被扩充了，那么约束的时候可以即有方法又有类型吗？

```go
type anys interface {
	int
	me()
}
```
答案也是不行，要么是只有类型要么是只有方法

这是只有方法的例子
```go
package main

import "fmt"

func main() {
	var a A
	var dd DD[A]
	dd.TT(a)
}

type A struct{}

func (a A) me() {}

type anys interface {
	me()
}

type DD[T anys] []T

func (dd *DD[T]) TT(t T) {
	fmt.Println(t, len(*dd))
}
```

约束跟接口是一样的也是可以嵌套的

```go
type ComparableHasher interface {
	comparable
	Hash() uintptr
}

// or

type ImpossibleConstraint interface {
	comparable
	[]int
}
```
这里的意义就是 and的意思 就是说这个约束是**可以比较的**还是必须得支持`hash()uintptr` 这个函数,不过有个界定，就可以嵌套的只能是内置的约束，自定义的无法嵌套


```go
//无法运行
type ans interface{

	int | string
}
type ImpossibleConstraint interface {
	ans
	[]int
}
```
解释一下，我们知道使用`|`表示约束的类型是或的意思，那么如果使用嵌套，就是和的意思，使用any或者comparable都是有意义的，比如说在一个方法的约束里，嵌入了一个comparable，意思就是实现这个方法的并且是可以比较的类型，就会满足这个约束，而自定义的通常不满足这些意义，比如说 一个约束里面是`int | string` 再嵌套一个 `float64` 这就是说 `int ｜ string` 得同时再满足 `float64`显然这是不可能的。

那么这里有一个疑问，可以给约束嵌入泛型吗？例如

```go
type EmbeddedParameter[T any] interface {
	T 
}

```
`cannot embed a type parameter` 答案是否定的，go不允许这么做。不过如果约束里面是方法就可以这么做，比如说

```go
type EmbeddedParameter[T any] interface {
	me() T 
```

## 使用方法
下面展示一下go泛型的基本使用方法
```go
package main

import "fmt"

func main() {
	fmt.Printf("%v",Age[int](12))
}

func Age[T any](t T) T{
	return t
}
```
这是函数使用泛型的写法，当函数使用泛型的时候，需要在变量前使用中括号标注泛型的具体约束，然后后面才能使用这个泛型类型，使用泛型函数的时候，中括号是可以省略的`Age(12)` 系统会自动推算泛型的具体实现。顺便说一下，泛型类型使用`%v`作为占位符，泛型类型无法进行断言，这一点跟`interface{}`不同。

当然了，我么也可以不用any，自定义一个约束

```go
package main

import "fmt"

func main() {
	Age[int](12)
}
type st interface{
  int | string
}
func Age[T st](t T) {
	fmt.Println((t))
}
```

看完了在函数内的泛型，我们在看一下在方法中如何使用泛型
```go
package main

import "fmt"

func main() {
	new(Age[int]).Post(12)

	var dd DD[int]
	dd.TT(12)
}

type Age[T any] struct {
	I T
}

func (a *Age[T]) Post(t T) {
	fmt.Println(a.I, t)
}

type DD[T any] []T

func(dd *DD[T])TT(t T){
	fmt.Println(t,len(*dd))
}
```
在 age 结构体声明的时候，声明了一个泛型 T ，在struct体内就可以使用这个T，值得注意的是，这个结构体方法内部仅可以使用定义在这个结构体对象上的泛型

下面是一个**错误案例**

```go
func (a *Age[T])Post[B any](t T,b B) {
	fmt.Println(a.I, t)
} 
```
`syntax error: method must have no type parameters`
## 实现原理
泛型的第一种方法是在编译这个泛型时，使用一个字典，里面包含了这个泛型函数的全部类型信息，然后当运行时，使用函数实例化的时候从这个字典中取出信息进行实例化即可，这种方法会导致执行性能下降，一个实例化类型`int, x=y`可能通过寄存器复制就可以了，但是泛型必须通过内存了（因为需要map进行赋值），不过好处就是不浪费空间

还有一种方法就是把这个泛型的所有类型全部提前生成，这种方法也有一个巨大的缺点就是代码量直线上升，如果是一个包的情况下还能根据具体的函数调用去实现该实现的类型，如果是包输出的的情况下，那么就得不得不生成所有的类型。

所以将两者结合在一起或许是最好的选择。

这种方法是这样的，如果类型的内存分配器/垃圾回收器呈现的方式一致的情况下，只给它生成一份代码，然后给它一个字典来区分不同的具体行为，可以最大限度的平衡速度和体积
## 跟其它语言的泛型进行对比
- c语言：本身不具有泛型，需要程序员去实现一个泛型，实现复杂，但是不增加语言的复杂度（换言之只增加了程序员的）
- c++和rust：跟go基本保持一种方式，就是增加编译器的工作量
- Java：将泛型装箱为object，在装箱和拆箱擦除泛型的过程中，程序执行效率会变低
## 参考资料
- https://coolshell.cn/articles/21615.html
- https://go.dev/doc/tutorial/generics
- https://colobu.com/2021/08/30/how-is-go-generic-implemented/
- https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md