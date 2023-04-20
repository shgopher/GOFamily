# go语言常量

go使用const来声明常量，常量表达式的运算在编译期就可以完成了，并且go常量是类型安全的，编译器优化友好的元素。

常量中的数据类型只可以是布尔、数字（整数、浮点和复数）以及字符串型。

```go
const(
  startOne = -2
  startTwo = 1.0
  startThree = 3.14
  isTrue = true
  hi = "hi"
)
```
go推崇无类型常量，这主要是因为go不支持隐式的类型转化，只能显式转换，所以一旦常量给定类型，那么它跟变量之间或许就要将常量转化类型才能进行运算。

```go
package main

import "fmt"

const (
	PI = 3.14
)

func main() {	
	var a float32 = PI
	var b float64 = PI
}

```
可以看到，如果是没有类型的常量，可以非常灵活的赋予float32或者float64都可以。

## 局部常量

在 Go 中，局部常量在编译期常量折叠（compile-time constant folding）时会被编译器处理成字面值，因此局部常量在编译时会被编译。

例如，下面的代码定义了一个局部常量 x，并使用它来初始化一个变量 y：

```go
func main() {
    const x = 42
    var y = x * 2
    fmt.Println(y)
}
```
在编译时，x 将被处理成字面值 42，因此编译器将 y 初始化为 84，如下所示：

```go
func main() {
    var y = 84
    fmt.Println(y)
}
```
可见，局部常量在编译期已经被折叠成了字面值，因此在运行时不会再进行计算了。
## 常量不可包含计算
在 Go 中，常量必须在编译时就可以确定其值。因此，常量的值必须是一个编译时的常量表达式，不能包含运行时的计算。

不过，常量表达式可以包含一些简单的运算，例如加、减、乘、除、取模等算术运算，以及位运算、逻辑运算等。例如：

```go
const (
    x = 2 + 3            // 常量表达式，结果为 5
    y = (x * 2) % 10     // 常量表达式，结果为 0
    z = x == 5 || y == 0 // 常量表达式，结果为 true
)
```
常量表达式还可以使用一些内置函数，例如 `len`、`cap`、`make`、`new` 等，以及可以在编译期计算的一些标准库函数，例如   `math.Sin`、`math.Pi` 等。

总之，常量必须在编译期就可以确定其值，因此不能包含运行时的计算，但可以包含一些简单的算术、位运算、逻辑运算等。
## 枚举常量iota

go语言提供隐式重复前一个非空表达式的机制
```go
package main

import "fmt"

const (
	PI = 3.14
	a
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println(PI, a, b, c, d, e, f)
}
```

这段代码将会全部输出3.14

iota 是go语言的一个预定义标识符，它表示const声明块中，每一个常量所处位置的偏移量，它本身也是一个无类型的常量，它的初始值是0，意思是说此处的iota跟第一行比偏移了0个位置

```go
const(
  a = 1 << iota
  b
  c
  d
  e
  f
)
```
同一行的常量 iota值是一样的，可以理解为偏移量相同

```go
const(
  a,b = iota
  c,d
)
```

如果要让iota的初始值是1，那么可以这么做

```go
const(
  _ = iota
  a
  b
  c
  d
)
```
iota只需要在一个const群中出现一次即可，出现多次跟出现一次的效果一样

```go
const(
  a = iota
  b = iota
  c = iota
)
// 一样
const(
  a = iota
  b
  c
)
```
当iota群中出现异类该如何处理

```go
const(
  a = iota
  b = 12
  c = iota
  d
  e
  f
  g
)
```
答案就是 异类输出自己的，其它的常量不受影响，比如这里的输出就是 `0 12 2 3 4 5 6` ，只要记住iota是偏移位置就可以理解为什么是这么输出的了。

如果不考虑常量的灵活性，极致追求安全性，那么也可以给iota常量加上类型

```go
const(
  a int = iota
  b
  c
  d
)
```
这种就要求变量类型必须是int才能被这些枚举类型接纳，但是一般常量还是用无类型的较为常见。
