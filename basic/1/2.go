// 实现单向链表的反转
package main

func (l *LinkedList) Reverse() {
	now := l.head.next
	var pre *sNode = nil
	for now != nil {
		tem := now.next // 已经记录好了下一位
		now.next = pre // 将 now.next 改写成 前一个now
		pre = now  // pre 就是前一个now
		now = tem // now 等于下一位
	}
	l.head.next = pre
}
