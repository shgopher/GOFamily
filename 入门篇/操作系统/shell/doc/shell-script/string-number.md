# 字符串和数字
## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
## 参数展开
```bash
fool=12
echo "${fool}_file"

12_个人

```
熟不熟悉？

这就跟js或者是某些模板引擎的"${} #{}是一样的"
## 访问位置参数
我们知道一般访问只能访问到9但是使用`${}`我们可以往后面访问。

## ${parameter:-word}

若 parameter 没有设置（例如，不存在）或者为空，展开结果是 word 的值。若 parameter 不为空，则展开结果是 parameter 的值

## ${parameter:=word}

若 parameter 没有设置或为空，展开结果是 word 的值。另外，word 的值会赋值给 parameter。 若 parameter 不为空，展开结果是 parameter 的值。

> 位置参数或其它的特殊参数不能以这种方式赋值。


## ${parameter:?word}

若 parameter 没有设置或为空，这种展开导致脚本带有错误退出，并且 word 的内容会发送到标准错误。若 parameter 不为空， 展开结果是 parameter 的值。

## ${parameter:+word}

若 parameter 没有设置或为空，展开结果为空。若 parameter 不为空， 展开结果是 word 的值会替换掉 parameter 的值；然而，parameter 的值不会改变。

## ${!prefix*}  ${!prefix@}

这种展开会返回以 prefix 开头的已有变量名。根据 bash 文档，这两种展开形式的执行结果相同

## ${#parameter}

展开成由 parameter 所包含的字符串的长度。通常，parameter 是一个字符串；然而，如果 parameter 是 @ 或者是 * 的话， 则展开结果是位置参数的个数。

## ${parameter:offset}

## ${parameter:offset:length}
这些展开用来从 parameter 所包含的字符串中提取一部分字符。提取的字符始于 第 offset 个字符（从字符串开头算起）直到字符串的末尾，除非指定提取的长度。

## ${parameter#pattern}

## ${parameter##pattern}


这些展开会从 paramter 所包含的字符串中清除开头一部分文本，这些字符要匹配定义的 pattern。pattern 是 通配符模式，就如那些用在路径名展开中的模式。这两种形式的差异之处是该 # 形式清除最短的匹配结果， 而该 ## 模式清除最长的匹配结果。

## ${parameter%pattern}

## ${parameter%%pattern}


这些展开和上面的 # 和 ## 展开一样，除了它们清除的文本从 parameter 所包含字符串的末尾开始，而不是开头。

## ${parameter/pattern/string}

## ${parameter//pattern/string}

## ${parameter/#pattern/string}

## ${parameter/%pattern/string}

这种形式的展开对 parameter 的内容执行查找和替换操作。如果找到了匹配通配符 pattern 的文本， 则用 string 的内容替换它。在正常形式下，只有第一个匹配项会被替换掉。在该 // 形式下，所有的匹配项都会被替换掉。 该 /# 要求匹配项出现在字符串的开头，而 /% 要求匹配项出现在字符串的末尾。/string 可能会省略掉，这样会 导致删除匹配的文本。



知道参数展开是件很好的事情。字符串操作展开可以用来替换其它常见命令比方说 sed 和 cut。 通过减少使用外部程序，展开提高了脚本的效率。举例说明，我们将修改在之前章节中讨论的 longest-word 程序， 用参数展开 ${#j} 取代命令 $(echo $j | wc -c) 及其 subshell

## time 命令

测试文件执行效率

```bash

time ./zhankai.sh
12_个人

real	0m0.018s
user	0m0.003s
sys	0m0.003s

```

## 使用参数展开也就是${}的好处

优点很明显，就是可以提高执行速度，就如同正则一样，缺点就是不容易想到。。。需要熟练的运用。。。

## declare

- declare -u upper
- declare -l lower

转化变量 upper和lower 前者无论是什么均转换成大写，后者是小写。

### 大小写转换参数展开（当然效率更高）
|格式|	结果
|-|-|
|${parameter,,}|	把 parameter 的值全部展开成小写字母。
|${parameter,}	|仅仅把 parameter 的第一个字符展开成小写字母。
|${parameter^^}|	把 parameter 的值全部转换成大写字母。
|${parameter^}	|仅仅把 parameter 的第一个字符转换成大写字母（首字母大写）。

## 算术展开

它被用来对整数执行各种算术运算。它的基本格式是：

$((expression))

我们看过八进制（以8为底）和十六进制（以16为底）的数字。在算术表达式中，shell 支持任意进制的整型常量。

|表示法|	描述
|-|-|
|number|	默认情况下，没有任何表示法的数字被看做是十进制数（以10为底）。
|0number|	在算术表达式中，以零开头的数字被认为是八进制数。
|0xnumber|	十六进制表示法
|base#number	|number 以 base 为底

## 算术运算符
|运算符|	描述
|-|-|
|+|	加
|-|	减
|*|	乘
|/|	整除
|**|	乘方
|%	|取模（余数）

因为 shell 算术只操作整型，所以除法运算的结果总是整数：
```bash
echo $(( 5 / 2 ))
2
```
这使得确定除法运算的余数更为重要：
```bash
echo $(( 5 % 2 ))
1
```
## 更现代的[[]](())代替[]()的原因

 = 符号的真正含义非常重要。单个 = 运算符执行赋值运算。foo = 5 是说“使得 foo 等于5”， 而 == 运算符计算等价性。foo == 5 是说“是否 foo 等于5？”。这会让人感到非常迷惑，因为 test 命令接受单个 = 运算符 来测试字符串等价性。这也是使用更现代的 [[ ]] 和 (( )) 复合命令来代替 test 命令的另一个原因。


### 赋值运算符
更现代的bash
例如+= \*= ++ 都可以实现了
## 注意⚠️

**但是要记得请放到更现代的容器中执行
 也就是执行的时候要$(())**


## 比较运算符

|运算符|	描述
|-|-|
|<=|	小于或相等
|>=|	大于或相等
|<|	小于
|>|	大于
|==|	相等
|!=|	不相等
|&&|	逻辑与
|![p](../../picture/fu.png)|逻辑或|
|expr1?expr2:expr3|	条件（三元）运算符。若表达式 expr1 的计算结果为非零值（算术真），则 执行表达式 expr2，否则执行表达式 expr3。

##  bc - 一种高精度计算器语言
```bash
bc -q
1 + 3
4
12 + 122121
122133
```
