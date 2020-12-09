# 从上到下的模式
## [目录](https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
## 函数写作方式
```bash

# 第一种方式

function name {
    commands
    return
}

# 第二种方式

name () {
    commands
    return
}

```

## 重要

通过在变量名之前加上单词 local，来定义局部变量。

```bash
name () {
  local apple='12'
  echo apple
}
```
通过这种方式，我们定义的变量apple就不会被其他的函数所运用了。

### 小技巧

当我们想知道一个程序执行的是否成功我们可以这样写
```bash
ls /1

# 当然我们知道这个是错误的地址
# 然后我们执行 echo $?

我们就可以发现输出的值是非零

```

也就是说，加入你执行一个程序，那么立即判定$?是否是0就可以判断这个函数执行的是否正确
