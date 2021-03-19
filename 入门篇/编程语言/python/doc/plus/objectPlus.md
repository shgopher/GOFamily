# 面向对象进阶

## 本单元是Python3中面向对象的一个进阶

上一章只是简单的介绍了一下类的继承，多态等特种本章介绍内容如下:

- 多重继承
- 元类
- 定制类

## 具体介绍

### 接下来演示，如何给类或者是实例对象添加新的属性和方法

实例对象

```py

class Person(object):
    pass
a = Person()
a.name = '12'
```
这样就在这个实例对象上增加了一个属性name，但是类并没有增加，所以其它 的实例对象也没有这个属性

如果是想在这个类上添加也是很容易的

```py
class Person(object):
    pass
Person.name = '12'
```
这样的话它的实例对象都可以取到这个属性了。

### __slots__

也就是说这个变量指向一个tuple，在这个tuple中是允许这个对象可以添加的新的属性，如果不在这个序列中，是无法绑定的，
并且它仅仅对于本类的实例对象有效，对于继承它的类来说并没有什么用。

```py

__author__ = 'ThomasHuke'

class Person(object):
    __slots__ = ('name', 'year')

p = Person()

p.name = 12
p.year = 12
p.bbt = 12


./slots.py
Traceback (most recent call last):
  File "./slots.py", line 13, in <module>
    p.bbt = 12
AttributeError: 'Person' object has no attribute 'bbt'

```

假如我们取消了p.bbt

```py
12 12
```
一切恢复正常。

这既是__slots__的功能，可以限制类的实例对象添加的属性。

### @property

Python内置的@property装饰器就是负责把一个方法变成属性调用

```py

class Person(object):
    @property
    def width(self):
        return self.width
    @width.setter
    def width(self, value):
        self.width = value

```

```py
p = Person()
p.name = 12
p.name
>>> 12

```
如果再使用p.name= xxx 就会报错。


其实也就是说让你在给类或者是实例对象添加属性的时候使用@property 添加的时候有所限制。

## 多重继承

### 概念

如tile，

```py

class Animal(object, World):
    pass
```

其实这就是多重继承。

使用都充继承就避免了继承树那种庞大的继承关系，需要什么功能继承就是了，简单明了。

联系到js中的多重继承,多么的不优雅，，，，但是 好特么灵活，它的实现是这样的原理

a要继承 bcd类，那么先封装一个x类，x包括了bcd，使用的时候直接让a继承x就相当于继承了bcd。我曹。反正我感觉这特么的是嘛嘛嘛。。。。好吧，太过于灵活的语言就是这个吊样。。。

## __str__

管理打印

## __iter__

```py
class Fib(object):
def __init__(self):
    self.a, self.b = 0, 1 # 初始化两个计数器a，b

def __iter__(self):
    return self # 实例本身就是迭代对象，故返回自己

def __next__(self):
    self.a, self.b = self.b, self.a + self.b # 计算下一个值
    if self.a > 100000: # 退出循环的条件
        raise StopIteration()
    return self.a # 返回下一个值
class Fib(object):
def __init__(self):
    self.a, self.b = 0, 1 # 初始化两个计数器a，b

def __iter__(self):
    return self # 实例本身就是迭代对象，故返回自己

def __next__(self):
    self.a, self.b = self.b, self.a + self.b # 计算下一个值
    if self.a > 100000: # 退出循环的条件
        raise StopIteration()
    return self.a # 返回下一个值
```

如果一个类想被用于for ... in循环，类似list或tuple那样，就必须实现一个__iter__()方法，该方法返回一个迭代对象，然后，Python的for循环就会不断调用该迭代对象的__next__()方法拿到循环的下一个值，直到遇到StopIteration错误时退出循环。


## __getitem__

要表现得像list那样按照下标取出元素，需要实现__getitem__()方法

## __getattr__

当调用不存在的属性时，比如score，Python解释器会试图调用__getattr__(self, 'score')来尝试获得属性，这样，我们就有机会返回score的值

## __call__

任何类，只需要定义一个__call__()方法，就可以直接对实例进行调用,当然其实就是调用的class中定义的那个__call__

## 枚举类
```py
from enum import Enum

Month = Enum('Month', ('Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'))

```

这样我们就获得了Month类型的枚举类，可以直接使用Month.Jan来引用一个常量，或者枚举它的所有成员

## type()

类比js 你可以认为是 object.creat()创建一个新的对象。

也就是说不需要你使用
```py
class ClassName(object):
    """docstring for ."""
    def __init__(self, arg):
        super(, self).__init__()
        self.arg = arg

```
这种style了，直接使用

`type(ClassName, 继承的父class, 自己的属性方法)`,好吧，我不是很喜欢这么生成一个类。

## 动态定义class----metaclass

按照默认习惯，metaclass的类名总是以Metaclass结尾，以便清楚地表示这是一个metaclass

```py
class ListMetaclass(type):
    def __new__(cls, name, bases, attrs):
        attrs['add'] = lambda self, value: self.append(value)
        return type.__new__(cls, name, bases, attrs)
```

有了ListMetaclass，我们在定义类的时候还要指示使用ListMetaclass来定制类，传入关键字参数metaclass：

```py
class MyList(list, metaclass=ListMetaclass):
    pass
```

__new__()方法接收到的参数依次是：

- 当前准备创建的类的对象；

- 类的名字；

- 类继承的父类集合；

- 类的方法集合。
