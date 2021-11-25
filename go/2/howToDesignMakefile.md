# Makefile 编写规范
> 一般 Makefile 要大写，这样会更容易看到这个文件

## Makefile的基本语法

基本语法
```bash
target:prerequisites
    command
```
- target ,可以是一个目标文件，也可以是一个可执行文件，还可以是一个标签，多个目标用空格分割
- prerequisites ,代表生成该target的依赖项，多个项目，可以用空格分割
- command ,代表target要执行的命令

只要tragets不存在，或者prerequisites中有一个以上的文件比targets文件新，那么commond就会执行，从而产生我们需要的文件，或者执行我们期望的操作

小tips：
- `@` 不会打印出要执行的命令
- `-` 忽略命令的出错

Makefile中的通配符：
- `*`
- `?`
- `~`

伪目标：不会为该目标生成任何文件的目标，例如
```bash
clean:
    rm hello.c
```
伪目标，make无法生成他的依赖关系,使用 `.PHONY:clean`

```bash
.PHONY:clean
clean:
    rm hello.c
```
如果不是伪目标，make可以生成依赖关系的

```bash
hello:hello.o
    gcc -o hello hello.o
hello.o: hello.c
    gcc -c hello.c
```
这里，上面的`:hello.o` 和下面的`hello.o:` 这就是依赖关系。

order-only 依赖

意思就是 prerequisites不止一个，我们需要某些prerequisite 改变才会重新构造target，不是全部，所以才会说是 order-only

```bash
hello: normal-prerequisites| order-only-prerequisites
    rm hello.c
```
只有第一次构造targets的时候，才会使用 order-only-prerequisites 后面，即使 order-only- 发生改变的时候，也不会重构targets

所以说这里的，normal-prerequisites 发生改变的时候，才会重新构造targets

make的变量赋值：
- 形式1: `ROOT_HERE = github.com/shgopher/GOFamily`
- 形式2: `ROOT_HERE := github.com/shgopher/GOFamily ` 建议使用这种方法代替 `=` 因为 `=` 容易出现赋值风险
- 形式3: 如果该变量没有被赋值，那么赋予等号后面的值 `ROOT_HERE ?= linux_amd64` 等于一种default的给定值的形式
- 形式4: `ROOT_HERE += linux_arm64` 表示，将后面的值添加到

使用变量的方法是: `@()` 或者 `@{}`

make的多行变量，我们通过define 设置多行变量，**这也是function的定义方法**，只不过function应用的时候使用call 在前面

```bash
define DO:
    Options:
        v
        h
        d
    endef
```
如果想让这个变量让子make引用的话，使用export即可 `export DO`

make还有自己内置的定义变量

- MAKE 当前make解释器的文件名称
- MAKECMDGOALS 命令行中指定的目标名
- CURDIR 当前make解释器的工作目录
- NAKE_VERSION make解释器的版本
- MKEFILE_LIST
- .DEFAULT_GOAL
- .VARIABLES
- .FEATURES
- .INCLUDE_DIRS

make中的自动化变量
- $@ 表示规则中的目标文件集
- $% 表示目标成员名 比如 foo.a(bar.o) ,`$%`:`bar.o` ,`%@`: `foo.a`
- $< 
- $?
- $^
- $+
- $| 所有的order-only依赖目标的集合，用空格分割
- $* 这个命令用的最多的，例如 `path/here.doiop.do` 目标模式是path/here%.d,那么$*的值就是  path/here

条件语句：

```bash
ifeq ($(Here),)
    echo xxx
else
    echo xxx
endif
```
- ifeq 表示相等
- ifneq 表示不等
- ifdef 表示定义
- ifndef 表示没有定义

make预定义函数：
- $(origin <variable>)
- $(addsuffix <suffix>,<names>)
- $(addprefix <prefix>,<names>)

等

引入其他的makefile：

```bash
include scripts/make-rules/common.mk

include scripts/make-rules/golang.mk
```

- 绝对路径，相对路径 `./script/i.mk`
- 指定路径，`make -I /xxx.mk`
- 使用目录路径 <prefix>/include

当然，也可以使用 `*` 通配符来引入一个文件夹中的所有文件

使用 `- include xdxx` 可以忽略这引入错误

## Makefile 的设计布局

原则：分层设计，大概意思就是在root路径有一个主 Makefile,在其它路径下有子 makefile，主 makefile 调用子makefile

makefile也经常要调用shell script，如果脚本比较复杂就使用shell脚本，如果比较简单的逻辑，直接编辑makefile就可以了。

下面给出makefile的基本布局：

![](https://gitee.com/shgopher/img/raw/master/makefile.png)

## Makefile 编写技巧

1. 多用通配符和自动变量
    
    例如： tools.go.%
    - tools.go.do
    - tools.go.swagger
    - tools.go.run

    当我们定义了tools.go.%,我们给它设置一个规则的时候，下面三个都可以去套用一个规则

2. 善于用函数

    - 优先使用make的自带函数
    - 如果可以，也可以自己去编写函数

3. 依赖需要用到的工具

    你的某个步骤，需要某个工具，你需要先检查是否有这个工具，如果没有就启动下载

4. 把常用功能放在 /Makefile中，不常用的放在子类的makefile中

5. 编写可扩展的makefile

    - 在不改变 makefile结构的情况下添加新功能
    - 扩展项目的时候，新功能可以自动纳入 makefile 现有的逻辑中

    前者可以规划布局搞定，后者，我们就需要制定大量的例如通配符，自动变量，函数等技巧来实现这个功能

6. 将所有的输出存放在一个目录下，便于查找和删除
    
    - 一般放在 _output 类似这种目录中

7. 使用带有层级的命名方式

    例如 tools.verify.swagger

    这样可以很好的分组

8. 做好目标拆分

    比如一个操作步骤，拆分成两个，比如说 将 “使用某工具” 和 “检查某工具是否存在” 分开 那么就可以灵活的运用，当检测到有工具的时候就不用安装了，可以大大的加快速度

9. 设置 options

    将一些可变的功能，通过options来控制

10. 定义环境变量


