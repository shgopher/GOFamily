package main

import (
	"fmt"
	"testing"
)

var (
	b = new(BinarySearchTree)
)

// 测试插入数据
func TestInsert(t *testing.T) {
	// 测试会不会生成root
	a1 := Insert(b, 1)
	fmt.Println("测试a1", *a1)
	// 测试普通的插入
	a11 := Insert(b, 20)
	a12 := Insert(b, 13)
	a13 := Insert(b, 4)
	a14 := Insert(b, 15)
	a2 := Insert(b, 6)
	fmt.Println("测试a2", *a2, *a11, *a12, *a13, *a14)
	// 测试是否会返回 -1
	a3 := Insert(b, 6)
	fmt.Println("测试a3", *a3)
}

// 测试 搜索
func TestSearch(t *testing.T) {
	t1 := Search(b, 3)
	fmt.Println("测试search", *t1)
}

// delete 写的不对
// 测试 删除数据
// func TestDelete(t *testing.T) {
// 	Delete(b, 15)
// 	t1 := Search(b, 15)
// 	fmt.Println("测试删除数据5应该不在了", t1)
// }

// func TestDelete1(t *testing.T) {
// // 	Delete1(b, 15)
// // 	t1 := Search(b, 15)
// // 	fmt.Println("结果应该是0,因为已经无法搜索到15了", t1)
// // }

func TestMRange(t *testing.T) {
	fmt.Println(MRange(b))
}
