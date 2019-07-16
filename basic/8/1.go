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

func main() {
	fmt.Println(ER([]int{1, 45, 66, 777, 8999}, 45))
}
