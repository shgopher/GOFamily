# 流程控制
## [目录](./summary.md)
## 流程控制用到的几个命令

- while
- until

## `while`

关于while的这个命令，我们先看基础用法：

```bash
while [[ condition ]]; do
  #statements
done
```

举例说明：

```bash
#!/usr/bin/env bash
# 测试while
read title
while [[$title=~'^(Thomas)(.jpg$)']] ; do
  echo "error！~！"
done
```
## 使用break和continue 类似一般的编程语言
## util

until 命令与 while 非常相似，除了当遇到一个非零退出状态的时候， while 退出循环， 而 until 不退出
> 也就是正确码是0非零是错误码，就算是遇到了错误码了util仍然执行。
