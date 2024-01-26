<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2024-01-26 17:35:54
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-01-26 17:52:57
 * @FilePath: /GOFamily/基础/helloWorld/README.md
 * @Description: 
 * 
 * Copyright (c) 2024 by shgopher, All Rights Reserved. 
-->
# hello world
## 诞生背景
Go 语言是 Google 的 Robert Griesemer,Rob Pike 和 Ken Thompson 于 2007 年开始设计开发的一门编程语言。Go 语言吸收了 C 语言的表达能力，并添加了像动态语言中的并发性等特性，目的是让编程更简单高效。
## 语言哲学
Go 语言编程哲学如下：

- 面向接口编程
- 组合替换继承
- 简单优于复杂
- 内置并发支持

## 语言特点

Go 语言的主要特点有：

- 语法简单，容易学习使用
- 编译非常迅速，执行速度较快
- 静态类型语言 - 变量和函数需要声明类型
- 垃圾回收 - 自动的内存管理，不需要手动释放内存
- 并发支持 - 原生支持 goroutine 和 channel
- 简洁规范 - 语法简单一致，容易阅读理解
- 代码风格统一 - 提供了标准的代码风格

## 安装 Go

要安装 Go 语言，可以去官网 https://go.dev 下载最新的安装包。

安装非常简单，一般直接使用默认设置即可完成安装。

安装完成后，可以使用 `go version` 命令检查安装是否成功。
## hello world
```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("Hello World")
		ch <- struct{}{}
	}()
	<-ch
}
```
