package main

import (
	"fmt"
	"testing"
)

var (
	arr = new(Box)
)

func TestInsertA(t *testing.T) {
	InsertM(arr, 10, "1")
	InsertM(arr, 10, "2")
	InsertM(arr, 10, "3")
	InsertM(arr, 10, "4")
	InsertM(arr, 10, "5")
	InsertM(arr, 10, "6")
	InsertM(arr, 10, "7")
	InsertM(arr, 16, "__")
	InsertM(arr, 16, "$##$F")
	InsertM(arr, 16, "/fsdf")
	InsertM(arr, 16, "rfdsd")
	InsertM(arr, 16, "fsdfef")
	InsertM(arr, 7, "71")
	InsertM(arr, 7, "74")
	InsertM(arr, 7, "756")
	InsertM(arr, 7, "7666")
	fmt.Println("测试 数组:", *arr)
}
func TestFind1(t *testing.T) {
	for _, v := range find1(arr, 10) {
		fmt.Println(v.k1, v.k2)
	}
	for _, v := range find1(arr, 7) {
		fmt.Println(v.k1, v.k2)
	}
	for _, v := range find1(arr, 16) {
		fmt.Println(v.k1, v.k2)
	}
}
