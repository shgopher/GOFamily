# go语言的条件和逻辑语句

## 条件语句
`if else` 是go里面的第一个条件语句，看一个例子就能一目了然：
```go
func fast(a int){

    if a > 0 {
        fmt.Println("YES")
    }else if a == 0 {
        fmt.Println("YES!")
    }else {
        fmt.Println("no")
    }

}
```
可以看出，if后面跟一个判断的语句，并且，不加括号。

switch是第二个判断语句。

```go
func fast(a int){
    switch a {
        case 1:
            fmt.Println("yes")
        case 2 :
            fmt.Println("yes!")
            fallthrough
        default:
             fmt.Println("no")
    }
}
```
这里面，switch后面是a，a是int类型，那么case后面也得是int类型,这里有个fallthrouth 语句，意思是，走进这个case后不是直接出去，而是直接往下继续走。这跟其它语言不同，其它语言是默认继续走，然后使用了break才能出来，go是默认就是break了。

switch还有另一种跟断言结合的使用方法

```go
func fast(a interface{}){
    switch a.(type){
        case int:
        case float64:
    }
}
```
断言的意义就是说搞清楚这个空接口具体的类型，所以叫做断言，断言的语句就是 1. `一个空接口类型变量.(type)`【只能在switch中使用】，另一种是 `b,ok := 一个空接口类型变量.(int)` 后面必须加具体类型，然后前面是第一个参数是返回的具体类型，第二个参数是布尔类型，判断是否断言成功。
## 逻辑语句
for 是go里面唯一的逻辑语句，在go里没有while和do while。这两个都是不合法的。

```go
// 第一种表示一直for循环
for {

}
// 第二种表示，只要a>12就循环
for a > 12 {

}
// 第三种就比较传统了。
for i := 0;i < 12; i++ {

}
```
## 遍历语句

range 是go唯一的遍历语句

```go
    b := make([]int,10)

    for a := range b {
    fmt.Println(a)
    }

```
这里要特别的指出来，a是对于b遍历的数据的值复制，也就是说下文中对于a的任何处理都不会影响原本的b里面的数据。