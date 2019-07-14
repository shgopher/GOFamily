package main

import (
	"fmt"
	"testing"
)

var (
	chaC = []int{
		1, 65, 34, 664, 1214, 44, 2, 876, 344, 2221, 6555, 1, 433, 1234, 7654, 8776, 11, 446, 643, 12345, 434, 8765, 123, 85223, 2217654332, 11, 678, 12654,
	}
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
