# 启动一个项目
## [目录](https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)

## 变量
定义一个变量

title="apple"

使用这个变量

$title

- 变量名可由字母数字字符（字母和数字）和下划线字符组成。

- 变量名的第一个字符必须是一个字母或一个下划线。

- 变量名中不允许出现空格和标点符号

> 在编程领域一般来说 全部大写的都代表了常亮例如PI 非全大写是常亮例如 apple-nice

关于变量是否使用“” ‘’引起来，参考引起来后他们的价值区别
分别是‘’转译所有符号 “”转译一部分 无--- 不转译。如果执意使用无那么可以使用\来转译


mv $filename ${filename}1

通过添加花括号，shell 不再把末尾的1解释为变量名的一部分。
也就是说类似于js中的\`\` \`#{app}\
1\` 就是这个意思。

```bash
cat << _EOF_
<HTML>
         <HEAD>
                <TITLE>$TITLE</TITLE>
         </HEAD>
         <BODY>
                <H1>$TITLE</H1>
                <P>$TIME_STAMP</P>
         </BODY>
</HTML>
_EOF_
```
这个意思就是把信息传递给了_EOF_

注意这个 token 必须在一行中单独出现，并且文本行中 没有末尾的空格
