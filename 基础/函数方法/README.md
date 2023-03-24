# 函数和方法
重点内容前情提要

- 函数初始化的顺序
- init 函数
- 函数式编程
- defer的运用
- 变长参数
- 方法集合决定接口的实现
- 类型嵌入以及方法集合
- define类型的方法集合
- 类型别名的方法集合

go 语言中函数的基本使用方法如下：

```go
// 函数的一般用法, go 拥有多个返回值
func Get(i int)(int, error) {
  return 0, nil
}
// 返回值具有变量的返回模式
func Post(i int)(value int){
  value = 1
  return
}

// 函数作为value值赋值给一个变量变量
var Get0 = func(i int) int {
  return 0
}
```

## 函数的初始化顺序
![](./gofunc.svg)

## init 函数
## defer 函数
## 变长参数
## 函数式编程


## 参考资料
- https://book.douban.com/subject/35720728/ 170页 - 243页
- https://coolshell.cn/?s=GO+%E7%BC%96%E7%A8%8B%E6%A8%A1%E5%BC%8F


