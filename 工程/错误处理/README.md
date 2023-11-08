<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-11-08 10:26:54
 * @FilePath: /GOFamily/工程/错误处理/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 错误处理
## 错误处理的基本认识
在 go 语言中，没有传统编程语言的 `try - catch` 操作，go 语言中一切错误都需要显式处理：
```go
if err := readFile("./x"); err != nil {
	return err
}
```
通常，我们规定函数返回的最后一个数据是错误接口：
```go
func age(v int)(int, error){
	if v > 10 {
		return 10, nil
	}
	return -1, fmt.Errorf("错误x")
}
```
我们直接返回一个 err 是一种简单的做法，如果错误比较初级也可以这么做，但是如果想要带有更精确的提示信息，可以在返回的时候 wrap 一层信息：

就以上文的读取数据为例
```go
if err := readFile("./"); err != nil {
	return fmt.Errorf("在读取数据的时候发生了错误，错误信息是：%w", err)
}
```
wrap 一层信息，对于错误的定位更加高效

## Error 的本质是什么？
错误处理的核心就是下面这一个 error 的接口
```go
type error interface{
	Error()string
}
```
所以只要我们的自建类型实现了这个接口，就可以使用很多的错误处理的方法。
### 自定义 error
我们使用 errors.New() 的时候其实就是返回了一个 go 自建的，叫做 errorString 的实现了 error接口的结构体。
这就是自建error的方法
```go
func main() {
	e := errors.New("a")
	println(e)
}
// (0x47fd48,0xc00003e730)
```
为了防止在比较错误的时候发生错误一致的情况，所以自建 error，返回的实际上是一个指针。
> 下文会提用什么方法进行比较 err ，实际上就是 “两个接口类型是否相等 --- 类型一致，值一致”，如果返回的值是指针，那么值肯定就不可能一样了。

```go
// go 源代码

func New(s string) error {
	return &errorString{s}
}
```

## errors.Is()
上文我们说到，错误可以使用wrap的方式进行封装，那么如果我们想判断封装的这么多层的错误中，有没有哪一层错误等于我们要的值，可以使用 这个函数进行判断：

```go
func main(){
	err := errors.New("err is now")
	err1 := fmt.Errorf("err1:%w",err)
	err2 := fmt.Errorf("err1:%w",err1)
	if errors.Is(err2,err) {
		fmt.Println("是一个错误")
	}
}
```
## errors.As()
这个方法跟上文的 Is() 类似，只不过它判断的是类型是否一致。
```go
type errS struct {
	a string
}

func (t errS) Error() string {
return t.a
}

err := errS{"typical error"}
err1 := fmt.Errorf("wrap err: %w", err)
err2 := fmt.Errorf("wrap err1: %w", err1)

var e errS

if !errors.As(err2, &e) {
panic("TypicalErr is not on the chain of err2")
}
println("TypicalErr is on the chain of err2")
println(err == e)
```
## errors.Join()
这个方法是将几个错误结合在一起的方法：
```go
	err1 := fmt.Errorf("err1:%w",err)
	err2 := fmt.Errorf("err1:%w",err1)
	err := errors.Join(err1,err2)
```
## 当错误处理遇到了defer
```go
func age() (int, error){
	if xx {
		return 0 ,err
	}
	defer f.close
}
```
这段伪代码的意思是说，当 有条件之后，返回一个错误，但是defer 的内容发生在这个err被固定之后，所以defer中如果再有错误将不会被处理。

那么我们该怎么更改呢？

我想你一定想到了前文我们说过，使用带有变量的返回值是可以将defer的值进行返回的：

```go
func age() ( i int, e error){
	if xx {
			e = xx
			i = xx
		return 
	}
	defer f.close
}
```
那么这种写法，defer中如果发生了错误就会覆盖掉了 程序执行中的err，所以这种方法也是不行的，即使它能照顾到了defer中的错误处理。

