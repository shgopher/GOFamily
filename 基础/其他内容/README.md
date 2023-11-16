# 其他内容
## 指针
go 语言的指针类型是 unsafe.Pointer，`uintptr` 是指针的计算类型，也就是说，地址必须通过 unsafe.Pointer 提取以后，再转为 uintptr 才能进行计算，uintptr 的实质是一个整数类型，并且完全可以容纳所有的指针的数据。

go 语言使用 `*` 符号用来表示指针类型，以及取指针类型实际数据这个操作，使用 `&` 取变量的指针 (地址)，我们看一下操作
```go
func main(){
	// 使用 内置函数 new() 去取结构体的地址
	var p *People  = new(People)
	*p = People{
		"1"
	}
	// 使用 取值符号去取得 变量 p 的地址
	fmt.Println(&p)
}
type People struct{
	name string
}

```
我们看一下如何直接计算指针类型
```go

func main() {
	// a ，string 类型
	var a string
	
	// b ，a变量的地址
	var b *string = &a
	fmt.Println("打印初始a变量的地址", b)
	
	// c，转为可计算的指针类型之后的变量
	c := uintptr(unsafe.Pointer(b)) // uintptr(unsafe.Pointer())
	fmt.Println("打印可计算指针类型c", c)
	c++
	fmt.Println("打印可计算指针类型c++", c)
	
	// 将c再转化为 a 的地址
	b = (*string)(unsafe.Pointer(c)) // (*string)(unsafe.Pointer())
	fmt.Println("打印转换后的a的指针", b)
}
```
```go
打印初始a的指针地址 0x14000096230
打印可计算指针类型c 1374390149680
打印可计算指针类型c++ 1374390149681
打印转换后的a的指针 0x14000096231
```
## go 的引用类型和非引用类型

|引用类型|非引用类型|
|:---:|:---:|
|slice, interface, chan, map|其余|

引用类型的实质其实就是 fat pointer 即：胖指针，整个类型使用 struct 作为底层数据，data 是一个指针地址，它指向的是要使用的数据。

## 全局变量和局部变量

全局变量，引用类型分配在堆上，值类型分配在栈上

局部变量，一般分配在栈上，当局部对象过大的时候分配在堆上，如果对局部变量做逃逸分析，发现它逃逸到了堆上，那么就将其分配到堆上

## go init 函数的执行顺序
init 函数在一个包内的执行顺序：对同一个 `go` 文件的 `init()` 调用顺序是从上到下的，对同一个 `package` 中的不同文件，将文件名按字符串进行 “从小到大” 排序 (数字排在前面)，之后顺序调用各文件中的 `init()` 函数

对于不同的包，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 `init()` 函数，如果包存在依赖，例如：导入顺序 main –> A –> B –> C，则执行顺序为 C –> B –> A –> main

**go 会先执行全局变量再执行 init**，当然多包全局变量的初始化跟 init 的执行顺序是一致的

## go 可比较类型
> https://go.dev/ref/spec#Comparison_operators

不可以的

- 切片，map func (这几种都是因为自己本身变来变去，同样的变量，不同时间，内部值就变了，所以他们不可以参与 == 的比较) 变量无法参与一般的比较，但是他们可以和 nil 作对比

可以的：

- 数字类型，bool，string 这种常见类型可比较
	- `a := make(chan []int)` 即使是这样的内部含有不可比较的通道变量本身也是可以比较的。
- chan 当两个 chan 类型的变量进行比较时，只有它们都为 nil 或者指向同一个通道时才返回 true，否则返回 false。
- 内部字段都必须是可比较类型的数组和结构体可以比较
- 指针可以比较，指针的实质是一个整数类型
- 接口类型是可比较类型

所以，接口变量是可以作为 map 的 key 值的，因为接口可以比较
```go
package main

import "fmt"

func main() {
	b := map[interface{}]int{}
	var s Some
	b[s] = 1
	fmt.Println(b[nil])

}

type Some interface {
	methods()
}
```
## go 可寻址类型

以下内容是**不可寻址**的量

> 字面量的解释 var a int = 12，12 就是字面量，也就是所谓的那个值本身；结果值的解释：就是这个结果这个 value 的值

- 常量 `const a  = 12` &a 不可寻址

- 基本类型值的字面量 `a := 12` &12 不可寻址

- 算术操作的结果值 `a := 12  b := 12    &(b+a)`

- 对各种字面量的索引表达式和切片表达式的结果值。不过有一个例外，对切片字面量的索引结果值却是可寻址的

	```go
	a := map[int]int{1: 1}
		for k := range a {
			fmt.Println(&a[k]) // 无法寻址，这个数据属于临时的，可变的数据
		}

		b := []int{1,2,3}

		for k := range b {

			// 这个可以寻址，每个切片值都会持有一个底层数组，
			// 而这个底层数组中的每个元素值都是有一个确切的内存地址
			fmt.Println(&b[k]) 
		}
	```
- 切片字面量的切片结果值为什么却是不可寻址？`a := []int{1,2}, &(a[:1])` 原因是切片表达式字面量的切片属于临时值跟字面量的结果值不一样，后者属于结果值有效值

- 对字符串变量的索引表达式和切片表达式的结果值。`a := "a" &("a") &a[0]`

- 对字典变量的索引表达式的结果值。`a := map[int]int{1:1} &(a[1])`

- 函数字面量和方法字面量，以及对它们的调用表达式的结果值。`&(func ())`

- 结构体字面量，见下面例子

- 类型转换表达式的结果值

- 类型断言表达式的结果值

- 接收表达式的结果值

我们看一道**面试题**：

这道题就是因为中间变量无法获取地址造成的 bug

```go
package main

func main() {
	// 此处go会自动调用值的指针来运行 SetName 但是因为 return Dog{name} 是一个临时的值，所以无法获取到指针
	New("nihao").SetName("monster")

}

func New(name string) Dog {
	return Dog{name}
}

type Dog struct {
	name string
}

func (d *Dog) SetName(n string) {
	d.name = n
}
```

我们可以这么改

```go
package main

func main() {
	
a := 	New("nihao")
a. SetName("monster")

}

func New(name string) Dog {
	return Dog{name}
}

type Dog struct {
	name string
}

func (d *Dog) SetName(n string) {
	d.name = n
}

```


另外自增 ++ 自减 -- 左边的表达式都必须是可寻址的类型，否则也是无法操作的。
> 字典字面量和字典变量索引表达式的结果值是个例外例如 ma [“12”] ++

在赋值语句中，赋值操作符左边的表达式的结果值必须可寻址的，但是对字典的索引结果值也是可以的

总结一下：

- `常量` + string 这种无法更改的数据无法寻址，函数通常来说也可以算作 “常量”，应该它就是一段代码，不可更改
- `结果值/字面量` 因为其无法更改所以寻址将没有意义
- `中间值` 或者 `临时对象` 比如说 &(a + b) 这类临时的变量的内存地址没有意义
- `不安全的操作` 比如 map 中的 k-v 经常要从一个哈希桶迁移到另一个桶，所以你获取地址，它经常会改变，外界还不得而知，所以获取到这个 key-value 的地址是不安全的
