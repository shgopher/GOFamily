# go语言零值

go在声明时期就会为变量提供默认值。

下面是所有的原生go类型的零值：

- 整数类型：0
- 浮点数：0.0
- 布尔类型：false
- 字符串类型：""
- 指针，接口，切片，channel，map，function：nil

至于复合类型，比如数组，和结构体的声明过程，就是`递归的，针对于它所有的子元素的初始化`

## 零值可用

go奉行零值可用，最经典的案例是切片的append过程，即是你最初的切片变量只是声明，并未make赋予底层数组，那么go系统仍然处理的结果是可用，并非Panic

```go
var s []int
s = append(s,1)
fmt.Println(s)
```
下吗再看一个例子
```go
func main() {
	var b *bytes.Buffer
	fmt.Println(b)
}
```
`fmt.Println(b)` 是调用的 `b.String()` ,那么为什么这里输出的是`<nil>`不是Panic呢？

```go
func (b *Buffer) String() string {
	if b == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}
	return string(b.buf[b.off:])
}
```
这就是答案，它在实现的String中实现了零值可用, `原因一:` **nil的对象是完全可以调用属于它身上的方法的**，`原因二：`这个方法内部实现的时候又是直接return一个 `<nil>` 避免了取nil索引值的情况。所以说基于这两点，它零值可用。

再看一个经典的buffer例子：

```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var a bytes.Buffer
	a.Write([]byte("12"))
	fmt.Println(a.String())
}
```
为什么只是声明了buffer就可以直接往里面写值呢？

```go
// 结构体
type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}

// 取自 Write() 相关的代码 这段代码保证了即便是初始化的buf是nil也可以零值可用。
if b.buf == nil && n <= smallBufferSize {
  b.buf = make([]byte, n, smallBufferSize)
  return 0
}

```
另外，当切片是nil的时候，也是完全可以取它的`[:]`切片的，只要不超过index都没问题，这也是满足零值可用的一个小tips

```go
var s []int
s = s[:] // 或者 s = s[:0]  or  s = s[0:] or  s= s[0:0]
```
## 零值不可用

- slice 零值赋值
- map 零值赋值
- 互斥锁的值复制

```go
var s slice
s[0] = 1 // 错误 ❌

var m map[int]int
m[1] = 12 // 错误 ❌
```
再看互斥锁的案例
```go
var mu sync.Mutex
mu.Lock()
mu.Unlock()
```
这段代码是可以正常使用的

但是，如果队mu进行值的复制就不能使用了

```go
package main

import (
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	foo(mu)
	mu.Unlock()
}

func foo(mu sync.Mutex) {
	mu.Lock()
	mu.Unlock()
}

```
这个问题的解释是这样的：互斥锁是带有状态的，就是说，当你复制的时候本来是a的状态，然后复制过去还是a的状态，但是这是一个新的对象了按道理应该是初始状态，所以就会出现错误，这也是传说中的重入锁(go不支持)，因为go的互斥锁是带有状态的，所以这种复制的方法就会出现错误。

## 参考资料
- https://github.com/golang/go/blob/037b209ae3e0453004a4d57e152aa522c56f79e4/src/bytes/buffer.go#L117
