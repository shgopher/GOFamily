package main

import "testing"

func TestDFS(t *testing.T) {
	g := NewGraph(100)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 4)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)
	g.DFS(0, 5)
}
