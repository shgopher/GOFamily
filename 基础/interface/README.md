<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-04-26 21:47:31
 * @FilePath: /GOFamily/基础/interface/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 接口
> 有关泛型相关的约束内容均放置在[泛型](../泛型)这一章节，这里不再赘述

重点内容前情提要
- 实现接口
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
## 接口嵌入
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


然而函数的使用是宽松的。当直接使用函数，以及return 函数的 的时候，引用类型(slice map func,interface,chan)是不需要显式转换的，只有非引用类型比如int，bool string strcuct 这种需要。

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
***interface如何判断nil***

***eface 和 iface的区别***

***如何查找interface中的方法***

***interface 设计的优缺点***
## 参考资料
- https://book.douban.com/subject/35720728/ 246页 - 286页 
- https://mp.weixin.qq.com/s/6_ygmyd64LP7rlkrOh-kRQ
- https://github.com/golang/go/blob/master/src/runtime/runtime2.go
- https://github.com/golang/go/blob/master/src/runtime/type.go