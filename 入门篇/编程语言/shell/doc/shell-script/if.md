
# if
## [目录](https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)

## 用法
if commands; then
     commands
[elif commands; then
     commands...]
[else
     commands]
fi

---

如果 if 之后跟随一系列命令，则将计算列表中的最后一个命令：
```bash
if false; true; then echo "It's true."; fi

```

这里的话，我们if后面有两个命令，false或者是true那么它用true不使用false。

## 文件表达式

先看个例子

```bash

#!/bin/bash
# test-file: Evaluate the status of a file
FILE=~/.bashrc
if [ -e "$FILE" ]; then
    if [ -f "$FILE" ]; then
        echo "$FILE is a regular file."
    fi
    if [ -d "$FILE" ]; then
        echo "$FILE is a directory."
    fi
    if [ -r "$FILE" ]; then
        echo "$FILE is readable."
    fi
    if [ -w "$FILE" ]; then
        echo "$FILE is writable."
    fi
    if [ -x "$FILE" ]; then
        echo "$FILE is executable/searchable."
    fi
else
    echo "$FILE does not exist"
    exit 1
fi
exit
```
我们可以看到在if后面接了一个`[]`这个东西里面就是文件表达式

下面给大家全部的文件表达式的参数及其代表意思

|表达式|	如果下列条件为真则返回True|
|-|-|
|file1 -ef file2	|file1 和 file2 拥有相同的索引号（通过硬链接两个文件名指向相同的文件）。
|file1 -nt file2	|file1新于 file2。
|file1 -ot file2	|file1早于 file2。
|-b file	|file 存在并且是一个块（设备）文件。
|-c file	|file 存在并且是一个字符（设备）文件。
|-d file	|file 存在并且是一个目录。
|-e file	|file 存在。
|-f file	|file 存在并且是一个普通文件。
|-g file	|file 存在并且设置了组 ID。
|-G file	|file 存在并且由有效组 ID 拥有。
|-k file	|file 存在并且设置了它的“sticky bit”。
|-L file	|file 存在并且是一个符号链接。
|-O file	|file 存在并且由有效用户 ID 拥有。
|-p file	|file 存在并且是一个命名管道。
|-r file	|file 存在并且可读（有效用户有可读权限）。
|-s file	|file 存在且其长度大于零。
|-S file	|file 存在且是一个网络 socket。
|-t fd	|fd 是一个定向到终端／从终端定向的文件描述符 。 这可以被用来决定是否重定向了标准输入／输出错误。
|-u file	|file 存在并且设置了 setuid 位。
|-w file	|file 存在并且可写（有效用户拥有可写权限）。
|-x file	|file 存在并且可执行（有效用户有执行／搜索权限）。

也就是说在shell中的if判断时，文件表达式还是很常见的，并且能表达的意思也很多。如果是真的就返回0否则就返回非零，只有是0的时候shell才会认为是对的。

## `fi`
我们可以看到有几个if就有几个fi，其实fi就是反过来的if啊，所以说他们就是一个正 的一个反过来的，代替了`{}`
  ```bash
  if true;then
    echo "hello-world"
  elif false; then
    echo "wrong"
  fi
  ```

### 第二种：

  ```bash
  if commands;then
    commands
  else
    commands
  fi
  ```
  我们可以看出，不仅是有fi并且fi还要和if在同一行，这样看起来才标准，并且要注意并不是我们常用的`else if`它直接是`elif`并且每个if或者是elif后面都要加上 then并且then和if并不是一行（一般写成一行加上;即可）

###  测试字符串表达式

|表达式|	如果下列条件为真则返回True
|-|-|
|string|	string 不为 null。
|-n string|	字符串 string 的长度大于零。
|-z string|	字符串 string 的长度为零。
|string1 = string2  <br/>string1 == string2|string1 和 string2 相同。 单或双等号都可以，不过双等号更受欢迎。
|string1 != string2	|string1 和 string2 不相同。
|string1> string2|	sting1 排列在 string2 之后。
|string1 < string2|	string1 排列在 string2 之前。

```bash
if [ -n "string" ]; then
  echo "nice"
fi
```
##测试整数表达式

|表达式|	如果为真
|-|-|
|integer1 -eq integer2|	integer1 等于 integer2。
|integer1 -ne integer2|	integer1 不等于 integer2。
|integer1 -le integer2|	integer1 小于或等于 integer2。
|integer1 -lt integer2|	integer1 小于 integer2。
|integer1 -ge integer2|	integer1 大于或等于 integer2。
|integer1 -gt integer2|	integer1 大于 integer2。

```bash
if [ 29 -eq $nice ]; then
  echo "ddd"
fi
```

## if和正则联系在一起

使用方式

```bash
if [[ "aa"=~正则 ]]; then
  #statements
fi
```
也就是有两个关键的地方
- [[]]
- something=~正则

> 小知识：正则表达式中 +和*的区别，
> \+ ：{1,}
> \*：{0,}
> 要记得这个差别。

### [[ ]]添加的另一个功能是==操作符支持类型匹配，

```bash
if [[ $something == *.jpg ]]; then
  echo "aa"
fi
```
或者是：
```bash
if [[ $something == fool.* ]]; then
  commands
fi
```
## (())

(( ))被用来执行算术真测试。如果算术计算的结果是非零值，则一个算术真测试值为真


因为复合命令 (( )) 是 shell 语法的一部分，而不是一个普通的命令，而且它只处理**整数**， 所以它能够通过名字识别出变量，而不需要执行展开操作

## command1 && command2 command1 || command2

```bash
mkdir number1 && mkdir number2
```

理解这些操作很重要。对于 && 操作符，先执行 command1，如果并且只有如果 command1 执行成功后， 才会执行 command2。对于 || 操作符，先执行 command1，如果并且只有如果 command1 执行失败后， 才会执行 command2。

> 我们要注意||在shell和其他编程语言的不同，
> 在其它语言中||表示任意一个对就可以，就算是全对了也是可以的，但是shell中必须是第一个错第二个对才可以，

## 复习一个知识点

重定位方式：
```bash
number> filename
```
- 0 输入流
- 1 输出流
- 2 错误流

```bash
cd /1212 2> appp.text
cat 0< app.le  1> app.md

```
其实这样也是一样复制文件内容的方式。

## ||

 [ -d temp ] || mkdir temp

 只有前面错误后面才会执行

 这点跟一般的|| 有着很大的不同，一般情况下的|| 只要双方有一个是正确的就是整体正确了，这个是shell中||跟一般编程语言中很大的不同

 ## 重要提示⚠️

类似于 [] 是一个表达式，其计算结果为真或假。这个[[ ]]命令非常 相似于 [] 命令（它支持所有的表达式），并且她**增加了**一个重要的新的字符串表达式
```bash
[[$string=~正则]]
```
