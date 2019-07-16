package main

import (
	"fmt"
	"testing"
)

func TestER(t *testing.T) {
	fmt.Println(ER([]int{1, 3, 5, 7, 8, 9, 10}, 5))
}
func TestERD(*testing.T) {
	fmt.Println(ERD([]int{1, 3, 5, 7, 8, 9, 10}, 5))
}

//  8.05 ns/op
func BenchmarkER(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ER([]int{1, 3, 5, 7, 9, 11, 15, 16, 156}, 16)
	}
}

func BenchmarkERD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ERD([]int{1, 3, 5, 7, 9, 11, 15, 16, 156}, 16)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/googege/AMAC/basic/8
// BenchmarkER-4    	200000000	         8.29 ns/op
// BenchmarkERD-4   	100000000	        16.3 ns/op
