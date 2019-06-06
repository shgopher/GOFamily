package main

var a = make([]int, 10)

func R(r int) {
	a[r]++
}

func D() {
	sorts.Ints(a)
	a = a[1:]
}
