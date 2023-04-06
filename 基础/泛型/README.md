# 泛型
**导读：**
- 约束
- 使用方法
- 实现原理 
- 跟其它语言的泛型进行对比
- 用例子学泛型
- issues

> 泛型需满足 `go1.18+`
## 约束
go使用interface作为约束，约束的意思是约束了这个泛型都具有哪些实际类型。所以可以理解为，go将interface的职责给扩展了，让接口不仅仅作为接口 --- 解耦的，抽象化的结构体，还具有了约束，对于类型的约束作用。

go可以将所有的接口（包括经典接口，和泛型以后的接口）都用作约束，但是可以不代表应该，要有选择的去使用约束，但是约束并不是都可以作为传统的接口来使用的，例如传统的接口只能存在方法，并不能存在类型，**只要存在类型就自动归纳为约束**

综上所述，第一，约束的概念大于接口，只有传统接口可以作为抽象类型去使用，约束只能存在于函数或者类型之中，不能单独使用。第二，不要把传统接口用在约束的地方，这种用法是不合适的，通常来说传统接口的模式就用作抽象类型的解耦即可（用了也不错，只是不建议）。

```go
type st interface{
  int | string
}
```
这里 st约束拥有int和string，请注意这里的st是约束，不是泛型类型

go内置了很多约束，比如说 any 和 comparable ，意思是任何类型和可以比较的类型。以后 ***应该*** 会有一个新的内置约束的包叫做`package constraints` 例如any comparable ，Ordered 等等约束都会内置到标准库中

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
所有说这里就可以看出来，在引入泛型之后，go的interface的功能扩充了。


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
这里的意义就是 **and**的意思 就是说这个约束是可以比较的还是必须得支持`hash()uintptr` 

下面这种方法也是可以的

```go
type NumericAbs[T any] interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
	Abs() T
}
```
上面的类型意思是满足数字类型，下面的意思是满足这个方法，所以最终实现这个约束，就是一个对象是数字类型，并且实现了这个接口

当结构体中使用泛型的时候，泛型不能直接作为嵌入使用
```go
type Lockable[T any] struct {
	T // 正确的方法应该是 T T
	mu sync.Mutex
}

```
错误提示： `embedded field type cannot be a (pointer to a) type parameter`

我们再看一下当泛型结构体嵌入到其它结构体中如何使用

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
}

type A[T any] struct {
	T T
}

type B[T any] struct {
	// 这里也可以使用 A[T] 前者是输入了实际类型，
	//后者虽然输入的是泛型类型T，但是同样属于一个具体类型
	A[C] 
	T T
}

type C struct {
}
```

可以看出关键点，泛型结构体被嵌入其它结构体的时候，泛型要给实际的类型才可以

这句解释是跟方法中的泛型做对比的

```go
type A[T any] struct{}

func(a A[T])Get(){}
```
我们知道方法里的泛型是使用这种方法，方法里的泛型是可以直接从struct定义的地方继承这个泛型T的，但是结构体的嵌套或者使用，就不能继承了，所以得给实际的类型不然就会报错  `embedded field type cannot be a (pointer to a) type parameter`

约束里的泛型同样不能直接嵌入使用

```go
type B[T any] interface {
	T
}
```
错误提示： `cannot embed a type parameter`

只有泛型类型是方法的参数时才可以这么做，比如说

```go
type EmbeddedParameter[T any] interface {
	~int | ~uint 
	me() T 
```
如果想使用这种约束，可以这么使用

```go
func Abs[T EmbeddedParameter[T]](t T)T{}
```
解释一下，中括号里面泛型的两个T表达的意思是不一样的，后面的T表达的是 约束里的泛型，表示 any，前面的T表示的是满足后面的这个约束的类型T，但是这里注意，后面T虽然之前定义的时候是any但是这里被改变了，改变为了必须满足约束 `EmbeddedParameter`的类型，如果说的通俗点，从any变成了，满足 `int ｜ uint and 实现 me()T方法 `后文会有代码进行解释。

当然了，后面的T没有也行，如果没有后面的T就是相当于不改变后面的T的约束类型了

```go
type Differ[T1 any] interface {
	Diff(T1) int
}

func IsClose[T2 Differ](a, b T2) bool {
	return a.Diff(b) < 2
}

```

当使用了泛型之后，是无法使用断言的，这是非法的，那么如果一定要在运行时的时候去判断类型怎么办呢？答案就是转变成`interface{}`即可，因为我们知道任何对象都已经实现了空接口，那么就可以被空接口去转化

```go
func GeneralAbsDifference[T Numeric](a, b T) T {
	switch (interface{})(a).(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64:
		return OrderedAbsDifference(a, b) 
	case complex64, complex128:
		return ComplexAbsDifference(a, b) 
	}
}
```

下面看一下别名的真实类型是泛型的情况

```go
type A[T any] []T

type AliasA = A // 错误 ❌

type AliasA = A[int] // 正确
```

其中错误的问题是 别名不能直接使用泛型类型 `cannot use generic type A[T any] without instantiation`，它需要泛型的实例化
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

接下来我们看一下，如何使用 `type a[T any] interface{} 有类型也有方法`的泛型结构

```go
package main

import "fmt"

func main() {
	var d DDD
	var i DDD
	d = 1
	i = 2
	io := AbsDifference[DDD](d, i)
	fmt.Println(io)
}

type DDD int

func (ddd DDD) Abs() DDD {
	return ddd + ddd
}

type NumericAbs[T any] interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
	Abs() T
}

