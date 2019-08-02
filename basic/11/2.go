package main

// RK 就是再增加一个 哈希算法，将所有的字符串转化为一个一个的哈希值，这样直接进行比较就很快了。
var (
	ttt map[string]int
)

func bt(ttt map[string]int) {
	ttt["a"] = 1
	ttt["b"] = 2
	ttt["c"] = 3
	ttt["d"] = 4
	ttt["e"] = 5
	ttt["f"] = 6
	ttt["g"] = 7
	ttt["h"] = 8
	ttt["i"] = 9
	ttt["j"] = 10
	ttt["k"] = 11
	ttt["l"] = 12
	ttt["m"] = 13
	ttt["n"] = 14
	ttt["o"] = 15
	ttt["p"] = 16
	ttt["q"] = 17
	ttt["r"] = 18
	ttt["s"] = 19
	ttt["t"] = 20
	ttt["u"] = 21
	ttt["v"] = 22
	ttt["w"] = 23
	ttt["x"] = 24
	ttt["y"] = 25
	ttt["z"] = 26
}
func Hash(a string) int {
	tii := make(map[string]int)
	bt(tii)
	ary := len(a) // 表示进制
	var result int
	for i := 0; i < len(a); i++ {
		b := tii[string(a[i])]
		result += b * exponent(ary, i)
	}
	return result
}

func exponent(a, n int) int {
	result := int(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}

func Pp(m, p string) bool {
	for i := 0; i <= len(m)-len(p); i++ {
		var a int
		if i+len(p) >= len(m) {
			a = Hash(string(m[i:]))
		} else {
			a = Hash(string(m[i : i+len(p)]))
		}
		if a == Hash(p) {
			return true
		}
	}
	return false
}
