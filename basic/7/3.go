// 基数  桶
package main

import (
	"fmt"
	"sync"
)

// 桶排序，这里使用 快排最为基础桶内排序。
// 桶排序

// 获取待排序数组中的最大值
func getMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func BucketSort(a []int) {
	num := len(a)
	if num <= 1 {
		return
	}
	max := getMax(a)
	buckets := make([][]int, num) // 二维切片

	index := 0
	for i := 0; i < num; i++ {
		index = a[i] * (num - 1) / max                // 桶序号
		buckets[index] = append(buckets[index], a[i]) // 加入对应的桶中
	}
	wg := sync.WaitGroup{}
	// tmpPos := 0 // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				QuickSort(buckets[i]) // 桶内做快速排序
			}(i)
			// copy(a[tmpPos:], buckets[i])
			// tmpPos += bucketLen
		}
	}
	wg.Wait()
	fmt.Println(buckets)

}

// 基数排序
// func RadixSort(arr []int, maxSize int) []int {
//
// }
