// 快排 归并 希尔排序
package main

func XiEr(x ...int) []int {
	var r []int
	for i := 0; i < len(x); i = i + 10 {
		r = Cha(x[0:i]...)
		if i < len(x) && len(x)-i < 10 {
			r = Cha(x...)
		}
	}

	return r
}

func ShellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}

// 归并排序

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 { // 当这个数组中的数据已经只有一个的时候，就可以返回了。
		return arr
	}
	middle := length / 2                            // 从中间劈开 分治的思想，一直将这些数据从中间劈开。
	left := arr[0:middle]                           // 从中间劈开的左侧的数据
	right := arr[middle:]                           // 右侧的数据
	return merge(mergeSort(left), mergeSort(right)) // 合并 左侧的和右侧的数据，将左右两侧的数据开始 排序。
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
