# go数字类型
**导读：**
- int
- float
- complex
- issues

## int

int 是go语言的整数类型，一共分为下面这么几类

|类型|描述|
|:---:|:---:|
|int|有符号的整数类型，具体占几个字节要看操作系统的分配，不过至少分配给32位|
|uint|非负整数类型，具体占几个字节要看操作系统的分配，不过至少分配给32位|
|int8|有符号的整数类型，占8位bit，1个字节。范围从负的2的7次方到正的2的7次方减1|
|int16|有符号的整数类型，占16位bit，2个字节。范围从负的2的15次方到正的2的15次方减1|
|int32|有符号的整数类型，占32位bit，4个字节。范围从负的2的31次方到正的2的31次方减1|
|int64|有符号的整数类型，占64位bit，8个字节。范围从负的2的63次方到正的2的63次方减1|
|uint8|无符号的正整数类型，占8位，从0到2的8次方减1.也就是0到255|
|uint16|无符号的正整数类型，占16位，从0到2的16次方减1|
|uint32|无符号的正整数类型，占32位，从0到2的32次方减1|
|uint64|无符号的正整数类型，占64位，从0到2的64次方减1|
|uintptr|无符号整数类型。它大到足以容纳任何指针|
|rune|int32的别名，代表一个 UTF-8 字符，比如'中''国'rune是两个文字符|
|byte|uint8别名，用于称呼字节符，比如 '中''国'，使用byte表示就不只两个字符，因为它代表了ASCII 码的一个字符|

go使用`'\x12'` 或者使用 `0x12`来表示16进制

go使用`'\012'`或者使用`012`来表示8进制

go不能直接显示2进制，使用`fmt.Printf("%b",12)` `1000` 来输出一个二进制
## float
|类型|描述|
|:---:|:---:|
|float32|浮点型，包括正负小数，IEEE-754 32位的集合，提供大约 6 个十进制数的精度，math.MaxFloat32 表示 float32 能取到的最大数值,math.SmallestNonzeroFloat32表示最小值|
|float64|浮点型，包括正负小数，IEEE-754 64位的集合，提供约 15 个十进制数的精度，math.MaxFloat64 表示 float64 能取到的最大数值,math.SmallestNonzeroFloat64表示最小值|

浮点数使用 `fmt.Printf("%.4f\n", math.Pi)` ，`%.nf` 来控制保留几位小数

## complex 复数
|类型|描述|
|:---:|:---:|
|complex64|复数，实部和虚部是float32|
|complex128|复数，实部和虚部都是float64|

使用go语言内置的函数 `real()`和`imag()`来分别获取到复数的实部和虚部，构建复数的时候使用内置函数`complex()`
## issues
***问题一：*** go语言数字之间的类型转化是如何进行的

go语言的类型转换是显性转化的，数字类型之间可以使用这种方法
```go
var i int
f := float64(i)
```
***问题二：*** 能说说`uintptr`和`unsafe.Pointer`的区别吗

- `unsafe.Pointer`是通用指针类型，它不能参与计算
- `uintptr`是指针运算的工具，但是它不能持有指针对象（意思就是它跟指针对象不能互相转换）
- `unsafe.Pointer`和`uintptr`可以相互转换，`unsafe.Pointer`和指针对象可以相互转换，但是`uintptr`和指针对象不能相互转换
- unsafe.Pointer是指针对象进行运算（也就是uintptr）的桥梁

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 指针对象
	v := new(int)
	// 将指针对象转化为通用指针类型
	vp := unsafe.Pointer(v)
	// 将通用指针类型转换为指针对象
	vo := (*int)(vp)

	//将通用指针对象转化为uintptr
	uv := uintptr(vp)
	//将uintptr转换为通用指针对象
	vpp := unsafe.Pointer(uv)
	fmt.Println(v, vp, vo, uv, vpp)

	// 对指针对象的地址进行计算
	t := new(string)
	// 首先先将t转化为unsafe.Pointer类型
	pt := unsafe.Pointer(t)
	// 然后将pointer再转化为 uintptr
	rt := uintptr(pt)
	//进行计算
	npt := rt + uintptr(1)
  // 计算完毕后，再将uintptr转化为unsafe.Pointer再转换为*string类型
	nt := (*string)(unsafe.Pointer(npt))
	fmt.Println(t, pt, rt, npt, nt)
}

```

***问题三：*** rune和byte的区别

rune是int32，byte是uint8，相比byte来说，rune可以容纳的字符个数要多很多，所以utf8编码的字符使用rune，而ascii使用byte

> unicode是一种字符编码，让每个字符和一个数字对应起来，仅此而已，至于这个数字如何存储它就不管了。utf8就是定义了如何具体存储这个编码数字的一种方法

## 参考资料
- https://zhuanlan.zhihu.com/p/145220416
- http://www.manongjc.com/article/50416.html
- http://c.biancheng.net/view/18.html