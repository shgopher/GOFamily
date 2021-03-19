# 模块

## 模块的引入

python


```python
from xx import xx
```

javascript 原生


```javascript
import xx from xx
```

nodejs


```javascript
require('')

当然目前的nodejs也支持了js中原生的方式

import xx from xx
```

golang


```golang
import(
  xx
)
```

## 模块的建立

每一个.py 文件都是一个模块，模块很好的避免了变量的命名问题，如果不是一个包内的文件，即便是明明是一样的也无所谓。

假如模块名冲突了怎么办？没事情我们还有包。一般而言我们的很多模块都在一个包内。这一点要注意，在golang中我们的包内部有一个或者多个文件，但是这些文件都是包，并没有跟Python一样的两极问题，也就是包，包的下级是模块，golang中一律都是模块。

```python

- tools
    - __init__.py
    - add.py
    - plus.py
    - max.py



```
这里面,文件__init__.py是必须的，如果没有就会被解析为普通的目录，并且只有它的模块名就是包名，其余的都要用tools.add来使用，注意两点

- 模块名后面是没有.py的例如上面提到的tools.add并不是tools.add.py
- \__init__这个模块，并不能使用tools.\__init__,因为他的作用是判别包，所以它自己的模块名和包名一致，使用的时候就是tools，并且它可以是有代码的，当然也可以是空的都是可以的。

包还可以包内包

```python
- tools
  - web
      __init__.py
      - example.py
  -plus
      __init__.py
      - max.py
  __init__.py
```


当然我们也能看出来，如果作为包内包每个包都要有一个__init__.py来申明"我是包"这件事情。同样的这个文件的模块名就是包名，并且内容可有可无。

```python
# 使用的时候

tools.web.example
````

这就是三级的包。


- 正常的函数和变量名是公开的（public），可以被直接引用，比如：abc，x123，PI等；

- 类似_xxx和__xxx这样的函数或变量就是非公开的（private），不应该被直接引用，比如_abc，\__abc.之所以我们说，private函数和变量“不应该”被直接引用，而不是“不能”被直接引用，是因为Python并没有一种方法可以完全限制访问private函数或变量，但是，从编程习惯上不应该引用private函数或变量。

- 类似__xxx__这样的变量是特殊变量，可以被直接引用，但是有特殊用途，比如上面的__author__，\__name__就是特殊变量，hello模块定义的文档注释也可以用特殊变量__doc__访问，**我们自己的变量一般不要用这种变量名**；


```python
def _name(arg):
  return 1

def _names(arg):
  return 2

def file():
    return _names + _names

这样就避免了内部的names和name被暴露在外部。
```

## __main__

```python
if __name__ == '__main__':
    test()
```

当我们在命令行运行hello模块文件时，Python解释器把一个特殊变量__name__置为__main__，而如果在其他地方导入该hello模块时，if判断将失败，因此，这种if测试可以让一个模块通过命令行运行时执行一些额外的代码，最常见的就是 **运行测试**。

我们可以用命令行运行hello.py看看效果：

```python

$ python3 hello.py
Hello, world!
$ python hello.py Michael
Hello, Michael!
```


如果启动Python交互环境，再导入hello模块：

```python
$ python3
Python 3.4.3 (v3.4.3:9b73f1c3e601, Feb 23 2015, 02:52:03)
[GCC 4.2.1 (Apple Inc. build 5666) (dot 3)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> import hello
>>>

```
导入时，没有打印Hello, word!，因为没有执行test()函数。

调用`hello.test()`，才能打印出Hello, word!

在 Python 的官方推荐的代码样式中，还有一种单下划线结尾的样式，这在解析时并没有特别的含义，但通常用于和 Python 关键词区分开来，比如如果我们需要一个变量叫做 class，但 class 是 Python 的关键词，就可以以单下划线结尾写作 class_

双下划线开头双下划线结尾的是一些 Python 的“魔术”对象，如类成员的 __init__、__del__、__add__、__getitem__ 等，以及全局的 __file__、__name__ 等。 Python 官方推荐永远不要将这样的命名方式应用于自己的变量或函数，而是按照文档说明来使用



这个跟上面有点类似。\_用作被丢弃的名称。按照惯例，这样做可以让阅读你代码的人知道，这是个不会被使用的特定名称。举个例子，你可能无所谓一个循环计数的值：
```py
n = 42
for _ in range(n):
    do_something()
```
这点跟go很像。

是因为导入时解释器确实对单下划线开头的名称做了处理。如果你这么写from <module/package> import \*，任何以单下划线开头的名称都不会被导入

以双下划线做前缀的名称（特别是方法名）并不是一种惯例；它对解释器有特定含义。Python会改写这些名称，以免与子类中定义的名称产生冲突。Python documentation中提到，任何__spam这种形式（至少以两个下划线做开头，绝大部分都还有一个下划线做结尾）的标识符，都会文本上被替换为_classname__spam，其中classname是当前类名，并带上一个下划线做前缀。

## 总结_ 和 __
- 单下划线是私有，不可被其他包引用。
- 双下划线是不可被子类继承

也就是说__ 是private _ 是 protect，也就是说`__`私有性要大于`_`


## 关于导入的进一步解释

- 如果是系统内部的模块，直接使用`import xx`即可，因为它存在于系统会自动查找的路径中，并且通过pip下载的第三方模块也会被加入可寻找到的路径中。

- 如果是本地模块
## from module import sm 和import的区别

区别就是 前者只是引入了一小部分，后者完全引入。

## sys.path.append('')

首先要明白它是运行时添加，运行过后就没了，所以也不存在什么缓存什么，

还有使用它的时候要先引入sys包。


## \__name__属性

一个模块被另一个程序第一次引入时，其主程序将运行。如果我们想在模块被引入时，模块中的某一程序块不执行，我们可以用__name__属性来使该程序块仅在该模块自身运行时执行。

```py
#!/usr/bin/python3
# Filename: using_name.py

if __name__ == '__main__':
   print('程序自身在运行')
else:
   print('我来自另一模块')
运行输出如下：
$ python using_name.py
程序自身在运行
$ python
>>> import using_name
我来自另一模块
>>>

```
**说明： 每个模块都有一个__name__属性，当其值是'__main__'时，表明该模块自身在运行，否则是被引入。**


下面插入一行来自注明Python项目 requests的代码，看看是如何导入模块的。

```py

import copy
import time
import calendar
import collections

from ._internal_utils import to_native_string
from .compat import cookielib, urlparse, urlunparse, Morsel

try:
    import threading
except ImportError:
    import dummy_threading as threading


class MockRequest(object):
```
我们注意看，前几个都是在路径中可以取到的模块，也就是系统内置模块或者是pip下载的第三方模块。

下面的`from .`或者是 `.compat`意思都是在本包内部，从一个模块引入另一个模块.
