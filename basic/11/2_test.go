package main

import (
	"fmt"
	"testing"
)

func TestPp(t *testing.T) {
	fmt.Println("测试hash", Pp("abcdefghijklmnopqrstuvwxyz", "abcde")) // 头
	fmt.Println("测试hash", Pp("abcdefghijklmnopqrstuvwxyz", "wxyz"))  // 尾
	fmt.Println("测试hash", Pp("abcdefghijklmnopqrstuvwxyz", "jklm"))  // 普通
	fmt.Println("测试hash", Pp("abcdefghijklmnopqrstuvwxyz", "bblm"))  // 没有
	// true
}

func BenchmarkPp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pp("abcdefghijklmnopqrstuvwxyz", "foew")
	}
}

func BenchmarkPp1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pp("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "foew")
	}
}

func BenchmarkPp2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pp("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "foew")
	}
}

// todo 哈希算法要优化，目前性能很糟糕。
