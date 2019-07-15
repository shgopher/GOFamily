package main

import (
	"fmt"
	"testing"
)

func TestXiEr(t *testing.T) {
	fmt.Println(XiEr(chaC...))
}

//33186 ns/op
func BenchmarkXiEr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		XiEr(chaC...)
	}
}

func BenchmarkCha2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cha(chaC...)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort(chaC)
	}
}

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort(chaC))
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(chaC)
	}
}
func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort(chaC))
}

//  4.58 ns/op
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(chaC)
	}
}
