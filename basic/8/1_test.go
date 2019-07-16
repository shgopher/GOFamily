package main

import (
	"fmt"
	"testing"
)

func TestER(t *testing.T) {
	fmt.Println(ER([]int{1, 3, 5, 7, 8, 9, 10}, 5))
}

//  8.05 ns/op
func BenchmarkER(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ER([]int{1, 3, 5, 7, 9, 11, 15, 16, 156}, 16)
	}
}
