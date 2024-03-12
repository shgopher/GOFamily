<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2024-03-12 14:21:48
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-03-12 15:09:07
 * @FilePath: /GOFamily/工程/内存对齐实践/README.md
 * @Description: 
 * 
 * Copyright (c) 2024 by shgopher, All Rights Reserved. 
-->
# 内存对齐实践
```go
type Person struct {
  Sex bool // 性别
  Height uint16 // 身高
  Addr byte
  Postion byte
  age int64
  weight uint16
}
```
我们分别算一下数据实际的大小和在内存中的的偏移量以及整个结构体因为内存对齐所一共占有的内存数据

实际大小
```go
type Person struct {
  Sex bool //1
  Height uint16 //2
  Addr byte //1
  Postion byte //1
  age int64 //8
  weight uint16//2
}
```
偏移量
```go
type Person struct {
  Sex bool // 0
  Height uint16 // 2
  Addr byte//4
  Postion byte // 5
  age int64 // 8
  weight uint16 // 16
}
```

实际在内存中的排序：
```bash
sex * height height addr postion * * age age age age age age age age weight weight 
```

在 Go 语言中，每种数据类型在内存中的字节对齐方式是不同的，主要遵循以下规则：

- 每个变量都会占用一块内存区域，并从这块区域的起始地址开始存储值。
- 结构体的起始地址必须是其最大字段的对齐值的倍数。
- 每个字段的起始地址必须是其类型对齐值的倍数。
- 内存对齐后的结构体大小必须是其最大字段的对齐值的倍数。
- 如果存在内存对齐的需求，编译器会在具体字段之间插入对齐字节 (padding)。


我们都知道取内存地址的时候是通过下标去获取的，所以当 height 位于 sex 后面，但是 sex 只占有一个字节，但是 height 有两个字节的时候，如果要精确获取 height 的起始位置，那么必须要按照 uint16 的大小的倍数去获取，也就是 2 的倍数，那么 sex 后面肯定要跟一个空值才行，所以 height 的偏移量就是 2

position 也是一样，它的大小是 8，那么要获取它的位置，至少是 8 的倍数才行，所以前面要补充两个空值

具体的每个字段的偏移量具体分析如下：

- Sex bool 占 1 字节，起始偏移为 0
- Height uint16 需要以 2 字节对齐，所以从偏移 2 开始
- Addr byte 占 1 字节，但前面已插入 1 字节 padding，所以从偏移 4 开始
- Position byte 紧接着 Addr 的下一字节，所以偏移 5
- age int64 需要以 8 字节对齐，所以从偏移 8 开始
- weight uint16 需要以 2 字节对齐，并且要跟在 age 之后，所以从偏移 16 开始

也就是说整个的结构体按照空值来算一共是 18 个字节的占有量，那么整个结构体是多少呢？

结构体的大小计算方法是最大值的倍数，这里就是 age 的 8 的倍数，必须要大于实际值 18，所以这里取 24

## 下面对这个结构体进行优化
### 顺序优化去除 padding
```go
type Person struct {
  Sex bool // 0
  Addr byte//1
  Postion byte // 2
  Height uint16 // 4
  weight uint16 // 6
  age int64 // 8
}
```
现在我们看一下具体的内存排布
```bash
sex addr postion * height height weight weight age age age age age age age age
```
可以看到这个时候的一共需要的数据只有 16，刚好结构体的大小是 8 的倍数，那么这个时候结构体的大小就变成了 16 字节

### 数据类型调整
我们可以使用更小的类型去承载数据，比如使用 uint8 去替代 uint16，和 int64

```go
type Person struct {
  Sex bool // 0
  Addr byte//1
  Postion byte // 2
  Height uint8 // 3
  weight uint8 // 4
  age uint8 // 5
}
```
整个结构体的大小就变成了 6 个字节
### 测试优化顺序的结构体大小

```go
package main

import (
	"fmt"
	"unsafe"
)

// 原始结构体
type OriginalPerson struct {
	Sex     bool   // 0  offset
	Height  uint16 // 2
	Addr    byte   // 4
	Postion byte   // 5
	Age     int64  // 8
	Weight  uint16 // 16
}

// 优化后的结构体
type OptimizedPerson struct {
	Sex     bool   // 0 offset
	Addr    byte   //1
	Postion byte   // 2
	Height  uint16 // 4
	weight  uint16 // 6
	age     int64  // 8
}

type OptimizedPerson1 struct {
	Sex     bool  // 0
	Addr    byte  //1
	Postion byte  // 2
	Height  uint8 // 3
	weight  uint8 // 4
	age     uint8 // 5
}

func main() {
	// 测试原始结构体大小
	fmt.Println("原始结构体大小:", unsafe.Sizeof(OriginalPerson{}))
	// 测试优化后结构体大小
	fmt.Println("优化后结构体大小:", unsafe.Sizeof(OptimizedPerson{}))
	// 测试优化后结构体大小1
	fmt.Println("优化后结构体大小:", unsafe.Sizeof(OptimizedPerson1{}))
}
```
```bash
原始结构体大小: 24
优化后结构体大小: 16
优化后结构体大小1: 6
```