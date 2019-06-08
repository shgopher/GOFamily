package main

func main() {

}

type Stack struct {
	data []interface{}
	top  int //表示栈顶 因为slie会自动扩容，而且需要减少不必要的损失（下文的使用top不直接变slice）所以我们使用top表示栈最后一个，最右边的那个，也可以说是
	// 最上面一个，其实怎么说都可以 我就说最上面那个吧。反正就是最新的那个 它始终是top，它最先出去。
}

//放入
func (s *Stack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top++
	}
	if s.top < len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

// 放出
func (s *Stack) Pop() interface{} {
	if s.top < 0 {
		return nil
	}
	s.top--
	return s.data[s.top]
}
