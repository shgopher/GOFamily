# 洗牌算法

算法的基本思想是，每次从一组数中随机选出一个数，然后与最后一个数交换位置，并且不再考虑最后一个数。

```go
// 生成随机数
func In(r *rand.Rand,i)int{
  return r.Intn(i)
}

func main(){

  r := rand.New(rand.NewSoure(time.Now().UnixNano()))
  a := []int{12,534,332,2,4,4,23,}
  for i:= len(a)-1;i >=0;i-- {
    suiji := In(r,i-1)
    swap(a,i,suiji)//交换 每次都把数字往前提一位，那么就会成功的洗牌。
  }  
}


func swap(a []int,i,j int){
a[i],a[j] = a[j],a[i]
}
```
