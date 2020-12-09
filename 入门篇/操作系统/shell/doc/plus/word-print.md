# 文本
## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
## 这一章我们要掌握的内容是

cat – 连接文件并且打印到标准输出

sort – 给文本行排序

uniq – 报告或者省略重复行

cut – 从每行中删除文本区域

paste – 合并文件文本行

join – 基于某个共享字段来联合两个文件的文本行

comm – 逐行比较两个有序的文件

diff – 逐行比较文件

patch – 给原始文件打补丁

tr – 翻译或删除字符

sed – 用于筛选和转换文本的流编辑器

aspell – 交互式拼写检查器

nl – 添加行号

fold – 限制文件列宽

fmt – 一个简单的文本格式转换器

pr – 让文本为打印做好准备

printf – 格式化数据并打印出来

pr —— 转换需要打印的文本文件

lpr —— 打印文件

lp —— 打印文件（System V）

a2ps —— 为 PostScript 打印机格式化文件

lpstat —— 显示打印机状态信息


lpq —— 显示打印机队列状态

lprm —— 取消打印任务

cancel —— 取消打印任务（System V）

groff – 一个文件格式系统

总之内容是很多，但是无外乎都是关于文本的处理和打印输出这方面的，所以放到一章也没什么问题
## sort

|选项|	长选项|	描述
|-|-|
|-b|	--ignore-leading-blanks|	默认情况下，对整行进行排序，从每行的第一个字符开始。这个选项导致 sort 程序忽略 每行开头的空格，从第一个非空白字符开始排序。
|-f|	--ignore-case	|让排序不区分大小写。
|-n|	--numeric-sort|	基于字符串的数值来排序。使用此选项允许根据数字值执行排序，而不是字母值。
|-r|	--reverse|	按相反顺序排序。结果按照降序排列，而不是升序。
|-k|	--key=field1[,field2]|	对从 field1到 field2之间的字符排序，而不是整个文本行。看下面的讨论。
|-m|	--merge	|把每个参数看作是一个预先排好序的文件。把多个文件合并成一个排好序的文件，而没有执行额外的排序。
|-o|	--output=file|	把排好序的输出结果发送到文件，而不是标准输出。
|-t|	--field-separator=char	|定义域分隔字符。默认情况下，域由空格或制表符分隔。

通常情况下，它会根据ASCII来排序，也就是小写在前，大写在后，并且，按照字母排序(因为字母的排序恰好就是ASCII上面规定的顺序)

## 关于du
du 命令可以 确定最大的磁盘空间用户

```bash
du pathName
# 默认就是pwd。
```

但是我们需要使用到`sort 和 head`
```bash
du | sort -nr | head
```
首先显示了一堆的磁盘用户，然后使用sort -n 对于数字(-n)进行反向(-r)排序，然后取出来头部的10个数据。

如果每行有好几个数据需要按照其中一个排序的花，我们可以这样用
```bash
sort -k n

# 举例子

sort -k 5

dr-xr-xr-x   3 root        wheel  4226 Sep 17 13:58 dev
drwxrwxr-x+ 76 root        admin  2584 Sep 17 14:59 Applications
drwxr-xr-x@ 63 root        wheel  2142 Jul 23 10:36 sbin
drwxr-xr-x+ 62 root        wheel  2108 Aug 28 21:08 Library
drwxr-xr-x@ 38 root        wheel  1292 Jul 23 10:36 bin
```
> n 从1开始计算。


```bash
sort --key=1,1 --key=2n distros.txt

```
意思是第一个字符用英语排列，第二个key第二个字符使用数字排列

## uniq
忽略或者是显示重复行

其实uniq相当于只有一个命令的sort，很简略的命令

> control + d 终止标准输入
## cut

这个 cut 程序被用来从文本行中抽取文本，并把其输出到标准输出。它能够接受多个文件参数或者 标准输入。

|-c char_list|	从文本行中抽取由 char_list 定义的文本。这个列表可能由一个或多个逗号 分隔开的数值区间组成。
|-|
|-f field_list|	从文本行中抽取一个或多个由 field_list 定义的字段。这个列表可能 包括一个或多个字段，或由逗号分隔开的字段区间。
|-d delim_char|	当指定-f 选项之后，使用 delim_char 做为字段分隔符。默认情况下， 字段之间必须由单个 tab 字符分隔开。
|--complement|	抽取整个文本行，除了那些由-c 和／或-f 选项指定的文本。

cut 命令最好用来从其它程序产生的文件中 抽取文本，而不是从人们直接输入的文本中抽取

## paste
这个 paste 命令的功能正好与 cut 相反。它会添加一个或多个文本列到文件中，而不是从文件中抽取文本列。 它通过读取多个文件，然后把每个文件中的字段整合成单个文本流，输入到标准输出。类似于 cut 命令， paste 接受多个文件参数和 ／ 或标准输入
