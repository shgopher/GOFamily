//二叉搜索树的查找 插入 删除
package main

import "fmt"

type Box struct {
	left  *Box
	right *Box
	value []*node
}
type node struct {
	k1 int // 构建关于k1的二叉查找树，将k1相同的放在同一个box节点。(k2 可以不同)
	k2 string
}

func InsertM(b *Box, value int, k2 string) {
	if b == nil {
		fmt.Println("aaaa")
		b = &Box{
			value: []*node{
				{k1: value, k2: k2},
			},
		}
	}
	if b.value == nil {

		b.value = []*node{
			{k1: value, k2: k2},
		}
	}
	for { // 循环的额目的是为了找到left或者nil等于nil的地方
		if value > b.value[0].k1 {
			if b.right == nil { // 如果等于nil那么我们只需要给他一个新的node节点即可。
				n := new(Box)
				n.value = []*node{
					{k1: value, k2: k2},
				}
				b.right = n
				break // 然后break出来
			} else {
				b = b.right // 如果下一个节点不是nil那么我们应该找下个节点。
			}
		} else if value < b.value[0].k1 { // 这里为什么可以直接使用value[0] 因为 不管是root是nil的情况(我们给他了0值)还是下一步为nil的状态如果是nil我们给了它新的node，如果不是nil那么它一定是拥有数组的。所以我们可以随意使用[0]
			if b.left == nil {
				n := new(Box)
				n.value = []*node{
					{k1: value, k2: k2},
				}
				b.left = n
				break
			} else {
				b = b.left
			}

		} else {
			b.value = append(b.value, &node{k1: value, k2: k2})
			break
		}
	}
}

func find1(b *Box, value int) []*node {
	if b == nil || b.value == nil {
		return nil
	}
	for b != nil {
		if b.value == nil {
			return b.value
		}
		if b.value[0].k1 == value {
			return b.value
		} else if b.value[0].k1 > value {
			b = b.left
		} else {
			b = b.right
		}
	}
	return nil
}
