# 一些工具

## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)

- ## 这些命令有
    - c

    - which显示某个文件的安装位置

    - apropos显示比较合理的后续的命令猜测

    - info显示info命令

    - whatis显示一个命令的简单简洁

    - alias创建命令的别名
- ## type

    - 用法：`type command`会显示出来某个命令的具体类型。
    ```bash
    type cd
    # cd is a shell builtin
    ```
    所以也就是说type是可以显示命令类型的命令
- ## which

    - which含义就是显示某个文件所在的位置

    - 用法：`which somefile`

    - which不对别名起作用

- ## help

    - help就是帮助你看到每个命令的命令参数

    - help用法`help commond`

    ```bash
    help cd
    # cd: cd [-L|-P] [dir]
    # Change the current directory # to DIR.  The variable $HOME is # the
    # default DIR.  The variable # CDPATH defines the search path # for
    # the directory containing DIR.  # Alternative directory names in # CDPATH
    # are separated by a colon (:).  # A null directory name is the # same as
    # the current directory, i.e. # `.'.  If DIR begins with a # slash (/),
    # then CDPATH is not used.  If # the directory is not found, # and the
    # shell option `cdable_vars' is # set, then try the word as a # variable
    # name.  If that variable has a # value, then cd to the value of # that
    # variable.  The -P option says # to use the physical directory # structure
    # instead of following symbolic # links; the -L option forces # symbolic links
    # to be followed.

    ```
    - 注意表示法：出现在命令语法说明中的方括号，表示可选的项目。一个竖杠字符 表示互斥选项。后面的dir表示可能出现的选项例如一个路径等等。

    - 有些命令可以直接使用`commond --help`来使用不过并不是所有的命令都有要注意了。

- ## man

    - man显示某个程序的手册，可以理解为type的加强版。

    - man显示的某个命令的全部信息，信息非常大非常全，但是不能作为教材，只能作为字典。如果忘记了某个程序的某个参数

    ```bash
    man cd
    # 就会出现一大堆的解释文段。
    ```
    - `man commond file`其中这个file是匹配文件，它会找到第一个匹配的地方。


- ## apropos

    - 类似于搜索引擎。进行关键字的搜索 是对于命令的搜索

    - 相对的来说grep是对于文件里的搜索，find是找文件名。这三者的区别要自己拿捏。

- ## whatis
    - 程序显示匹配特定关键字的手册页的名字和一行命令说明
- ## README

    - gzip 软件包包括一个特殊的 less 版本，叫做 zless，zless 可以显示由 gzip 压缩的文本文件的内容

    - 在/usr/share/doc中通常有文件安装的readme

- ## alias

    - 创建命令的别名
    ```bash
    alias th='cd ~/Desktop; ls; cd ./code'
    # 这样我们发现就设置了一个快捷方式
    ```
    - 所以alias设置别名其实就是设置快捷方式，并且要记得两点：
        - 使用''
        - shell中如果不是语意问题不要设置空格
        ```bash
        echo ""#此时空格需要因为是语义问题
        alias=''#此时不要空格，因为不存在语义问题，如果有空格报错。
        ```

        ```bash
        thomasdeMacBook-Air:code thomashuke$ type th

        # 查看语句看到详细的说明内容在上面。
        th is aliased to `cd ~/Desktop; ls; cd ./code'
        ```
- ## unalias
    - 我想你会知道它的意思的。😄

    - `unalias th`删除就好了。
## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
