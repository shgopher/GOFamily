# 错误处理

我们要知道这几种错误
- bug
- 用户带来的错误
- 无法预知的系统带来的错误

第一种不用说一定要解决，第二种，想办法检测出来指引用户改正，第三种比如说网络断了，这种错误我们还是要有一定的处理方式的。

## try except finally体制

当然你也能在js中看到这套错误处理机制。这在高级语言中是很常见的错误处理机制。不过js中是catch

```py
try:
    print(....)
except ZeroDivisionError as e:
    print(e)
finally:
    print(...)
print('fdfdfd')

```
这个就是大致的过程。

## logging 模块

内置的logging模块，相当强大，不错不错。

## raise

raise的意思是抛出，其实可以类比 throw。

```py

class FooError(ValueError):
    pass

def foo(s):
    n = int(s)
    if n==0:
        raise FooError('invalid value: %s' % s)
    return 10 / n

foo('0')

```

```py
def bar():
    try:
        foo('0')
    except ValueError as e:
        print('ValueError!')
        raise

bar()

```

在这个print后面我们还用了raise，这是为什么呢？其实就是我们在打印了错误后向上层抛出了错误。

## pdb

启动Python的调试器pdb，让程序以单步方式运行，可以随时查看运行状态。

$ python3 -m pdb xx.py


## pdb.set_trace()

这个方法也是用pdb，但是不需要单步执行，我们只需要import pdb，然后，在可能出错的地方放一个pdb.set_trace()，就可以设置一个断点

## 文档测试

当我们编写注释时，如果写上这样的注释：

```py

def abs(n):
    ''' # 使用这个符号，里面的信息只有测试的时候才会被调用。一般的时候这些代码并不会执行。
    Function to get absolute value of number.

    Example:

    >>> abs(1)
    1
    >>> abs(-1)
    1
    >>> abs(0)
    0
    '''
    return n if n >= 0 else (-n)
```

Python内置的“文档测试”（doctest）模块可以直接提取注释中的代码并执行测试。
