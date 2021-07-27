# golint 静态代码扫描检查工具

静态代码扫描检查简单的说就是，在不运行这段代码的情况下，通过扫描代码段，来分析代码的质量和可能出现的问题，以及不规范的地方。我们使用go lint去检查写好的代码。

比如说以下这几种常见的错误都可以被检查出来

1. 逻辑混乱
```go
if a != false {}
```
明显就不如
```go
if a{}
```
更加易懂以及规范

2. 或者是实际上的错误比如说

```go
fmt.Printf("this is %d","my dear")
```
这种明显的错误也会被检查出来。

这里推荐一个项目是叫做 [golangci-lint](https://github.com/golangci/golangci-lint)

使用方法：

在项目根目录下新建一个`.golangci.yml`文件 

```zsh
linters:
  disable-all:true  # 关闭其它的linter
  enable: # 下面是开启的linter（全局全关，下面的意思就是豁免的）
    - deadcode # 找到没有用的代码
    - errcheck # 有些错误并没有被检查是否是错误【if err != nil 】所以这段代码就是寻找到没有那句话的地方，因为所有的错误都应该被检查
    - gosimple # 简化代码
    - govet # 报告可疑结构，例如参数与格式字符串不一致的 Printf 调用
    - ineffassign # var没有被使用
    - staticchek # 静态分析检查
    - structcheck # 寻找那些没有被使用到的struct字段
    - typecheck # 就像 Go 编译器的前端一样，解析和类型检查 Go 代码
    - unused # 检查 Go 代码中未使用的常量、变量、函数和类型
    - varcheck # 查找未使用的全局变量和常量
    - scopelint # Scopelint 检查 go 程序中未固定的变量
    - golint # 执行 Effective Go 和 CodeReviewComments 中提出的风格约定
liters-settings:
  govet: # 对govet做出某些具体的规定
    check-shadowing:true
    check-rangeloops:true
    check-copylocks:true
```

然后在根目录下执行命令golangci-lint run 即可。

给出一个具体的例子：

代码：
```go
package main

import "fmt"

func main() {

}

// age is a differ
func age() {
	var a bool

	if a != false {
		fmt.Printf("this is a %d", "12")
	}

}

```

结果：

```zsh
golangci-lint run

main.go:9:6: `age` is unused (deadcode)
func age(){
     ^
main.go:12:4: S1002: should omit comparison to bool constant, can be simplified to `a` (gosimple)
if a != false {
   ^
main.go:13:2: printf: fmt.Printf format %d has arg "12" of wrong type string (govet)
	fmt.Printf("this is a %d","12")
	^
```

