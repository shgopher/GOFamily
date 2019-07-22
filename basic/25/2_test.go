package main

import (
	"fmt"
	"testing"
)

var (
	arr = new(Box)
)

func TestInsertA(t *testing.T) {
	InsertM(arr, 10, "1")
	InsertM(arr, 10, "2")
	InsertM(arr, 10, "3")
	InsertM(arr, 10, "4")
	InsertM(arr, 10, "5")
	InsertM(arr, 10, "6")
	InsertM(arr, 10, "7")
	InsertM(arr, 16, "__")
	InsertM(arr, 16, "$##$F")
	InsertM(arr, 16, "/fsdf")
	InsertM(arr, 16, "rfdsd")
	InsertM(arr, 16, "fsdfef")
	InsertM(arr, 7, "71")
	InsertM(arr, 7, "74")
	InsertM(arr, 7, "756")
	InsertM(arr, 7, "7666")
	fmt.Println("测试 数组:", *arr)
}
func TestFind1(t *testing.T) {
	for _, v := range find1(arr, 10) {
		fmt.Println(v.k1, v.k2)
	}
	for _, v := range find1(arr, 7) {
		fmt.Println(v.k1, v.k2)
	}
	for _, v := range find1(arr, 16) {
		fmt.Println(v.k1, v.k2)
	}
}

// 10 1
// 10 1
// 10 2
// 10 3
// 10 4
// 10 5
// 10 6
// 10 7
// 7 71
// 7 74
// 7 756
// 7 7666
// 16 __
// 16 $##$F
// 16 /fsdf
// 16 rfdsd
// 16 fsdfef

//  从结果来看，这种可以加入重复变量的二叉查找树，可以快读读出某个值相同的一类对象，然后我们可以得到这些对象的其它的值，这些值是不一样的，这就是它的意义。g
