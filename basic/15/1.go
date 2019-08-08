package main

func Bei(YWeight int, a []int, i int, maxWgith int, result *int) {
	if i == len(a) || YWeight == maxWgith {
		if YWeight > *result {
			*result = YWeight
		}
		return
	}
	// 开始分叉
	// 左叉 要求这个货物不加入背包
	Bei(YWeight, a, i+1, maxWgith, result)

	// 右叉 要求 货物加入背包，并且剪枝技巧，判断是否超重，超重的话这一叉子就算了。
	// 剪枝
	if YWeight+a[i] <= maxWgith {
		Bei(YWeight+a[i], a, i+1, maxWgith, result)
	}

}
