# 讲述`list`

## list的基本描述
list的话，我想你可以把它当做是array，也就是数组

### 关于list的常用的几个方法。
- insert 在某个位置插入某个字符
- pop 删除，末尾的字符
- append 在末尾增加字符

另外值得说明的是跟golang这种语言不通，Python list也就是数组是可以有不一个类型的数据的，跟js也是一样的。
另外当要使用某个元素的时候使用`value[keyNumber]`这种格式。

### 我们看一个Python和js的对比

```js
thomasdeMacBook-Air:~ thomashuke$ node
> const a = [1,2,3]
undefined
> a[-1]
undefined
> a[1]
2

```
```Python
a = [1,2,3]
a[-1]
3
```

这一点是很不同的，在原生js中它并没有提供这个语法，不过我们可以变相的实现

```js
const a = (array,num) => {
  if (array[0] !== 'undefined' || array !== 'undefind' || array[num] !== 'undefind'){
    if (num < 0){
        return array[array.length + num]
    }
   return array[num]
  }
  return `❌❌❌`
}

```
