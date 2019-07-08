//使用数组
// 要实现的功能 包括 1 顺序入出 2 循环式入出 3 阻塞式 4 并发式（包括了两种 原子操作和加锁）
package main

import "sync"

// 使用循环slice来实现一个队列
type QueueSlice struct {
	lock sync.Mutex
	trail  int
	head   int
	length int
	value  []interface{}
}

