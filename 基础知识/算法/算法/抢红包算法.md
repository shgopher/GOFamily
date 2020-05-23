# 抢红包算法
- 二倍均值法
每次给一个随机值，番位是 0,（总数/剩余人数 ）* 2， 乘2 主要是为了公平。
- 线段分割法
这个简单，把钱数，映射到长度上，然后在一个比如10块的金额上，我们假设有一个10的长度的棍子
然后我们随机生成比如说要求是分为3份，我们就将这个棍子分为三分，然后每一段的长度/总长度 * 金额 就是随机值。

```go
func redPacket(m,n int)[]int{
ma := make(map[int]int)
re := make([]int,n-1)
for i := 0;i < n-1;i++ {
      t := rand(0,m)
      if _,ok := ma[t];ok {
        continue
        i --
      }
      re = append(re,t)
}
sort.Ints(re)
result := make([]int,n)
for i := 0;i<len(re)-1;i++ {
  result = append(result,re[i+1]-re[i])
}
return result
}
```
解释：

假如 10块分为3份算法解释如下：

我们在0-10块这个区间选2个点 比如说  7 9 吧

然后我们取  7-0 也就是7  9-7 也就是2 10-9 也就是1 所以最后结果是 7 2 1 