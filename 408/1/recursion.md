# 递归

## 斐波那契数列
`1 1 2 3 5 8 13` 这个斐波那契数列的递归方程是`F(n) = F(n-1)+F(n-2)`其中n是角标。这个函数，当遇到n是0或者是1的时候就可以返回值了，为啥呢，因为当n=1orn=0的时候他们的值都是1，而且是最小的两个数字了。
```go
func fi(n)(value int){
  if n == 1|| n == 2{
    return 1
  }
  return fi(n-1)+ fi(n-2)
}
```
这种计算方法其实使用到了大量的重复的数据，所以说这种简单的递归，会对性能造成很大的影响，时间复杂度也不低，
它的时间复杂度是`2^n`怎么算出来的？很简单吗，每次都被分裂为2，2的每一个又是分裂成2这就是细菌的传染过程嘛
那么细菌的传染过程就是指数级别的，所以这个时间复杂度就是2的指数级别。

## 递归优化的重要性

我们还是拿斐波那契数列来出说

上文是普通的拥有大量的重复计算的斐波那契数列

这是一个记录了不同的n的计算结果的优化版的斐波那契数列
```go
var (
	ma = map[int]int{}
)
func fi(n int)int{
	if n ==1 || n == 2 {
		return 1
	}
	if v,ok := ma[n];ok { // 使用了哈希表来记录结果
		return v
	}
	value := fi(n-1)+fi(n-2)
	ma[n] = value
	return value
}

```

我们看一下当n等于40的时候，时间能差距多少

```go
package main

import (
	"testing"
)

var (
	n  = 40
)

func BenchmarkFi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fi(n)
	}
}
func BenchmarkFi2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fi2(n)
	}
}
```

输出结果

```go
goos: darwin
goarch: amd64
pkg: github.com/googege/test
BenchmarkFi
BenchmarkFi-4    	63431078	        17.8 ns/op
BenchmarkFi2
BenchmarkFi2-4   	       2	 636617175 ns/op
PASS

Process finished with exit code 0
```

同样的结果，不同的算法差距是 35765009.8 倍

对于递归，或者是重叠子问题的优化，这也是动态规划的核心之一，将大量重叠的计算省去，可以节省大量的时间。

## 向动态规划去引申

如果使用备忘录的方式，我们去做题是从上到下去做题的，也就是说先有n然后n-1 然后n-2 等等。从n往1这么向下做的，然而动态规划是从1到n这样的从下向上的

我们看斐波那契的动态规划样式的解法

```go
func fi1(n int) int {
	if n <3 {
		return 1
	}
	ma := make([]int, n+1)
	ma[1], ma[2] = 1, 1
	for i := 3 ;i <= n;i++ { // 从下往上
		ma[i] = ma[i-1]+ma[i-2]
	}
	return ma[n]
}
```
经过benchmark的测试，跟备忘录的方式是一样的时间复杂度，(o(n))

```go
BenchmarkFi
BenchmarkFi-4    	72419725	        15.9 ns/op
BenchmarkFi1
BenchmarkFi1-4   	68605408	        15.9 ns/op

```
所以说使用备忘录或者是动态规划，他们的时间复杂度是基本上相同的，但是他们都是需要额外的空间的，
使用备忘录通常是哈希表，动态规划通常是数组

**通过观察可以得知这里的状态转移方程其实只需要前面俩数字所以我们可以省去这个数组**

```go
func fi1(n int) int {
	if n <3 {
		return 1
	}

	pre1, pre2 := 1, 1
	value := 0
	for i := 3 ;i <= n;i++ {
		value  = pre1+pre2 // 使用两个pre 就可以代替这个数组了，更加节省空间，也可以说是双指针吧。
		pre1,pre2 = pre2,value

	}
	return value
}
```
这只是优化到了o(n) 其实还可以优化到log n 具体的请看[这里](../数据结构/矩阵.md)

如果动态规划中，是使用的1维数组，基本上就可以把数组省略掉改成两个指针，标准的操作题目
就是斐波那契数列。
