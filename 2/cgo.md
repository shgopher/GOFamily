# 动态插件，cgo等go可引入的其它插件系统
## go使用动态库
通过`plugin`包，可以将动态外库打包进程序里，而不是非要使用静态链库。
## go使用c语言的代码
- cgo：将c代码嵌入到go代码中，
- 或者将c代码嵌入到go的源代码中，一个写到go文件中，一个单独的`.c`存在

常见的用法是将c代码嵌入到go的代码中：

```go
package main

/*
#include <stdio.h>

void print(){
	printf("hello 科科人神");
}
*/
import "C"

func main() {
	C.print()
}
```
导入的c代码必须使用 `/**/`注释起来，并且，必须显示的`import "C"` 才可以。`C`是大写。 

## 还可以将c文件引入到go中：

```go
package main
/*
#include "hello.h"
*/
import "C"
func main() {
	C.hello()
}
```

同一个路径下有个 `hello.h`

除此之外还可以使用[静态库和动态库](https://studygolang.com/articles/28307?fr=sidebar)的方式引入c语言的文件
## 元编程

意思就是将代码看成数据，代码在执行过程中也是可以改变，更新，删除等特性的

常见的 反射 其实就是元编程的一种，它提供给我们在运行时改变代码的能力

除此之外，go还提供了编译器改变代码的能力：go generate