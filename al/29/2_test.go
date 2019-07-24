package main

import (
	"fmt"
	"testing"
)

func TestTOPK(t *testing.T) {
	h := NewHeap(11)
	h.Insert1(4)
	h.Insert1(23)
	h.Insert1(45)
	h.Insert1(33652)
	h.Insert1(564)
	h.Insert1(643)
	h.Insert1(34)
	h.Insert1(234)
	fmt.Println("heap---???????????>>>", TOPK(3, h))
}
