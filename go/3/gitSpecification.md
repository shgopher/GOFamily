# git 规范
- 集中式工作流
- 功能分支工作流
- git flow 工作流
- forking
## 集中式工作流
大家 clone 远程的 master分支，然后本地开发，然后提交到远程服务器，如果遇到无法提交的错误了，就自己修改，避免冲突。这种方法回发现出现很多commit，并且很杂乱
## 功能分支工作流
大家 从master分支上，开辟新的分支 `git checkout -b xxx `，然后各自开发，最后提交requests，然后再merge到master
## git flow 工作流
> 适合非开源，内部项目，不需要fork，只需要开辟分支即可。

git flow 工作流，拥有五种模式，分别是：
- master 主分支，发布状态
- develop 开发状态的分支，不能做开发工作，这是一个合并用的分支
- feature 研发阶段用来作功能开发，通常来说是从develop上fork的，并且起名字是feature/xxx 这种方式，提交到develop分支
- release 发布阶段的预发布分支，通常是fork 自 develop分支的，然后提交到master和develop分支
- hotfix 紧急bug修复分支 ，从master分支进行fork的，然后提交到master和develop分支
## forking
> 适合开源项目

从远程仓库，fork一份到自己的仓库，