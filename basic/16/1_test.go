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
