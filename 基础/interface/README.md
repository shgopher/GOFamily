<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-05-06 22:46:02
 * @FilePath: /GOFamily/基础/interface/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 接口
> 有关泛型相关的约束内容均放置在[泛型](../泛型)这一章节，这里不再赘述

重点内容前情提要
- 方法集合决定接口的实现
- 接口的嵌入
- 接口类型的底层
- 空接口
- 接口的相等性
- 小接口的意义
- 接口提供程序的可扩展性
- 接口提供程序的可测试性
- 接口的严格对比函数的宽松

接口，go语言提供的，用于抽象以及数据解耦的组件，**在操作接口时，go语言要求的严格程度远大于函数和方法。**

在 go1.18+ 中接口的概念从包含抽象的方法变成了类型，但是我们在外部可以操作的接口仍然只能是经典接口，不过经典接口是可以当做约束在内部使用的，不过官方不推荐经典接口当做一般约束那种使用方式。

go语言为了区分，将传统接口称之为接口，将扩展的接口称之为约束，实际上传统接口是约束概念的子集。扩展的约束并不能在函数或者方法以及类型之外使用。

这一章我们只介绍经典接口的基本概念和使用方法。

## 方法集合决定接口的实现
通常，我们使用下面这种形式去完成接口的实现
```go
type Writer interface {
	Write([]byte) (int, error)
}
type BufferWrite struct {
	value bytes.Buffer
}

func(b *BufferWrite) Write(p []byte) (int, error){}
```
我们知道，go语言规定，不可跨包去实现类型的方法，所以我们只讨论自定义的类型与接口之间的关系，首先抛出结论：松耦合关系。

go语言使用一种叫做鸭子理论的方式去实现接口：只要实现了接口的方法（go泛型, go1.18+, 理论从方法改为了类型）就算是实现了这个接口，这属于隐式接口实现。

通常来说，例如es6，Java等都是要显式的说明实现了哪个接口的，但是go不需要，它只需要实现方法即可，而且，它还可以实现多个接口，毕竟多实现几个方法就可以算是实现了这个接口。由此还可以推断出，go可以多实现方法，而不会影响对接口的实现。

这里要说明一下，go语言对于定义在指针类型上的变量有语法糖：

```go
type Writer interface {
	get()
	set()
}
type Student struct{}
func(s Student) get() {}
func(s *Student) set() {}

func main(){
	var t Student
	var tp = new(Student)
	var w Writer 
	// ❌
	w = t
	// ✅
	w = tp

}
```
我们可以看到，get和set分别是一个值类型和一个指针类型上实现的，这里我们的结论是：当实现接口时，**类型的指针变量**在实现方法上可以包括***定义在类型指针上的方法以及定义在值类型上的方法***，但是**值类型变量**只包含定义在**值类型**上的方法

这里是提示信息：

```bash
cannot use t (variable of type Student) as Writer value in assignment: Student does not implement Writer (method set has pointer receiver)
```
通常来说，这不应该成为程序员的烦恼，所以想用谁就好好的定义在谁上面的方法即可，完全不会出错。

## 接口嵌入
### 在接口中嵌入接口类型
```go
type Writer interface {
	Write([]byte) (int, error)
	error()
}
type Error1 interface {
	error()
}
type WriterError interface {
	Writer
	Error1
}
```
这样的组合就可以组合成一个新的接口，并且嵌入的接口还可以有方法上的交集，go是不介意的（go1.14+）。

### 在结构体中嵌入接口类型

## 接口类型的底层

## 空接口的使用

## 接口的相等性

## 小接口的意义

## 接口提供程序的可扩展性

## 接口提供程序的可测试性

## 接口的严格和函数的宽松对比

接口的实现是严格的：在实现接口的时候函数需要显示转换

```go
func main() {
  http.ListenAndServer(":8080", http.HandlerFunc(hi))
}

func hi(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hi")
}

```
这是正确的用法，不能因为 hi 跟 http.HandlerFunc 底层一样，就认为它俩相等，因为http.HandlerFunc实现了接口，并不代表 hi实现了接口。

实际上是不等于的关系，需要显式的转换一下。

接下来让我们看一下刚才代码的底层原理实现：

```go
type b interface{
  add(int,int)int
}

type b1 func(int,int)int

func(b b1)add(x,y int)int{
  return b(x,y)
}

func t (x,y int)int{
  return x+y
}

func main() {
  var bb b = b1(t)
  bb.add(1,2)
}
```


然而函数的使用是宽松的。当直接使用函数，以及return 函数的 的时候，引用类型(slice, map, func, interface, chan)是不需要显式转换的，只有非引用类型比如int，bool string strcuct 这种需要。

```go
// 不需要显示的转换

// 函数类型 return 
type b func(string) int

func get2() b {
	return func(s string) int {
    println(s)
		return len(s)
	}
}

// 函数类型 直接使用

func main() {
	get(func(int)string{
		return "hello"
	})
}

type N func(int)string

func get(n N){}

// 当然你如果显式的转换一下也是没有问题的
func get3() b {
	return b(func(s string) int {
		return len(s)
	})
}
```

interface 看似也可以解释这个问题但是本质上是不同的，interface是实现，并不是底层 type a int 这种类型,实际上 动态类型实现接口以后，就等于接口类型那个类型了。
```go
package main

func main() {
	var h hi = NewHi()
	h.get()
}

type hi interface {
	get()
}
type hey struct{}

func (hey) get() {
	println("hi")
}

func NewHi() hi {
	return hey{}
}


```
不过除了函数等引用类型之外的非引用类型还是必须显示的转换的

```go
// 整数类型
type a int

func get1() int{
	var a1 a
	a1 = 12
	return int(a1)
}

// struct 也需要显示的转换
type b struct {
  i int
}

type b1 b

func get7()b1{
  return b1{i:1}
}

// 或者是

func get8() b1 {
  return b1(b{i:1})
}

```

## issues
`问题一：` ***interface如何判断 nil***

只有类型是 nil + value 是nil的接口类型才是nil，否则它不等于nil
```go
package main

import (
	"fmt"
)

func main() {

	var a1 a
	// true
	fmt.Println(a1 == nil)
	var b1 *b
	a1 = b1
	// false
	fmt.Println(a1 == nil)

}

type a interface {
	get()
}

type b struct{}

func (*b) get() {}
```

`问题二：` ***如何判断 go 接口类型是否相等***


`问题三：` ***eface 和 iface的区别*** 


`问题四：` ***如何查找interface中的方法***


`问题五：` ***interface 设计的优缺点***

## 参考资料
- https://book.douban.com/subject/35720728/ 246页 - 286页 
- https://mp.weixin.qq.com/s/6_ygmyd64LP7rlkrOh-kRQ
- https://github.com/golang/go/blob/master/src/runtime/runtime2.go
- https://github.com/golang/go/blob/master/src/runtime/type.go