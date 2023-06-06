<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-06-07 03:55:09
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
- 如何判断接口类型的相等
- 论述“nil error != nil”的原因
- 小接口的意义
- 不可滥用空接口
- 接口作为程序水平组合的连接点，提供程序的可扩展性
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
这是接口的底层数据：

一般接口：
```go
// interface
type iface struct {
	tab *itab
	data unsafe.Pointer
}
```
空接口：
```go
// empty interface
type eface struct {
	_type *_type
	data unsafe.Pointer
}
```
data 表示的意思一样，值是动态类型的地址。我们比较value时，比较的是地址指向的数据是否相同而不是地址本身。

空接口并没有定义接口的方法，因此`_type` 定义的均为动态类型的元数据
```go
type _type struct {
	size uintptr
	ptrdata uintptr
	hash uint32
	tflag tflag
	align uint8
	fieldalign uint8
	kind uint8
	alg *typeAlg
	gcdata *byte
	str nameOff
	ptrToThis typeOff

}
```
一般接口因为本身定义了方法，因此它需要定义自己的方法，以及动态类型的数据，因此它除了 `_type` 外，还定义了 `interfacetype` 用来存储自己定义的方法元数据。 
```go
type itab struct {
	// 非空接口本身的信息
	inter *interfacetype
	// 动态类型数据
	_type *_type
	// _type.hash的copy，用于 switch 判断类型
	hash  uint32
	_     [4]byte
	// 动态类型已实现接口方法的调用地址数组
	fun   [1]uintptr
}
```
注意这里的fun数组，这里定义的 `[1]uintptr` 在实际使用时，可能不是[1],这里的数据时可变的，如果是2，就表示实现了两个方法。

原文的注释是这样的 `// variable sized. fun[0]==0 means _type does not implement inter.`
```go
type interfacetype struct {
	// 接口本身的类型信息
	typ _type
	// 接口所在的包路径
	pkgpath name
	// 接口方法集合
	mhdr []imethod
}

```

## 如何判断接口类型的相等
当接口类型未被赋予动态类型时，它的两个字段，即：动态类型字段和动态类型value字段均为nil，那么这个未初始化的接口变量就恒等于`nil`

当接口类型被赋予了动态类型，那么如果判断这时候的接口类型，必须为类型相同以及值相同，接下来我们看一个案例：

***两个非空非nil接口变量比较：***
```go
func main() {
    printNonEmptyInterface1()
}

type T struct {
    name string
}
func (t T) Error() string {
    return "bad error"
}
func printNonEmptyInterface1() {
    var err1 error    // 非空接口类型
    var err1ptr error // 非空接口类型
    var err2 error    // 非空接口类型
    var err2ptr error // 非空接口类型

    err1 = T{"eden"}
    err1ptr = &T{"eden"}

    err2 = T{"eden"}
    err2ptr = &T{"eden"}
    println("err1:", err1)
    println("err2:", err2)
    println("err1 = err2:", err1 == err2)             // true
    println("err1ptr:", err1ptr)
    println("err2ptr:", err2ptr)
    println("err1ptr = err2ptr:", err1ptr == err2ptr) // false
}
```
```go
err1: (0x104c959a8,0x1400004c748)
err2: (0x104c959a8,0x1400004c728)
err1 = err2: true
err1ptr: (0x104c95988,0x1400004c738)
err2ptr: (0x104c95988,0x1400004c758)
err1ptr = err2ptr: false
```

