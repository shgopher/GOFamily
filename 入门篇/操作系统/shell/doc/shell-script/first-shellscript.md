# 第一个脚本
## [目录](./summary.md)
## 解释什么是shell script

最简单的解释，一个 shell 脚本就是一个包含一系列命令的文件。shell 读取这个文件，然后执行 文件中的所有命令，就好像这些命令已经直接被输入到了命令行中一样

Shell 有些独特，因为它不仅是一个功能强大的命令行接口,也是一个脚本语言解释器。我们将会看到， 大多数能够在命令行中完成的任务也能够用脚本来实现，同样地，大多数能用脚本实现的操作也能够 在命令行中完成。

## 怎么写脚本

- 编写一个脚本。 Shell 脚本就是普通的文本文件。所以我们需要一个文本编辑器来书写它们。最好的文本 编辑器都会支持语法高亮，这样我们就能够看到一个脚本关键字的彩色编码视图。语法高亮会帮助我们查看某种常见 错误。为了编写脚本文件，vim，gedit，kate，和许多其它编辑器都是不错的候选者。

- 使脚本文件可执行。 系统会相当挑剔不允许任何旧的文本文件被看作是一个程序，并且有充分的理由! 所以我们需要设置脚本文件的权限来允许其可执行。

```bash
-rwxr--r--
# 就OK了。或者直接设置
chmod a+x file_path
```


- 把脚本放置到 shell 能够找到的地方。 当没有指定可执行文件明确的路径名时，shell 会自动地搜索某些目录， 来查找此可执行文件。为了最大程度的方便，我们会把脚本放到这些目录当中。

```bash
#!/usr/bin/env bash
```
这个 shebang 被用来告诉操作系统将执行此脚本所用的解释器的名字。每个 shell 脚本都应该把这一文本行 作为它的第一行。

例如文中的这个shebang就是说使用bash作为shell解释器，当然你也可以设置其它的解释器。

## 如果让你的脚本变成可执行文件
- 使用chmod
- 讲它加入到path变量中

```bash
echo $PAHT

# 查看变量

/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
一共有5个地址。
```

## 变成可执行文件
- 使用chmod
- 自我设置几个目录讲其添加到path环境变量中

 ```bash
 export PATH=path_name:"$PATH"
 ```
 还记得吗 export的作用
 是讲shell变量提取出来添加到环境变量中，其实它的作用就是讲你自己写的命令添加到Linux环境中。实现从用户到系统的过程。

 - mv file /usr/local/bin

讲文件转移到这个文件夹中

我们看一个例子

```bash
cd ~/Desktop
mkdir bin
cd bin
touch text.sh
chmod 766 ./text.sh
export PAHT=~/Desktop/bin:"$PATH"#不要用‘’,否则$就别翻译为普通的文本了
text.sh #就可以执行了
```
或者是
```bash
touch text.sh
chmod 766 ./text.sh
./text.sh #可以执行
```
```bash
touch text.sh
chmod 766 text.sh
mv text.sh /usr/local/bin
text.sh #可以执行了
```

## 长短选项

当在命令行中输入选项的时候，短选项更受欢迎，但是当书写脚本的时候， 长选项能提供可读性。


## 例子

```bash

find playground \
    \( \
        -type f \
        -not -perm 0600 \
        -exec chmod 0600 ‘{}’ ‘;’ \
    \) \
    -or \
    \( \
        -type d \
        -not -perm 0711 \
        -exec chmod 0711 ‘{}’ ‘;’ \
    \)
```
里面很清楚末尾都有一个\\它的作用就是注释掉换行符，当然括号前面的\就是真的为了讲命令中的（转译为真的（
