# map
导读：

- map 的基本操作
- map 基础知识
- map 底层知识
- iussues
## map 的基本操作
因为 go 使用拉链法去构造 go 语言的 map，所以只要内存不被消耗光，map 中的元素是无限的。跟 slice 不同，map 不分 lenth 和 cap，在 make 中**只有**一个位置，这个位置表示的含义就是 cap

map 的基础操作
```go
func main() {
	// 声明一个新的map
	var m map[string]int
	// 给引用类型map分配底层数据结构
	m = map[string]int{} // or m = make(map[string]int,100)
	// 插入数据
	m["hello"] = 1
	if v, ok := m["hello"]; ok {
		fmt.Println(v)
	}
	// 在初始化的时候无法初始化两个一样的key，这个检查是编译器就开始的
	m = map[string]int{"a":1,"a":2} // error
	// 在初始化的时候无法初始化两个一样的key，但是如果是一个变量的话是可以的，
	// 因为编译器时无法获取变量的值的，所以这种方式通过。
	m = map[string]int{a:1,a:2} // 这样是正确的。
	// 遍历数据
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	// 删除数据,不存在这个数据不会panic
	delete(m, "hello")
	// 删除 map 所有条目
	clear(m)

}
```
## map 基础知识
映射的第一步就是把键值转化为哈希值 (通常是一个很大的整数)，然后根据哈希值的映射，把 key-value 本体，以及对应的哈希值存储在哈希筒中，go 使用链表法充当哈希筒，当寻找值的时候，通过哈希计算，以及取模映射，先查找到哈希筒，然后使用哈希值的方式去寻找有没有符合的哈希值，如果找到了，那么再使用 key 原来的值二次比较，这一步主要是为了规避哈希碰撞，即：俩 key 算出来的哈希值是一样的这种行为。

因为存 `==` 这种行为，所以 go 语言 map 的键值是有局限的，**不可以是函数类型、字典类型和切片类型**，如果是接口类型，传入的实际类型也不能是函数，字典和切片，例如

```go
package main

func main() {
	_ = map[interface{}]int{
		[]int{1, 2}: 2, // panic
		"1": 1,
	}

}

```
**虽然，map 的 key 可以设置为接口，但是最好不要这么干，因为会在运行时引入风险，比如上述代码就是风险。同样的，如果 key 是数组类型，亦或者是 struct，参数值都不能是函数，切片和 map。**无论被埋藏的有多深，都不能出现切片，map，函数这几种类型，比如说 `[10][2][]string`，运行时都会看出来。

那么用什么类型的作为 key 值是比较推荐的呢？

这里有两个关键词：求哈希的速度，判断相等的速度，基本原理是越简单的类型速度越快，比如 bool，int8，就比 int64 快，因为 int8 单个值只占了一个字节，int64 是 8 个字节，所以越复杂的越慢，比如一个 struct，因为求一个 struct 的哈希值，需要对里面的字段都进行哈希计算，然后合并起来，大大影响了速度，所以优先选用数值类型和指针类型 (因为指针类型也就是一个 16 进制的正整数而已)，如果选 string，最好短一些的比较好

在内存不爆炸的情况下，map 的 key 是无限量的，随意添加。

map 有一个最佳实践是使用形如 `value,ok := map[key]` 的 “comma OK” 的方法去获取 map 的值，OK 的意义就是为了获得 key 是否存在这个 map 里，因为就算不存在，map 也不会报错或者 Panic，返回的是这个值的零值，比如 int 就返回 0，那如果某一个 key 刚好结果是 0 就说不清了对吧，所以引入了这个 “comma ok” 机制，另外还存在一个不存在也不会 Panic 的操作，就是使用 `delete(map,key)` 的方法去删除 key 值

map 的遍历跟 slice 一样，使用 m := range map 的方法，但是输出的顺序是不固定的，如果想输出稳定的值，可以将 range 改成普通的 for 循环，然后 key 值使用一个切片存储，这样的话，读取 key 值的时候顺序就是固定的了。

## map 底层知识

当 map 使用字面量的方式进行初始化的时候，数量小于25时，它会首先启用 make 关键字创建一个 map，然后使用 `m["k"] = v` 的方式进行初始化，当超过 25 的时候，其实也差不多，它会使用两个切片分别存储 key 和 value，然后使用 `m[k] = v` 的方式进行初始化，实际上 go 的复合类型使用字面量进行初始化基本上都是这种方式，基本上的流程就是使用关键字创建底层数据，然后使用最简单的方式进行 k-v 赋值，所以你也可以把字面量的赋值当作一种语法糖。



map 的语法在运行时会转化为另一套对应关系，这个转化是在编译器搞定的。

