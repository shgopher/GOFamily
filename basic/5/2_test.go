package main

import (
	"fmt"
	"testing"
)

func TestPait(t *testing.T) {
	h := NewHeap(20)
	fmt.Println("dddd", HeapPai([]int{0, 1, 3, 4, 5, 6, 2, 89, 87, 54}, h))
}
