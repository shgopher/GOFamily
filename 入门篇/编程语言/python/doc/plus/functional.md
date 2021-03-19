# 函数式编程

## 解释含义

大概就是允许你将函数传来传去。

## 解释高阶函数

意思就是说，函数名可以像一般的变量一样可以传递，看例子：

```python
abs(-1) #求绝对值的全局函数
f = abs
f(-1)
>>> 1
```
这说明了，函数名可以像变量一样正常的传递。于此向对应的js

```javascript
const a = () => {
  console.log('hello world')
}
let b = a
b()
```
效果是一样的，这就是高阶函数的方式(只不过js中的this是有问题的，不过=>并没有this，所以也就没有这个烦恼了)

### 举个例子

```python

def add(x, y, f):
    return f(x) + f(y)

```

```javascript
 function age(f, z) {
   z(f)
 }
```

也就是传说中的回调函数。



`age(12,x => {something}`


把函数作为参数传入，这样的函数称为高阶函数，函数式编程就是指这种高度抽象的编程范式

## map && reduce

- map

```python
def file(x):
    return x * x

a = map(file,[1, 2, 3])

```

说明：

第一个参数是一个函数

第二个函数式一个interable

返回值是一个interator

在这里我们要介绍另一个函数list，它可以将一个惰性数据变成一个list

```python
list(a)
[1, 4, 9]

```

类似js中的array中的map等高级函数
`[].map((r) => {})`

#### 小技巧

如果从数字变成字符串？

```python
a = [1, 2, 3]
list(map(str,a))


>>> a = [1, 2, 3]
>>> list(map(str,a))
['1', '2', '3']

```

## 强调⚠️

`less is more ,或者是 enough is more`是一个很重要的理论。代码应该精炼，用最少的字做更多的活就是目的。

## reduce

直接看例子

```python

def file(x, y):
    return x + y

reduce(file,[1, 2, 3, 4, 5])

>>> 15


```

解释它的作用

前面有个函数，然后选取后面list中的 0 1位的数字进行计算，然后将返回的结果作为新一轮的第一个参数，选取3位置的参数，继续计算，当然这个函数要求前面的那个函数必选传入两个参数。并且要有返回值。

## 返回函数
```python

def count():
    fs = []
    for i in range(1, 4):
        def f():
             return i*i
        fs.append(f)
    return fs

f1, f2, f3 = count()

```

`结果是 999`原因是，当把count赋值给 f1 f2
f3时，每次赋值i的值都会变化，但是赋值的这个时刻，并非立即执行，所以当执行f1()f2()f3()他们的时候，i都已经循环了3次所以最后结果都输出了999

```python
分析一下好了，这个函数返回值是一个tuple,所以结果是[f, f ,f]

那么运行了三次i不就是3了吗，所以一起输出的时候立即执行i就是3了呀，所以输出的就是999了
```

也就是说
f1, f2, f3 = (f, f, f),并且f也要返回 i * i的因为是闭包，所以可以运行原来函数中的变量，刚好经过了三次赋值，那么i就变成了3，所以结果就是999

而且 这里面还有个结构赋值，也就是说可以使用

a, v, c = [1, 2, 3]

这种方式进行赋值，当然或者是

a, v, c = (1, 2, 3)

## 匿名函数lambda

在js中匿名函数通常也就是取消函数名的函数，但是在Python中需要一个名词来引用那就是lambda

```Python

lambda x: x * x

就等于

file(x):
return x * x

```

但是 Python对匿名函数的支持有限，只有一些简单的情况下可以使用匿名函数。

## 装饰器

现在，假设我们要增强某个函数的功能，比如，在函数调用前后自动打印日志，但又不希望修改某个函数的定义，这种在代码运行期间动态增加功能的方式，称之为“装饰器”（Decorator）。

- @语法

```Python

def log(func):
    def wrapper(*args, **kw):
        print('call %s():' % func.__name__)
        return func(*args, **kw)
    return wrapper

观察上面的log，因为它是一个decorator，所以接受一个函数作为参数，并返回一个函数。我们要借助Python的@语法，把decorator置于函数的定义处：

@log
def now():
    print('2015-3-25')
调用now()函数，不仅会运行now()函数本身，还会在运行now()函数前打印一行日志：

>>> now()
call now():
2015-3-25
把@log放到now()函数的定义处，相当于执行了语句：

now = log(now)
```
其实可以这么理解
@log其实就是说log单做容器，让下面正在定义的函数当做参数传入。

并且我们要做的还是调用原来的定义的那个函数就OK了。并没有别的异常。

## 偏函数

这里说的偏函数，跟数学中的偏函数并不是一个意思，所以有必要解释一下含义：

```python

import functools

int2 = functools.partial(int, base=8)

# 这样int2就获得了int的本领，并且呢，它的默认参数还是8
# 当然也就是八进制
# 这个偏函数跟下面的这个意思是一样的

def int2(x, base=8):
    return int(x, base)


```
其实也就是这个意思：

所以，简单总结functools.partial的作用就是，把一个函数的 **某些参数给固定住（也就是设置默认值）**，返回一个新的函数，调用这个新函数会更简单。

注意到上面的新的int2函数，仅仅是把base参数重新设定默认值为8，但也可以在函数调用时传入其他值。当然也就是所谓的更改默认值，不过这个也是无所谓的，当然是可以更改了呀，😝 。
