package main

// 求 ！
func G(x int) int {
	if x == 1 {
		return 1
	}
	return G(x-1) * x
}

func G2(x int) int{
	if x == 1 {
		return 1
	}

	t := G2(x -1) * x // 这一步 就是 入栈 和出栈
	return t
}

func Gi(x int)(x1 int){
	if x <=1 {
		return 1
	}
	return Gi(x-1) + Gi(x-2)
}
// 递归要做的就两件事 1 找出 最里面的栈底 2 找出 每个栈变量的计算公式 比如这个题 计算公式 F（n） = x * (x -1) 这里面的栈底
// 就是 当x==1 返回1 即可。这也是出栈的条件。如果没有这个条件，将会 Stack Overflow 也就是栈溢出。
