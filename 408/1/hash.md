# 哈希表
## 哈希表
value通过哈希函数的计算，然后找到一个对应的数值，把这个数值存入到这个数组中，当然这个数组中可以是储存的一个链表这样的一个结构就是哈希表
## 哈希函数
把value计算成数值的这么的一个函数

比如举个例子，我们有一个value然后我们取这个string的每一个的字符，然后找到它对应的utf8对应的数字，然后我们将他们加和，然后把和跟30或者多少的取摸，得到的数字就是代表这个value要去的数组的位置。
```go
func HashFunc(s string)rune{
	var value rune
	for _,v := range s {
		value +=v
	}
	return value % 30
}
```
## 哈希碰撞
假如我们有value1 value2 然后通过这个哈希函数计算出来的值是相同的，那么这个就是哈希碰撞,所以我们该怎么办呢？

我们可以将这个数组中存放一个链表的头，然后存入这些相同的数字代表的k-v值，只要这个链表不长就基本上还是o（1）不过
假如有人要攻击这个hash，把链表搞的很长就可以拖垮这个服务器了。
## 针对hash的攻击
假如没有对string的大小作限制
```go

func BenchmarkHashFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashFunc(s) // s  很大很大
	}
}
func BenchmarkHashFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashFunc("0")
	}
}


```

output:

```go
23574 ns/op //
4.65 ns/op // 2
```
可以看出 如果有一个超级大的string很可能就直接拖垮了服务器
## map
这个就是最常见的哈希表了
## set
还记得redis中的set把。它就是使用的hash表来实现的 其中set和zset都不能有重复的数据，原因就是因为哈希表不能有重复的key
## hashmap hash set tree map tree set
hashmap和treemap或者是hashset和treeset他们的核心就是使用了hash表来实现还是使用二叉树来实现。

所以说map不一定是是哈希表来实现的，也可能底层是一个二叉树同理集合也是一个意思。

他们的区别也很容易，首先hash的方式是无序的，但是使用树的结构是可以 **有序的排列的** 但是哈希的查找的时间复杂度是o(1)，树的 查找的时候的时间复杂度是O(log2n),所以追求效率的话，肯定是哈希了，不过假如你要你的储存结构是有序的，那么可以采用树了

话说回来，其实map的储存是有序的反而很容易出现安全问题，所以说假如你需要安全的有序请参考redis中的zset，哈希+跳表 双结构就可以实现又安全又可以排序。

## 不加密的哈希函数算法

MurmurHash 是一种为了非加密的快速哈希函数，只要用途比如布隆过滤器。
