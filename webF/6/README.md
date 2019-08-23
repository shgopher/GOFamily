# 水平居中和上下居中

- 竖直方向：

父元素为 td 或 th 时，vertical-algin即可。
或者设置为table-cell，要么就使用一个一劳永逸的办法设置元素的positon然后设置top值，只要值设置为百分比。
就不怕大小变来变去。

- 水平方向

1 设置postion 2 设置元素的左右margin为auto也是可以的。

这是一般盒子模型的设置模式，如果想使用flex，模型或者是 grid双维度模型可以参考本单元的flex使用方法。
