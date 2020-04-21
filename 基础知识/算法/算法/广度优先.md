# 广度优先搜索 BFS

首先明确一个概念，广度优先搜索的最大适用范围是图，单常常用在树的上面。

广度优先搜索的概念就是按层级搜索，以广度的方式搜索
```go
     8
  7     9
1  4  5    19
```
按照 8 7 9  1 4 5 19的方式开始搜索，按照这种方式搜索就是广度优先搜索，描述一下基本的原理就是先想问root节点，然后把子节点按照一定的顺序加入到队列中，每次遍历root的时候都加入他们的子节点（但不加入孙子节点）这样就只需要遍历这个队列即可。

下面我们看一下搜索的代码原理解释

```go
func (n *Node) BFS(value interface{}) *Node {
	// 定义
	ma := make(map[*Node]int)
	result := &Node{}
	queue := NewQueue(10)
	// 首行辅助任务完成
	queue.Push(n)
	ma[n]++
	// 开始循环寻找
L:
	for queue.Length() > 0 {
		v := queue.Pop()
		val := v.(*Node)
		ma[val]++
		if val.Value == value {
			result = val
			break L
		}
		for _, v := range val.Child {
			if _,ok := ma[v];!ok {
				queue.Push(v)
			}
		}

	}
	return result
}
```
