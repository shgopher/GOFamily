# 链表

链表使用的非常广泛，比如你又一个需求，你并不知道到底有多少个数据，而且需要数据不断的叠加，你可以使用链表来处理，比特币或者区块链中，一个一个的区块连接的起来，本质上其实就是一种链表。链表形如一个一个的火车中间被铁链子连接起来。而且，我们通常要指定这个链表中的头和尾部的元素。

**链表插入操作**

```go
type Linkedlist struct {
  Head *Node
  Trail *Node
  Length int

}
type Node struct {
  Value interface{}
  Pre *Node
  Next *Node
}
func (l *Linkedlist)Insert(value interface{}) {
// ...
    next := now.Next
    now.Next = new(Node)
    now.Next.Value = value
    now.Next.Pre = now
    now.next.next = next.Pre
// ...
}
```
**链表删除操作**

```go
func(l *Linkedlist)Del(){
  // ...
next := now.Next.Next
now.Next = next
next.Pre  = now
// ...
}
```

对于链表来说插入和删除的确是1的时间复杂度，但是，查找链表的时候耗费的时间就是n的时间复杂度来，所以说查找数组快，插入删除链表快。
