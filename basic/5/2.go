// 堆排序
package main

func topToDownPaixu(h *Heap, i int) {
	for {
		max := i
		if 2*i <= h.count && h.value[i] < h.value[2*i] {
			max = 2 * i
		}

		if 2*i+1 <= h.count && h.value[max] < h.value[2*i+1] {
			max = 2*i + 1
		}
		if max == i {
			break
		}
		h.value[i], h.value[max] = h.value[max], h.value[i]
		i = max
	}

}
func pai(a []int) []int {

	h := NewHeap(10)
	h.value = a
	h.count = len(a) - 1
	for i := h.count / 2; i >= 1; i-- {
		topToDownPaixu(h, i)
	}
	return h.value

}
