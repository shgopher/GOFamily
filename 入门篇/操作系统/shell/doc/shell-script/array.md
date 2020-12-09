# 数组
## [目录](./summary.md)
## 基本用法

数组变量就像其它 bash 变量一样命名，当被访问的时候，它们会被自动地创建。这里是一个例子

```BASH
a[1]=foo

echo ${a[1]}
foo
```
当然也可以使用declare -a a来创建数组。

## 赋值

- name[subscript]=value
- name=(value1 value2)
