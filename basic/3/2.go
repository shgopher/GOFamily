// 使用链表
// 要实现的功能 包括 1 顺序入出 2 循环式入出 3 阻塞式 4 并发式（包括了两种 原子操作和加锁）
package main

import "sync"

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

