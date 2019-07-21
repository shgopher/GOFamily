// 二叉搜索树的查找 插入 删除
package main

type BinarySearchTree struct {
	value int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func Search(root *BinarySearchTree, x int) *BinarySearchTree {
	var result *BinarySearchTree
	if root == nil {
		return &BinarySearchTree{}
	}
	if root.value == x {
		return root
	} else if root.value < x {
		result = Search(root.right, x)
	} else {
		result = Search(root.left, x)
	}
	return result

}

func Search1(root *BinarySearchTree, x int) *BinarySearchTree {
	for root != nil {
		if root.value == x {
			return root
		} else if root.value < x {
			root = root.right
		} else {
			root = root.left
		}
	}
	return root

}
