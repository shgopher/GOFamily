# 数组
## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
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
