# 递归

什么是递归？

递归可以简单的理解是：自己调用自己，

我们也可以看成是 入栈和出栈 也就是FILO 最先进去的哪一层，却在最后才出来。这大概就是递归的一些思想。

举个简单的例子

```go
package example

func example(x int)int{
	if x == 1 {
		return 1
	}
	t := example(x-1) // 入栈的过程
	t++ //出栈后的操作
	return t //返回值
}

// 当然我们可以简写为

func exampleM(x int)int{
	if x == 1 {
		return 1
	}
	return exampleM(x -1) + 1 // 这一步 就是 将 入栈 和出栈 以及 返回 以及 计算 （也就是 每个栈单位都进行+1的计算）合为一体
}

```