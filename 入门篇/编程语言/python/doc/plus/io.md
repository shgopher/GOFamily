# io


操作系统运行用户进行读写，并不是直接的读写，而是提供给你一个接口，你通过这个对象来实现读 和 写。

## open()
```py
a = open('url', 'r')
```
url 就是你要传入的地址。opinions就是参数。

## read()

当你已经open了文件后 就开始使用 read来读文件了，然后在read()后，就需要人为的关闭close()文件。

## with语句

```py
with open('url',o) as f:
    print(f.read())
```
这时候你无须使用f.close来关闭文件了，with语句已经自动帮你解决了这个问题。


使用read(size)是每次读取一定的数量的数据，尽量不要直接使用read()这样有可能一次读取过多，内存爆掉。

readlines()是每次读取一行的数据。

> line.strip()) 把末尾的'\n'删掉

## 'rb'

后面的opinions如果是rb的话，那么读取的时候就是读二进制文件。

open还可以传入第三个数据，就是`coding='' ` 这样的如果不是utf8的话也是有办法处理好编码的。

## open的第四个参数errors=

指定发生错误的默认条件。


## 写文件

写文件和读文件是一样的，唯一区别是调用open()函数时，传入标识符'w'或者'wb'表示写文本文件或写二进制文件：

```py
>>> f = open('/Users/michael/test.txt', 'w')
>>> f.write('Hello, world!')
>>> f.close()
```
你可以反复调用write()来写入文件，但是务必要调用f.close()来关闭文件。当我们写文件时，操作系统往往不会立刻把数据写入磁盘，而是放到内存缓存起来，空闲的时候再慢慢写入。只有调用close()方法时，操作系统才保证把没有写入的数据全部写入磁盘。忘记调用close()的后果是数据可能只写了一部分到磁盘，剩下的丢失了。所以，还是用with语句来得保险：

```py
with open('/Users/michael/test.txt', 'w') as f:
    f.write('Hello, world!')

```
要写入特定编码的文本文件，请给open()函数传入encoding参数，将字符串自动转换成指定编码。


所以读取的话，都是

1. 打开文件(区别是opinions 是 r还是w)
2. 都可以使用with语句来简略close这句话。
3. 读是read() 写是write()

使用with语句操作文件IO是个好习惯~!

## stringIO和byteIO

顾名思义，stringIO和byteIO就是在内存之中读取信息前者是读取str后者是读取二级制信息。

写入。


```python
from io import StringIO
f = StringIO()
f.write('hello')
print(f.getvalue())
```

读取

```py
f = StringIO('dddd')
f.read()
```

ByteIO方法类似于StringIO

```py
from io import ByteIO
b = ByteIO()
b.write('abb'.encode('utf-8'))
print(b.getvalue())
```
读取时一样的。

```
b = ByteIO('b'xe4xb8xadxe6x96x87')
b.read(size)

```

## os

其实无论是shell还是nodejs还是golang还是如今的Python它调用的os都是系统底层给它的接口，是你的操作系统的作用，语言是没有这个能耐的。毕竟你的操作系统才是你运行的根本。

- os.name()能显示出来是类unix还是win

- os.uname()能准确的显示出来是你的操作系统，不过用之前应该先用os.name因为这个apiwin不提供。

- os.environ 提供了环境变量。注意这是一个变量并不是一个函数



操作文件和目录的函数一部分放在os模块中，一部分放在os.path模块中

### os.path

```py
# 查看当前目录的绝对路径:
>>> os.path.abspath('.')
'/Users/michael'
# 在某个目录下创建一个新目录，首先把新目录的完整路径表示出来:
>>> os.path.join('/Users/michael', 'testdir')
'/Users/michael/testdir'
# 然后创建一个目录:
>>> os.mkdir('/Users/michael/testdir')
# 删掉一个目录:
>>> os.rmdir('/Users/michael/testdir')

```

在编程过程中不要使用字符拼接，不论是Python或者是nodejs或者是golang shell等等，如果使用字符拼接的话，如果信息来自不同的操作很有可能会出错，所以尽量使用os.path.join()是比较好的方式。

同样拆分的时候使用os.paht.split()
```py
# 对文件重命名:
>>> os.rename('test.txt', 'test.py')
# 删掉文件:
>>> os.remove('test.py')

```

但是这里面并没有复制的功能，这怎么办？？？放心还有其它的模块提供这个功能呢~！~！~！

### shutil

```py

import shutil
shutil.copyfile()
```

大家来看一段代码

```py

[x for x in os.listdir('.') if os.path.isfile(x) and os.path.splitext(x)[1]=='.py']

```

这一段代码是一个生成器，不得不说Python就是优雅，
就这一行代码就有了很多的意思。

这个代码的意思是输出每个x变成这个list，并且，是x 便利 os.listdir('.')也就是本目录，后面的is是判断也就是筛查器，并且这个and也就是js中的&& 当然也是筛查器的一部分。这就是简介之处~！~！~！

## Python中的json

json,对于json的操作其实js中是最方便的，可以说json就是一个类js的东西，只不过它作为一个键值对的一种储存方式，被很多其它的变成语言也支持。

我们把变量从内存中变成可存储或传输的过程称之为序列化，在Python中叫pickling，在其他语言中也被称之为serialization，marshalling，flattening等等，都是一个意思。

序列化之后，就可以把序列化后的内容写入磁盘，或者通过网络传输到别的机器上。

反过来，把变量内容从序列化的对象重新读到内存里称之为反序列化，即unpickling。

### pickle模块来解决这个问题

```py
>>> import pickle
>>> d = dict(name='Bob', age=20, score=88)
>>> pickle.dumps(d)
b'\x80\x03}q\x00(X\x03\x00\x00\x00ageq\x01K\x14X\x05\x00\x00\x00scoreq\x02KXX\x04\x00\x00\x00nameq\x03X\x03\x00\x00\x00Bobq\x04u.'
```

pickle.dumps()方法把任意对象序列化成一个bytes，然后，就可以把这个bytes写入文件。或者用另一个方法pickle.dump()直接把对象序列化后写入一个file-like Object：

```py
>>> f = open('dump.txt', 'wb')
>>> pickle.dump(d, f)
>>> f.close()

```

下面反序列化

```py
>>> f = open('dump.txt', 'rb')
>>> d = pickle.load(f)
>>> f.close()
>>> d
{'age': 20, 'score': 88, 'name': 'Bob'}

```

要明白是取内存中的信息还是硬盘内部的，如果是内存中的使用StringIO和ByteIO,不过这个时机少，内存中一般都是非常重要或者是正在处理的文件，一般取文件都是在硬盘中所以open更常用。当然最好使用with语句。

其实序列化我们可以简单的理解为-----处理json文件。😄
