# tuple 元组

## 解释tuple 元组的意思

tuple名字叫做元组，
```bash
name = ( "number1", "number2", "number3" )

```
它也没有append()，insert()这样的方法。其他获取元素的方法和list是一样的，你可以正常地使用classmates[0]，classmates[-1]，但不能赋值成另外的元素。

不可变的tuple有什么意义？因为tuple不可变，所以代码更安全。如果可能，能用tuple代替list就尽量用tuple。

## 注意⚠️

括号()既可以表示tuple，又可以表示数学公式中的小括号，这就产生了歧义，因此，Python规定，这种情况下，按小括号进行计算，计算结果自然是1。

所以，**只有1个元素的tuple定义时必须加一个逗号`,`，来消除歧义**

其余的使用情况跟list是一样的。

## Python 对比golang
其实这个让我想起了golang
在golang中，数组是不能被改变的，但是slice 切片是可以改变的，也就是说
- list相当于 slice
- tuple相当于 array

不过js中，并没有不能改变的数组。
## 对比一下

##### Go

```Go
a0 := [12]int{1, 2, 3} array
a1 := [...]int{1, 2, 3} array
a2 := []int{1, 2, 3} slice
```
##### Python

```Python
#/usr/bin/env Python3
a = [1, 2, 3] list
a = (1, 2, 3) tuple
```
##### js

```js
const a = [] OK了没有什么限制，异常灵活，也可以说异常容易出错。
```
##### shell
```bash
declare -a a
a 数组就被创建了

或者是

a[0]=121
直接使用，不使用创建，这样也是创建的一种办法。
```
