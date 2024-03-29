# 切片

导读：

- 切片的基本操作
- 切片和数组的基本概念
- 切片数组的底层数据结构
- issues

## 基本操作
```go
// 初始化一个切片,设置长度为5，容量为10
a := make([]int,5,10)
// 获取一个新的切片,遵循左闭右开的原则，取值时并且只看长度不能看容量
// b切片，长度为1，容量为8（下文有讲，向右原则）
b:= a[2:3]
// 判断切片长度为0
len(a) == 0
// append 数据
a = append(a, 1)
// 将切片内容清空（element设置为初始值，比如int就是0，string就是""）
clear(a)
//
// 将一个切片转为一个数组,注意，新数组不能超过切片的长度，且是数据的深度拷贝
var arr =  [4]int(a) 
```
## 切片和数组的基本概念
数组是拥有一段连续内存的数据结构，切片是存储了这个数据结构地址，长度以及容量的 struct，这里俗称这种数据结构 (类似切片) 叫引用类型

生成数组有两种方式 `[3]int` 和 `[...]int` 但是后者其实只是一种语法题，go 会自动推断出容量，因为推断是在编译期搞定的，所以并不会影响数组运行时的效率。

不同长度的相同数据类型数组不是一个类型，比如 `[1]string{}` 和 `[2]string{}` 就是两个类型

```go
t.Extra = &Array{Elem: elem, Bound: bound}
```
这段代码可以解释一下，它来自 go 的源代码，可以看出，生成数组的是一个 struct，那么显而易见了，里面的各项参数都必须一致的情况下 struct 才是一致的，所以，必须类型和容量都满足才行

但是切片没有这个烦恼，只要数据类型一致就是一种类型，因为它在编译期间的结构体中只有类型，并没有数量，数量需要在运行时才能确定

“切片的切片” 的容量是和 “切片的” 容量是不一致的 (比如这里的 a 和 b)，我们来看一个例子：

```go
package main

import "fmt"


func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:5]
	fmt.Println(len(a), cap(a), len(b), cap(b))
}
```

output：`6, 6, 3, 4`

我猜你肯定以为 b 的容量也是 6，但是不是，go 规定，***切片只能向右看***，不能向左看，我们来说一下上面这个例子，a 就不用说了，a 的底层数据结构数组就是 6 个长度，所以 a 自然长度和容量都是 6，但是 b，它是切片的切片，遵从左闭右开的规则，它的长度是从 index 的 2 到 4，也就是说是 `[3 4 5]`，自然它的长度就是 3，这毫无疑问，又因为切片遵从 “只能向右看容量” 的规则，它的容量从 index 2 开始往后算，也就是 `6 - 2 = 4` 所以它的长度是 3 容量是 4


## 切片，数组的底层数据结构
严格意义来说，go 的切片不存在扩容，如果切片想要的数据量大于底层数组的容量时，那么系统会做两件事，开辟新的数组，给这个数组生成新的切片，之前的数组和切片并没有任何的改变，而且如果没有被引用了，还会被 gc 掉

数组在数量小于等于 4 的时候，直接分配在栈内存里，如果大于四且没有逃逸到堆上时，变量就会在静态存储区初始化然后拷贝到栈上
> 静态存储区：内存在程序编译的时候就已经分配好，这块内存在程序的整个运行期间都存在。它主要存放静态数据、全局数据和常量。

>栈区：在执行函数时，函数内局部变量的存储单元都可以在栈上创建，函数执行结束时这些存储单元自动被释放。栈内存分配运算内置于处理器的指令集中，效率很高，但是分配的内存容量有限。任何临时变量都是处于栈区的，包括在 main () 函数中定义的变量

>堆区：亦称动态内存分配。程序在运行的时候用 malloc 或 new 申请任意大小的内存，程序员自己负责在适当的时候用 free 或 delete 释放内存。动态内存的生存期可以由我们决定，如果我们不释放内存，程序将在最后才释放掉动态内存。但是，良好的编程习惯是：如果某动态内存不再使用，需要将其释放掉，否则，我们认为发生了内存泄漏现象。

