# 和标准输入流的交流

## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
> 也就是说如何跟shell输出的时候做一些交互

## 用到的命令
我们使用带有 -n 选项（其会删除输出结果末尾的换行符）的 echo 命令

## read

该命令式读取一行标准输入流

|选项	|说明|
|-|-|
|-a| array	把输入赋值到数组 array 中，从索引号零开始。我们 将在第36章中讨论数组问题。
|-d| delimiter	用字符串 delimiter 中的第一个字符指示输入结束，而不是一个换行符。
|-e|	使用 Readline 来处理输入。这使得与命令行相同的方式编辑输入。
|-n| num	读取 num 个输入字符，而不是整行。
|-p| prompt	为输入显示提示信息，使用字符串 prompt。
|-r|	Raw mode. 不把反斜杠字符解释为转义字符。
|-s|	Silent mode. 不会在屏幕上显示输入的字符。当输入密码和其它确认信息的时候，这会很有帮助。
|-t| seconds	超时. 几秒钟后终止输入。read 会返回一个非零退出状态，若输入超时。
|-u| fd	使用文件描述符 fd 中的输入，而不是标准输入。

**在read的默认变量中它的内容会输送给 `$REPLY`**

## IFS

IFS 的默认值包含一个空格，一个 tab，和一个换行符，每一个都会把 字段分割开

我们可以自定义这个IFS，然后通过我们自己设置的ifs 我们就可以很好的解决比如输入命令是有空格的但是又是一个文件的那种行为了

例子：
```bash
read int1 int2 int3 int4
IFS=":"
```
```bash
IFS=":" read user pw uid gid name home shell <<< "$file_info"
```
这里面，使用了`<<<`

这个 <<< 操作符指示一个 here 字符串。一个 here 字符串就像一个 here 文档，只是比较简短，由 单个字符串组成

### 提示⚠️
IFS 的默认值是等于这个环境变量
```bash
IFS="$OLD_IFS"
```
