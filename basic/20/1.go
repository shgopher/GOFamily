package main

// 1 使用链表来实现栈
// 2 使用数组（切片）来实现栈

// 总之只要满足 FILO就ok了。

//首先我们要实现 栈的 出栈和入栈，以及查阅栈顶这三个功能，然后我将使用两种方式进行选择，然后我们再实现 前进和后退这个功能

// 数组形式的栈
type StackSlice struct {
	value  []interface{}
	length int
}

// 节点
type node struct {
	value interface{}
	next  *node
}

// 链表形式的栈
type StackLinkedList struct {
	length int
	head   *node
}

type Bower interface {
	In()   // 将数据压进去
	Out()  // 将数据取出来
	Next() // 前进
	Back() // 后退
}

func GO(b Bower) {
	b.Next()
}

func Back(b Bower) {
	b.Back()
}

// new 一个新的SatckSlice
func NewSlice(c int) *StackSlice {
	return &StackSlice{
		length: 0,
		value:  make([]interface{}, c, c), // 设置应有的容量大小。
	}
}

// new 一个新的StackLinkedList
func NewList() *StackLinkedList {
	return &StackLinkedList{
		length: 0,
		head: &node{ // 第一个是哨兵，使用哨兵机制。
			value: 0,
			next:  nil,
		},
	}
}

// 数组类型实现栈
func (sli *StackSlice) In() {
	sli.length++

}
func (sli *StackSlice) Out() {
	if sli.length > 0 {
		sli.length--
	} else {
		sli.length = 0
	}

}
func (sli *StackSlice) Next() {

}
func (sli *StackSlice) Back() {

}

// 链表类型实现栈
func (sli *StackLinkedList) In() {

}
func (sli *StackLinkedList) Out() {

}
func (sli *StackLinkedList) Next() {

}
func (sli *StackLinkedList) Back() {

}
