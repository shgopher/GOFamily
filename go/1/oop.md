# go语言中的面向对象

其实go语言并没有完整的面向对象，go语言可以说是面向接口编程的一个语言，go里面最重要的一个组件就是接口，接口让go语言变得非常灵活。
## 函数
go语言中的多返回值，使用的方式是栈，这不同于c的寄存器模式，优点是结构简单，可以从右向左压入栈中，计算的时候从栈中取出就会从左往右的排列了，而且可以天生支持多重返回，缺点就是速度比从寄存器模式慢了很多。
##  go里面的继承，使用的是结构体的内嵌

```go
type People struct {
    face int

}
type Student struct {
    People
    year int
}
// 当然也可以选择不内嵌，将结构体当成一个类型也行

type Student struct {
    People People
    year int
}

```

在struct中也可以将接口当作一个类型加进去

```go
type Talk interface {
    Say()
}
type People struct {
     Talk
}
```

看到这里你应该发现了把，在go里面变量的首字母的大小写是有不同的含义的，如果首字母是小写的，那么这个变量的使用权就是这个包的内部，如果是大写那么这个变量就可以被其它包使用。所以它的界限是以出包不出包为界线的。

接下来我们谈一个非常重要的东西 --- 方法

## 方法

```go
type Student struct {
    year int
}
func(s *Student)Show(){
    fmt.Println(s.year)
}
```

这里我们讲解一下知识点，首先是方法的使用方式，如图所示，方法使用就是`func(变量 对象)方法名称(){}` ,这里注意一下，并不是只有structure才能拥有方法，任何只要是以下方法的对象都可以拥有

```go
type Student int
func(s *Student)get(){

}
type People map[string]string
```

只要是这种形式的都可以拥有方法，哦对了接口不行。因为接口没人任何的实际意义，它也就不可能拥有方法。

这里还有一个知识点，就是go语言中的指针，

```go
// 这是指针类型的声明，我声明了，a是int类型，并且是一个int类型的指针类型。也就是说a装的不是int，是某个int的地址。
var a *int

// 赋值, 这里注意哈，new后面只能是类型 type ，所以这里可以使用new(int)因为int是一个类型。
a = new(int)
// 这样也是OK的。这个时候类型就是b了，b和int不是一个类型。
var a *b

type b int
a = new(b)
 
```
这里稍微谈一下  类型和底层类型的使用。

```go
type b int
func fast()b{
    return 1
}
```

可以非常明显的看出，b类型的底层是int，但是b不是int，int只能是b的底层数据，就跟rune和int32，byte和uint8一样。
但是这个时候重点来了，这里有个返回值，返回的是b，那么具体要返回什么值才能满足b呢？也就是返回它的具体值即可。

更多例子：

```go

type b map[string]string

func fast()b {
    return map[string]string{}
}
```
 这里还要说一下，在go里面，值类型上的方法和指针类型上的方法是可以直接调用的。

 举个例子：

 ```go
 type a sturct {

 }

 func(a1 a)get(){

 }
// 这样百分之一百OK
 func main(){
     var aa a
     aa.get()

 }

 // 这样也是OK的

 func main(){
     var aa  = new(a)
     aa.get()
 }

 func(a1 *a)set(){}

 // 这样一定OK

 func main(){
     var aa = new(a)
     aa.set()
 }

 // 但是这样也OK

 func main(){

    var aa a
    aa.set()     
 }
 ```

## 接口

接口的本质实际上是代码的解耦，接口的使用，跟接口的实现在两个部分。

go实现了隐式接口，只要实现了接口的方法就等于实现了这个接口，为此，有些特立独行的接口为了防止被实现就内置包内方法，这样外部包就无法访问了。

任何变量都实现了空接口 `interface{}` 所以说，只要是空接口的地方，所有的变量类型都可以放进去。当然了，内部需要断言一下。


### 接口的实现
上面说到了定义在值和指针类型上的方法可以互换，或者说可以语法糖式的调用，但是在接口的实现上不能打马虎眼。一定不行就是不行；
```go
type Tallk interface {
    say()
}
type Student struct {

}
func(s *Student)say(){

}

func main(){
    s := new(Student)
    // 这里必须是指针类型才可以，值类型就不行， var s Student，就是错的。
    // 但是，如果是值类型实现的方法，那么指针类型的初始化是可以正常使用的，也就是说指针类型拥有比值类型更高的等级。
    Td(s)
}
func Td(t talk) {

}
```
### 接口的继承

```go
func main() {
	var c b
	c.set()
	c.get()
}

type a interface {
	get()
}

type b interface {
	a
	set()
}
```