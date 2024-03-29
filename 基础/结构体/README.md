<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-03-31 19:05:02
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-01-13 17:49:43
 * @FilePath: /GOFamily/基础/结构体/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 结构体
## 简单介绍 struct
```go
type People struct {
  Addr string
  name string
  year int
}
```
- Peopel 首字母大写，根据我们所学的首字母大写可导出的知识，它是包级可导出结构
- Addr 首字母大写，所以它是可导出字段，name 和 year 都是小写，所以他们俩不可导出

下面看一下结构体的使用
```go
// 初始化
func main(){
  var p People
  p.Addr = "北京"
  p.name = "小明"
  p.year = 2000

  var p1 = People{
    Addr: "北京",
    name: "小明",
    year: 2000,
  }

  var p2 = &People{
    Addr: "北京",
    name: "小明",
    year: 2000,
  }
}
type People struct {
  Addr string
  name string
  year int
}

```

```go
// 结构体的调用

func main(){
  var p  = &People{
    Addr: "北京",
    name: "小明",
    year: 2000,
  }
  // 这里是语法糖，p虽然是people的指针，
  // 但是它却可以直接调用 Addr
  // 实际上就是 (*p).Addr 的省略
  p.Addr = "上海"
  p.name = "小红"
  p.year = 2001
  fmt.Println(p.Addr, p.name, p.year)

  // 我们还可以使用new来代替&
  // 这种写法跟上文中的 var p = &People{} 
  // 一个意思
  var p1 = new(People)
  fmt.Println(p1)
}
type People struct {
  Addr string
  name string
  year int
}
```
## 解耦结构体声明和调用
当我们实现结构体的时候，如果显示写出结构体的字段变量名称，就可以不按照顺序，以及可以不完全实现全部字段，这样的话，结构体的声明和实现就可以完全解耦，当然可以隐藏实现的结构体变量，那么你不得不要按照顺序，以及实现全部字段，满足这两者才可以。

```go
// 显式实现
type People struct {
  Addr string
  name string
  year func(int)int
}
func main(){
  var p = People{
    Addr: "北京",
    year: func(a int)int{
      return 2000 + a
    }
  }
 
}
```
上述代码就是显示的写出了字段的变量名称，你看，name 并没有被写上，这种情况下，name 就会被命名为一个初始值，即 “”

这样，即使结构体本身有什么增加字段的行为，实现结构体的逻辑代码也不用改变了。

如果是隐式的话，那么必须按照顺序，以及数量进行实现，建议在字段不变以及字段数量非常少的时候使用。

## 匿名 struct

我们使用匿名 struct，可以完全将另一个结构体嵌入到这个结构体中。

```go
type Student struct{
  People
  score int
}
type People struct {
  name string
  year string
  addr string
}
```

使用匿名函数，等于将这个匿名函数的字段完全给予了这个新的结构体，这跟使用这个结构体当作一个字段的类型是不同的，我们看一个例子
```go
type Student struct{
  People People
  score int
}
type People struct {
  name string
  year string
  addr string
}
```
这段代码跟上文看起来很像，但是一个是将 people 直接嵌入，一个是当成了它的一个字段类型。让我们看一下这两者使用起来的区别

```go
type Student1 struct{
  People People
  score int
}
type Student struct{
  People
  score int
}
type People struct {
  name string
  year string
  addr string
}

func main(){
  var s = new(Student)
  var s1 = new(Student1)
  s.name = "小明"
  s.year = "2000"
  s.addr = "北京"
  s.score = 100
  fmt.Println(s.name, s.year, s.addr, s.score)
  s1.people.name = "小红"
  s1.people.year = "2001"
  s1.people.addr = "上海"
  s1.score = 200
  fmt.Println(s1.people.name, s1.people.year, s1.people.addr, s1.score)
}
```
我们发现匿名函数的字段直接赋予了新的结构体，它不再需要显式的调用 `s1.People.name`

实际上，不止一个 struct 可以当作匿名函数，内置类型也是可以的

```go
type Student struct{
  string
  People
  int
  score int
}
type People struct {
  name string
  year string
  addr string
}

func main(){
  var s = new(Student)
  s.People = People{
    name: "小明",
    year: "2000",
    addr: "北京",
  }
  s.int = 12
  s.string = "12"
  s.score = 100
  fmt.Println(s)
}
```
函数也可以将指针类型当作类型以及直接嵌入，与非指针相比，我们不能直接调用嵌入的结构体的字段，因为它目前还是 nil，这个时候我们必须先将结构体初始化，然后再进行操作，下文的代码有演示。


