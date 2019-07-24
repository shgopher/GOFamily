// 堆排序
package main

func topToDownPaixu(h *Heap, i int) {
	for {
		max := i
		if 2*i <= h.count && h.value[i] > h.value[2*i] {
			max = 2 * i
		}

		if 2*i+1 <= h.count && h.value[max] > h.value[2*i+1] {
			max = 2*i + 1
		}
		if max == i {
			break
		}
		h.value[i], h.value[max] = h.value[max], h.value[i]
		i = max
	}

}

// 这是 从上 往下的建立堆的方式 其实跟 一个push一个 然后 从下往上一样，一个效果 这个结果是建立堆完成。
func BuildHeapTopToLow(a []int, h *Heap) {
	h.value = a
	h.count = len(a) - 1
	for i := h.count / 2; i >= 1; i-- { // 一定要 n/2 才能完全的建立 堆
		topToDownPaixu(h, i)
	}

	// 开始排序

}

func HeapPai(a []int, h *Heap) []int {
	BuildHeapTopToLow(a, h)
	k := h.count
	for k > 1 {
		h.value[k], h.value[1] = h.value[1], h.value[k]
		k--
		h.count--
		topToDownPaixu(h, 1) // 其实就是 排序堆顶的那一个值罢了，因为我们只要找到最大或者最小 其它值不用管。
	}
	return h.value

}
