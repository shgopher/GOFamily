package main

import (
	"fmt"
	"testing"
)

func TestP(t *testing.T) {
	fmt.Println(P("abcdefghijklmnopqrstuvwxyz", "abcde")) // 头
	fmt.Println(P("abcdefghijklmnopqrstuvwxyz", "wxyz"))  // 尾
	fmt.Println(P("abcdefghijklmnopqrstuvwxyz", "jklm"))  // 普通
	fmt.Println(P("abcdefghijklmnopqrstuvwxyz", "bblm"))  // 没有
	// true
	// true
	// true
	// false

}

func BenchmarkP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P("abcdefghijklmnopqrstuvwxyz", "foew")
	}
}

func BenchmarkP1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "foew")
	}
}

func BenchmarkP2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "foew")
	}
}

// BenchmarkP-4    	10000000	       167 ns/op
// BenchmarkP1-4   	 5000000	       342 ns/op
// BenchmarkP2-4   	 2000000	       837 ns/op
