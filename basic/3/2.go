// 使用链表
// 要实现的功能 包括 1 顺序入出 因为链表可以没有一个范围，所以链表实现的队列叫做无边界队列
package main

import (
	"fmt"
	"sync"
)

// 使用 数组和链表 公用的一个interface。
type Queue interface {
	// 进去，进队列
	In(v interface{})
	// 出队列
	Out() (v interface{})
}

// 链表
type QueueLinkedList struct {
	lock sync.Mutex
	trail  *node
	length int
	head   *node
}

type node struct {
	value interface{}
	next  *node
	last  *node
}
// 初始化一个链表式的队列
func NewQueueLinkedList() *QueueLinkedList{
	return &QueueLinkedList{
		length:0,
		head:&node{
			value:0,
			last:nil,
			next:nil,
		},
		trail:nil,
	}
}
func(q *QueueLinkedList)In(v interface{}){
	q.length++
}

func(q *QueueLinkedList)Out()(v interface{}){
	if q.length <=0 {
		return fmt.Errorf("错误，链表队列中没有数据。")
	}
	q.length--
}

