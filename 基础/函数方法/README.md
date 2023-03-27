# 函数和方法
重点内容前情提要

- 函数初始化的顺序
- init 函数
- 函数式编程
- defer的运用
- 变长参数
- 方法集合决定接口的实现
- 类型嵌入以及方法集合
- define类型的方法集合
- 类型别名的方法集合

go 语言中函数的基本使用方法如下：

```go
// 函数的一般用法, go 拥有多个返回值
func Get(i int)(int, error) {
  return 0, nil
}
// 返回值具有变量的返回模式
func Post(i int)(value int){
  value = 1
  return
}

// 函数作为value值赋值给一个变量变量
var Get0 = func(i int) int {
  return 0
}
```

## 函数的初始化顺序
![](./gofunc.svg)

- main 函数是所有go可执行函数的起始位置
- 在main函数运行前，需要先运行导入子包的流程
- 在导包的过程中，最先运行的是最深层的子包
- 按照 **常量，全局变量，init函数** 的顺序进行初始化的操作
- 子包初始化之后，开始返回前一层的包中进行常量，全局变量，init函数的初始化，图示流程非常清晰的画出了这个过程。
- 多个相同的包只会导入一次，并且取相同包不同版本的可满足的最小版本导入，例如 `v1.1.0 v1.2.0`，只导入`v1.2.0`，不会默认导入最新的版本。

## init 函数
在一个包中，以及在一个包的某个文件中，可以存在多个init函数，一般来说，同一个包的不同file，按照文件字符串比较的“从小到大”的顺序先被编译file中的init先执行，同一个file中的init函数按照先后顺序执行，但是go语言规范告诉我们，不要依赖init的执行顺序。

init 函数无法主动执行，它的执行是系统自动执行的，所以显式的去调用 init 函数会发生报错。

通常来说，init的目的就是系统的初始化，因为它一定会被执行，下面我们介绍一下init函数的几个用途。

***重置包级变量值。***

```go
// src/context/context.go

var closedchan = make(chan struct{})
func init(){
  close(closedchan)
}
```
我们的包级变量是一个chan，并且需要它是一个已经关闭的chan，我们使用init来确保它提前一定处于一个已关闭的状态。

***对包级变量进行初始化，保证其后续可用。***

```go
var (
  OutBox []int
  InBox  chan int
)

func init(){
  OutBox = make([]int, 10)
  InBox = make(chan int, 10)
}
```

***init 函数中的注册模式。***

这种模式在go中有两处经典的案例，其一是 `database/sql` 包的使用，其二是image包的使用。

```go
// databsae/sql 包的使用

import(
  "database/sql"
  _ "github.com/lib/pq" 
)

func main() {

  db, err := sql.Open("postgres", "")
}
```
这里使用`import _` 就是只有导包的过程，并没有任何除了init 函数之外的函数调用。那么这么做的原因是什么呢？

`工厂方法 + 导包的顺序`

下面我们看一下这两个包的源码

```go
// github.com/golnag/go/src/database/sql/sql.go
var drivers = make(map[string]driver.Driver)

func Open(driverName, dataSourceName string) (*DB, error) {
	driveri, ok := drivers[driverName]
}

func Register(name string, driver driver.Driver) {
	drivers[name] = driver
}
```
- sql 包有一个全局的map，这里存放了所有的driver，这个map的key是driver的名字，value 是 driver.Driver 的接口类型，通常来说存入的都是实现了这个接口的动态类型。
- open 函数会直接判断是否拥有含有这个key的value值
- register 函数，也就是注册函数，这个函数通常由实现了sql.driver的数据库第三方库的init函数去进行最终的注册。

```go
// github.com/lib/pq

func init() {
	sql.Register("postgres", &Driver{})
}
```
- 这里就是实现了sql.Driver 的 pq.Driver 结构体进行了名为“postgress”的注册。

下面我们进行整体的分析：

-  在pq包实现的时候，它引入了"database/sql" ，"database/sql/driver" 这两个包，它将实现的数据库实例注册到sql的map中
- 用户调用的时候，我们调用了 `sql.Open` 使用不同的name就可以使用不同的数据库，这里是工厂方法设计模式的运用

这种方法可以使pq这个包完全没有暴露到用户面前，仅需要使用sql包就可以间接的使用pq的代码，很好的做了隔绝。

下面看一下 image包的运用。

```go
package main

import(
  "image"
  _ "image/png"
  _ "image/gif"
  _ "image/jpeg"
  "os"
)

func main(){

  width,height,error := image(os.Args[1])
  if error != nil{
    println(error)
  }
  println(width, height)

}

func image(f string)(int,int,error){
  file,err := os.Open(f)
  defer file.Close()

  img, _, err := image.Decode(file)
  if err != nil {
    return 0,0,err
  }

  b := img.Bounds()
  return b.Max.X,b.Max.Y,nil
}
```

我们发现，image.Decode(f) 应该就如同上文提到的 Open 函数一样的功能，也就是工厂方法中的工厂。

下面我们看一下 image/gif 包中的init函数：

```go
//https://github.com/golang/go/blob/master/src/image/gif/reader.go

func init() {
	image.RegisterFormat("gif", "GIF8?a", Decode, DecodeConfig)
}
```

可以看到，这个gif包也是调用了image包，并且通过 init函数，将动态类型注册到了image提供的注册器中。

***init 函数中检查失败的处理方法。***

使用init去检查某些数据的正确性，相当于做最后的质检工作，一旦发生了错误，直接使用Panic，快速停掉程序。

