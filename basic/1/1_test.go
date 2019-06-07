package main

import (
	"testing"
)

var list = newLinkedList(0)

func TestRange(t *testing.T) {
	list.AddTrail(12)
	list.AddTrail(12222)
	list.Range()
	list.Delete(list.Search((1)))
	list.Range()
	list.Add(list.Search(0), "哈哈哈哈")
	list.Range()
}

var listD = newListD(00)

func TestLinkedList_Range(t *testing.T) {

	listD.addTrail(1)
	listD.addTrail(2)
	listD.addTrail(3)
	listD.Range()
	listD.Delete(listD.Search(1))
	listD.Range()
}
