# go语言常量

go使用const来声明常量，常量表达式的运算在编译期就可以完成了，并且go常量是类型安全的，编译器优化友好的元素。

```go
const(
  startOne = 0
  startTwo = 1
  startThree = 2
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
答案就是 异类输出自己的，其它的常量不受影响，比如这里的输出就是 `1 12 3 4 5 6 7` ，只要记住iota是偏移位置就可以理解为什么是这么输出的了。

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
## 参考资料
- 图书：go语言精进之路