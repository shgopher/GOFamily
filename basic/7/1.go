// 冒泡 选择 插入
package main

// 冒泡

func Mao(x ...int) []int {
L:
	for i := 0; i < len(x); i++ {
		y := 0
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
				y++
			}
		}
		if y == 0 { // 这就就是 优化，如果 某次冒泡 一次都没有动，那么说明已经拍好了 不用再动了。
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

// 选择