```go
m := make(map[string]int, 10) -> m := runtime.makemap(maptype, 10, m)// maptype 下文有解释

v := m["hello"] -> v := runtime.mapaccess1(maptype,m , "hello")  // 这里实际上引入的是 "hello"的 指针

v,ok := m["hello"] -> v,ok := runtime.mapacess2(maptype,m , "hello")

m["hello"] = 12 -> v := runtime.mapassign(maptype,m , "hello")

delete(m,"hello") -> runtiem.mapdelete(maptype,m , "hello")

```

hmap 是一个 struct，这个结构体拥有众多字段，它用来描述这个 map 底层应用具有的所有字段。可以理解为它是描述 map 的头部文件，存储了所有的字段，但是并不存储真实的 body。

```go
type hmap struct {
	count     int 
	flags     uint8
	B         uint8  
	noverflow uint16 
	hash0     uint32 
	buckets    unsafe.Pointer  // 重点
	oldbuckets unsafe.Pointer 
	nevacuate  uintptr        
	extra *mapextra 
}

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap
	nextOverflow *bmap
}

// 这是一个桶本身的数据结构,但是在运行时阶段会重塑这个结构体，添加更多的字段
type bmap struct {
	tophash [bucketCnt]uint8
}
// 这是添加字段以后的样子
// 所以三个数组就完成了数据的存储，也可以避免padding的发生
type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
``` 
- count：当前 map 的 value 个数，len 返回的就是这个值
- flags:map 的状态标志 iterator oldterator hashwriting samesizegrow
- B：`2 ^ B = 桶数量`
- noverflow：指的是 overflow 的桶的数量
- hash0：哈希函数的种子值，用作哈希函数的参数，引入随机性
- buckets：指向桶数组的指针
- oldbuckets：在 map 扩容阶段指向前一个桶的指针
- nevacuate:map 扩容阶段充当扩容进度计数器，所有下标号小于 nevacuate 的桶都是已经完成了数据排空和迁移的操作的
- extra：【此字段是可选字段】如果有 overflow 的桶出现，该字段保证 overflow 的桶不会被 gc，具体操作就是该字段存储所有指向 overflow 的桶的指针

当声明一个 map 的时候，运行时就会为这个 map 具体的变量生成一个实例，上门那个是 map 的描述符号，这里说的这个数据结构是定义这个 map 中所有元信息的。

```go

type maptype struct {
	typ    _type
	key    *_type
	elem   *_type
	bucket *_type // 表示hash bucket 内部的类型
	// function for hashing keys (ptr to key, seed) -> hash
	hasher     func(unsafe.Pointer, uintptr) uintptr
	keysize    uint8  // size of key slot
	elemsize   uint8  // size of elem slot
	bucketsize uint16 // size of bucket
	flags      uint32
}
```
go 的 map 只有一套 method，只要使用这个 maptype 的具体不同实现就可以满足操作。比如 map[int]int 和 map[int]string，他们都使用 maptype 来定义自己的元信息，但是操作这这个元信息的函数是一个，只要取不同的 typ key 这种指针类型取操作就可以了。

存储数据的 body 本体是由一个 `类似二维数组` + `链表` 组成的，具体的图像如下：

![](./hash.svg)

tophash 区域就是为了寻找 key 和 value 使用空间换时间的原理做的索引，当我们把 hashcode 的高位值逐一比较的时候就可以确定 key 和 value 的位置。

key 和 value 都是连续的内存区域，key 和 value 在一个桶中只能存在 8 个，多余的在不满足扩容的情况下就会存储在溢出桶中，寻找 key 和 value 使用 tophash 的位置即可，go 使用分开存储 tophash，key，value 的方式，所以 go 就避免了数据 padding，满足了内存对齐，避免了内存的浪费。需要注意的是，当 key 和 value 大于 128 个字节的时候，存储的就是他们的指针

map 的扩容有一个具体的衡量指标，负载因子，当 hmap 中的 `count > 负载因子 * 2^B` (map 的负载因子为 6.5)，或者溢出桶过多的时候，就会扩容。当因为溢出桶太多的时候，创建的新 map 的桶和现有的桶一样，当因为不满足负载因子导致的扩容时，会创建两倍于现有 bucket 的新 bucket，但是旧的桶 data 并不会立刻被迁移到新的 bucket 中，在 map 进行插入和删除的过程中旧的内容会被逐步迁移到新的桶中。原来旧的内容就会被 gc 掉。

普通 map 的扩容并不是原子性的，所以 map 的扩容过程会去检查是否已经处于扩容中。

