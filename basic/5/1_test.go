package main

import (
	"fmt"
	"testing"
)

var (
	h *Heap
)

func TestInsert(t *testing.T) {
	h = NewHeap(100)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	h.Insert(6)
	h.Insert(7)
	fmt.Println(h.value)

}

func TestRange(t *testing.T) {
	fmt.Println(h.Range())
}
