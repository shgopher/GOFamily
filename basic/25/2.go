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
	for {
		if value > b.value[0].k1 {
			if b.right == nil {
				n := new(Box)
				n.value = []*node{
					{k1: value, k2: k2},
				}
				b.right = n
				break
			} else {
				b = b.right
			}
		} else if value < b.value[0].k1 {
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
