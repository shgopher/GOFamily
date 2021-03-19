# 函数

## 函数的调用

filename()

- max()函数（或者说是方法）返回参数中最大的那个，其实用其它语言实现起来也很容易。

- int() str() 等这几个函数也很常见
- 函数的名是可以传递的，它就是一个变量罢了，也就是换个还是指向那个门的新钥匙也是无所谓的。

## 函数的定义

```python
def filename():
    commonds
    return
```
## 语句占位符

pass

```python
if x > 0:
    pass

def hu():
    pass
```
## 函数模块的导入

```python
from filename import anothername # filename不要加.py 两者都不要加''
```
## 函数返回多个值

如同golang中一样Python中是可以同时返回多个值的
```python
return x, y, z
```
```golang
return x, y, z
```


原来返回值是一个tuple！但是，在语法上，返回一个tuple可以省略括号，而多个变量可以同时接收一个tuple，按位置赋给对应的值，所以，Python的函数返回多值其实就是返回一个tuple，但写起来更方便。

但是golang中的多个值就是多个值，所以go语言才是真实的返回多个返回值的语言。

## 默认参数
```python

def power(x, n=2):
    s = 1
    while n > 0:
        n = n - 1
        s = s * x
    return s
```
这样，当我们调用power(5)时，相当于调用power(5, 2)

这就是给定默认值，在js中也有几乎类似的实现

```javascript
function name({x, y = 12} = {1, 34}){

}
```
在js中甚至你所有参数都可以不要，所有的参数都给你有默认值，
也就是说你可以单一的有默认值，也可以一起有默认值，当然这个算是结构赋值的用法了。

必选参数在前，默认参数在后，否则Python的解释器会报错，因为如果你默认在前面的话难道你要搞个`(,nam)`这样的样子吗？怎么可能呢。。。。

如果有默认值，那么默认值一定是排列在队尾的。
## ⚠️

定义默认参数要牢记一点：默认参数必须指向不变对象

如果是可变的，那么变量就处于一直在变化的环境中。

```python
def file([x, y, z]):
    comm
```
事实证明这种行为Python不能接受，它顶多可以接受

```python

def file(l = []):
    # COMBAK:
        return
```


## 不变对象 str None

设计str、None，不变对象一旦创建，对象内部的数据就不能修改，这样就减少了由于修改数据导致的错误。此外，由于对象不变，多任务环境下同时读取对象不需要加锁，同时读一点问题都没有。我们在编写程序时，如果可以**设计一个不变对象，那就尽量设计成不变对象**。

## 可变参数

注意哦是可变的参数，

我们把函数的参数改为可变参数：

```python
def calc(*numbers):
    sum = 0
    for n in numbers:
        sum = sum + n * n
    return sum
```

定义可变参数和定义一个list或tuple参数相比，仅仅在参数前面加了一个*号。在函数内部，参数numbers接收到的是一个tuple，因此，函数代码完全不变。但是，调用该函数时，可以传入任意个参数，包括0个参数：

```python

>>> calc(1, 2)
5
>>> calc()
0

```

## 关键字参数

可变参数允许你传入0个或任意个参数，这些可变参数在函数调用时自动组装为一个tuple。而关键字参数允许你传入0个或任意个含参数名的参数，这些关键字参数在函数内部自动组装为一个dict
```python
def person(name, age, **kw):
    print('name:', name, 'age:', age, 'other:', kw)

person('thomashuke', 30, {'nu':12, 'rr':1221})
# 或者是
a = {'sd':12, 'dsdsds':1122}
person(1, 2, a)
```
要注意这个关键字参数在内部是组合为一个dict然而前面那个可变参数被组合成一个tuple，或者是list

## 命名关键字参数

对于关键字参数，函数的调用者可以传入任意不受限制的关键字参数。

```Python
def file(x, y, *, job, city)
```
那么你后面只能使用job和city

```Python

def file(x, y, *name, job, city) # x,y是位置参数，name是可变参数(list tuple) job和city是关键字参数
```
## 参数组合

参数需要一定的顺序
必选参数，默认参数，可变参数，命名关键字参数，关键字参数

> 当中它们有的可以省略，但是顺序要按照这个顺序来排列。

举个例子

```Python
def file(x, y = 12, *a,**j):
    pass
    return
a = [1, 2, 2]
或者是 a = (1, 2, 3)
j = {'a':12}
file(12, 12, 12,j)
```
file(12, a, j)是错误的中间不能间隔
另外例如*a
可以不传入一个list或者tuple可以直接传入 file(1,2,'12','12',city='12',name='12')

```Python

def file(x,b = 12, *, a, **a):
file(12, 45, a=12, d=12, c=12)
```

**最神奇的是通过一个tuple和dict，你也可以调用上述函数**

```Python
def file(a, b = 12, *arg, **args):
    print(a, b, arg, args)
    return
arg = (1, 2, 3, 4)
args = {'a':12, 'b': 12}
    file(*arg, **args)
    1, 2, (3, 4) {'a':12, 'b':12}
    ```
    说起来函数的参数这方面Python也是相当的灵活。

    
