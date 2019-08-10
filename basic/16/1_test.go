package main

import (
	"fmt"
	"testing"
)

func TestDPPackage(t *testing.T) {
	ZW = 12
	a = []int{1, 2, 3, 4, 5}
	result = 0
	status := make([][]int, len(a))
	for i := 0; i < len(status); i++ {
		status[i] = make([]int, ZW+1)
	}
	DPPackage(ZW, a, &result, status)
	fmt.Println(result)
}

// BenchmarkDPPackage-4   	10000000	       156 ns/op
func BenchmarkDPPackage(b *testing.B) {

	ZW = 12
	a = []int{1, 2, 3, 4, 5}
	result = 0
	status := make([][]int, len(a))
	for i := 0; i < len(status); i++ {
		status[i] = make([]int, ZW+1)
	}
	for i := 0; i < b.N; i++ {
		DPPackage(ZW, a, &result, status)
	}
}
