// 堆 的增删改查
package main

import "fmt"

type Heap struct {
	// 数据结构堆中的数据
	value []int
	// 一共的容量
	cap int
	// 已经储存的量
	count int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		// cap+1 的原因是堆的底层数组第一个0下标的不储存数据
		value: make([]int, cap+1, cap+1),
		cap:   cap,
		count: 0,
	}
}

// 插入是从下 往上
func (h *Heap) Insert(data int) error {
	if h.count >= h.cap {
		return fmt.Errorf("堆已满，请勿再次插入数据，本次插入的数据无效")
	}
	// 满足条件的就将数据+1
	h.count++
	h.value[h.count] = data

	// 下面开始堆化 从下往上的堆化。
	downToTop(h)
	return nil
}

// 插入是从下 往上
func (h *Heap) Insert1(data int) error {
	if h.count >= h.cap {
		return fmt.Errorf("堆已满，请勿再次插入数据，本次插入的数据无效")
	}
	// 满足条件的就将数据+1
	h.count++
	h.value[h.count] = data

	// 下面开始堆化 从下往上的堆化。
	downToTop1(h)
	return nil
}

func (h *Heap) Range() []int {
	result := make([]int, 0)
	if h.count > h.cap {
		return []int{}
	}
	i := 0
	for i*2+1 <= h.count {
		result = append(result, h.value[2*i], h.value[i], h.value[2*i+1])
		i++
	}
	return result
}

// 删除是从上往下
func (h *Heap) Delete() error {
	h.value[1], h.value[h.count] = h.value[h.count], h.value[1] // 交换
	h.count--
	// 开始堆化
	topToDown(h)
	return nil
}

func topToDown(h *Heap) {
	max := 1
	i := max
	for {
		max = i
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

func downToTop(h *Heap) {
	i := h.count

	// 在这一步之前其实并没有形成一个堆，因为你只是机械的往这个数组中push数据而已，然后通过下面的堆化，通过i和i/2然后才演变为了堆
	for i/2 > 0 && h.value[i] > h.value[i/2] {
		h.value[i], h.value[i/2] = h.value[i/2], h.value[i]
		i = i / 2
	}
}

func downToTop1(h *Heap) {
	i := h.count

	// 在这一步之前其实并没有形成一个堆，因为你只是机械的往这个数组中push数据而已，然后通过下面的堆化，通过i和i/2然后才演变为了堆
	for i/2 > 0 && h.value[i] < h.value[i/2] {
		h.value[i], h.value[i/2] = h.value[i/2], h.value[i]
		i = i / 2
	}
}
