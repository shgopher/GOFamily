# go项目编程规范

- 非编码类规范
    - 开源规范
    - 文档规范
    - 版本规范
    - Commit规范
    - 发布规范

- 编码类规范
    - 目录规范
    - 代码规范
    - 接口规范
    - 日志规范
    - 错误码规范 

## 开源规范
![](https://gitee.com/shgopher/img/raw/master/openSourceAgreement.png)

开源协议的各种规范，各种限制看这张图就足够了。

下面是项目，特别是开源项目，需要具备的规范。
- 项目结构合理
- 严格遵守代码规范，比如变量的命名，
- 代码的质量：比如良好的算法，运用得当的设计模式，合理的代码长度
- 单元测试覆盖率要高
- 版本发布要规范，遵守版本的标号，例如 v1.1.3 ， 要有版本号文件，通常是 CHANGELOG.md
- 向下兼容，高版本要兼容低版本的东西
- **详细的文档说明** 
- 安全的代码，不要出现涉密的内容，比如内部的密码等
- 完善的example，好的example胜过文档说明
- 良好的Commit 提交规范
- 保证项目的可用性，每一个版本都要经过严格的测试才能放出来
- 及时处理 requests
- 做好推广以及讨论小组
- 遵守git规范
## 文档规范
### README 规范
1. 项目名称
2. 功能特征
3. 软件的架构图【可选】
4. 快速开始例子
5. 项目依赖包
6. 构建项目的方法
7. 运行项目的方法
8. 使用项目的方法
9. 如何贡献
10. 社区
11. 关于作者
12. 谁在用
13. 许可证
### 项目文档规范
文档通常放置在 docs/ 里面，下面是规范文档的内容

![](https://gitee.com/shgopher/img/raw/master/%E6%96%87%E6%A1%A3%E8%A7%84%E8%8C%83.png)
### API接口规范
API 作为一个文件夹进行存放内容，其中大概是包含了以下内容：

- README.md 基本介绍
- CHANGELOG.md 变更历史说明log
- generic.md 保存通用的基本参数，返回参数，等通用数据
- struct.md 列出文档中使用的数据结构
- user.md secret.md policy.md restful资料
- err_code.md 错误码的描述文件

其中 user.md secret.md 存放的就是最重要的具体接口描述，它包含了以下[内容](https://github.com/googege/iam/blob/master/docs/guide/zh-CN/api/user.md)：

- 接口的描述
- 请求方法，比如 `GET` or  `POST`等，以及path
- 输入参数
- 输出参数
- 请求的例子

## 版本规范
标准是：主版本号.副版本号.修改版本号，例如 v1.13.2
- 主版本号：大版本，假如是不再兼容，可以换大版本号，例如python 2 和 3，当主版本号更改的时候，副版本号和修改号必须清零。
- 副版本号：偶数稳定，奇数开发，兼容性的修改，当副版本号更改的时候，修改号（bug号）必须清零
- 修改号：兼容性的修复bug
## Commit规范
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
### Header的不同类型：
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

### scope的意思
大概可以理解为分类，比如不同的组建分类，不同的功能分类等。
比如我们使用不同的功能作为scope，例如说：
- docs
- apiserver
- changelog
- readme
### subject
subject是这个commit的tile描述的东西，开头必须用动词，并且用一般现在时，结尾不要加标点符号
### body
body就是这个commit的具体描述了
### footer
footer并不是必须的，一般来说明不兼容的改动，和关闭的issue，如果是涉及到了不兼容的改动，一定要加上`BREAKING CHANG: `

关闭的issue，新建一行，写上关闭的issue号码即可。

例如：

```bash
fix(apiserver): change some xx

change one ,change two, change three

BREAKING CHANG: change xxxxxxxxxxxx

Closes #1178
```
### 特殊commit
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
### commit提交频率以及合并提交
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

1
3
## 发布规范

## 目录规范

## 代码规范

## 接口规范

## 日志规范

## 错误码规范 

## 参考资料
- [go语言项目开发实战](http://gk.link/a/10ADE)