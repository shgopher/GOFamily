package main

import (
	"fmt"
	"testing"
)

func TestP(t *testing.T) {
	fmt.Println(P("tggigjrwefewfokjogrejgernvdmsfkewfoewfpogjhyryevfndjjd", "foew"))
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

//
// BenchmarkP-4    	 2000000	       807 ns/op
// BenchmarkP1-4   	 1000000	      1465 ns/op
// BenchmarkP2-4   	  300000	      4295 ns/op
