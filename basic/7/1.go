// 冒泡 选择 插入
package main

// 冒泡

func Mao(x ...int)[]int{
	for i := 0;i < len(x);i++ {
		for j := 0;j < len(x) - i -1  ;j++  {
			if x[j] > x[j+1] {
				x[j],x[j+1] = x[j+1],x[j]
			}
		}
	}
	return x
}

func Mao1(x...int) []int{
	for i := 0;i < len(x);i++ {
		for j := 0;j < len(x) - i -1  ;j++  {
			if x[j] < x[j+1] {
				x[j],x[j+1] = x[j+1],x[j]
			}
		}
	}
	return x
}


// 选择