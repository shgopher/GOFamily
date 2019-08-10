// 动态规划下的背包问题。
package main

var (
	ZW     int   // 载重
	a      []int // 重量的数组
	result int   // 最后的结果

)

func DPPackage(zw int, a []int, result *int, status [][]int) {
	status[0][0] = 1
	for i := 1; i < len(a); i++ { // 设置这个矩阵的x
		// 设置这个矩阵的Y 每次是加入这个背包
		for j := 0; j <= zw; j++ {
			if status[i-1][j] == 1 {
				status[i][j] = status[i-1][j]
			}
		}
		// 设置这个矩阵的Y 每次是不加入这个背包
		for t := 0; t <= zw-a[i]; t++ {
			if status[i-1][t] == 1 {
				status[i][t+a[i]] = 1
			}
		}
	}
	for n := zw; n >= 0; n-- {
		if status[len(a)-1][n] == 1 {
			*result = n
			returng
		}
	}
}
