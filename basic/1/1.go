//实现单向,双向，循环链表以及增删操作
package main

import "fmt"

//单链表（以及单循环链表的节点对象）
type sNode struct {
	next  *sNode
	value interface{}
}

// 单链表本身
type LinkedList struct {
	head   *sNode
	length int
}

// new 一个新的链表
func newLinkedList(ele interface{}) *LinkedList {
	return &LinkedList{
		head: &sNode{
			nil,
			ele,
		},
		length: 0,
	}
}

// 新增加一个节点
func (l *LinkedList) Add(thisNode *sNode, element interface{}) {
	newOne := new(sNode)
	newOne.value= element
	newOne.next = thisNode.next
	thisNode.next = newOne
	l.length++
}

// 新增节点到末尾
func (l *LinkedList) AddTrail(element interface{}) {
	now := l.head
	for now.next != nil {
		now = now.next
	}
	newOne := new(sNode)
	newOne.value = element
	now.next = newOne
}

// 删除一个节点
func (l *LinkedList) Delete(thisNode *sNode) {
	now := l.head
	for now.next != nil {
		if now.next == thisNode {
			break
		}
		now = now.next
	}
	thisNode = now.next
	if thisNode.next != nil {
		now.next = thisNode.next
	} else {
		now.next = nil
	}
	l.length--
}

// 按照len查找一个节点
func (l *LinkedList) Search(len int) *sNode {
	now := l.head
	lenNow := 0
	for lenNow < len {
		now = now.next
		lenNow++
	}
	return now
}

// 按照pre element 寻找它的下一个的节点
func (l *LinkedList) PreSearch(element interface{}) *sNode {
	now := l.head
	for now != nil {
		if now.value == element {
			return now
		}
		now = now.next
	}
	return nil
}

// range 出整个的linkedList

func (l *LinkedList) Range() {
	now := l.head
	for now.next != nil {
		fmt.Print(now.value)
		now = now.next
	}
	fmt.Println(now.value)
}

/*

下面👇  是双链表的基本操作


*/

// 双链表 （以及循环双链表的节点对象）
type dNode struct {
	prev  *sNode
	next  *sNode
	value interface{}
}

// 双链表本身
type LinkedListD struct {
	head   *dNode
	trail  *dNode
	length int
}
