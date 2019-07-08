// 如何在递归的过程中使用重复值。
package main

var (
	e = make(map[int]int)
)

func T(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	} else if v, ok := e[x]; ok {
		return v
	}
	//e[x] = T(x) // 这里引起了Stack Overflow
	t := T(x-1) + T(x-2)
	e[x] = t
	return t
}

func T1(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 2
	}
	return T1(x-1) + T1(x-2)

}
