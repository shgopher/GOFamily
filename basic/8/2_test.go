package main

import (
	"fmt"
	"testing"
)

var (
	chaC = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
)

// 8
func TestER1(t *testing.T) {
	fmt.Println("ER1", ER1(chaC, 5))
}

// 9
func TestER2(t *testing.T) {
	fmt.Println("ER2", ER2(chaC, 5))
}

// //查找第一个大于等于给定值的元素
//8
func TestER3(t *testing.T) {
	fmt.Println("ER3", ER3(chaC, 5))
}

// //查找最后一个小于等于给定值的元素
//9
func TestER4(t *testing.T) {
	fmt.Println("ER4", ER4(chaC, 5))
}
