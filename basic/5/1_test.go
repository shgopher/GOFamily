package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	h := NewHeap(7)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	h.Insert(6)
	h.Insert(7)
	fmt.Println("插入", h.value[1:h.count+1])

}

func TestDelete(t *testing.T) {
	h := NewHeap(7)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	h.Insert(6)
	h.Insert(7)
	fmt.Println("删除前", h.value[1:h.count+1])
	h.Delete()
	fmt.Println("删除后", h.value[1:h.count+1])
}

func TestInser1(t *testing.T) {
	h := NewHeap(31)
	for i := 1; i <= 31; i++ {
		h.Insert(i)
	}
	fmt.Println(h.value[1 : h.count+1])
}

func TestInser2(t *testing.T) {
	h := NewHeap(31)
	h.Insert1(10)
	h.Insert1(1)
	h.Insert1(3)
	h.Insert1(4)
	h.Insert1(19)
	h.Insert1(5)
	h.Insert1(6)
	h.Insert1(9)
	h.Insert1(65)
	h.Insert1(654)
	h.Insert1(44)
	h.Insert1(32)
	h.Insert1(2)
	h.Insert1(1)

	fmt.Println("....", h.value[1:h.count+1])
}

// 插入 [7 4 6 1 3 2 5]
// 删除前 [7 4 6 1 3 2 5]
// 删除后 [6 4 5 1 3 2]

//      56.1 ns/op
func BenchmarkInsert(b *testing.B) {
	h := NewHeap(b.N)
	for i := 0; i < b.N; i++ {
		h.Insert(i)
	}
}

//  178 ns/op
func BenchmarkDelete(b *testing.B) {
	h := NewHeap(b.N)
	for i := 0; i < b.N; i++ {
		h.Insert(i)
	}
	for i := 0; i < b.N; i++ {
		h.Delete()
	}
}
