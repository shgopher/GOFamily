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

用到的git命令有：

- 切换/添加分支的 git checkout  和 git checkout -b
- push 功能的 ,添加tag push的 git push --tag,和push 分支的 git push origin (分支名称)
- merge功能的 git merge --no-ff (来源分支)
- 还有普遍的 git commit ,git add 
- tag功能的 git tag -a 

通常来说，我们的开发流程是这样的：

1. 在master主分支上先新建一个develop的分支
2. 切换到develop分支，然后再新建一个开发分支，比如叫做feature/hello-world
3. 在开发分支上开发代码，开发完成以后，再次fetch 远程代码，然后commit（通常可以合并commit：使用 git rebase -i <要修改的commit id 父id>），发布到GitHub。
4. 提交一个requests，经过代码的review以后，经过上级领导merge到develop上
5. 然后在develop上，这个时候将要发布版本了，那么就从这个时候的develop上新建一个分支是release/v1.11.0 类似这样的分支，然后在这个分支上开始测试工作，这个过程中可以修改，然后最终完成以后，merge到 develop和maser分支，
6. 删除 feature分支和release分支

## forking
> 适合开源项目

从远程仓库，fork一份到自己的仓库，开发完成后，提requests