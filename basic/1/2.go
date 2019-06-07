// 实现单向链表的反转
package main

func(l *LinkedList ) turn(){
	now := l.head
	for now.next != nil {
		now = now.next
	}
}