// AbsDifference computes the absolute value of the difference of
// a and b, where the absolute value is determined by the Abs method.
func AbsDifference[T NumericAbs[T]](a, b T) T {
	d := a - b
	return d.Abs()
}
```
## 实现原理
泛型的第一种方法是在编译这个泛型时，使用一个字典，里面包含了这个泛型函数的全部类型信息，然后当运行时，使用函数实例化的时候从这个字典中取出信息进行实例化即可，这种方法会导致执行性能下降，一个实例化类型`int, x=y`可能通过寄存器复制就可以了，但是泛型必须通过内存了（因为需要map进行赋值），不过好处就是不浪费空间

还有一种方法就是把这个泛型的所有类型全部提前生成，这种方法也有一个巨大的缺点就是代码量直线上升，如果是一个包的情况下还能根据具体的函数调用去实现该实现的类型，如果是包输出的的情况下，那么就得不得不生成所有的类型。

所以将两者结合在一起或许是最好的选择。

这种方法是这样的，如果类型的内存分配器/垃圾回收器呈现的方式一致的情况下，只给它生成一份代码，然后给它一个字典来区分不同的具体行为，可以最大限度的平衡速度和体积
## 跟其它语言的泛型进行对比
- c语言：本身不具有泛型，需要程序员去实现一个泛型，实现复杂，但是不增加语言的复杂度（换言之只增加了程序员的）
- c++和rust：跟go基本保持一种方式，就是增加编译器的工作量
- Java：将泛型装箱为object，在装箱和拆箱擦除类型的过程中，程序执行效率会变低
## 用例子学泛型
理论学习完了，不使用例子进行复习的话会忘的很快的。跟着我看几个例子吧

***例子一：*** 函数泛型 `map-filter-reduce`

```go
package main

import (
	"fmt"
)

func main() {
	vM := Map[int]([]int{1, 2, 3, 4, 5}, func(i int) int {

		return i + i
	})
	fmt.Printf("map的结果是：%v", vM)
	vF := Filter[int]([]int{1, 2, 3, 4, 5}, func(t int) bool {
		if t > 2 {
			return true
		}
		return false
	})
	fmt.Printf("filter的结果是:%v", vF)
	vR := Reduce[Value, *Result]([]Value{
		{name: "tt", year: 1},
		{name: "bb", year: 2},
		{name: "7i", year: 3},
		{name: "8i", year: 4},
		{name: "u4i", year: 5},
		{name: "uei", year: 6},
		{name: "uwi", year: 7},
		{name: "uti", year: 8},
	}, &Result{}, func(r *Result, v Value) *Result {
		r.value = r.value + v.year
		return r
	})
	fmt.Println("reduce的结果是：", vR.value)

}

// Map:类似于洗菜，进去的菜和出来的菜不一样了所以需要两种种类
func Map[T1, T2 any](arr []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(arr))
	for k, v := range arr {
		result[k] = f(v)
	}
	return result
}

