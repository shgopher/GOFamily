// 使用指针的方式来实现二叉树,百分之99%的树结构都是用这个方式，只有完全二叉树（例如堆）使用数组更好，因为可以节省了指针的空间
// 但是其他的非完全二叉树用数组的话会浪费非常多的空间，使用指针的方式更好。
package main

type BT struct {
	value interface{}
	left  *BT
	right *BT
}

func pre(x *BT) []interface{} {
	if x == nil { // 说明遍历的已经到头了
		return []interface{}{0}
	}
	if x.left == nil && x.right == nil {

		return []interface{}{x.value}
	}
	var res []interface{}
	res = append(res, x.value)
	res = append(res, pre(x.left)...)
	res = append(res, pre(x.right)...)
	return res
}

func in(x *BT) []interface{} {
	if x == nil { // 说明遍历的已经到头了
		return []interface{}{0}
	}
	if x.left == nil && x.right == nil {
		return []interface{}{
			x.value,
		}
	}
	res := in(x.left)
	res = append(res, x.value)
	res = append(res, in(x.right)...)
	return res
}

func last(x *BT) []interface{} {
	if x == nil { // 说明遍历的已经到头了
		return []interface{}{0}
	}
	if x.left == nil && x.right == nil {
		return []interface{}{
			x.value,
		}
	}
	res := last(x.left)
	res = append(res, last(x.right)...)
	res = append(res, x.value)
	return res
}
