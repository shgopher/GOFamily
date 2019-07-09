package main

import (
	"fmt"
	"testing"
)

func TestMao(t *testing.T) {
	fmt.Println("common",Mao(2,4,6566,66,344343,435,663,44,45,5,6,66,65,45))
	fmt.Println("0",Mao(0,0,0,0,0))
}

// 95 ns/op
func BenchmarkMao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mao(1,1,1,1,1,1,12,3,4,5,6,98,5)
	}
}
// 118
func BenchmarkMao1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mao1(1,1,1,1,1,1,12,3,4,5,6,98,5)
	}
}

func TestMao2(t *testing.T) {

	fmt.Println("fan",Mao1(1,443,2333,556,4344,66534,4434,))
}