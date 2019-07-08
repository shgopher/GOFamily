//使用数组
package main

type Queue interface {
	// 进去，进队列
	In(v interface{})
	// 出队列
	Out() (v interface{})
}

// 使用循环slice来实现一个队列
type QueueSlice struct {
	trail  int
	head   int
	length int
	value  []interface{}
}

// 使用循环链表即可
type QueueLinkedList struct {
	trail  *node
	length int
	head   *node
}

type node struct {
	value interface{}
	next  *node
	last  *node
}
