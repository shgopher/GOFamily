package main

func P(m, p string) bool {
	for i := 0; i <= len(m)-len(p); i++ {
		var a string
		if i+len(p) >= len(m) {
			a = string(m[i:])
		} else {
			a = string(m[i : i+len(p)])
		}
		if a == p {
			return true
		}
	}
	return false
}
