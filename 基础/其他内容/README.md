# 其他内容
## go的引用类型和值类型

|引用类型|值类型|
|:---:|:---:|
|slice, interface, chan, map, func|array, struct, 数字类型, bool|

## 全局变量和局部变量

全局变量，引用类型分配在堆上，值类型分配在栈上

局部变量，一般分配在栈上，当局部对象过大的时候分配在堆上，如果对局部变量做逃逸分析，发现它逃逸到了堆上，那么就将其分配到堆上

## go init函数的执行顺序
init函数在一个包内的执行顺序：对同一个 `go` 文件的 `init()` 调用顺序是从上到下的，对同一个 `package` 中的不同文件,将文件名按字符串进行“从小到大”排序(数字排在前面),之后顺序调用各文件中的`init()`函数

对于不同的包，如果不相互依赖的话，按照main包中import的顺序调用其包中的`init()`函数，如果包存在依赖，例如：导入顺序 main –> A –> B –> C，则执行顺序为 C –> B –> A –> main

**go会先执行全局变量再执行init**，当然多包全局变量的初始化跟init的执行顺序是一致的

## go可比较类型
> https://go.dev/ref/spec#Comparison_operators


- 数字类型，bool，string ，指针，通道，可比较
	- `a := make(chan []int)` 即使是这样的内部含有不可比较的通道变量本身也是可以比较的。	
- 内部字段都必须是可比较类型的数组和结构体可以比较
- 切片，map func 变量无法参与一般的比较（但是变量的指针可以，因为指针可以比较），但是他们可以和 nil 作对比
- 接口类型是可比较类型
	- 空接口类型的实例是可以比较的，结果为true (俩实例都是nil，并且类型一样，所以相等)
	- 空接口的实例比较，赋值后（任何类型都可以赋值给空接口），会判断值的类型是否相等（比如说是否都是int类型赋值给了空接口），再判断相同类型的值是否相等，只有都满足才是true
	- 带有方法的接口，方法一样（证明类型一致），结果是true
	- 带有方法的接口，当被赋值以后（就是实现了方法，并且赋值给接口），只要赋值给俩接口类型的是相同的类型，那么就是true，否则是false
		```go
		func main() {
			var a A
			var a1 A
			var b B
			var c B // 如果 var c C 下面的结果就是false 
			a = b
			a1 = c
			fmt.Println(a == a1) // true
		}

		type A interface {
			d()
		}	

		type B int

		func (b B) d() {}

		type C int

		func (c C) d() {}

		```
	- 一个含有方法的接口实例和一个空的接口实例比较，并且双方实例均未赋值，那么两者相等。但是，对空接口赋值以后双方不想等
		```go
		package main

		import "fmt"

		func main() {
			var a A
			var b B
			fmt.Println(a == b) // true

		}

		type A interface {
			d()
		}

		type B interface{}

		```

所以，接口是可以作为map的key值的，因为接口可以比较
```go
package main

import "fmt"

func main() {
	b := map[interface{}]int{}
	var s Some
	b[s] = 1
	fmt.Println(b[s])

}

type Some interface {
	methods()
}
```
## go可寻址类型

以下内容是**不可寻址**的量

> 字面量的解释 var a int = 12 , 12就是字面量，也就是所谓的那个值本身；结果值的解释：就是 这个结果 这个value 的值

- 常量 `const a  = 12` a 不可寻址

- 基本类型值的字面量`a := 12` 12 不可寻址

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
- 切片字面量的切片结果值为什么却是不可寻址 `a := []int{1,2}, &(a[:1])` 原因是 切片表达式字面量的切片属于临时值跟字面量的结果值不一样，后者属于结果值有效值

- 对字符串变量的索引表达式和切片表达式的结果值。`a := "a" &("a") &a[0]` 

- 对字典变量的索引表达式的结果值。`a := map[int]int{1:1} &(a[1])`

- 函数字面量和方法字面量，以及对它们的调用表达式的结果值。`&(func ())`

- 结构体字面量，见下面例子

- 类型转换表达式的结果值

- 类型断言表达式的结果值

- 接收表达式的结果值

我们看一道**面试题**：

这道题就是因为中间变量无法获取地址造成的bug

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


另外 自增 ++ 自减 -- 左边的表达式都必须是可寻址的类型，否则也是无法操作的。
> 字典字面量和字典变量索引表达式的结果值 是个例外 例如 ma["12"] ++ 

在赋值语句中，赋值操作符左边的表达式的结果值必须可寻址的，但是对字典的索引结果值也是可以的

总结一下：

- `常量` + string 这种无法更改的数据无法寻址，函数通常来说也可以算作“常量”，应该它就是一段代码，不可更改
- `结果值` 因为其无法更改所以寻址将没有意义
- `中间值`或者`临时对象` 比如说 &（a + b） 这类临时的变量的内存地址没有意义
- `不安全的操作` 比如map中的k-v 经常要从一个哈希桶迁移到另一个桶，所以你获取地址，它经常会改变，外界还不得而知，所以获取到这个key-value的地址是不安全的
