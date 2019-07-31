// 深度优先搜索
// 深度优先搜索的原理就是回溯，一股道走到底，走不通了就往上返回一级，如果全部都行不通就再往上返回一级。
package main

// 深度优先搜索
func (g *Graph) DFS(a, b int) {

	if a == b {
		return
	}
	pre := make([]int, g.v)
	for i := range pre {
		pre[i] = -1
	}

	visited := make([]bool, g.v)
	visited[a] = true
	isFound := false
	g.R(a, b, pre, visited, isFound)
	g.Range(pre, a, b)

}

func (g *Graph) R(a int, b int, pre []int, visited []bool, isFound bool) {
	if isFound {
		return
	}
	visited[a] = true
	if a == b {
		isFound = true
		return
	}
	list := g.adj[a]
	for i := list.Front(); i != nil; i = i.Next() {
		value := i.Value.(int) // 链表中的内容
		if !visited[value] {
			pre[value] = a
			g.R(value, b, pre, visited, false) // 这里其实是一个多叉树，当一个顶点走投无路的时候就停止了（走投无路就是它里面的链表中的顶点都被访问过了，那么这一个叉就gg了）
		}
	}
}
