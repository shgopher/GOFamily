# for循环
## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
## for循环是比while循环更常用的东西
> 这一点从各大编程语言也能看出来
## 基本用法

**两种方案**

第一种

```bash
for variable [in words]; do
    commands
done
```
第二种

```bash
for (( i = number; i < numer; i++ )); do
  #statements
done
```
```bash

例如

for (( i = 0; i < 10; i++ )); do
  echo $i
done

# ./for.sh
# 0
# 1
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```

第二中就是普遍的方法，我们只说说第一种

## 方法一

for 命令真正强大的功能是我们可以通过许多有趣的方式创建 words 列表

```bash
for i in {1...9};do
done

或者是

for i in $*;do
done
都可以啊
for i in didd*.jpg;do
done
这种通配符的这种也是支持的。

 for i in $(strings $1); do
 done

 这种“return”的方式也是支持的啊！
```
其实这种方法是不是也很熟悉，没错 在js中 for（let i in somethong）也是这个道理，只不过这个里面都是kye而已。
