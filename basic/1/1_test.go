package main

import (
	"testing"
)

var list = newLinkedList(0)

func TestRange(t *testing.T) {
	list.AddTrail(1)
	list.AddTrail(2)
	list.AddTrail(3)
	list.AddTrail(4)
	list.AddTrail(5)
	list.Range()
	list.Reverse()
	list.Range()
}

var listD = newListD()

func TestLinkedList_Range(t *testing.T) {

	listD.addTrail(1)
	listD.addTrail(2)
	listD.addTrail(3)
	listD.Range()
	listD.Delete(listD.Search(1))
	listD.Range()
}