```go
type Student struct{
  Name *Name
  *People
  int
  string
  score int
}
type Name string{
  fist string

  last string
}
type People struct {
  name string
  year string
  addr string
}

func main(){
  var s = new(Student)

  s.Name = new(Name)
  s.People = new(People)

  s.Name.first = "小明"
  s.Name.last = "小红"
  s.name= "小明"
  s.year = "2000"
  s.addr = "北京"
  s.int = 12
  s.string = "12"
  s.score = 100

  fmt.Println(s)
}

```

通过这个演示我们发现，其实嵌入更像是一种语法糖，它就跟将结构体当作类型，前面的结构体字段变量跟这个类型保持一致这种操作是同样的作用，只不过直接嵌入可以将嵌入的字段当作自己的字段，省略了中间的变量罢了，它有下面的等于关系

```go
type Student struct{
  //People *People
  // 等于 
  // *People


  // People People
  // 等于 
  // People
  

}
type People struct {
  name string
  year string
  addr string
}
```
当一个结构体拥有了其它嵌入的时，也一起拥有了在他们身上的方法，不过，非嵌入的那种，还是需要带上中间的字段变量名称比如：

```go
package main

type Student struct {
	Name Name
	*People
	int
	string
	score int
}

type Name struct {
	first string
	last  string
}

type People struct {
	name string
	year string
	addr string
}

func (s *Name)S() {
	println("stuend-s")
}

// 如果这里不是定义在指针上的方法，这段代码将报错
// 原因是 people嵌入的时候还是nil，如果这里不是指针上的方法，是值上的方法
// 底层是需要 *People （取值）  操作的，但是这时候是nil，所以就会报错。
func (s *People)SetS() {
	println("people-setS")
}

func main() {
	var s = new(Student)
  // 这里注意，因为不是嵌入的操作，所以它必须带上中间的 Name
	s.Name.S()
	s.SetS()
}

```
## 空结构体

通常，当我们需要一个临时的变量时，我们可以会想到设置一个 bool 类型，因为我们潜意识中感觉一个 bool 类型是比较小的，但是一个空的变量才是最小的，下面让我们看一个例子，这个例子发生在使用 channel 传递信息这个场景。

```go
func main() {
  sig := make(chan bool)
  go func(){
    time.Sleep(time.Second)
    sig <- true
  }()
  <- sig
  fmt.Println("任务已经完成")
}
```
没错，这段代码中，使用一个 bool 类型的 channel 是没什么问题的，你甚至也可以使用 int string 都可以，因为只是传递一个信息，信息的内容不重要，但是当我们站在优化的角度来考虑，这里的 bool 就不完美了，我们改成空的结构体即可：
```go
func main() {
  sig := make(chan struct{})
  go func(){
    time.Sleep(time.Second)
    sig <- struct{}{}
  }()
  <- sig
  fmt.Println("任务已经完成")
}
```
需要注意的是，一个空的结构体，表示它类型的方式是 `struct{}`，而使用这个空结构体的方式就是 `struct{}{}`，前面的大括号是跟 struct 一起的整体表示空结构体，后面的大括号表示一个空结构体类型的结构体调用

## 直接嵌套还是作为字段

```go
type Pool struct {
  wg sync.WaitGroup
  JobQueue chan Job
  dispatcher *dispatcher
}
// or
type Pool struct {
  sync.WaitGroup
  
  JobQueue chan Job
  dispatcher *dispatcher
}

```

结构体嵌套可能带来的问题：

- 名称冲突
如果 Pool 结构体中还定义了 Add/Done/Wait 方法，和嵌套的 WaitGroup 中的方法就会产生冲突。

- 不必要的方法
嵌套整个 WaitGroup 会让 Pool 结构体拥有 Add/Done/Wait 等方法，但 Pool 可能只需要 Wait 就够了。

- 使结构体臃肿
嵌套整个 WaitGroup 会让 Pool 结构体看起来很臃肿，包含许多其实用不到的方法。

- 继承关系不清晰
Pool 并不是 WaitGroup 的一种，嵌套整个 WaitGroup 让人可能误以为它是在继承 WaitGroup。

所以，作为一个字段，可以避免上述问题，又可以保持必要的功能。

总结一下，嵌套要慎用，只有当：

- 被嵌套的类型方法不会与现有方法冲突
- 被嵌套的所有方法都会被用到
- 两者之间有逻辑上的继承关系

时，才更适合使用嵌套 (组合) 的方式。

否则，使用字段的方式可以获得必要的功能，而不引入嵌套的潜在问题。