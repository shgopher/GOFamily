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

// 广度优先搜索
func (g *Graph) BFS(a, b int) {
	if a == b {
		return
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
