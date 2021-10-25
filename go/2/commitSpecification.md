# Commit规范
commit遵循以下原则：

1. 清晰的知道每一个commit的变更内容
2. 可以基于commit进行过滤查找，比如git log --oneline --grep "^feat|^fix|^perf"
3. 基于规范的 commit 生成 commit log
4. 根据规范的某种类型的 commit 触发发布或者构建的自动化流程
5. 根据commit的不同类型，可以自动生成版本号

下面介绍一下Angular的commit规范，此规范较为标准。
> `<>` 代表必选项，`[]`代表了可选项，空行是必须的，`:`后面的空格也是必须的，限制字符在50，72，或者100以内。
```bash
<type>[(optional scope)]: <subject description>
// 空行
[optional body]
// 空行
[optional footer]

```
示例：

```bash
fix(apiserver): test for kafka

this is a test word

this is a footer
```
## Header的不同类型：
> Production: 新增功能类型，这种requests一定要好好review
> Development类型，不改变功能，只是构建工具的改变等类似这种免测试的改变

|类型|类别|说明|
|:---:|:---:|:---:|
|feat|production|新增功能|
|fix|production|bug修复|
|perf|production|提高代码的性能|
|style|development|代码格式变化|
|refactor|production|其它类型，例如简化代码，重命名变量，删除冗余代码等优化工作|
|test|development|新能测试例子|
|ci|development|持续集成或者部署相关|
|docs|development|文档类的更新|
|chore|development|其它类型，例如构建流程，依赖管理，或者辅助工具的改变等工作|

快速定位commit的类型：

![](https://gitee.com/shgopher/img/raw/master/angular-commit-header-type1.png)

![](https://gitee.com/shgopher/img/raw/master/angular-commit-header-type2.png)

## scope的意思
大概可以理解为分类，比如不同的组建分类，不同的功能分类等。
比如我们使用不同的功能作为scope，例如说：
- docs
- apiserver
- changelog
- readme
## subject
subject是这个commit的tile描述的东西，开头必须用动词，并且用一般现在时，结尾不要加标点符号
## body
body就是这个commit的具体描述了
## footer
footer并不是必须的，一般来说明不兼容的改动，和关闭的issue，如果是涉及到了不兼容的改动，一定要加上`BREAKING CHANG: `

关闭的issue，新建一行，写上关闭的issue号码即可。

例如：

```bash
fix(apiserver): change some xx

change one ,change two, change three

BREAKING CHANG: change xxxxxxxxxxxx

Closes #1178
```
## 特殊commit
如果当前的commit还原了之前的commit，那么应该这么写：

```bash
revert: 要还原的commit header

This reverts commit hash值
```

举例：

```bash
revert: docs(docs): xxxxxx

This reerts commit k79380c8cfc8308a8a6e18f9yc8b8114febc9b48a
```
## commit提交频率以及合并提交
提交频率：
- 只要有改动就提交commit
- 按照每天的固定时间点提交commit

git rebase -i 命令，可以合并commit，这样可以看起来更简洁一点。

它的命令如下：

|命令|目的|
|:---:|:---:|
|p,pick|不对该commit做任何处理|
|r,reword|保留该commit，但是修改提交信息|
|e,edit|保留该commit，但是rebase的时候会暂停，让你去修改这个commit|
|s,squash|保留该commit，将这个commit合并到上一个commit|
|f,fixup|相同于squash，但是这个commit信息不保留|
|x,exec|执行其它shell命令|
|d,drop|删除该commit|

## 切换分支

`git checkout `

其中，如果是创建分支就是

`git checkout -b xxxx`

不加`-b` 就是切换的意思了，加了就是添加的意思。

切换完分支以后别忘了用 `git log --oneline` 去查看一下当前的commit情况，其中 `--oneline` 就是把只取commit的tile，省略大部分详细内容，变成一行。类似这样的

```bash
9a713f8 Update README.md
8859ec6 Merge pull request #13 from shgopher/n1
b7a044e Update README.md
f8028fb Update README.md
f38197a Update README.md
901eaf5 Update README.md
70c49e4 Create README.md
22d4f0f Delete README.md
```

如果不加那么就会变成这样

```bash
Author: 科科人神 <shgopher@aliyun.com>
Date:   Wed Oct 13 20:18:44 2021 +0800

    docs(go): test rebase
    
    docs(go): test rebase1
    
    docs(go): test rebase2

commit 9e63c4f2c6af46d06a4ada2b4bdc4034247e4b1b
Author: 科科人神 <shgopher@aliyun.com>
Date:   Wed Oct 13 19:53:31 2021 +0800

    docs(go): add some subject rules
    
    rules:  commit rules, README.md rules, doc file rules

```
## 撤销之前的commit，并添加新的commit
如果之前的commit过于多，那么也不用非要rebase，可以直接先取消，然后再添加。

```bash
git reset HEAD~3

git add .

git commit -am "feat(go): add xxx"
```
注意HEAD~3 就是取消的个数。
## 变更commit信息
- 如果是只变更最近的一次 `git commit -amend`
- 如果是变更好早之前的，那么就只能使用 `git rebase -i 父commit_id`这种方法了。这种方法变化commit信息后，变化的commit的id就会改变了，不是原来的id了。
## commit 信息自动化
- [commitizen-go](https://github.com/lintingzhen/commitizen-go) 交互式的自动生成commit message
- commit-msg 检查commit message，这个属于自己写脚本搞定的事情，
- [go-gitlint](https://github.com/llorllale/go-gitlint) 检查历史提交的commit message是否符合规范
- [gsemver](https://github.com/arnaud-deprez/gsemver) 语意化版本自动生成工具
- [git-chglog](https://github.com/git-chglog/git-chglog) 根据commit msg 自动生成 CHANGELOG.md 文件
