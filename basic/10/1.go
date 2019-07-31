// 广度优先搜索
package main

import (
	"container/list"
	"fmt"
)

// a表示启示顶点 b表示终止顶点，搜索算法查找的是ab两个顶点之间的路径

// adj -- adjacency table 邻接表
type Graph struct {
	v   int          // 顶点的个数
	adj []*list.List // 邻接表-
}

// 初始化这个图
func NewGraph(v int) *Graph {
	adj := make([]*list.List, v, v)
	for i := 0; i < v; i++ {
		adj[i] = list.New()
	}
	return &Graph{
		v:   v,
		adj: adj,
	}
}

// 往图里增加边
// 增加 a b 之间的边
func (g *Graph) AddEdge(a, b int) {
	g.adj[a].PushBack(b)
	g.adj[b].PushBack(a)
}

// 广度优先搜索 关键词 visited pre queue
func (g *Graph) BFS(a, b int) {
	var queue []int
	if a == b {
		return
	}
	pre := make([]int, g.v)
	visited := make([]bool, g.v)
	for k := range g.adj { // 将所有的顶点都设置为未访问
		pre[k] = -1
	}
	visited[a] = true
	isFound := false

	for len(queue) > 0 && !isFound {
		top := queue[0]                                 // 第一个顶点。
		list := g.adj[top]                              // 这个是获取第一个顶点的链表
		queue = queue[1:]                               // 队列队头出列
		for i := list.Front(); i != nil; i = i.Next() { // 这个是为了遍历每个顶点中的链表中的顶点，然后将这些个顶点加入这个队列中。
			value := i.Value.(int)
			if !visited[value] {
				pre[value] = top
				if value == a {
					isFound = true
					break
				}
				queue = append(queue, value)
				visited[value] = true
			}

		}
	}
	if !visited {
		fmt.Println("找不到这个路径")
	} else {
		g.Range(pre, a, b)
	}

}

// 深度优先搜索
func (g *Graph) DFS() {

}

// 遍历搜索路径
func (g *Graph) Range(pre []int, a, b int) {
	if a == b || pre[b] == -1 {
		fmt.Println(b)
	} else {
		g.Range(pre, a, pre[b]) // 往前找路径。递归到a == b 或者 pre[b] = -1 就会跳出来。
		fmt.Println(b)
	}
}
