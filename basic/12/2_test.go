package main

import (
	"fmt"
	"testing"
)
//0
func TestT(t *testing.T) {
	fmt.Println(T(40))
}

//0.55s
func TestT1(t *testing.T) {
	fmt.Println(T1(40))
}
//12.9 ns/op
func BenchmarkT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T(10)
	}
}
//290 ns/op
func BenchmarkT1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T1(10)
	}
}


// 不重复运算和重复运算的 递归对比，相差了 可以说是将近 30倍。