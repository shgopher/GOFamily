//LRU 链表实现方法：
package main

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	length int
	head   *Node
}



func (l *LinkedList) MoveTOHead(this *Node) {
	next := this.next
	now := l.head.next
	for now.next != nil {
		if now.next == this {
			break
		}
	}
	now.next = next
	this.next = l.head.next
	l.head.next = this
}

func (l *LinkedList) Delete() {
	cur := l.head.next
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
}
