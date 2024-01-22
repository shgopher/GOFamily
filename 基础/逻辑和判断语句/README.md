<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-01-22 17:17:14
 * @FilePath: /GOFamily/基础/逻辑和判断语句/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 逻辑和判断语句
go 语言基本上继承了 c 语言的逻辑判断语句，但是仍有些不同，比如：

- 在循环语句中，仅保留了 `for` 语句，do-while，while 均不支持，这也是满足了 go 小而美的原则 -- 单一功能仅保留一种方式。
- break continue 后面添加 label 功能
- switch 中 case 语句不会自动执行下一个 case，需要显示的使用 `fallthrough`，当然了在 select 中的 case，并没有 fallthrough 功能
- switch 中 case 支持表达式列表
- switch 增加 type 模式，让类型也可以作为选择的条件
- 跟 c 语言最大的不同，增加了 select - case 功能。
## if
go 语言中的控制语句基础语法：
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
go 语言推崇的控制语句是 “快乐路径”：

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
for 循环的基础语言和 c 类似，for range 语句中，range 后面接数组，指向数组的指针，切片，字符串，和拥有读权限的 channel
```go
for i:= 0;i< 10;i++ {}

for k,v := range sliceValue{}

```
其中普通循环语句中，i 值属于整个 loop 所有，这一点需要好好注意，我们在作用域那一章节中也提到过。

for-range 语句中，有两点需要注意，首先跟普通循环一样，变量属于整个 loop 所有。

其次，k,v 变量是切片或者 map 的 index 和 value 复制体，当遇到比如 slice 更改数据的时候，切勿直接更改 k v，可使用 ` arr[k] = v+1` 这种方式直接改变切片本身，另外，当使用 for-range 语句时，如果第二个变量没有使用的价值，可以不写，并且无需使用占位符 `_`：

```go

for k := range sliceValue{}

// 如果是省略第一个，那么还是需要占位符的


for _,v := range sliceValue{}
```

切片在运行时，len 是会变化的，因为确定切片的 len 是 runtime 的责任，而数组的 len 是在编译期确定的，for - range 后面的数组或者切片，真正处理的其实是这个数据的复制，下面看一个 bug
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

这里 number 输出的是 5，因为我们 append 添加的长度是原切片的长度，但是循环体中存储的 len，还是最初的长度 5，所以只能循环 5 次。

for-range 中还有一个容易出 bug 的事情，比如 range 后面跟一个数组，因为 range 的时候实际上是数组的复制品，看下面这段代码：

```go
package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	fmt.Println(a)

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println(r, a)

}
```

本来期望的 r 值是 1 12 13 4 5，因为在 i == 0 时就更改了数组 a 的值，但是最终输出的却是 1 2 3 4 5，原因很简单，因为 v 读取的是 a 这个数组的复制品，也就是说，实际上这个代码的隐藏含义是 ` range  a'` 这里的 `a'` 就是 a 的复制品，所以更改了 a，a 的复制品也不会被改变，解决方法就是取这个数组的切片就可以了，这样即便是复制了，也只是复制了一份儿指针而已。

### string
如果 for-range 中后面接的是 string，每次迭代不是按照 byte 来进行的，而是按照 rune 来进行，比如 `"你好"` 每次的迭代输出的就是你和好，而不是 byte
### map
for range 后面是 map 时，无法保证 map 的输出顺序，但是 map 和 slice 一样都是胖指针，所以时可以对 map 进行直接操作的。如果在循环体中新创建一个 map 项，那么这个项目在 range 时有可能会被输出，不能保证一定。
### channel
channel 的本质也是一个胖指针，所以它也可以直接被操作本体，channel 在 range 时是阻塞式的读取，如果不关闭 channel，这个 range 会一直阻塞。

```go
var c chan int

for v := range c {

}
```
这段代码就会造成一直阻塞，进而触发系统的 Panic

## switch && select
在 go 语言的 switch 和 select 的 case 中，我们经常会使用 break 来跳出循环，默认跳出的就是最小的那个循环单位，比如下面这个例子

```go
func a(){
  for {
    select {
      case <- time.After(time.Second):
      break
    }
  }
}
```
这段代码是有 bug 的，因为它无法跳出 for 这层循环，只能跳出 select 这里。那么我们该怎么做呢？这个时候应该使用 label 了，所谓 label 就是标签的意思，意思就是指定好了 break 的位置，它的作用就是这个。

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

这个时候就可以 break 到 Now 指定的层级了。

我们应避免使用 fallthrough 来执行多条件表达式，比如这样的代码

```go
switch v.(type){
  case int: fallthrough
  case int8: fallthrough
    println("yes")
}
```
这种代码也很烂，我们可以直接使用多个 case 并列的方式：

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

这段代码使用的就是判断 type 的断言模式，固定用法就是 ` value.(type) ` 前面是要判断的 value，后面是固定用法 type，必须是这个单词才行。

我再带你回忆一下普通的断言，`a.(int)` 除了在 switch 中的断言方式，普通的断言就跟这段代码是一致的，一个 any 类型 (interface {}) 后面跟具体的类型。

经过上面的初步介绍，接下来，我们深入看一下 for range 的一些底层原理
## 汇总一下 for 和 for range 中最容易迷惑的几段代码

第一段代码：

```go
func main() {
  arr := []int {1,2,3}
  for _,v :=  range arr {
    arr = append(arr,v)
  }
  fmt.Println(arr) 
}
```
这段代码 for 循环不会一直循环，原因是，arr 会在 range 一个复制一份儿，这个复制体的 len 在最初的 range 中的开头已经确定是 3，后面继续追加的 arr，并不会改变这个最初读取的 `len == 3 ` 这个结果。

不过，如果你使用的是传统的循环，那么这种写法就会出现 bug：

```go
package main

import "fmt"

func main() {
	n := 12
	for i := 0; i < n; i++ {
		n++
		fmt.Println("i")
	}
}
```
使用这种传统的 for 循环，因为 n 在循环体和循环内部都是同一个，所以循环不会结束

因此你应该将这种代码改写为 for - range 模式：
> go 1.22 增加了对于整数的 for range，之前只有 chan slice map，不过整数的 for range 前面只有一个变量，并且跟其他 for range 一致，n 为复制值，并且 for i := range n，i 也是复制值。这跟其他的 for range 保持一致

```go
package main

import "fmt"

func main() {
	n := 12
	for range n {
		n++
		fmt.Println("i")
	}
}
```

第二段代码：

```go
arr := []int{1,2,3}
result := [] *int{}
for ,v := range arr {
result = append(result, &v)
}
```
这段代码是存在 bug 的，&v，首先，根据作用域可知道，这个 v 是 loop 级作用域，那么这个 result 中存在的&v 就是同一个值，所以这个代码是错误的。

改正的方式就是放入正确的切片中数据的指针即可：

```go
arr := []int{1,2,3}
result := [] *int{}
for k := range arr {
result = append(result, &arr[k])
}
```

第三段代码：

```go
for i,_ := range result {

  result[i] = 0
}
```

这是一段对 int 切片进行归零的方法，很多人会觉得这要循环一次会非常浪费时间，其实不会，因为在编译器中，它不会真的循环，编译器会优化这个操作，直接给它内存清零。
