package main

func main() {

}

type Stack struct {
	data []interface{}
	top  int //表示栈顶
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
