package main

import (
	"fmt"
	"testing"
)
//0
func TestT(t *testing.T) {
	fmt.Println(T(40000))
}

//0.55s
func TestT1(t *testing.T) {
	fmt.Println(T1(40))
}
//12.9 ns/op
func BenchmarkT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T(5000)
	}
}
//290 ns/op
func BenchmarkT1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T1(50)
	}
}


// 不重复运算和重复运算的 递归对比，相差了 可以说是将近 30倍。

// 如果是 5000这个数量级 非重复计算的 只有21 但是 重复计算的  肯定是 瘫痪了，这就是 算法 和数据结构的魅力 啊 !!!!