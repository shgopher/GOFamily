# 复合字面量作为构造器

go的复杂类型经常采用复合字面量的方式进行初始化，例如 struct，数组，slice map，比如：

```go
a := Some{
  name:"hi there",
  year:12,
  }
arr := [...]int{1,2,3,4}

sl := []int{1,2,3}

m  := map[int]int{
  1:1,
  2:2,
}
```

当struct需要使用 `{key :value}` 来构建时，即使可以按照顺序省略掉key，也千万不要这么做，加上key就可以将结构体的**使用和设计解耦**，因为可以乱序进行struct的赋值，并且即使某些字段没有赋值，那么系统也会自动的赋予它的初始值。当获取struct的指针时，最好是在字面量上取`&` ，比取变量的指针地址更加符合具体的含义，即：获取的事结构体的指针类型并且复制给一个变量。比如`a  := &Some{}`就比`a := Some{} , &a`更好，另外当我们要获取一个结构体的初始值的时候使用`a := Some{}` 比 `a = new(Some)` 更加常用。

数组和切片的复合体跟struct不同，它使用下标来作为key，它们几乎不会使用key：value 的形式，不过下面两种情况还是会出现的：

- 为了省略中间元素
- 为了显著的体现下标

```go
//为了省略中间元素
func main() {
	a := [...]int{12, 4: 3}
  b := [...]int{'a':0,'b':1}
	fmt.Println(a,b)
}
```

```go
//为了显著的体现下标
var data = []int{ 0:-10, 1:-5, 2:0, 3:1, 4:2, 5:3 }
```

当slice和array含有复合类型的时候，可以直接省略复合类型的类型名称，直接用`{}`即可
```go
package main

import "fmt"

func main() {
	a := []Some{
		{
			name: "12", // 这里尽量不要省略字段名称，方便解除跟结构体设计时的耦合
			year: 12,
		},
		{
			name: "12",
			year: 12,
		},
	}
	fmt.Println(a)

}

type Some struct {
	name string
	year int
	p    bool
}
```

map字面量作为初始值的时候，非常自然的就会使用key-value的方式，需要注意两点
- 当value或者key是复合类型的时候（key必须是[可比较的](../%E5%85%B6%E4%BB%96%E5%86%85%E5%AE%B9/README.md#go可比较类型)）可以省略复合类型的类型名称
- 如果是指针类型，那么连指针也可以省略

## 参考资料
- 图书:go语言精进之路
