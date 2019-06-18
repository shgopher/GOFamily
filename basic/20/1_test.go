package main

import (
	"testing"
)

// 测试的原则 1 空值 2 边界左右↔️ 3 错误值 4 普通值
func TestGo(t *testing.T) {
	// 0值
	b1, b2 := NewSlice(), NewSlice()
	Back(b1, b2)
	Back(b1, b2)

	// 普通值
	b21, b22 := NewSlice(), NewSlice()
	GO(b21, b22, "https://google.com")
	GO(b21, b22, "https://nudao.xyz")
	Back(b21, b22)
	GO(b21, b22, "12")


}

func TestBack(t *testing.T) {

}
