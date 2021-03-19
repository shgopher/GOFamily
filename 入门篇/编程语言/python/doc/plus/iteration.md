# 迭代

## 什么是迭代？

例子：

```python
for a in b.value():
    pass

for x, y in de.items():
    pass


```

## 字符串也可以😀

```python

for a in 'abcv':
    pass
```

## 如何判断是否可以迭代

请看代码

```python
from collections import Iterable

isinstance('abc', Iterable)
True
```

### 这里就引出了模块这个东西，稍后会继续写，这里先简单的说一下

使用方法请看
```python
from source import something

```

注意

- source 不能加后缀名

## 双变量

```python
for x, y in list:
    pass
```
这种行为在Python中很常见，就如同Python可以返回多个return 参数一样(返回的其实是一个省略的tuple)，Python的双变量也是很常见，并且很好用的行为。

#### 类比golang

```golang

for i := 0; i < count; i++ {

}

但是同时还可以，

for _, r := range s {
fmt.Printf("%c,", r)
    }
}
这也是一种多变量遍历的实例

```
## 迭代生成式

```python

list(range(1, 11))
[1, 2, 3, ......11]

```

```python

[x * x for x in range(1, 11)]

[1, 4, 9 ....]

```

第二种方法是 前面的是对于后面x的一种简单的计算，就是对于迭代的每个数字进行什么计算，然后返回回来就是这个意思。

当然你也可以使用


```python

[x + x for x in range(1, 12)]

[2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22]

```
这都是可以的，就如同js中

```javascript
[1, 2 ,....].filter(() => {
  x + x
})

```
是一样的意思，只不过，形式不同，仅此而已。


**甚至你还可以这样**

```python

[x + y for x in 'abc' for y in 'bbadc' ]
```

```python
[x + '=' + y for x, y in d.items()]

['x=y']
```
这些也都是可以的。可以说Python是非常灵活，并且资源非常丰富的一个编程语言。

总结一下这个迭代的写法

```python

[x... for x in ...]

```

注意最前面不论x怎么样一定要有x，就算是直接输出x也要有x

## 举个例子

```python

[s.lower() for x in ['add', ' DDff']]

```
## 结合for
```python

l = ['adDd',1, 'SDD']
[s.lower() for x in l if isinstance(x, str)]

```
