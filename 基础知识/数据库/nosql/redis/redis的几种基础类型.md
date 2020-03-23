# redis的几种基本的类型
## string
redis的string是动态数组，也就是说它字符串占的长度，要小于它底层数组的容量，这一点跟go里面的slice的扩容还是有点像的，它也有类似的原理，它扩容的时候会判断字符串的长度，当长度小于1m的时候每次扩容都是容量翻倍，当大于1m的时候，每次扩容增加一个1m的空间，但是有大小限制，redis中的字符串最大值是512m

redis中采用这种冗余的预处理机制来扩容主要是为了防止频繁的内存申请，内存的分配是很浪费时间的。

**基础操作**

```shell
// 设置一个string // 可以同时两次set一个string，就是把它直接颠覆了，
set googege redis
// 得到一个string
get googege
// 获取字符串的长度
strlen googege
// 对string做切片,getrange string fitst-index end-index
// 并且必须提供后面的两个起始和结束的index，否则报错。
getrange googege 1 2
// 覆盖字符串 setrange string first-index new string.
setrange googege 1 ddddddd
// 追加字符串 append string xxx
append googege ttt

```
从以上操作发现了一个问题，并没有提供字符串的 **子串** 插入（提供的只有覆盖）和删除的操作，意思就是比如在3 5之间插入一个子串，或者是删除5-7之间的子串。

**计数器**

当字符串是整数的时候，可以将它当成计数器

```shell
// 设置一个计数器 set key int
set aa 1
// 得到一个计数器 get key
get aa
// 增加一个固定的值 incrby key number
incrby aa 100
// 减去一个固定的值 decrby key number
decrby aa 100
// 每次增加1 incr key
incr aa
// 每次减去1 decr key
decr aa
```
注意计数器有返回，它的区间是 `[long.MIN,long.Max]`全闭区间
> [-9223372036854775807,9223372036854775807]

错误提示是` ERR increment or decrement would overflow`

**字符串的删除和自动删除(过期)**
```shell
//删除字符串 del key
del aa
// 设置过期时间(单位是s) expire key time
expire aa 60
// 查看寿命 ttl key
ttl aa
```
## list
list的底层是一个双向链表，所以可以使用这个链表来实现队列或者stack的功能
```shell
// rpush rpop lpush lpop 分别是 右边推入 右边推出 左边推入 左边推出

// 实现栈的功能
rpush aaa go  java tt
rpop aaa
rpop aaa
rpop aaa
//or
lpush aaa go java tt
rpop aaa
rpop aaa
rpop aaa
// 实现队列的功能
rpush aaa go java tt
lpop aaa
lpop aaa
lpop aaa
//or
lpush aaa go java tt
rpop aaa
rpop aaa
rpop aaa
// 获取链表的长度
//llen key
llen aaa
```
我们可以获取list的子链

```shell
rpush aaa go rust dart2
lindex aaa 1 // 获取aaa的第2个元素
lrang aaa 0 2 // 获取子链
lrang aaa 0 -1 // -的意思就是倒着数
// 遍历全部的数据
lrang 0 -1 // 这个时候使用-数不用使用llen了，也是极好的。
```
**修改和插入元素**
```shell
rpush aa go java rust
// 修改
lset aa 1 t // 变成了  go t rust

// 插入数据
linsert aa before go tt// tt go java rust
linsert aa after go tt // go tt java rust
```
这里要注意，插入数据不是根据的下标的顺序，因为redis经常用在分布式的环境中，那么分布式中的下标就没有意义了，所以特别的指定在某个元素后面或者前面插入xx元素。
**删除元素**
```shell
rpush aa go rust dd
lrem aa 1 go // 删除的时候不仅仅要写出来个数还要写出来元素的值
```
**固定长度的列表**
```shell
rpush aa 1 2 3 4 5 6 7 8 9 10 // 往队列里添加数据
ltrim 0 8 // 标注起始和结束即可。
lrange aa 0 -1
1) "1"
2) "2"
3) "3"
4) "4"
5) "5"
6) "6"
7) "7"
8) "8"
// 小技巧
ltrim -1 0 // 因为end比start小，那么这个list就会被删除光。
```
### 探究list的链表的真实底层
![p](./1.1.png)

当数据很少的时候只是用ziplist，当数据起来了才会采用quicklist，Ziplist 是由一系列特殊编码的内存块构成的列表
### ziplist

## hash
## set
## zset
