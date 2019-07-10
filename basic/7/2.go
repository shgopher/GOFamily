// 快排 归并 希尔排序
package main

func XiEr(x ...int)[]int {
	var r []int
	for i := 0;i < len(x);i = i+10{
		r = Cha(x[0:i]...)
		if i < len(x) && len(x)- i < 10 {
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
