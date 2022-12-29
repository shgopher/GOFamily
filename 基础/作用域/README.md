# 作用域

先给出两个经典的案例：

if级变量
```go
func main() {
  if a :=0;false{
  
  } else if b :=0;false{

  }else if c :=0;false{

  }else{
    println(a,b,c)
  }
}
```
a b c 这三个变量的作用域属于 if 这整个结构，只要前面初始化过，那么后面就可以使用，所以你会发现print可以输出三个变量。

loop级变量
```go
func main() {
  for i:=0;i<10;i++ {
    go func(){
      println(i)
    }()
  }
  time.Sleep(time.Second)
}
```
因为 i属于loop级别，通常来说for执行过程是快于新开辟一个goroutine的，所以导致这是个goroutine输出的都是最后一个i，即都输出10

根据go 1.20的表述，以后这个loop级别的变量非常有可能会被修改成局部变量。如果没有被修改，我们可以开辟一个局部变量

```go
func main() {
  for i:=0;i<10;i++ {
    i := i
    go func(){
      println(i)
    }()
  }
  time.Sleep(time.Second)
}
```
## 代码块的类型
代码块是代码执行的基本单位，代码执行流总是从一个代码块流转到另一个代码块。

- 显式代码块
  - 函数体
  - for 循环体
  - if语句
- 隐式代码块
  - 宇宙代码块 ：所有的go源码都属于这个代码块
  - 包代码块：每一个包都有一个代码块，这个代码块包括了整个包的源码
  - 文件代码块：每一个文件都有一个代码块，这个代码块包括了整个文件的源码
  
预定义的标识符，比如 make new cap len 作用域是宇宙代码块

包级变量，包级常量的作用域是包代码块

go源代码中导入的包名称作用域是文件代码块

方法接收器，函数参数，返回值，对应的作用域是函数作用域

函数体内部声明的变量和常量作用域仅限于函数体内部的局部作用域
```go

func main(){
  {
    const A = 12 // A的作用域就仅限于这个大括号里面

  }

}
```
## if
if包含了一个隐式的代码块：
```go
if a:=1;a<2 {
  println(a)
}
```
其实等于：
```go
{
  a:=1
  if a<2{
    println(a)
  }
}
```

所以，在if中的大括号里，是可以输出a的值的。




## for
## switch 
## select