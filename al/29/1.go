package TN

import (
	"fmt"
)

// 快排
func TOPKQuickSelect(x []int, k int) int {
	left := 0
	right := len(x) - 1
	t := quickSort(left, right, x, k)
	fmt.Println("测试TOP", t)
	return t
}

func quickSort(left int, right int, arr []int, k int) int {
	if left < right {
		t, m := pattion(left, right, arr, k)
		fmt.Println("测试pattion", t, m)
		if t == -1 {
			return -1
		} else if t == 0 {
			return m
		} else if t == 1 {
			quickSort(left, m-1, arr, k)
		} else if t == 2 {
			quickSort(m+1, right, arr, k)
		}
	}
	return -110
}
func pattion(left int, right int, arr []int, k int) (int, int) {
	p := left
	index := p + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[p] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[p], arr[index-1] = arr[index-1], arr[p]
	if k == index {
		return 0, index - 1
	} else if k > index {
		return 1, index - 1
	} else {
		return 2, index - 1
	}
}
