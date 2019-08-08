package main

import (
	"fmt"
	"testing"
)

var (
	maxWgith1 = 17
	maxWgith2 = 14
	a1        = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2        = []int{0, 7, 5, 8}
)

func TestBei(t *testing.T) {
	result := 0
	Bei(0, a1, 0, maxWgith1, &result)
	fmt.Println(result)
}

func TestBei1(t *testing.T) {
	result := 0
	Bei(0, a2, 0, maxWgith2, &result)
	fmt.Println(result)
}

// 17
// 13
