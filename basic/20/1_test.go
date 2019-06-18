package main

import (
	"fmt"
	"testing"
)

// 测试的原则 1 空值 2 边界左右↔️ 3 错误值 4 普通值
func TestBower(t *testing.T) {
	// 测试 使用数组的方式
	//
	// 普通值
	a1 := &StackSlice{}
	a2 := &StackSlice{}
	b := NewBower(a1, a2)
	b.GO("1")
	b.GO("2")
	b.GO("3")
	b.GO("4")
	b.GO("5")
	b.GO("6")
	b.GO("7")
	b.GO("8")
	b.Back()
	b.Back()
	b.Back()
	b.Back()
	b.Back()
	fmt.Print(b.Back())
	fmt.Println(b.GO("测试"))

	// 空
	a21 := &StackSlice{}
	a22 := &StackSlice{}
	b2 := NewBower(a21, a22)
	fmt.Println(b2.Back())

	//
	// 测试 使用 单链表的方式

	// 普通值
	b1 := &StackLinkedList{}
	b2 := &StackLinkedList{}
	c := NewBower(b1, b2)
	c.GO("1")
	c.GO("2")
	c.GO("3")
	c.GO("4")
	c.GO("5")
	c.GO("6")
	c.GO("7")
	c.GO("8")
	c.Back()
	c.Back()
	c.Back()
	c.Back()
	c.Back()
	fmt.Print(c.Back())
	fmt.Println(c.GO("测试"))

	// 空
	b21 := &StackSlice{}
	b22 := &StackSlice{}
	c2 := NewBower(b21, b22)
	fmt.Println(c2.Back())

}