// Filter:类似于摘菜，进去的菜和出来的菜是一种，不过量减少了
func Filter[T any](arr []T, f func(T) bool) []T {
	var result []T
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce:类似于做菜，将菜做成一道料理，所以需要两种类型
func Reduce[T1, T2 any](arr []T1, zero T2, f func(T2, T1) T2) T2 {
	result := zero
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}

type Value struct {
	name string
	year int
}
type Result struct {
	value int
}
```

`map的结果是：[2 4 6 8 10] filter的结果是:[3 4 5] reduce的结果是： 36`

***例子二：*** 方法上的泛型 `sets`
```go
package main

import (
	"fmt"
)

func main() {

	// 这里 Sets的具体类型和Make的具体类型都是int,所以可以正常赋值
	var s Sets[int] = Make[int]()
	//
	s.Add(1)
	s.Add(2)
	fmt.Println(s)
	fmt.Println(s.Contains(3))
	fmt.Println(s.Len())
	s.Iterate(func(i int) {
		fmt.Println(i)
	})
	fmt.Println(s)
	s.Delete(2)
	fmt.Println(s)
}

// Sets 一个key  存储对象
type Sets[T comparable] map[T]struct{}

// Make 实例化一个map
func Make[D comparable]() Sets[D] {
	// 泛型就像一个管道一样，只要实例化的时候管子里的东西一致，那么就是一根管子
	return make(Sets[D])
}

// Add 向这个sets添加内容
func (s Sets[T]) Add(t T) {
	s[t] = struct{}{}
}

// delete ,从这个sets中删除内容
func (s Sets[T]) Delete(t T) {
	delete(s, t)
}

//  Contains 播报t是否属于这个sets
func (s Sets[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

//Len sets拥有的长度

func (s Sets[T]) Len() int {
	return len(s)
}

// Iterate 迭代器，并且给予每个元素功能

func (s Sets[T]) Iterate(f func(T)) {
	for k := range s {
		f(k)
	}
}
```
`map[1:{} 2:{}]
false
2
1
2
map[1:{} 2:{}]
map[1:{}]
`

***例子三：*** 外部定义的约束 `实现一个sort接口类型`

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
// ~ 代表只要底层满足这些类型也可以算满足约束
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uintptr |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}
type orderedSlice[T Ordered] []T

func (s orderedSlice[T]) Len() int           { return len(s) }
func (s orderedSlice[T]) Less(i, j int) bool { return s[i] < s[j] }
func (s orderedSlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func OrderedSlice[T Ordered](s []T) {
	sort.Sort(orderedSlice[T](s))
}
```
## issues
***问题一：*** 关于泛型中的零值

在go里面对泛型的零值并没有一个所谓的泛型零值可以使用，需要根据不同的实践去实现，比如

```go
package main

import "fmt"

func main() {
	
}

type Aget[T any] struct {
	t *T
}
// 根据实际判断，如果a的t不等于nil再返回，如果是nil就返回一个T类型的nil（意思就是只声明）
func (a *Aget[T]) Approach() T {
	if a.t != nil { 
		return *a.t
	}
	var r T
	return r
}

```
实际上目前，还没一个确切的泛型的零值，那么我们要做的只能是按照实际来具体分析，按照提案，以后有可能使用`return ...` `return _` `return ` `return nil` `return T{}` 这些都是可能的结果，我个人比较喜欢 `return T{}` 来表示泛型的零值，或许在go1.19或者go2的时候能实现，拭目以待吧。

***问题二：*** 无法识别使用了底层数据的其它类型

```go
type Float interface {
	~float32 | ~float64
}

func NewtonSqrt[T Float](v T) T {
	var iterations int
	switch (interface{})(v).(type) {
	case float32:
		iterations = 4
	case float64:
		iterations = 5
	default:
		panic(fmt.Sprintf("unexpected type %T", v))
	}
	// Code omitted.
}

type MyFloat float32

var G = NewtonSqrt(MyFloat(64))
```
这里约束 Float拥有的约束类型是 `~float32` 和 `float64`当在switch中定义了float32和flaot64时，无法识别下面的新类型 MyFloat即使它的底层时 float32 ，go的提议是以后在switch中使用 `case ~float32:` 来解决这个问题，目前尚未解决这个问题

***问题三：*** 即便约束一致，类型也是不同的

```go
func Copy[T1, T2 any](dst []T1, src []T2) int {
	for i, x := range src {
		if i > len(dst) {
			return i
		}
		dst[i] = T1(x) // x 是 T2类型 不能直接转化为 T1类型
	}
	return len(src)
}

```
T1,和T2 虽然都是any的约束，但是啊，它不是一个类型啊！

```go
Copy[int,string]() // 这种情况下，你能说可以直接转化吗？？？
```
这种代码可以更改一下

```go
dst[i]= (interface{})(x).(T1)
```
确认是一种类型以后才能转化

## 参考资料
- https://coolshell.cn/articles/21615.html
- https://go.dev/doc/tutorial/generics
- https://colobu.com/2021/08/30/how-is-go-generic-implemented/
- https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/generics
- https://mp.weixin.qq.com/s/zKnh_iPm8skxWv3rxaOscw
- https://mp.weixin.qq.com/s/BFsoQPvrog_sMKMTEofZyQ