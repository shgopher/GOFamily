package main

import (
	"fmt"
	"testing"
)

func TestMao(t *testing.T) {
	fmt.Println("common",Mao(2,4,6566,66,344343,435,663,44,45,5,6,66,65,45))
	fmt.Println("0",Mao(0,0,0,0,0))
}

func BenchmarkMao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mao(1,32,23,23,2344,5534,23,1,1,323,323,22,55,13,67,7442,24245,5424,)
	}
}

func BenchmarkMao1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mao(0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,)
	}
}