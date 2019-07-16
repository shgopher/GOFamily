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
	fmt.Println("ER1-0", ER1(chaC, 0))
	fmt.Println("ER1-1", ER1(chaC, 1))
	fmt.Println("ER1-2", ER1(chaC, 2))
	fmt.Println("ER1-3", ER1(chaC, 3))
	fmt.Println("ER1-4", ER1(chaC, 4))
	fmt.Println("ER1-5", ER1(chaC, 5))
	fmt.Println("ER1-6", ER1(chaC, 6))
	fmt.Println("ER1-7", ER1(chaC, 7))
}

// 9
func TestER2(t *testing.T) {
	fmt.Println("ER2-0", ER2(chaC, 0))
	fmt.Println("ER2-1", ER2(chaC, 1))
	fmt.Println("ER2-2", ER2(chaC, 2))
	fmt.Println("ER2-3", ER2(chaC, 3))
	fmt.Println("ER2-4", ER2(chaC, 4))
	fmt.Println("ER2-5", ER2(chaC, 5))
	fmt.Println("ER2-6", ER2(chaC, 6))
	fmt.Println("ER2-7", ER2(chaC, 7))
}

// //查找第一个大于等于给定值的元素
//8
func TestER3(t *testing.T) {
	fmt.Println("ER3-0", ER3(chaC, 0))
	fmt.Println("ER3-1", ER3(chaC, 1))
	fmt.Println("ER3-2", ER3(chaC, 2))
	fmt.Println("ER3-3", ER3(chaC, 3))
	fmt.Println("ER3-4", ER3(chaC, 4))
	fmt.Println("ER3-5", ER3(chaC, 5))
	fmt.Println("ER3-6", ER3(chaC, 6))
	fmt.Println("ER3-7", ER3(chaC, 7))
}

// //查找最后一个小于等于给定值的元素
//9
func TestER4(t *testing.T) {
	fmt.Println("ER4-0", ER4(chaC, 0))
	fmt.Println("ER4-1", ER4(chaC, 1))
	fmt.Println("ER4-2", ER4(chaC, 2))
	fmt.Println("ER4-3", ER4(chaC, 3))
	fmt.Println("ER4-4", ER4(chaC, 4))
	fmt.Println("ER4-5", ER4(chaC, 5))
	fmt.Println("ER4-6", ER4(chaC, 6))
	fmt.Println("ER4-7", ER4(chaC, 7))
}

// ER1-0 -1
// ER1-1 0
// ER1-2 2
// ER1-3 4
// ER1-4 6
// ER1-5 8
// ER1-6 10
// ER1-7 -1
// ER2-0 -1
// ER2-1 1
// ER2-2 3
// ER2-3 5
// ER2-4 7
// ER2-5 9
// ER2-6 11
// ER2-7 -1
// ER3-0 0
// ER3-1 0
// ER3-2 2
// ER3-3 4
// ER3-4 6
// ER3-5 8
// ER3-6 10
// ER3-7 -1
// ER4-0 -1
// ER4-1 1
// ER4-2 3
// ER4-3 5
// ER4-4 7
// ER4-5 9
// ER4-6 11
// ER4-7 11
//

// 所有情况都测试完毕，代码正确
