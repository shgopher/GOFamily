//使用数组
// 要实现的功能 包括 1 顺序入出 2 循环式入出 3 阻塞式 4 并发式（包括了两种 原子操作和加锁）
package main

import (
	"fmt"
	"sync"
	"time"
)

// 顺序
type QueueSliceSequence struct {
	lock   sync.Mutex
	trail  int
	head   int
	length int
	value  []interface{}
}

func NewQueueSliceSequence(num int) *QueueSliceSequence {
	value := make([]interface{}, num+1, num+1)
	return &QueueSliceSequence{
		trail:  0,
		head:   0,
		length: 0,
		value:  value,
	}
}

func (q *QueueSliceSequence) In(v interface{}) {
	if q.length >= len(q.value)-1 {
		return
	}
	if q.trail == q.head && q.head != 0 {
		return
	}
	q.length++
	q.value[q.length] = v
	q.trail++
}
func (q *QueueSliceSequence) Out() (v interface{}) {
	if q.head == q.trail {
		return fmt.Errorf("没有数据")
	} else if q.length == 1 && q.head == 0 {
		out := q.value[q.length]
		q.trail = 0
		q.length--
		return out
	}
	now := q.value[q.head+1]
	q.head++
	q.length--
	return now
}

// 使用循环slice来实现一个队列
type QueueSliceCycle struct {
	lock   sync.Mutex
	trail  int
	head   int
	length int
	value  []interface{}
}

func NewQueueSliceCycle(num int) *QueueSliceCycle {
	now := make([]interface{}, num+1, num+1)
	return &QueueSliceCycle{
		trail:  0,
		head:   0,
		length: 0,
		value:  now,
	}
}
func (q *QueueSliceCycle) In(v interface{}) {
	if (q.trail+1)%len(q.value) == q.head {
		return
	}
	q.length++
	now := (q.trail + 1) % len(q.value)
	q.value[now] = v
	q.trail = now

}
func (q *QueueSliceCycle) Out() (v interface{}) {
	if q.head == q.trail {
		return fmt.Errorf("无法取出数据，这个循环队列空了")
	}
	now := (q.head + 1) % len(q.value)
	q.length--
	q.head = now
	return q.value[now]
}

// 阻塞
type QueueSliceCrimp struct {
	lock   sync.Mutex
	trail  int
	head   int
	length int
	value  []interface{}
}

func (q *QueueSliceCrimp) In(v interface{}) {
L:
	for {
		if (q.trail+1)%len(q.value) == q.head {

		} else {
			break L
		}
		time.Sleep(time.Second / 10)
	}

	q.length++
	now := (q.trail + 1) % len(q.value)
	q.value[now] = v
	q.trail = now

}
func (q *QueueSliceCrimp) Out() (v interface{}) {
L:
	for {
		if q.head == q.trail {

		} else {
			break L
		}
		time.Sleep(time.Second / 100)
	}

	now := (q.head + 1) % len(q.value)
	q.length--
	q.head = now
	return q.value[now]

}

// 并发加锁
// 只需要 加上 sync.Mutext 即可。

// 并发原子操作

// 使用循环slice来实现一个队列
//type QueueSliceAtom struct {
//	lock   sync.Mutex
//	trail  int
//	head   int
//	length int
//	value  []interface{}
//}
//
//func (q *QueueSliceAtom) In(v interface{}) {
//
//}
//func (q *QueueSliceAtom) Out() (v interface{}) {
//
//}
