<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-06-11 08:43:57
 * @FilePath: /GOFamily/工程/错误处理/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 错误处理

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