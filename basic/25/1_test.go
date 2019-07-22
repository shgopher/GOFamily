package main

import (
	"fmt"
	"testing"
)

var (
	b = new(BinarySearchTree)
)

func TestInsert(t *testing.T) {
	// 测试会不会生成root
	a1 := Insert(b, 1)
	fmt.Println("测试a1", *a1)
	// 测试普通的插入
	a11 := Insert(b, 2)
	a12 := Insert(b, 3)
	a13 := Insert(b, 4)
	a14 := Insert(b, 5)
	a2 := Insert(b, 6)
	fmt.Println("测试a2", *a2, *a11, *a12, *a13, *a14)
	// 测试是否会返回 -1
	a3 := Insert(b, 6)
	fmt.Println("测试a3", *a3)
}

func TestSearch(t *testing.T) {
	t1 := Search(b, 3)
	fmt.Println("测试search", *t1)
}
