<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-04-06 19:03:14
 * @FilePath: /GOFamily/基础/interface/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 接口
> 有关泛型相关的约束内容均放置在[泛型](../泛型)这一章节，这里不再赘述

重点内容前情提要
- 接口类型的底层
- 空接口
- 接口的相等性
- 小接口的意义
- 接口提供程序的可扩展性
- 接口提供程序的可测试性
- 接口的严格对比函数的宽松

接口，go语言提供的，用于抽象以及数据解耦的组件，**在操作接口时，go语言要求的严格程度远大于函数和方法。**

在 go1.18+ 中接口的概念从包含抽象的方法变成了类型，但是我们在外部可以操作的接口仍然只能是经典接口，不过经典接口是可以当做约束在内部使用的，不过官方不推荐经典接口当做一般约束那种使用方式。

go语言为了区分，将传统接口称之为接口，将扩展的接口称之为约束，实际上传统接口是约束概念的子集。扩展的约束并不能在函数或者方法以及类型之外使用。

这一章我们只介绍经典接口的基本概念和使用方法。

## 接口类型的底层

## 空接口的使用

## 接口的相等性

## 小接口的意义

## 接口提供程序的可扩展性

## 接口提供程序的可测试性

## 接口的严格和函数的宽松对比

## issues
***interface如何判断nil***

***eface 和 iface的区别***

***如何查找interface中的方法***

***interface 设计的优缺点***
## 参考资料
- https://book.douban.com/subject/35720728/ 246页 - 286页 
- https://mp.weixin.qq.com/s/6_ygmyd64LP7rlkrOh-kRQ
- https://github.com/golang/go/blob/master/src/runtime/runtime2.go
- https://github.com/golang/go/blob/master/src/runtime/type.go