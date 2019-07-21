package main

import (
	"fmt"
	"testing"
)

func chaC(now *BT, x []int, i int, j int) {
	left := new(BT)
	right := new(BT)
	now.value = x[0]
	now.left = left
	now.right = right
	if i*2 <= j {
		chaC(left, x[2*i-1:], 2*i, j)
	} else {
		return
	}
	if i*2+1 <= j {
		chaC(right, x[2*i:], 2*i+1, j)
	} else {
		return
	}
}

func TestPre(t *testing.T) {
	b := new(BT)
	chaC(b, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 10)
	fmt.Println(pre(b))
}

func TestIn(t *testing.T) {
	b := new(BT)
	chaC(b, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 10)
	fmt.Println(in(b))
}

func TestLast(t *testing.T) {
	b := new(BT)
	chaC(b, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 10)
	fmt.Println(last(b))
}
