# 编码

## 编码要闻

所谓编码，是因为我们要把汉字，英文单词等，转化为二进制的数字，因为计算机只认识数字。
最开始的编码是`ascii`,当时还只能储存英文和某些字符，但是因为中国等其它国家并不使用英语，所以我们开始有了自己的编码，但是这个时候就出现了问题，因为它会不能避免的出现一些重合的东西，这个时候就出现了乱码，然后世界开始使用了`unicode`，不过它有个缺点就是它总是两个字节的储存东西，所以会造成资源的浪费，这个时候又出现了著名的`utf-8`它有个很好的地方就是兼容了ASCII，英语就使用ASCII如果是汉字的花就使用多字节储存，所以现在世界通用的编码就是`utf-8`。

通常在计算机的内存中使用的是Unicode编码，在硬盘中使用的是utf-8的编码。

```html
<meta charset="utf-8">
```
这个意思就是说，当浏览器去处理这个DOM树的时候使用utf-8的编码去处理。

## 那么在脚本语言Python3中又用了什么编码呢？

答案是**Unicode**

Python提供了2个方法来管理字符和数字

- chr( )--- 函数把编码转换为对应的字符
- ord( )--- 函数获取字符的整数表示

```Python

print(ord('托'))

25176

print(chr(25176))

托
```

## 改变

由于Python的字符串类型是str，在内存中以Unicode表示，一个字符对应若干个字节。

如果要在网络上传输，或者保存到磁盘上，就需要把str变为以字节为单位的bytes。

Python对bytes类型的数据用带b前缀的单引号或双引号表示：
```bash
x = b'ABC'

```
以Unicode表示的str通过`encode()`方法可以编码为指定的bytes，

反过来，如果我们从网络或磁盘上读取了字节流，那么读到的数据就是bytes。要把bytes变为str，就需要用`decode()`方法

- encode() 从`str`字节变成`bytes（b''）`
- decode() 从`bytes`变成`字节 ('')`


要计算str包含多少个字符，可以用`len()`函数：
```bash
>>> len('ABC')
3
>>> len('中文')
2
```
> 注意len是计算多少字符并不是多少字节。

## 注意 ⚠️

len()函数计算的是str的字符数，如果换成bytes，len()函数就计算字节数：
```bash
>>> len(b'ABC')
3
>>> len(b'\xe4\xb8\xad\xe6\x96\x87')
6
>>> len('中文'.encode('utf-8'))
6

```

也就是说 英语中一个字符就是一个字节，中文中一个字符通常是3个字节。所以中文占用的空间要大一些。
## 保存文件
由于Python源代码也是一个文本文件，所以，当你的源代码中包含中文的时候，在保存源代码时，就需要务必指定保存为UTF-8编码。当Python解释器读取源代码时，为了让它按UTF-8编码读取，我们通常在文件开头写上这两行：

```bash
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
```
前者让类Unix系统知道，这是个Python3执行文件

后者是说这个文件中使utf-8去读取源代码。

## 字符的替换

```bash

'hello %s' % '11'
'hello 11'

```
这是在Python命令行中演示的，通常情况下它不使用print也可以，并且使用这个方式来进行替换

## 注意 ⚠️
但是直接使用` "somethings..." 这种方式` 在Python非脚本使用中是无法输出的。要记得。
```bash
print('dfd%s' % 'ss')

dfdss
```
要使用print（）进行输出
## 常见的占位符有：

|%d|	整数
|-|-|
|%f|	浮点数
|%s|	字符串
|%x|	十六进制整数

一般情况下直接使用%s就OK了，统统转换成字符，其实并没有什么问题。

看这个

```bash
'hus %s %%' % 'dff'

输出

'hus dff %'
```

这个就表明了：如果想让输出%只要使用%%这种方式，那么使用\%可以吗？我们试试

```bash

'hus %s \%' % 'sd'
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: incomplete format

```

错误类型是格式不完整，所以说，这种方法不可以。

```bash
r = (85 - 72) / 72

>>> print("sdfds: %.2f" % r)
sdfds: 0.18

```
那么如果我们直接将r写入这个print中可以吗？

我们试试

```bash
print("sdfdsf:%.2f"  % (85 - 72) / 72)

提示错误


TypeError: unsupported operand type(s) for /: 'str' and 'int'

```
结果识别类型错误

## 我们看看js中如果要嵌入一个东西该怎么使用

```js
console.log(`sdsssdds ${(12-1)/123} dedsdsd`)
sdsssdds 0.08943089430894309 dedsdsd
```
没错使用`${}`即可 如果是jade使用`#{}`即可。就可以进行替换符了。
