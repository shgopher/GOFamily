package main

var a = make([]int, 10)

func A(ele int) { // 第一个进去的
	a = append(a, ele)
}
func D() { // 第一个出去就ok了。
	a = a[:len(a)-1]
}
