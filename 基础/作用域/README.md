# 作用域

先给出两个经典的案例：

if级变量
```go
func main() {
  if a :=0;false{
  
  } else if b :=0;false{

  }else if c :=0;false{

  }else{
    println(a,b,c)
  }
}
```
a b c 这三个变量的作用域属于 if 这整个结构，只要前面初始化过，那么后面就可以使用，所以你会发现print可以输出三个变量。

loop级变量
```go
func main() {
  for i:=0;i<10;i++ {
    go func(){
      println(i)
    }()
  }
  time.Sleep(time.Second)
}
```
因为 i属于loop级别，通常来说for执行过程是快于新开辟一个goroutine的，所以导致这是个goroutine输出的都是最后一个i，即都输出10

根据go 1.20的表述，以后这个loop级别的变量非常有可能会被修改成局部变量。如果没有被修改，我们可以开辟一个局部变量

```go
func main() {
  for i:=0;i<10;i++ {
    i := i
    go func(){
      println(i)
    }()
  }
  time.Sleep(time.Second)
}
```