我们可以将错误处理都放在defer中处理就可以了
```go
func age() ( i int, e error){
	if xx {
		i = xx
		e = xx
		return 
	}
	defer func(){
		e1 := f.close()
		if e != nil {
				if  e1 != nil {
					log(e1) 
				}
				return 
		}
		err = e1
	}()
}
```
这样，两种错误都能处理到了

## 错误处理实战的三种方式
### 经典的错误处理方式
### 屏蔽过程中的错误处理
### 利用函数式编程延迟错误处理
### 分层架构中的错误处理方法
### errgroup的使用技巧
## 错误处理实战技巧
这里会介绍在实战过程中用到的诸多技巧
### 使用 errors.New() 时要写清楚包名
```go
package age

ErrMyAge := errors.New("age: ErrMyAge is error")
ErrMyAddress := errors.New("age: ErrMyAddress is error")
```
### 使用 error处理一般错误，使用panic处理严重错误（异常）
使用这种模型就避免了类似 Java 那种所有错误都一样的行为，Java 的使用 try-catch 的方式导致任何错误都是一个方式去处理，非常有可能让程序员忽略错误的处理

然而 go 不同，**错误**使用 error，**异常**使用 panic 的方式去处理。
- 错误 ： error
- 异常：panic
### 基础库，应该避免使用 error types 
因为这种写法容易造成代码的耦合，尤其是在我们写的基础库中，非常容易造成改动代码来引入的不健壮性。
```go
package main

import (
	"fmt"
)

type ErrMyAge struct {
	errv string
}

func (e *ErrMyAge) Error() string {
	return fmt.Sprintf("age: %s", e.errv)
}

func main() {
	err := &ErrMyAge{"err age is hi"}
	fmt.Println(err)
}
```
或者使用 errors.New()
```go
ErrAge := errors.New("age: ErrAge is error")
ErrAddress := errors.New("age: ErrAddress is error")
```

实际上他们都是 error types ,如果别人使用了这个基础库，那么势必这些错误就会跟使用者的代码耦合，我们改动了代码，第三方的代码就会因此受到影响。



## 错误码的设置

## issues
`问题一：` **请说出下列代码的执行输出***

```go
package main

import "fmt"

func main() {
	defer func() {
		println("3")
	}()
	defer func() {
		println("2")
	}()
	defer func() {
		println("1")
	}()
	fmt.Println("hi1")
	panic("oops")
	// 这里的defer将不会进栈，所以也就不会执行了。
	defer func() {
		println("x")
	}()
	fmt.Println("hi2")
}

```
答案是
```go
hi1
1
2
3
oops
panic: oops

goroutine 1 [running]:
main.main()
	/tmp/sandbox1932632082/prog.go:18 +0xa7

Program exited.
```
解释：Panic，意味着恐慌，意思等于return，所以panic下面的数据是无法执行的，defer不同，他们是顺序的将这些defer函数装入函数内置的defer栈的，所以在return之后，defer栈会执行，所以这里的defer 1 2 3 可以执行，Panic前面的 hi1 可以执行，但是Panic之后，相当于return后面的hi2 就无法执行了。

`问题二：` **看一段代码，分析答案**

```go
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("starting f")
	g(2)
	fmt.Println("behind  g") //会终止执行
}

func g(i int) {
	defer fmt.Println("Defer in g", i)
	fmt.Println("Panicking!")
	panic(fmt.Sprintf("%v", i))
	fmt.Println("Printing in g", i) //终止执行
}
```
答案是
```js
starting f
panicking
Defer in g 2
recoverd in f
```

解释一下，首先执行的是 f函数的代码，然后开始执行g，在g中遇到了Panic，所以panic后面的 parinting in g 就无法执行了，所以执行了 defer in g
这个时候 f中的 g(2) 后面的数据也无法执行了，因为整个f也陷入了恐慌，所以它只能return 进入defer了，defer中刚好有recover，所以执行了recover信息后，就退出了函数。

## 参考资料
- https://mp.weixin.qq.com/s/EvkMQCPwg-B0fZonpwXodg
- https://mp.weixin.qq.com/s/D_CVrPzjP3O81EpFqLoynQ
- https://time.geekbang.org/column/article/391895
- 极客时间《go进阶训练营》