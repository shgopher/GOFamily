package main

//

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
