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

go内置了很多约束，比如说 any 和 comparable ，意思是任何类型和可以比较的类型。

约束不仅仅可以单独写出来，还可以内置于函数内部。

```go
func Age[T int| string,B float64| string](i T,j B){}
```
这种形式下，T 和 B 的约束就是仅限此函数使用 

下面我们看一种形式，这种情况下约束的不仅仅是string和int，而是包含了底层是他们的所有数据，比如说 `type DD int` 也符合这个约束

```go
type st interface{
	~string | ~int
}
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
这是函数使用泛型的写法，当函数使用泛型的时候，需要在变量前使用中括号标注泛型的具体约束，然后后面才能使用这个泛型类型，使用泛型函数的时候，中括号是可以省略的`Age(12)` 系统会自动推算泛型的具体实现。顺便说一下，泛型类型使用`%v`作为占位符

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
}

type Age[T any] struct {
	I T
}

func (a *Age[T]) Post(t T) {
	fmt.Println(a.I, t)
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