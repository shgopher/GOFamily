// 实现两个有序链表合并为一个有序链表
package main


func Together(list1,list2 *LinkedList)*LinkedList{
	now1 := list1.head
	now2 := list2.head
	l := new(LinkedList)
	now := l.head
	for now1.next != nil && now2.next != nil {
		value1 := now1.value.(int)
		value2 := now2.value.(int)
		if value1 <= value2 {
			now.next = now1
			now1 = now1.next
		}else{
			now.next = now2
			now2 = now2.next
		}
		now = now.next
	}
	if now1.next != nil {
		now.next = now1
	}else if now2.next != nil {
		now.next = now2
	}
	return l
}
