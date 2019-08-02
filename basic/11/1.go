package main

func P(m, p string) bool {
	b := make([]string, len(p))
	for i := 0; i < len(m)-len(p); i++ {
		if i+len(p) > len(m) {
			b = append(b, string(m[i:]))
		} else {
			b = append(b, string(m[i:i+len(p)]))
		}
	}
	for i := 0; i < len(b); i++ {
		if p == b[i] {
			return true
		}
	}
	return false
}
