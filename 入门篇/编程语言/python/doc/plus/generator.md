# 生成器
## [目录](./summary)
## generator 也就是生成器的英文拼写，它的主要作用是生成大批量的数据

- 方法一 (x for x in ['a', 'v'])

其实也就是把上一章迭代方法中的[]换成了()，那么返回的对象就不同了，前者是生成了一个list后者是生成了一个生成器。

其实跟js中的generator是一样的，打印出来这个生成器的内容只需要使用`next()`方法就OK了

```python
l = ( x for x in ['1','a','b'])
next(l)
>>> 1
next(l)
>>> a
next(l)
>>> b

```

直到计算到最后一个元素，没有更多的元素时，抛出StopIteration的错误。

```python
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
StopIteration
```

## generator 函数

如果一个函数定义中包含yield关键字，那么这个函数就不再是一个普通函数，而是一个generator：

> js中需要将函数名前面添加 * 例如 function \*age(){}

```python
def file():
    while n < 100:
        yield

```


## 定义一个generator，依次返回数字1，3，5：

```python

def odd():
    print('step 1')
    yield 1
    print('step 2')
    yield(3)
    print('step 3')
    yield(5)
    ```


调用该generator时，首先要生成一个generator对象，然后用next()函数不断获得下一个返回值：

```python

>>> o = odd()
>>> next(o)
step 1
1
>>> next(o)
step 2
3
>>> next(o)
step 3
5
>>> next(o)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
StopIteration

```

也就是说它每次使用next方法当做容器并且运行到yield就停止，再次调用next接着上次的继续运行。跟js中是一样的。

定义的函数，其次每次返回的都是generator对象。


我们基本上从来不会用next()来获取下一个返回值，而是直接使用for循环来迭代

```python

def file():
    print('step1')
    yield 1
    print('step2')
    yield 2
    return '应该停止了'
    print('2')
    yield 123


for a in file():
    pass

```

证明了几点

- return后面肯定要停止。
- yield要停止一次
- for循环可以直接遍历generator无需再使用next
- 使用for循环无法取得return和yield返回的值

对于最后一个问题，我们可以使用这个办法解决。

```python

while true:
    try:
        x = file()
        print('数据是：',x)
    except StopIteration as e:
        print('数据是:',e.value)
        break

```


## 迭代器


我们已经知道，可以直接作用于for循环的数据类型有以下几种：

- 一类是集合数据类型，如list、tuple、dict、set、str等；

- 一类是generator，包括生成器和带yield的generator function。

这些可以直接作用于for循环的对象统称为可迭代对象：**Iterable**。

可以使用isinstance()判断一个对象是否是Iterable对象：

```python
from collections import Iterable

isinstance([],Iterable)
>>> True

```

可以被next()函数调用并不断返回下一个值的对象称为迭代器：Iterator。

检查是否是 interator
```python

from collections import interator

isinstance([],interator)
>>> false


```

这两个就是检查interable迭代器和生成器的差别。

- interable
- interator

生成器都是interator对象，但list、dict、str虽然是Iterable，却不是Iterator


将interable变成Interator,使用全局函数iter()即可


```python
a = iter(['1', '2'])

isinstance(a,Interator)

>>> True
```

### 你可能会问，为什么list、dict、str等数据类型不是Iterator？

这是因为Python的Iterator对象表示的是一个**数据流** ，Iterator对象可以被next()函数调用并不断返回下一个数据，直到没有数据时抛出StopIteration错误。可以把这个数据流看做是一个有序序列，但我们却不能提前知道序列的长度，只能不断通过next()函数实现按需计算下一个数据，所以Iterator的计算是**惰性的**，只有在需要返回下一个数据时它才会计算。

Iterator甚至可以表示一个无限大的数据流，例如全体自然数。**而使用list是永远不可能存储全体自然数的。**