当持续向 map 中写入数据，并且删除，然后继续写入删除，这个时候如果没有造成装载因子的超出，就会造成溢出桶过多的情况，这个时候就会造成内存的泄漏，所以会创建一个一样大的 map，存储没有被删除的数据，并且将那些被标记为 delete 的数据进行垃圾回收，没错，你进行 delete 的时候并没有真的 delete，只有 gc 以后才是真的 delete。

将旧的数据放在 runtime.hmap 中的 oldbuckets 字段上，然后将新的数据结构放在 buckets 上，溢出桶的设置也是一样的，因为 extra 指向的数据结构也是三个字段，老的，新的，正在使用的。

运行时会将一个旧桶的数据分流到两个新的桶子里，但是如果是内存泄漏的等量扩容时，就只会把一个旧桶的数据导入到一个新的桶子里，请注意这里的迁移指的是拷贝，在计数器计算数据已经被分流完全以后，旧的桶和旧的溢出桶的数据就会被 gc 掉

因为存在旧桶和新桶，所以在查找数据的时候，会先从旧桶查找数据，如果没有再去新的桶中查找，当我们写入和删除数据的时候，除了写入的新数据到新的桶中，也会把旧的桶中的一部分数据拷贝到新的桶中，当然了不会拷贝全部，这是为了效率考虑的，

## issues
`问题一：` **map 元素可以取地址吗？**

[不能](../其他内容/README.md#go可寻址类型)，map 元素 (例如 ma [“12”]) 属于结果值，所以无法获取地址

`问题二：` **map 可以并发读写吗？可以 recover 吗？**

map 是线程不安全类型，读写得加互斥锁；被 recover 的 Panic 有几项是不能的：

- 数据竞争 (比如：**对 map 进行并发读写**)
> 可以通过 go 的编译标记 race 对代码进行检测是否存在数据竞争
- 内存不足出现的 Panic
- 死锁出现的 Panic

`问题三：` **sync.Map 适合的场景，和 map 加锁的区别**

sync.Map 在读多写少性能比较好，否则并发性能很差

> 有关 sync.Map 的[详细内容](../../工程/go标准库/sync.md#syncmap)

map 不支持并发的读或者是写 (go1.6 以后就不支持并发的读和写了，之前的版本支持并发的读，但是不支持并发的写)，所以 map+锁性能在读多写少和读少写多，读和写一样多的情况下是一样的。

最优解是使用多把锁即：分段锁的方式，并发读写，大幅提高性能

`问题四：` **在值为 nil 的字典上执行读操作会成功吗，那写操作呢？**

答案是在值为 nil 的字典写会 Panic，读是没问题的。

```go
package main

import "fmt"

func main() {
	var a map[int]int
	fmt.Println(a[1]) // 结果是int的一个初始值 0 
  a[0] = 1 // panic
}

```
`问题五：` **为什么不用向下寻址式？**

我们知道解决哈希碰撞的问题有向下寻址法，链表法，这是因为哈希函数只能把数据尽可能分布的均匀，如果哈希函数的输出的范围大于输入的范围，这是不现实的，这就要求映射无穷多，这显然不可能，所以必然会有两个不同的 key 算出来的哈希值是相同的，那么如果很多 key 算出的哈希值都是一样的，这就出现了查找效率从 O(1) 下降到了 O(n) 这就是所谓的哈希冲突。

> 这里提到的哈希值一样，有可能是后几位一样，也就是部分一样，比如后 9 位相同

向下寻址法的意思是：依次向下一位去探究是否是要找的哈希对，当然插入的时候也是这，向没有数据的地方插入，所以这种方法必须要使用数组这种结构，而且遍历的时候还得使用循环数组的这种思想 `index := hash("author") % array.len`，开放寻址法有一个数据指标叫做装载因子，就是说元素的个数/数组长度，一旦装载因子大于 70 %乃至 90 % 基本上就倒退为了 O(n) 的时间复杂度，并且底层是数组的情况下，必须使用连续的内存地址，并且数组长度是有限的，并且大概率会发生内存 padding 的情况，因为 kv 要存储在一起，这就又造成了更大的浪费。

总结一下：不使用向下寻址使用拉链法的原因在于，1 可以利用碎片式的内存，2 不用内存 padding 造成浪费 3 原则上链表长度无限，可以无限增加。

`问题六：` **map 如何操作真缩容？**

可以使用 cap 重新生成一个 map，然后使用遍历的方式将老的 map 数据导入到新的小的 map 中，如果你知道数据不会再次增大的情况下是可以这么做的。

## 参考资料
- https://github.com/golang/go/
- https://blog.csdn.net/EDDYCJY/article/details/120465701