**切片的底层数据结构。** 

```go
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```
实际上，切片就是一个 struct (struct 是 go 的基本组成单位，实际上很多东西都可以使用结构体来表达，比如接口底层也是结构体)

当我们使用字面量的方式去生成一个 slice 的时候，会在编译期间搞定，但是如果使用 make 去生成一个切片时，就会在运行时生成一个切片，在运行时生成切片，不仅仅会进行边界检查，而且还会看是否逃逸到堆上。

**slice 的扩容规律，**如果期望容量大于当前容量的两倍就会使用期望容量，如果当前切片的长度小于 [256](https://github.com/golang/go/blob/fd0ffedae2dd9e202efc2dd7f7937baa08600d26/src/runtime/slice.go#L178) 就会将容量 x2，
如果当前切片的长度大于 256 就会逐步 (按照这个公式 `cap += (cap + 768) / 4`) 从 2 倍的倍增速率调整到 1.25 的倍率，增加到新容量大于期望容量。不过这还不是全部

slice 扩容具体操作。扩容的时候并不是严格按照这个规律来的，这只是一个大致的规律，实际运行的时候会进行 `padding` (填充)，也就是内存的对齐，将容量字节数尽量靠近 2 的次方，比如说期望的新容量是 5，这时期望分配的内存大小为 40 字节 (具体看一个位置多少个字节，这里是按照 8 举例子)，运行时会向上取整内存的大小到 48 字节 (取整到 go 的[推荐值](https://github.com/golang/go/blob/b634f5d97a6e65f19057c00ed2095a1a872c7fa8/src/runtime/sizeclasses.go#L84))，所以新切片的容量是 `48 / 8 = 6`

我们知道当切片扩容的时候，重新分配底层数组是会牵涉到内存的复制的，因此尽量减少内存复制就是我们要追求的事情了，当我们往后面追加数据的时候，如果可以提前预估要使用的容量，那么就不牵涉到多次的内存复制了

```go
func main() {
	const sliceSize = 1000
	m := make([]int, 0, sliceSize)
	for i := 0; i < sliceSize; i++ {
		m = append(m, i)
	}
}

```

**slice 的浅拷贝** `s := make([]int, 3,12)`，`s1 := s[1:2]`，那么 s1 和 s 是浅层拷贝，他们拥有一个共同的底层数组，更改任意一个另一个的值也会发生改变。

**slice 的深拷贝。**copy 内置函数会将源内存 data 直接一次性拷贝，所以慎用，很消耗资源。

**slice 的边界检查消除以及优化。**go 在编译数组或者切片的时候都会在编译期间进行边界越界的检查，不过也只是普通的检查，比如直接使用整数或者常量访问数组，对于**用变量去获取切片数组的情况**根本就检查不了，这个时候 go 的运行时就会起作用，发出 Panic，那么边界检查有什么缺点呢？就是会降低运行时的效率，当然这里指的是运行时的检查，因为编译器的检查首先本身能力有限，其次只会降低编译速度而已，所以 go 从 1.17 开始，就开始支持了边界检查的消除，为了就是在证明不会越界的情况下，提高代码的运行效率，下面我们看一下取消运行时边界检查的例子：

```go
package main

func main() {
	// 这种不能确定，所以还是会检查
	a := []int{1,2,3}
	for i := range a {
		a[i]
}

// 这样，就可以消除边界检查，因为运行时确定不可能出现问题
for i := len(a) - 1; i >= 0; i-- {
		_ = a[i]
	}
}

// 或者这样也可以

func A(i []int, b []byte) {

	if len(i) > 256 {
		i = i[:256] // 这里就是给运行时一个暗示，表示i的最大index不会超过 256
		for _, v := range b { // v的最大值也不会超过 256（byte最大值 255）
			_ = i[v]
		}
	}
}

```
## append 追加数据的操作
前文我们知道 slice 会发生扩容情况，那么这种扩容一般是在 append 操作的时候发生的，通常来说操作是这样的：
```go
a := make([]int,3,10)
a = append(a,1) // 注意，append操作是追加的意思，所以这里不是 100，而是0001
// 0001
```
当 slice 底层数组不够 append 的时候，就会发生扩容：
```go
a := make([]int,3,4)
a[0] = 1
a[1]= 2
a[2] = 3
// 这里还是相同的数组，因为容量没有超过
a = append(a,4) 

// 这里的a 指向的就不同于之前的那个数组了。
//底层就变成了一个新的数组  1 2 3 4 5 0 0 0 ;fmt.Println(cap(a)) : 8
a = append(a,5)
```
同样底层数组的不同区域的切片 append 的时候经常会发生意想不到的结果：

```go
func main() {
	a := make([]int,3,5)
	a1 := a[1:2]
	fmt.Println(a,a1)
	a1 = append(a1,2)
	fmt.Println(a,a1)
	a = append(a,4)
	fmt.Println(a,a1)
	a1 = append(a1,5)
	fmt.Println(a,a1)

}
//[0 0 0] [0]
// [0 0 2] [0 2]
// [0 0 2 4] [0 2]
// [0 0 2 5] [0 2 5]
```
从上面的案例可以说明，首先，a1 a 指向同一个底层数组，其次，a1 和 a 在 append 的时候，都是在各自切片的后面添加数据，他们会互相影响，在写代码的时候容易出现 bug
### 对 append 的优化
对一个未知大小的切片进行 append 操作的最佳选择是初始化一个 nil 切片：
```go
var s []string
s = append(s,"a")
```  
s 在初始化的时候没有分配内存，在 append 的时候分配了一个底层数组，下面这种方式就会浪费一次内存分配
```go
s := make([]string,0)
s = append(s,"a")
```
这种方式，在 s 初始化的时候会给予它一个底层长度为 0 的数组，即便长度为 0，go 并未分配实际的内存空间，但是仍然浪费了执行片段，append 的时候还要再次分配底层数组。

不过使用 make 的方式比较适合**已知容量**的场景。

除此之外还有两种初始化的方式：
- `[]int(nil)`
- `[]string("a")`

我们知道 `[]int{}` 和 `var s []int` 是两种皆然不同的初始化方式，虽然都是 0 长度，但是前者不是 nil，后者是 nil，不过 append 的时候不会介意是否是 nil，这也提醒了我们判断是否是空切片的方式，使用 len == 0 才是正确的方法，“是否等于 nil” 是错误用法。

`[]string(nil)` 的用法非常少见，通常来说，使用场景就下面这么一个：

```go
src := []string("hi","there","!")
s := append([]string(nil), src...)
```
我们这里使用一个值为 nil 的切片主要是为了符合类型的需求。
## 切片的 copy
copy 是 go 语言的内置函数，全局使用，`copy(a,b []Type)`，copy 是深度拷贝，它将后者的数据完全拷贝给前者。

要注意的是，将要被复制的元素能复制的量取决于前者的 length

比如下面这种情况，被复制的元素就是 0，但是并不会 panic
```go
src := []int{1,2,3}
var d []int
// []
copy(d,src)
```
一般来说，我们会使用相同的 length：

```go
src := []int{1,2,3}
d := make([]int,len(src))
copy(d,src)
```
或者直接使用 append 也能做到 copy
```go
src := []int{1,2,3}
d := append([]int(nil),src...)
```
## 切片未被合理 gc
当切片完成自己的使命时，我们希望它可以正常的被 gc 掉，通常来说，我们可以手动使用 `runtime.GC()` 来强制系统进行垃圾回收，下面我们看一种 bug，这种 bug 出现以后，我们手动的垃圾回收将会无效

```go
s := make([]int,10)
// 此处原本的想法是只取两个数据
// 但是造成了10个数据都不能垃圾回收，8个浪费
b := s[8:]
runtime.GC()
runtime.KeepLive(b)
```
本来我们期望 s 底层数组可以被垃圾回收，但是 b 也指向这个相同的底层数组，那么这个垃圾回收就无法执行。

正确的方法是
```go
copy(res,s[8:])
return res
```
当然了，我们深究 gc 的本质就会发现，三色 gc 算法是看某个对象是否还被变量持有，是否可以通过变量的方式追踪到，来作为回收标准的，要么我们像前者一样，不再让变量持有值，要么将值变成 nil 即可：
```go
type A struct {
	v []byte
}
s := make([]A,10)
// 此处原本的想法是只取两个数据
// 但是造成了10个数据都不能垃圾回收，8个浪费
for i:= 0;i<8;i++ {
	// 这里是手动的将前八个数据丢弃掉
	s[i].v = nil
}
b := s[8:]
runtime.GC()
runtime.KeepLive(b) 
```


这种多切片指向底层数组而造成的无法正常垃圾回收的行为很常见。在工作中还是应该检查好自己的代码，避免这种行为的发生。
## 切片中的 range 注意事项

range 时，我们直接修改返回的值是不会生效的，因为返回的值并不是原始的数据，而是数据的复制。

```go
type Student struct {
	year int
}

func main() {
	result := []Student{
		{12},
		{13},
	}
	for _, student := range result {
		student.year++
	}
	// 12 13
	fmt.Println(result)
}
```
也就是说，这里的 `student := ` student 的改变不会影响 result 中的任何数据，除非这里的 `[]Student` 是 `[]*Student`

下面我们演示一下，正确的在 range 时的操作方式：

```go
type Student struct {
	year int
}

func main() {
	result := []Student{
		{12},
		{13},
	}
	for i := range result {
		result[i].year++
	}
	// 13 14
	fmt.Println(result)
}
```
正确的方式就是直接使用 result 本身进行操作，就可以真正的去改变 result 了。

并且，在 range 的时候，range 后面的数据其实也是复制品，也就是说，这里的 ` := range result` result 也是复制品，原有的 result 如何变化都不会影响 range 结果。

```go
// 这里的 result
for i := range result {
	// 跟这里面的result不是一个值，
	//只有里面的result才是跟外面的result是一个值
		result[i].year++
	}
```

下面我们再看一个案例：
```go
s := make([]int,3)
for range s {
	s = append(s,10)
}
// [0 0 0 10 10 10]
```
你猜结果是多少呢？是会一直 range 吗？因为数据在一直添加啊，nonono，只会 range 3 次而已，因为 range 后面的 s 是固定不变的，它本身只是原有 s 的复制品而已。

综上：

- range 后面的数据是原有数据的复制品
- range 前面的 k v 更是后面复制品输出数据的复制品
- range 里面的数据才是跟外面的数据保持一致

第三点很关键，range 后面的数据跟 range 里面的数据并不是一个：
```go
s := make([]int,3)
// 这里的s
for range s {
	// 这里的s 跟外面保持一致
	s = append(s,10)
}
```
**不能**把 range 当做 function 来类比：
```go

func main(){
	s := make([]int,1)
	range(s)

}
func range(s any)(k,v any){
	s[0]++
}
```
如果是函数，函数体的变量 s 和函数内部的 s 就是同一个，显然，range 中，range 后面的 s 和 range 里面的 s 并不是同一个。
 
## 切片转化为数组
在 go 1.20 版本中，新添加了切片转为数组或者数组指针的操作，具体实现如下：
```go
s:= make([]int,2,4)
a := [1]int(s) 
a1 := [2]int(s)
a2 := (*[1]int)(s)
a3 := (*[2]int)(s)
```
`a[0] == s[0]`，也就是说，转化出来的就是底层的那个数组的复制，**注意并不是底层数组本身**，不过这里相同切片转化的数组指针是指向这个切片的底层数组的，所以 a2 和 a3，s 是公用一个数组的，如果更改了 a2[0]，那么 a3，s 也是会发生改变的，a1 a 已经是数据的复制了，他们有了分别的生活了。

下面我们看一个具体的代码演示：
```go
func main() {
	s := make([]int, 2, 4)
	a := [1]int(s)
	a1 := [2]int(s)
	a2 := (*[1]int)(s)
	a3 := (*[2]int)(s)
	fmt.Println(s, a, a1, a2, a3)
	a[0] = 1
	fmt.Println(s, a, a1, a2, a3)
	a2[0] = 12
	fmt.Println(s, a, a1, a2, a3)
}
//[0 0] [0] [0 0] &[0] &[0 0]
//[0 0] [1] [0 0] &[0] &[0 0]
//[12 0] [1] [0 0] &[12] &[12 0]

// 可以证明确实是指向同样的底层数据
fmt.Println(&a2[0], &a3[0], &s[0])
// 0x1400012a020 0x1400012a020 0x1400012a020
```

当一个切片转化为一个数组的时候，数组的长度不能大于切片的长度，而不是容量

那么下面这种代码就会出现 bug

```go
s:= make([]int,2,4)
a := [3]int(s)
a1 := (*[3]int)(s)
// panic: cannot convert slice with length 2 to array or pointer to array with length 3
```
下面我们看一下转化后的特殊值

```go
s:= make([]int,2,4)
a := [0]int(s) //  []
a1 := (*[0]int)(s)// &[]

//将非空切片转为长度为 0 的数组，得到的指针不是 nil，比如 b2
var j []int
b := [1]int(j) // panic
b1 := [0]int(j)//  [] 
b2 := (*[0]int)(j)//  <nil>
b3 := (*[1]int)(j)// panic

//将 nil 切片转为长度为 0 的数组，得到的指针为 nil
c := make([]int,0)
u := (*[0]int)(c) //&[] 
u1 :=[0]int(c) //  []
```
## issues
`问题一：` ***如果有多个切片指向了同一个底层数组，那么你认为应该注意些什么***

一定要避免 a 切片的更改造成的底层数据的改变，对 b 切片的结果造成影响，因为它们指向同一个数据底层

```go
a := []int{1,2,3}
b := a[:]

a[2] = 4

fmt.Println(a,b)
```
`[1 2 4] [1 2 4]`

`问题二：` ***怎样沿用扩容的思想对切片进行 “缩容***

```go
a:= []int{1,2,3}
b := a[:2] // b = [1,2]

// 如果确定a的数据多余的没有任何的用途了
nb := make([]int,2)
copy(nb,b)
```
所谓扩容的思想，就是创造一个新的底层数据

`问题三：` ***nil 切片和空切片 (比如 []int {}) 的区别***

最大的区别就是指向的底层数组的地址不一样

- nil 压根就没有地址
- 空切片是有正儿八经的地址的，**只不过这个地址指向的数组不占用空间**，这个数组叫做 zero 数组，并且所有的空切片指向同一个数组就是这个 zero 数组，也可以说在 go 里，zero 数组是唯一的存在，它存在的目的就是为了空切片

```go
a := [0]int{}
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a)) // 0
```
空的数据是不占内存空间的，还有类似的，比如空的 struct 也是一样的

```go
a := struct{}{}
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a)) // 0
```

`问题四：` ***slice 和 array 的不同使用场景是什么***

如果数据是固定的，可以用数组，否则还是切片更加灵活，实际上绝大多数情况下还是切片更好用
## 参考资料
- https://www.jianshu.com/p/9ea2fba64f06
- https://chai2010.cn/advanced-go-programming-book
- https://blog.csdn.net/kevin_tech/article/details/122138489
- https://blog.csdn.net/weixin_39927993/article/details/112099007
