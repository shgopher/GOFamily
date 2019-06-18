package main

// 1 使用链表来实现栈
// 2 使用数组（切片）来实现栈

// 总之只要满足 FILO就ok了。

//首先我们要实现 栈的 出栈和入栈，以及查阅栈顶这三个功能，然后我将使用两种方式进行选择，然后我们再实现 前进和后退这个功能

// 数组形式的栈
type StackSlice struct {
	value  []interface{}
	length int
}

// 节点
type node struct {
	value interface{}
	next  *node
}

// 链表形式的栈
type StackLinkedList struct {
	length int
	head   *node
}

type Stack interface {
	Top() (url string)
	In(url string)
	Out() (url string)
	Length() int
	Add()
	Delete()
}

type Bower struct {
	a1 Stack // 第一个 栈
	a2 Stack // 第二个栈
}

// new一个新的Bower
func NewBower(a1, a2 Stack) *Bower {
	return &Bower{
		a1: a1,
		a2: a2,
	}
}

// 前进
func (b *Bower) GO(url string) string {

	if b.a2.Length() <= 0 {
		b.a1.In(url)
		return url
	} else {
		url := b.a2.Out()
		b.a1.In(url)
		return url
	}
}

//后退
func (b *Bower) Back() string {
	if b.a1.Length() <= 0 {
		return ""
	} else {
		url := b.a1.Out()
		b.a2.In(url)
		return b.a1.Top()
	}
}
