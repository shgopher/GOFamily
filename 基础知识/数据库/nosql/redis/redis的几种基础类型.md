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
```bash
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
```
## hash
## set
## zset
