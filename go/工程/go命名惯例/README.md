# go命名惯例

go语言的命名哲学就是**尽可能的简单化**。

go要求的简单，不仅仅是简单化，还要关注上下文的连贯性，go喜欢一个简洁明了的命名+对字段的注释

```go
// userTypeInfo stores the information associated with a type the user has handed
// to the package. It's computed once and stored in a map keyed by reflection
// type.
type userTypeInfo struct {
	user        reflect.Type // the type the user handed us
	base        reflect.Type // the base type after all indirections
	indir       int          // number of indirections to reach the base type
	externalEnc int          // xGob, xBinary, or xText
	externalDec int          // xGob, xBinary or xText
	encIndir    int8         // number of indirections to reach the receiver type; may be negative
	decIndir    int8         // number of indirections to reach the receiver type; may be negative
}
```
比如这段代码，对于字段和结构体都进行了注释，但是结构体内部的字段都相对简洁

## 包

包名称一般使用小写单词进行命名，例如说：`context`,`http`,`net`,无需考虑包名称的重复化，因为在go引入包的过程中，实际上是引入的路径，并且go的包在引入过程中还可以进行包的重命名

```go
import (

  "example.com/hey/log"

  blog "example.net/bob/log"
)
```
包的路径名称最后一位尽量和包的名称保持一致性，不过不一致也是有不少情况的，比如说，现在有一个适应不同语言的log，那么命名的时候就变成了`github.com/shgopher/golog`,`github.com/shgopher/javalog`,但是他们实际上的包都叫log，并不是golog和javalog，那么使用的时候就需要对其中一个包进行重命名了。
## 变量，类型，函数，方法
因为go语言使用导出的变量时需要带上包的名称(除了. 省略包名称的方式，一般不提倡)，所以命名的时候变量不需要再加上包的称呼，比如说string包的Reader就不需要写成`string.StringReader`,写成`string.Reader`即可

go变量的命名使用驼峰的命名方式，go的文件名称使用下划线连接的方式，比如说`UpperName`,`lowName` 区别就是前者是大驼峰，导出的变量，后者是小驼峰，不可导出变量，缩略词如果首字母大写那么全部大写，比如`HTTP`，不能写成`Http`

go代码里使用了大量单字母或者短字母的命名方式。

for循环中以及if条件语句中使用单字母
```go
for k,v := range users{

}

if err := method1();err != nil {

}
```
函数和方法的参数或者返回值，常见单字母或单个单词的方式出现。

```go
func method1(name string){}
```
方法调用的时候会绑定类型信息，可以尽量使用单个单词的命名方法

函数和类型多以多单词驼峰的方法组合，比如`func MakeFile(){}`，`type gobEncoderType struct {}`

在命名变量的时候，不要带上类型信息，比如说 `names`比`nameSlice`要好，因为go强调命名和使用越近越好，没必要在字段上加上类型的名称。

go使用一致性来强调单字母或者单个单词的意义，意思就是说使用的单个字母在任何地方表示的意思都是一样的。

```go
for k,v := range names{}
for i,v := range users{}
```
从k v i 这三个字母就可以观察到，绝大多数情况下，在go语言中kv就是表示的key和value，i表示的就是index，那么即便是单个字母也不会造成识别障碍

另外`t` `b` `n` 也很常见，他们通常表示 time 和 byte 以及数量
## 常量
- go语言的常量并不要求全大写，通常只有本身是全大写的常量才是全大写，比如`PI`
- 常量不需要赋予类型，系统会根据根据使用时期左变量的类型，以及运算操作进行自动推断

一般常量
```go
const(
  keepHostHeader = false
)
```
全大写常量
```go
const PI = 3.14

const SIGABRT = Signal(0x6)
```
## 接口
拥有唯一方法的接口，或者内嵌多个拥有唯一方法的接口一律使用`单词+er`结尾，比如`Reader`或者`WriteReadCloser`

go语言推崇简单接口，和内嵌多个简单接口的接口，所以你会看到go语言标准库很多带有er结尾的接口

```go
type Reader interface{
  Reade()
}
type Writer interface{
  Write()
}
type WriteReadeCloser interface{
  Writer
  Reader
  Closer
}
```
## 两种命名的对比
简洁的并且考虑上下文一致性的代码：
```go
func RuneCount(b []byte)int{
  count := 0
  for i:= 0;i< len(b);{
    if b[i] < RuneSelf {
      i++
    } else{
      _,n := DecodeRune(b[i:])
      i+=n
    }
    count++
  }
  return count
}
```
见名知意的java式的代码:
```go
func RuneCount(buffer []byte)int{
  runeCount := 0
  for index := 0;index < len(buffer);{
    if buffer[index] < runeSelf {
      index++
    } else {
      _, size := DecodeRune(buffer[index:])
      index += size
    }
    runeCount++
  }
  return runeCount
}
```
## 参考资料
- 图书：go语言精进之路