## defer 函数
```go
func Get(){
  f,err := os.Open("test.txt")
  if err != nil{
    panic(err)
  }
  // 此处就是 defer 函数
  defer f.Close()  
}
```
defer函数有下面几个规则
- defer后面只能跟函数或者是方法
- 一个函数中的defer函数，会使用“栈”的方式运行
- defer后面的函数和方法，返回值会被直接舍弃

下面看一下defer函数的几种用法

***配合recover函数处理 panic***

在go里，如果要处理函数内的panic，有且仅有一种方法，那就是使用recover函数。而且recover函数必须运行在defer之中。那么让我们看一下具体的代码实现

```go
func Get() {
	defer func() {
		if err := recover(); err != nil {
				println("recoverd")
		}
	}()
	panic("panic")
}
```
***修改函数的具名返回值***

具名返回值，就是返回值带有变量的，比如：
```go
func hi(a,b int)(x,y int){
  defer func(){
    x ++
    y ++
  }()
  x = a+b
  y = a-b
  return  
}
// hi(1,3) 返回值为 (5,-1)
```
在分析这段代码的时候，我们要牢记一句话，在具名返回值的函数中，go的return并不会立刻return，它会等待defer执行完毕后再真的return。也就是说，它return的是 x y 的最后时刻。

那么，让我们看一下非具名的函数：

```go
func hi(a,b int)(int,int){ 
	var x,y int
  defer func(){
    x ++
    y ++
		println("defer 会执行的")
  }()
  x = a+b
  y = a-b
  return  x, y
}
// hi(1,3) 

// output:
// defer 会执行的 5 -1
// 4 -2
```
在非具名的函数中，return x，y 的时候就已经决定返回值的最终结果了，但是这个defer还是会执行，这个是真的，xy的值确实也是在defer里改变了，但是return的是之前的中间量，defer里面的xy，并不会左右之前已经设定好的x y 的复制值了。

所以输出的是 4 -2 ，而且defer里面的输出 “defer 会执行的 5 -1” 也执行了。

defer函数虽然是先入后出，在所有的正式指令执行完成以后执行，但是它的变量初始化可是顺序的：
```go
func hi(a, b int) (x, y int) {
	j := 0
	defer func(j int) {
		println(j)
	}(j)
	j++
	return
}
// output: 0
```
这里的 defer 顺序执行完变量的初始化，所以它里面的j变量完成了值的复制，为 0，那么下文的 j++ 就不会对上面的复制品起任何作用了，即便 prinln(j) 确实发生在 j++ 之后，总结一句话，defer 函数执行是最后，但是初始化是正常的先后顺序。

如果想得到 1 的答案，可以这么改：
```go
func hi(a, b int) (x, y int) {
	j := 0
	defer func() {
		println(j)
	}()
	j++
	return
}
```
这里的j**就会**受到下面j++的影响了。

还有一点要注意，defer函数的作用域也是正常的作用域，也就是说，上文hi函数演示内容，j:= 0 必须声明在defer 函数之前，虽然我们知道defer的内容最后执行，但是它也遵循正常的作用域。如果j := 0 发生在defer之后，它就会无法找到这个变量。

看了上面那么多容易搞混的defer用法，这里要说明一下，正常的使用方法是这样的：

```go
func hi(a,b int)(x,y int){
  defer func(){
    x ++
    y ++
  }()
  x = a+b
  y = a-b
  return  
}
```
也就是我们第一种使用的惯例，它的效果就是最后改变返回值，这也是这几种方式中最常用的方式。

***输出调试信息***

```go
func trace(s string)string{
  println("prepare",s)
  return s
}
func un(s string){
  println("out",s)
}
func a(){
  defer un(trace("a"))
  println("in a")
}
func b(){
  defer un(trace("b"))
  println("in b")
  a()
}
func main(){
  b()
}

// prepare b
// in b
// prepare a
// in a
// out a
// out b 
```

这是go文档提供的一个日志记录的例子，接下来我们分析一下

- 首先b的执行，b中的defer函数 un开始初始化，它的初始化就是执行trace("b"),所以最先执行的是 prepare b
- 然后开始执行 println("in b")
- 接下来执行a，跟b一样，先执行初始化的un中的trace("a")函数，然后执行 println（“in a”）
- 然后开始执行a的defer函数，所以a的out先执行，（因为它后入，先出）
- 接下来最后执行的是b中的defer，那么就是 out b

***还原变量的旧值***

```go
func init(){
  oldfile := firstfile
  defer func(){
    firstfile = oldfile
  }()
  firstfile = "README.md"
  //...
}
```


## 变长参数


## 函数式编程
函数就是一个普通的类型，它跟int，string，拥有相同的地位，所以你会发现函数式编程在go语言的代码里运用的很广泛。

比如：

```go
package main

func main() {
	 a, _ := Get("hello", func(s string) int {
		println(s)
		return len(s)
	})
	println(a)
}

func Get(s string, f func(string) int) (int, error) {
	return f(s), nil
}
```


这里还有关于函数式编程其它相关内容：

### 委托和反转控制

### map-reduce

### go generation

### 修饰器

### pipeline

### k8s visitor

## 方法集合决定接口的实现

## 类型嵌入以及方法集合

## define类型的方法集合

## 类型别名的方法集合

## 参考资料
- https://book.douban.com/subject/35720728/ 170页 - 243页
- https://coolshell.cn/?s=GO+编程模式
- https://github.com/golang/go/blob/06264b740e3bfe619f5e90359d8f0d521bd47806/src/database/sql/sql.go#L813
- https://github.com/lib/pq/blob/922c00e176fb3960d912dc2c7f67ea2cf18d27b0/conn.go#L60


