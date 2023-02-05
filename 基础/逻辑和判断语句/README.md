<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-02-06 00:51:34
 * @FilePath: /GOFamily/基础/逻辑和判断语句/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 逻辑和判断语句
go语言基本上继承了c语言的逻辑判断语句，但是仍有些不同，比如：

- 在循环语句中，仅保留了 `for` 语句，do-while ，while 均不支持，这也是满足了 go 小而美的原则 -- 单一功能仅保留一种方式。
- break continue 后面添加label功能
- switch 中 case 语句不会自动执行下一个 case，需要显示的使用 `fallthrough`，当然了在 select 中的 case，并没有 fallthrough 功能
- switch 中 case 支持表达式列表
- switch 增加 type 模式，让类型也可以作为选择的条件
- 跟c语言最大的不同，增加了 select - case 功能。
## if
go语言中的控制语句基础语法：
```go
if conditon {

} else if condition{

} else {
  
}
// 比如：

if a > 0 {
  println("a > 0")
}else if a < 12 {
  println("a > 0 并且 < 12")
}else {
  println("a > 12 或者是非正数")
}
```
go语言推崇的控制语句是 “快乐路径”：

- 当出现错误时，直接返回
- 成功的逻辑不要放在 if 语句中
- 返回值通常在此函数的最后一行

```go
err, value := handle()

if err != nil {
  return err
}

return value
```
## for
for 循环的基础语言和c类似，for range 语句中，range后面接 数组，指向数组的指针，切片，字符串，和拥有读权限的channel
```go
for i:= 0;i< 10;i++ {}

for k,v := range sliceValue{}

```
其中普通循环语句中，i 值属于整个loop所有，这一点需要好好注意，我们在作用域那一章节中也提到过。

for-range 语句中，有两点需要注意，首先跟普通循环一样，变量属于整个loop所有。

其次，k,v 变量是切片或者 map 的 index 和 value 复制体，当遇到比如 slice 更改数据的时候，切勿直接更改 k v ，可使用 ` arr[k] = v+1` 这种方式直接改变切片本身，另外，当使用 for-range语句时，如果第二个变量没有使用的价值，可以不写，并且无需使用占位符`_` :

```go

for k := range sliceValue{}

// 如果是省略第一个，那么还是需要占位符的


for _,v := range sliceValue{}
```

切片在运行时，len是会变化的，因为确定切片的len是runtime的责任，而数组的len是在编译期确定的,for - range 后面的数组或者切片，真正处理的其实是这个数据的复制，下面看一个bug
```go

	a := []int{1, 2, 3, 4, 5}
	var number int
	for i := range a {
		number++
		if i == 0 {
			a = append(a, 6, 7)
		}
	}
	println(number) // 5
```

这里number输出的是5，因为我们 append 添加的长度是原切片的长度，但是循环体中存储的len，还是最初的长度5，所以只能循环5次。在老的for-range中还有一个容易出bug的事情,比如range后面跟一个数组，因为处理的时候是数组的复制品，所以 `array[k] = v` 是无法更改老的数组的，除非使用数组的切片，或者数组的指针，但是最新的go中这个bug已经修复了，即便后面跟的是数组，那么也会默认去调用数组的指针。

### string
如果for-range中后面接的是string，每次迭代不是按照byte来进行的，而是按照rune来进行，比如 `"你好"` 每次的迭代输出的就是 你 和 好，而不是byte
### map
for range 后面是map时，无法保证map的输出顺序，但是map和slice一样都是胖指针，所以时可以对map进行直接操作的。如果在循环体中新创建一个map项，那么这个项目在range时有可能会被输出，不能保证一定。
### channel
channel的本质也是一个胖指针，所以它也可以直接被操作本体，channel在range时是阻塞式的读取，如果不关闭channel，这个range会一直阻塞。

```go
var c chan int

for v := range c {

}
```
这段代码就会造成一直阻塞，进而触发系统的Panic

## switch && select
在go语言的 switch 和 select 的 case 中，我们经常会使用 break 来跳出循环，默认跳出的就是最小的那个循环单位，比如下面这个例子

```go
func a(){
  for {
    select {
      case <- time.After(time.Second):
      break;
    }
  }
}
```
这段代码是有bug的，因为它无法跳出 for这层循环，只能跳出select这里。那么我们该怎么做呢？这个时候应该使用 label 了，所谓label就是标签的意思，意思就是指定好了 break 的位置，它的作用就是这个。

```go
func a(){
  Now:
  for {
    select {
      case <- time.After(time.Second):
      break Now
    }
  }
}
```

这个时候就可以break到 Now 指定的层级了。

我们应避免使用 fallthrough 来 执行多条件表达式，比如这样的代码

```go
switch v.(type){
  case int: fallthrough
  case int8: fallthrough
    println("yes")
}
```
这种代码也很烂，我们可以直接使用多个case并列的方式：

```go
package main

func main() {
	var i int16 = 12
	a(i)
}

func a(v any) {
	switch v.(type) {
	case int, int8:
		println("yes")
	default:
		println("no")
	}
}
```

这段代码使用的就是 判断type的断言模式，固定用法就是` value.(type) ` 前面是要判断的 value ，后面是固定用法 type，必须是这个单词才行。

我再带你回忆一下普通的断言，`a.(int)` 除了在switch中的断言方式之外，普通的断言就跟这段代码是一致的，一个 any 类型（interface{}）后面跟具体的类型。