> println ，预定义函数，在编译期间，会由编译器根据要输出的参数的类型，将println替换为特定的函数，这些预定义函数定义在 [runtime/print.go](https://github.com/golang/go/blob/96d16803c2aae5407e99c2a1db79bb51d9e1c8da/src/runtime/print.go#L255) 中，针对 eface和iface的打印函数是：
	```go
	func printeface(e eface){
		print(e._type, e.data)
	}
	```
	```go
	func printiface(i iface){
		print(i.tab, i.data)
	}
	```


如代码所示，我们要判断的是，非空接口，并且已经实现了动态类型的两组接口类型，答案已经写在代码里了，即：`err1 == err2` `err1ptr != err2ptr`

现在就让我们从源码出发来探究一下原因。

首先，我们知道**类型相同以及值相同才是真的相等**，err1 和 err2 的动态类型均为 `T` ，值也均为 `T{"eden"}` ，所以他们相等；err1ptr ，和err2ptr 的类型均为 `*T`，值均为`&T{"eden"}`，但是系统却判断他们不相等，从源码来看，二者的`tab *itab`，因为动态类型的元数据相同，这个字段一致，所以类型一致，从第二个字段 data 来说，***data存储了`&T{"edent"}`的地址，这个地址指向的内容仍为地址***，从内容上来说地址指向的地址值并不相同，所以这就可以解释为什么结果是false了。

***接下来我们看一下两个空非nil接口类型的比较：***

```go
func main() {
	var a any
	var b any
	a = &S{
		"1",
	}
	b = &S{
		"1",
	}
	//(0x102f8aaa0,0x1400004c758)
  //(0x102f8aaa0,0x1400004c748)
	println(a)
	println(b)
	// false
	println(a == b)
}

type S struct {
	name string
}
```
根据源码所知，a和b的 `_type` 是完全相同的，然而地址指向的地址不相同，所以结果是false，下面让我们稍微改动一下：
```go
func main() {
	var a any
	var b any
	a = &S{
		
	}
	b = &S{
		
	}
// (0x104422aa0,0x1400004c767)
// (0x104422aa0,0x1400004c767)
	println(a)
	println(b)
	// true
	println(a == b)
}

type S struct {
	
}
```
这个时候你惊奇的发现，结果竟然是true，这是为什么呢？不是说，地址指向的地址应该是不同的吗？nonono，并不是所有的情况都是那样，如果结果是空接口，那么空接口的所有变量指向的都是同一个地址，所以从结果上来说，data其实地址是相同的，指向的是同一个数据，所以答案是true。
>在 Go 中，空数据结构（比如 struct{}）不占用任何内存空间，因此在创建空数据结构时，它们实际上是指向同一个地址的。这是因为在 Go 中，每个变量都需要分配内存空间，以便可以存储它们的值。但是，由于空结构体没有任何字段，因此它们不需要分配任何内存空间。因此，在创建空结构体时，它们实际上是指向同一个已经分配的零大小内存块的指针。

***一个非空接口类型和一个空接口类型一定不相等吗？***

如果你根据源码来看，第一个字段本身就不一样，肯定不相等了，但是go在比较相等时，比较的是`_type`字段，并不是全部的tab数据，所以当两者 字段中的`_type`相同就表示类型相同:
```go
func main() {
	var a any = S{6}
	var b B  = S{6}
	//(0x1040dfae0,0x1400004c760) 
	//(0x1040e5a68,0x1400004c758)
  //true
	println(a,b)
	println(a == b)
}
type S struct{
	int
}
type B interface{
	get()
}
func(S)get(){}
```
所以从结果来看，_type 字段相同均为 S，data也是一致的，所以答案是 true

***nil接口类型：***
```go
func main() {
	var e error
	var a any
	//(0x0,0x0)
	println(a)
	//(0x0,0x0)
	println(e)
	println(e == nil)
	println(a == nil)
	println(a == e)
}
```
当一个接口是未给定动态类型的接口类型，它就是nil接口，那么它的类型和data值均为空，所以只要是nil接口，他们均相等，并且等于nil。

最后说明一下，***当接口类型获取动态类型的时候，绝大多数情况下，会将动态类型的值复制，并且放置在一个新的内存空间里，所以原始数据跟接口类型的数据再无瓜葛，指针类型除外***，不过为了节省空间，有一种情况，go编译器就会放弃这个动作，并不会每次都重新分配。

```go
func main() {
	var x any = 34
	var y any = x
	var z any = x
// (0x1023343e0,0x10232c1b0)
// (0x1023343e0,0x10232c1b0)
// (0x1023343e0,0x10232c1b0)
	println(x)
	println(y)
	println(z)
}
```
可以看到，go判断，x y z 三个空接口类型的动态类型，类型均相同都是int，并且 data 指向同一块内存地址。

非空接口也是一样：
```go
func main() {
	var a1 a = b{}
	var a2 a = a1
// 	(0x102791a48,0x1400004c758)
//  (0x102791a48,0x1400004c758)
	println(a1)
	println(a2)
}

type a interface {
	get()
}
type b struct {
	name string
}

func (b) get() {}
```
我们如果想获取关于接口的内部实现细节，可以看一下这个[项目](https://github.com/bigwhite/GoProgrammingFromBeginnerToMaster/blob/main/chapter5/sources/dumpinterface.go)，可以输出内部的信息
## 小接口的意义
***接口越小，抽象程度越高，使用范围也就越大***

一群飞禽走兽，我们可以给他们的行为抽象为“飞行”，一群能游泳的动物我们可以给他们的行为抽象为“游泳”。那么飞行和游泳涵盖的内容就会非常的多，使用范围就会很大，比如我们现在有一个函数，要求对所有能飞行的动物做出打印动作，那么众多飞行的动物就都可以使用这个函数。
```go
func main() {
	e1 := e{"大鹅"}
	y1 := y{"老鹰"}
	// {大鹅} 他们的具体行为模式是： 大鹅慢慢的飞
	PrintFlyer(e1)
	// {老鹰} 他们的具体行为模式是： 老鹰迅速飞行
	PrintFlyer(y1)
}

type Flyer interface {
	Fly() (flyMod string)
}

func PrintFlyer(f Flyer) {
	fmt.Println(f, "他们的具体行为模式是：", f.Fly())
}

// 大鹅
type e struct {
	name string
}

func (e) Fly() string {
	return "大鹅慢慢的飞"
}

// 老鹰
type y struct {
	name string
}

func (y) Fly() string {
	return "老鹰迅速飞行"
}
```

***易于实现和测试***

当接口的方法较少时，动态类型实现的方法就少了，必然容易实现以及容易测试。

***高內聚，易于复合组合***

我们抽象程度很高的接口，接口做的事情就很单一，比如飞行类的接口方法就是飞行，游泳动物的接口方法就只有游泳，当有会游泳也会飞行的动物时，我们只需要成立一个新的接口，将飞机类和游泳类的接口嵌入到新接口中就形成了一个全新的会飞行也会游泳的接口了。

如果一个接口涵盖了各种方法，那么当组合接口的时候，势必某些方法是被弃用的，所以综上所述，设置单一的，高内聚的方法是好的设计方案。

### 如何设计小接口
1. 先初步抽象出接口，这个时候可以有耦合，也可以不够高抽象，但是你得先定义出一个初步的接口出来，与此同时我们也得清楚，越是业务代码，抽象出一个高内聚的接口越难。

2. 将大接口拆分为小接口，使用一段时间以后，我们会发现某些操作是可以单出被提取出来的，比如 io 包的 writer 和 reader，那么我们就可以把这个动作单独抽象出来。抽象的最高程度就是只有一个方法，这就非常的內聚了，可以说，这种程度的抽象在日常业务中还是相对比较难的，需要在长时间的使用中，慢慢摸索。

综上所述：现搞出一个能用的大接口再说，以后慢慢解耦，形成抽象程度更高的小接口。

## 不可滥用空接口


## 接口作为程序水平组合的连接点，提供程序的可扩展性

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


函数的使用是宽松的。当直接使用函数，以及return 函数的 的时候，（其它引用类型也一样：slice, map, func, interface, chan）是不需要显式转换的，只有非引用类型比如int，bool string strcuct ... 需要。

```go
// 不需要显示的转换
// 或者也可以说：系统自动把这个匿名函数推导为了b类型。

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
	var a = func(int)string{
		return "hello"
	} 
	get(a)
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

`问题二：` ***论述 “ nil error != nil ” 的原因***

nil error通常可以用这种方法来输出:
```go
func main() {
	var a error = (*b)(nil)
	//(0x102205988,0x0)
	println(a)
	//false
	println(a == nil)
}

type b struct {
	error
}

```
可以发现，nil error 的类型并不是0x0，而 nil 接口变量是 0X0,0x0，所以这两者并不相同。
`问题三：` ***eface 和 iface的区别*** 

eface和iface的第二个字段相同均存储的是动态类型的地址，然而eface的第一个字段保存的是动态类型的元数据，即：_type 字段，而iface的第一个字段不仅仅保存了动态类型的元数据 _type ，还保存了自己的方法集合的相关数据，以及动态类型实现的方法地址等数据。

`问题四：` ***如何查找interface中的方法***



`问题五：` ***interface 设计的优缺点***



## 参考资料
- https://book.douban.com/subject/35720728/ 246页 - 286页 
- https://mp.weixin.qq.com/s/6_ygmyd64LP7rlkrOh-kRQ
- https://github.com/golang/go/blob/master/src/runtime/runtime2.go
- https://github.com/golang/go/blob/master/src/runtime/type.go