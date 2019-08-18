# GO的坑
1. go中的for循环要想退出一定要使用这种方式

```go
L1:
for{
  L2:
  for{
    L3:
    for{
      break L1// 看你是L几就是调到哪里
    }
  }
}
```
```go
for{
  for{
    for{
      break  // 并不能跳出全部，只能跳出最里面的那层，
    }
  }
}
```

2. go所有的数据都是数据的复制，并不是指针类型，如果想使用指针类型请明示

```go
func A(i int)int{
  i++
  return i
}
func main(){
  t := 2
  A(t)
  fmt.Println(t)//结果还是2，因为你在A中只是复制而已
}
```
3. go中所有的数据只要声明，就可以初始化，但是引用类型的初始化是nil，引用类型的初始化需要make或者明示 比如 `a=[]int{1,2,3,4,5}`

4. defer的执行顺序，和它初始化的顺序完全是两码事，具体你可以把它想成栈就ok了。defer虽然是FILO但是初始化还是要顺序初始化的。这也正常
5. 对于接口的时候非常严格，指针和非指针对象完全是不同的。
6. rang中的数据是对原有数据的复制，其实所有的地方使用的数据都是复制。所以要谨记这一条。


具体的信息可以参考 https://github.com/googege/blog/blob/master/go/go/important/README.md
