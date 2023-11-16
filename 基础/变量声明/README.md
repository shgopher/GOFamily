# 变量声明
## 变量声明符号 `var` 和 `:=`
在 go 语言中，我们使用 `var` 来表示声明一个变量
```go
var a string = "hello world"
```
这就是一个标准的变量声明方式，var 符号在最前面，接着就是变量 a，变量后面紧跟着是变量的类型，这里是 string 类型，也就是字符串类型，`=` 后面是要赋值的具体值

我们也可以直接声明不赋予初始值，go 语言默认，声明即赋予初始值，那么这里的字符串类型初始值就是一个空字符串 `""`
```go
var a string
```
go 语言具有类型推断能力，所以我们可以省略类型，让 go 语言的编译器去推断类型
```go
var a  = "hello world"
```
编译器会自动推断 a 为 string 类型

在函数体内部 (这是一个先决条件) 我们**也**可以使用省略的写法，就是使用一个符号 `:=` 来充当 `var` 的角色，也就是初始化的工作，比如说
```go
func main(){
  a := "hello world"
}
``` 
可以看到，整个的用法是 `[变量] [:=] [初始值]` 这三者缺一不可，而且还不能多，不能在 a 后面带有类型，不能省略初始值，且仅限于函数/方法内部使用

go 语言支持多变量同时赋值
```go
var a, b string

var c, d string = "1", "2"

a, b := "hello world",12
```
其中，使用 var 进行声明的时候，如果是多个变量同时声明，必须是相同类型；使用 `:= ` 进行多变量赋值时，多个变量可以不同类型，因为全靠编译器推断

## 常见的变量声明方式
从广义上来说，go 语言只有两种变量，包一级的变量和函数一级的变量
```go
var DefaultValue int

func NewMethod(n int){

}
```
其中 `DefaultValue` 就是包级变量，`n` 就是函数级变量 (也是形式变量)

下面我列举一些常见的声明方式
```go
var (
 a string
 b int32 = 1
)

var c map[int]string //包级变量

func hi(){
  d := 12 //仅限函数内部使用，变量后面不能有类型
  var c map[int]string // 函数级变量
  var e = "hello world" //自动推断变量类型
}

```
可以说声明的方式很多，不过呢，在一个项目中应该尽量保证声明方式的一致性，因为可以加强代码的统一性，减少理解代码的难度。
## go 语言可导出变量
go 语言跟一般的语言不同，它使用变量**首**字母的大小写来区分变量的可导出性质，大写 (如果使用中文作为变量名称，默认是可导出的) 代表可导出，小写代表仅限包内部使用 (包这一级，多个文件只要是同一个包就可以使用)
```go
package Example

var OutPutName = 0o77
var inName = 0x99
```
其中 `OutPutName` 是一个可导出的变量，`inName` 是不可导出变量

go 语言还拥有比如函数，方法，结构体，接口，等等各类型的组件，这里你可以先不懂到底是什么，你先有一个印象，但凡首字母是大写的都是可以导出的，小写就是包内使用，没错，go 语言就是这么简单
## 包级变量
`第一种声明形式`：**声明的同时，显式初始化**

- `var a = method("t")` go 编译器根据右侧的返回值自动确定左侧变量的类型
- `var a = 3.14` 在没有具体返回值，没有具体类型的情况下，go 赋予它默认类型，比如 float 的默认类型就是 float64，int 类型的默认类型就是 int
- 如果要显式的赋予类型，并且保证命名的一致性 `var a int = 12` 的行为应该避免，应该写成 `var a = int(12)` 用来保证一致性
  ```go
  var(
    a = 3.14
    b = int(12)
    e = errors.New("EOF")
  )
  ```

`第二种声明形式`：**声明但是延迟初始化**

只有 `var a int64` 这一种方式，不过呢，go 语言的声明是直接赋予零值的，比如说这里的 a 默认就是 `0`

go 语言变量声明的聚集和就近原则：将同一类型的放在一个 var() 内部；或者另一种分类方法：将有初始值的放在一个 var 里，将延迟初始化的放在另一个 var 里。

分类一

```go
var(
a int
b int
c int
)
var(
d string
e string
f string
)
```

分类二

```go
var(
  a = 12
  b = "string"
  c = "string"
)
var(
  d int
  c string
  f bool
)
```
接下来谈一谈就近原则，变量的声明和变量的使用尽量的近，不要都声明在头部，如果一个变量被全局大量使用，那么可以放在头部，如果就是仅仅使用少量的次数，还是应该在使用的前面就进进行声明

## 函数级变量
`第一种声明形式`：**延迟初始化**

在函数体内使用 var

```go
func Method(){
  var a int
}
```
如果变量特别多，也可以使用 `var()` 的方法在函数内部使用

`第二种声明形式`：**声明且显式初始化的局部变量**

```go
func method1(){
  a := 12
  b := int32(20) // 改变默认类型
}
```
## 小心 shadow 的变量
我们知道，当有两个两个以上的变量在赋值时，如果其中有一个未被提前声明，那么就需要使用 `:=`，这个时候系统会自动判断有哪些未提前声明，然而有一种场景下系统会发生误判，准确的来说这是一种歧义，系统的判断会跟程序员的心理不一致，出现了变量 shadow 的行为，让我们看一下代码：
```go
func WithName(){
  var a string
  // tracing 为 bool 类型
  if tracing{
    a,err := example.Method()
  } else {
    a,err := example.Method1()
  } 
}
```
在外层，a 已经提前声明，但是在 if 这个作用域中，由于 err 并未提前声明，所以使用了 `:=`，由于系统**无法获知**这里的 a 是否需要再次声明，所以 go 语言默认 a 是一个新的变量，这样外层的 a 就无法得到新的值，外层 a 也就被内层的 a 给 shadow 了。

如果想改变这种 bug，我们可以将 err 也提前声明：

```go
func WithName(){
  var a string
  var err error
  // tracing 为 bool 类型
  if tracing{
    a,err = example.Method()
  } else {
    a,err = example.Method1()
  } 
}
```
或者也可以改变内部的变量名称，来改变这种 shadow：
```go
func WithName(){
  var a string
  // tracing 为 bool 类型
  if tracing{
    ai,err := example.Method()
    a = ai
  } else {
    ai,err := example.Method1()
    a = ai
  } 
}
```
## 参考资料
- https://juejin.cn/post/7241452578125824061



