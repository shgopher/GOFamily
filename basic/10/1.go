// 广度优先搜索
package main

import "container/list"

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
