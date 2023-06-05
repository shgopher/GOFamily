<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-06-05 09:21:14
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
- 论述“nil error != nil”的原因
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

接口具有静态语言和动态语言的共同特点。

静态类型特点：
- 接口可以作为类型：`var e error`
- 支持在编译期进行类型检查：在编译器期间会对右边的变量进行类型的检查，检查是否实现了接口定义的方法类型中的所有方法

动态类型的特点：
- 在运行时可以对接口进行动态赋值
- 接口类型可以在运行时被赋予不同的动态类型变量，从而进行 “多态”

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
在结构体中嵌入一个接口，就相当于实现了这个接口，拥有了它的方法，但是注意，拥有的是这种抽象方法，那么我们为什么要将一个接口类型嵌入到一个结构体中呢？

因为go语言规定，嵌入接口，结构体相当于实现了这些接口，这里注意，这里必须是直接嵌入，如果是接口类型作为变量类型的方式是不能拥有接口的方法的。
```go
type A interface{
	get()
	set()
}

// ✅
type A1 struct{
	A
}
// ❌
type A2 struct{
	A A
}
// ❌
type A3 struct{
	a A
}
```

***当结构体本身也实现了方法时，优先调用结构体的方法。***

这个场景是这样的，在某个函数中，它的参数是一个接口类型，并且这个函数调用的只是这个接口类型的某个，或几个方法，并不是全部，那么我们作为结构体，想实现这个接口，又不想多实现额外的方法，那么这种方法就很好用了。

```go

package main

import "fmt"

func main() {
	var a A
	D(a)
}

type A struct {
	BI
}

type BI interface {
	get()
	set()
}

func (A) get() {
	fmt.Println("hi")
}

func D(b BI) {
	b.get()

}

```

注意上文提到了，接口中内嵌接口的时候，内嵌的接口方法可以重复，但是结构体中内嵌的接口，不允许出现方法重复的问题：

```go
// ❌ ： ambiguous selector a.get
type A struct{
	BI
	BI1
}
type BI interface{
	get()
	set()
}
type BI1 interface{
	get()
	err()
}
```
不过，要想解决这个问题，我们只需要让结构体实现这种重复的方法即可，这样，优先级就提升到了结构体，接口的方法就不会被调用了。比如：

```go
package main

import "fmt"

func main() {
	var a A
	D(a)
	D1(a)
}

type A struct {
	BI
	BI1
}
type BI interface {
	get()
	set()
}
type BI1 interface {
	get()
	err()
}

func (A) get() {
	fmt.Println("hi")
}

func D(b BI) {
	b.get()

}

func D1(b BI1) {
	b.get()
}

```
下面让我们看一下，这种用法在单元测试场景中的应用

让我们描述一个场景：

有一个函数，它接受一个接口类型作为参数，我们要对它进行单元测试，而且我们要伪造一些数据。

```go
// 函数体
func MaleCount(s stmt)(int,error){
	result,err := s.Exec("SELECT count(*) FROM exployee_tab WHERE gender=?","1")
	if err != nil {
		return 0,err
	} 
	return result.Int(),nil
}
// 抽象接口
type stmt interface {
	Close error
	NumInput()int
	Exec(stmt string,args ...string)(Result,error)
	Query(args []string)(Rows,error)
}
// 接口相关的一些数据
type Result struct{
	Count int
}
func(r Result) Int()int{return r.Count}

type Rows []struct{}
```

我们可以看到，要想对这个 MaleCount 函数进行处理，那么一个实现了stmt接口的动态类型必不可少，但是我们并不需要所有的方法，仅仅需要Exec方法。

所以我们第一步就是设置一个fake类型，并且将接口内嵌来完成“继承”。
```go
type fakeStmtForMaleCount struct{
	stmt
}
// 这里实际上只是简写，
//真正的测试要对smt和arg进行测试的
func(f fakeStmtForMaleCount)Exec(stmt string,args ...string)(Result,error){
	return Result{1},nil
}
```
当我们内嵌完成继承之后，我们相当于拥有了这些抽象方法，然后我们在这个接口体上自行实现Exec，这样就可以将结构体的Exec优先级提前。

那么让我们开始使用虚假数据 `Result{1}` 开始测试 MaleCount 函数

```go
func TestEmployeeMaleCount(t *testing.T) {
	fs := fakeStmtForMaleCount{}
	v,err := MaleCount(fs)
	if err != nil {
		t.Error("error is :",err)
	}
	if v != 1  {
		t.Errorf("we want %d, actual is %d",1,v)
	}
}
```
## 接口类型的底层

## 论述“nil error != nil”的原因
## 空接口的使用
## 接口的相等性
## 小接口的意义
## 接口提供程序的可扩展性
## 接口提供程序的可测试性
## 接口的严格和函数的宽松对比
接口的实现是严格的：在实现接口的时候函数需要显示转换

```go
func main() {
  http.ListenAndServe(":8080", http.HandlerFunc(hi))
}

func hi(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hi")
}

```
这是正确的用法，不能因为 hi 跟 http.HandlerFunc 底层一样，就认为它俩相等，因为http.HandlerFunc实现了接口，并不代表 hi实现了接口。

实际上是不等于的关系，需要显式的转换一下。

接下来让我们看一下刚才代码原理：

```go
type b interface{
  add(int,int)int
}

type a func(int,int)int

func(a1 a)add(x,y int)int{
  return a1(x,y)
}

func t (x,y int)int{
  return x+y
}

func main() {
  var bb b = a(t)
  bb.add(1,2)
}
```


然而函数的使用是宽松的。当直接使用函数，以及return 函数的 的时候，引用类型(slice, map, func, interface, chan)是不需要显式转换的，只有非引用类型比如int，bool string strcuct 这种需要。

```go
// 不需要显示的转换
// 或者也可以说，系统自动把这个匿名函数推导为了b类型。
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
// 或者说，你人为的帮它推导出了类型是b类型
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