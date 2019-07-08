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
	var now *node
	if q.trail != nil {
		now = &node{
			value: v,
			last:  q.trail,
			next:  nil,
		}
		q.trail.next = now

	}else {
		now = &node{
			value:v,
			last:q.head,
			next:nil,
		}
		q.head.next=now
		now.last = q.head
	}
	q.trail = now
}
//
func(q *QueueLinkedList)Out()(v interface{}){
	if q.length <=0 {
		return fmt.Errorf("错误，链表队列中没有数据。")
	}else if q.length == 1 {
		q.length--
		q.trail = nil
		return q.head.next.value
	}
	out := q.head.next
	now := out.next
	now.last = q.head
	q.head.next = now
	q.length--
	return out.value
}

