// 冒泡 选择 插入
package main

// 冒泡

func Mao(x ...int) []int {
L:
	for i := 0; i < len(x); i++ {
		y := false
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
				y = true
			}
		}
		if y { // 这就就是 优化，如果 某次冒泡 一次都没有动，那么说明已经拍好了 不用再动了。
			break L
		}
	}
	return x
}

func Mao1(x ...int) []int {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x
}

// 插入排序

func Cha(x ...int) []int {
	if len(x) == 0 {
		return make([]int, 0)
	}
	for i := 1; i < len(x); i++ {
		value := x[i]
		j := i - 1
	L:
		for ; j >= 0; j-- {
			if x[j] > value {
				x[j+1] = x[j] // 移动找位置
			} else {
				break L
			}
		}
		x[j+1] = value // 当后面的都往后移动了以后，那么就会在j+1的地方空出来一个位置。
	}
	return x
}


// 插入排序的链表实现

func ChaLinkedList(head interface{}) interface{}{
	
}

// 选择排序

func Xu(x...int)[]int{
	for i := 0 ;i < len(x)-1;i++ {
		min := i
		for j := i + 1; j < len(x); j++ {
			if x[min] > x[j] {
				min = j
			}

		}
		x[i],x[min] = x[min],x[i]
	}
	return x
}