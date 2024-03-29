# 字符串
- 字符串的基础操作
- 字符串的基础知识
- 字符串的底层操作
## 字符串的基础操作
```go
// 声明
var s = "你好"// s := "你好"
//读取
fmt.Print(s)
// 转化为unicode码点存储，单个字符， 例如'你'，也是使用的rune来存储的，、
// 不过string底层本身是[]byte的方式存储的
_ = []rune(s)
// 转化为字符存储
_ = []byte(s)
// 字符的简单拼接
s1 := "你"
s2 := "好"
s := s1 + s2
//多行字符
s = `
你好

世界
`
// 多行字符也通常用在屏蔽 ""作用的地方
s= `{"user": "shgopher", "links": ["https://github.com/shgopher"]}`
```
## 字符串的基础知识
- 字符串的数据是只读数据不可更改
- 字符串的零值是可用的，`var s string` 结果是 `""`
- 获取字符串长度的操作时间复杂度是 `O(1)` 因为它是不可变的只读数据，所以长度被保存在了字段中，直接读这个字段即可
- 字符串可以通过 `+`，`+=` 进行拼接
- 字符串可以使用 `> < >= <= == =!` 运算符，比较的顺序是：
  - 先比较长度
  - 再比较是否是指向一块内存地址
  - 如果都满足再比较具体数据
- 字符串原生支持 unicode 字符集，并且 go 默认支持 utf-8 的编码算法
  - rune 存储 unicode 的一个码点
  - byte 存储真实的底层字符 (比如 utf-8，三个字符来保存一个中文字符，rune 就只显示一个字符，byte 会显示三个)
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
- 最常规的使用 `+` 和 `+=`
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
根据 benchmark，得出以下结论：
- 带有预估 string 长度的 strings.Builder 最快
- 带有预估的 bytes.Buffer 和 strings.Join 性能第二档次
- 没有预估长度的 strings.Builder 和 bytes.Buffer 以及 + += 第三档次
- fmt.Sprintf 最差劲

那么：

- 当能给出预估的情况下，优选使用 strings.Builder
- strings.Joins 性能最稳，没有预估的情况下，使用这个稳定啊 (实际上这个 join 就是调用了 string.Builder，并且给出了预估长度)
	
```go
func Join(elems []string, sep string) string {
switch len(elems) {
case 0:
	return ""
case 1:
	return elems[0]
}
// 这里是搞定string长度的
n := len(sep) * (len(elems) - 1)
for i := 0; i < len(elems); i++ {
	n += len(elems[i])
}
// 使用了builder
var b Builder
b.Grow(n)
b.WriteString(elems[0])
for _, s := range elems[1:] {
	b.WriteString(sep)
	b.WriteString(s)
}
return b.String()
}

```
- 操作符 + += 最直观，并且在字符短，以及编译器知道连接的字符串个数时，这种方式还能得到编译器的优化
- fmt.Sprintf 用在多类型组成字符串的时候是最好的，虽然它效率很差，但是人家能力强啊

**综上所诉**：优先选 `strings.Join`
## 字符串的底层
### 数据结构
一个 string 的底层数据类似一个 slice，只不过这个 slice 是只读数据，它的底层不同于一般的 slice，是一个特别的 struct
```go
type stringStruct struct{
	str unsafe.Pointer
	len int 
	// 注意常规的slice这里是有一个cap的
	//但是string因为是只读的关系只有length的含义
}
```
`runtime/string.go` 中出现了这么一段代码
```go
// rawstring allocates storage for a new string. The returned
// string and byte slice both refer to the same storage.
// The storage is not zeroed. Callers should use
// b to set the string contents and then drop b.
func rawstring(size int) (s string, b []byte) {
	p := mallocgc(uintptr(size), nil, false)

	stringStructOf(&s).str = p
	stringStructOf(&s).len = size

	*(*slice)(unsafe.Pointer(&b)) = slice{p, size, size}

	return
}
```
我们仔细看注释的这句话，当一个 string 导入数据的时候，运行时会给定一个辅助的 slice，用来辅助的导入数据，然后当数据导入完毕之后，这个 slice 的描述符，也就是这个代表了这个 slice 的 struct 就会被删除掉，所以说其实 string 不能跟 slice 划上等号，也不能简简单单的说它是一个只读的 slice，实际上它压根就不是 slice，slice 在生成 string 的过程中只是起到了辅助作用

### 类型转换
字符串进行的转化只能是 string 和 `[]rune` or `[]byte` 互相转换
```go
package main

import "fmt"

func main() {
	a := "【你好】"
	b := []byte(a)
	c := []rune(a)
	// []byte 是可以直接使用fmt包直接输出为string的，但是[]rune需要显式进行转换。
	fmt.Printf("现在我们打印出原始数据%s,打印出[]byte转化后的数据%v，打印出[]rune转化后的数据%v,打印出逆转的数据%s 和 %s,", a, b, c, b, string(c))
}
```

上文我们提到的字符串的构造，例如删除一个字符，追加一个字符，都无一例外需要改变这个 string，那么很明显任何数据的处理都是**拷贝**的数据，原数据是不会有任何变化的，所以这就告诫我们字符串的处理要小心非常有可能浪费大量的内存。

我们看一下底层的转换代码：
```go
const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

func stringtoslicebyte(buf *tmpBuf, s string) []byte {
	var b []byte
	if buf != nil && len(s) <= len(buf) {
		*buf = tmpBuf{}
		b = buf[:len(s)]
	} else {
		b = rawbyteslice(len(s))
	}
	// 这里就发生了拷贝
	copy(b, s)
	return b
}

```

于此同时我们也能发现 string 的底层的真实存储是 `[]byte` 不是 `[]rune`
```go
package main

import "fmt"

func main() {
	a := "【你好】"
	var b []byte
	copy(b, a)
	var c []rune
	// invalid argument: arguments to copy c (variable of type []rune) 
	// and a (variable of type string) have different element types rune and byte
	copy(c, a)
	fmt.Println(b, c)
}

```


go 为某几种特别的情况优化了 string 和 slice 转换必须拷贝的情况，意思就是不需要拷贝就让这个 string 直接使用这个 slice 的底层，但是有个规定，只要是 slice 发生了改变，那么这个 string 立即失效

`b = []int{1,2}`

- string(b) 用在 map 的 key 中 `ma[string(b)]++`
- string(b) 在字符串的拼接句子中 “a” + string(b)
- for-range 中的 string 到[]byte 的转换

## for-range 字符串
因为对一个字符串使用 range 的时候，go 默认使用 utf8 的编码方式，但是 string 的底层是[]byte 的存储方式，所以直接 range 的时候，将这个时候的字符转化为字符就会发生乱码的情况
```go
s := "你好"
for k := range s {
	// äå 
	fmt.Printf("%c",s[k])
}
```
解决方法有两种：

1. 直接获取 value 值
```go
func main() {
	s := "你好"
	// 这里的 v 就直接是 rune
	for _, v := range s {
		fmt.Printf("%c", v)
	}
}
``` 
2. 将 s 转化为 []rune 来获取真正的 unicode 编码：
```go
func main() {
	s := "你好"
	sr := []rune(s)
	// 这里的 v 就直接是 rune
	for i := range sr {
		fmt.Printf("%c", sr[i])
	}
}
```

