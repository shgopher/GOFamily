<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-28 01:33:50
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-07-24 22:21:40
 * @FilePath: /GOFamily/基础/作用域/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher shgopher@gmail.com, All Rights Reserved. 
-->
# 作用域

先给出两个经典的案例：

if 级变量
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
a b c 这三个变量的作用域属于 if 这整个结构，只要前面初始化过，那么后面就可以使用，所以你会发现 print 可以输出三个变量。

loop 级变量
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
因为 **i 属于 loop 级别**，通常来说 for 执行过程是快于新开辟一个 goroutine 的，所以导致这是个 goroutine 输出的都是最后一个 i，即都输出 10

根据 go 1.20 的表述，以后这个 loop 级别的变量**非常有可能**会被修改成局部变量。如果没有被修改，我们可以开辟一个局部变量 (go 1.22 将这个 loop 级变量修改为局部变量)

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
  - if 语句
- 隐式代码块
  - 宇宙代码块：所有的 go 源码都属于这个代码块
  - 包代码块：每一个包都有一个代码块，这个代码块包括了整个包的源码
  - 文件代码块：每一个文件都有一个代码块，这个代码块包括了整个文件的源码
  
预定义的标识符，比如 make new cap len 作用域是宇宙代码块

包级变量，包级常量的作用域是包代码块

go 源代码中导入的包名称作用域是文件代码块

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
if 包含了一个隐式的代码块：
```go
c := 12
if a:=1;a<2 {
  println(a)
}else if b:=3;b < c {
  println(b)
}else {
  println(a,b)
}
```
其实等于：
```go
  c:= 12
{
  a:=1
  if a<2{
    println(a)
  }
  b := 3
  else if b < c {
    println(b)
  }else {
    println(a,b)
  }
}
```

所以，在 if 中的大括号里，是可以输出 a 的值的。
## for
for 循环的作用域可以使用下面的两种方式展现。for 循环中的变量都属于 loop 级，不属于每次的那个小循环的局部变量，所以在整个的 loop 过程中，变量是唯一的，就是那一个。只不过不同的时间点，这个变量的值有可能是不同的。

```go
for i:=0;i<10;i++ {

}

// 等价于

{
  i:=0
  for ;i<10;i++ {

  }
}
```
这里需要注意一下，一旦 go 在后续的版本中修改了 loop 级变量这个设定，这个等价就不成立了。上文有讲。

```go
for k,v := range slice {

}
// 等价于
{
  k,v := 0,0
  for k,v = range slice {

  }
}
```
## switch && select
这两种结构的等价是相同的，都是 case 级。

```go
switch expression:
case list1:
case list2:
default:

等价于

switch expression:
case list1:
{}
case list2:
{}
default:
{}
```
select 稍微有点不同，因为 select 中的 case 是可以新建一个变量的，除了我们说的每一个 case 是一个作用域之外，case 上新建的这个变量也只是属于下面这个小的作用域。