# 字符串
- 字符串的基础操作
- 字符串的基础知识
- 字符串的高效构造
- 字符串的底层操作
## 字符串的基础操作
```go
// 声明
var s = "你好"// s := "你好"
//读取
fmt.Print(s)
// 转化为unicode码点存储
_ = []rune(s)
// 转化为字符存储
_ = []byte(s)
// 字符的简单拼接
s1 := "你"
s2 := "好"
s := s1 + s2
// 多行字符
s = `
你好

世界
`
```
## 字符串的基础知识
- 字符串的数据是只读数据不可更改
- 字符串的零值是可用的，`var s string` 结果是`""`
- 获取字符串长度的操作时间复杂度是`O(1)`因为它是不可变的只读数据，所以长度被保存在了字段中，直接读这个字段即可
- 字符串可以通过`+`,`+=` 进行拼接
- 字符串可以使用`> < >= <= == =!` 运算符，比较的顺序是：
  - 先比较长度
  - 再比较是否是指向一块内存地址
  - 如果都满足再比较具体数据
- 字符串原生支持unicode字符集，并且go默认支持utf-8的编码算法
  - rune存储unicode的一个码点
  - byte存储真实的底层字符（比如utf-8，三个字符来保存一个中文字符，rune就只显示一个字符，byte会显示三个）
  ```go
  package main

  import "fmt"

  func main() {
    var s = "中"
    fmt.Print([]byte(s), []rune(s))
  }
  // output: [228 184 173] [20013]
  ```
- 使用``原生支持多行字符
## 字符串的高效构造
字符串的构造有以下这么几种
- 最常规的使用`+`和 `+=`
- fmt.Sprintf
- strings.Join
- strings.Builder
- bytes.Buffer

让我们分别给出代码：
```go
//最常规的使用`+`和 `+=`
package main

import "fmt"

func main() {
	s1 := "中"
	s2 := "国"
	fmt.Print(s1 + s2)
}
```
```go
//fmt.Sprintf
package main

import (
	"fmt"
)

func main() {
	s1 := "中"
	s2 := "国"
	fmt.Print(fmt.Sprintf("%s%s", s1, s2))
}
```
```go
//strings.Join
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "中"
	s2 := "国"
	fmt.Print(strings.Join([]string{s1, s2}, ""))
}
```
```go
//strings.Builder
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
  b.Grow(2) // 给出猜测最终的string长度
	s1 := "中"
	s2 := "国"
	b.WriteString(s1)
	b.WriteString(s2)
	fmt.Print(b.String())
}
```
```go
//bytes.Buffer
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	s1 := "中"
	s2 := "国"
	b.WriteString(s1)
	b.WriteString(s2)
	fmt.Print(b.String())
}
```
根据benchmark，得出以下结论：
- 带有预估string长度的strings.Builder 最快
- 带有预估的bytes.Buffer 和strings.Join 性能第二档次
- 没有预估长度的strings.Builder和bytes.Buffer以及 + += 第三档次
- fmt.Sprintf最差劲

那么：

- 当能给出预估的情况下，优选使用 strings.Buffer
- strings.Joins性能最稳，没有预估的情况下，使用这个稳定啊
- 操作符 + += 最直观，并且在字符短，以及编译器知道连接的字符串个数时，这种方式还能得到编译器的优化
- fmt.Sprintf 用在多类型组成字符串的时候是最好的，虽然它效率很差，但是人家能力强啊
## 字符串的底层

## 参考资料
- 图书: go精进之路
- https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-string/