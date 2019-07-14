package TN

// 快排
func TOPKQuickSelect(x []int, k int) {
	left := 0
	right := len(x) - 1
	quickSort(left, right, x, k)

}

func quickSort(left int, right int, arr []int, k int) {
	if left < right {
		b1, b2 := 0, 0
		t, m := pattion(left, right, arr, k)
		if t == 0 {
			// fmt.Println(arr[m])
			return
		}
		if t == 1 {
			b1, b2 = left, m-1
		} else if t == 2 {
			b1, b2 = m+1, right
		}
		quickSort(b1, b2, arr, k)
	}
}
func pattion(left int, right int, arr []int, k int) (int, int) {
	p := left
	index := p + 1
	for i := index; i <= right; i++ {
		if arr[i] > arr[p] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[p], arr[index-1] = arr[index-1], arr[p]
	if k == index {
		return 0, index - 1
	} else if k < index {
		return 1, index - 1
	} else {
		return 2, index - 1
	}
}
