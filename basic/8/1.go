package main

import "fmt"

func ER(x []int, value int) int {
	low := 0
	heigh := len(x) - 1
	for low <= heigh {
		middle := low + (heigh-low)>>1 // 这样为了防止 数字过大。
		if x[middle] == value {
			return middle
		} else if x[middle] > value {
			heigh = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}

// 	low := 0
// heigh := len(x)
// 使用递归的方法来实现查找。
func ERD(x []int, value int) int {
	low := 0
	heigh := len(x) - 1
	return eRD(x, low, heigh, value)
}
func eRD(x []int, low int, heigh int, value int) int {
	if low > heigh {
		return -1
	}

	middle := low + (heigh-low)>>1
	result := 0
	if low <= heigh {

		if value == x[middle] {
			result = middle
		} else if value > x[middle] {
			result = eRD(x, middle+1, heigh, value)
		} else {
			result = eRD(x, low, middle-1, value)
		}

	}
	return result

}

func main() {
	fmt.Println(ER([]int{1, 45, 66, 777, 8999}, 45))
}
