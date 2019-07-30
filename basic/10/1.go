// 广度优先搜索
package main

import "container/list"

// adj -- adjacency table 邻接表
type Graph struct {
	v   int          // 顶点的个数
	adj []*list.List // 邻接表
}
