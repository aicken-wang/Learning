### 分布式版本控制系统

- Git是分布式的SCM ,SVN是集中式的

- Git每一个历史版本存储完整的文件，SVN存储文件差异

- Git可离线完成大部分操作，SVN则相反

- Git有着更优雅的分支和合并实现

- Git有更强的撤销修改和修改版本历史的能力

- Git速度更快，效率更高

  #### Git 的安装 和简单配置

  - Git安装[https://git-scm.com/downloads](https://git-scm.com/downloads)

- 新手指南[https://git-scm.com/book/zh/v2](https://git-scm.com/book/zh/v2)

  ```shell
  which -a git
    git --version
    vim .bash_profile
    ```

#### 创建项目

```shell
#查看本地的远程仓库
git remote
#未添加远程仓库显示为空
#添加远程仓库
git remote add origin git@github.com:xxxx
#查看本地的远程仓库
git remote
#显示远程仓库origin
//origin
#将本地的master分支推送至远程origin仓库
git push origin master
#创建开发分支并切换至开发分支dev
git checkout -b dev
#推送一个dev分支至远程仓库origin
git push origin dev

#Download Github project
git clone git@github.com:xxxxx

```

#### 开发

```shell
#创建开发components分支开发模板
git checkout -b createComponents
#查看当前版本状态
git status
#添加到本地缓存区
git add .
# 提交到本地仓库
git commit -m "提交components"
#切换到开发分支
git checkout dev
#合并分支 --on-ff 会将合并的日志添加到git Log中，不加只会在dev中打上合并的tag
git merge createComponents --on-ff
# 查看日志
git log
# 推送至远程dev
git push origin dev
# 查看现有的分支
git branch
# 删除分支createComponets
git branch -d createComponets

# 模拟数据工具
# https://www.charlesproxy.com/
# https://software-download.microsoft.com/download/pr/MediaCreationTool1903.exe
```

- git 原理

  ````shell
  #查看 add 命令 
  git --help add
  # HEAD、master、与branch
  # HEAD：当前commit的引用
  # branch： 对commit的引用
  # 正向工作 本地到远程
  # 工作区  - git commit -a ->本地库 - git push ->远程仓库：origin
  # 工作区  - git add - > 暂存区 - git commit ->本地库 - git push ->远程仓库：origin
  # git add | git rm | git rm --cached | git rm -f
  # 反向工作 远程到本地
  # 工作区  <- git checkout | git diff - 暂存区 <- git diff --cached -本地库 <- git fetch | git pull->远程仓库：origin
  #工作区  <- git checkout HEAD | git diff HEAD -本地库
  ````

  

  git 目录

  ```shell
  # config  配置文件
  $ cat config
  # HEAD  分支信息 
  $ git branch
  $ cat HEAD
  # objects
  $ echo "test" | git hash-object -w --stdin
  9daeafb9864cf43055ae93beb0afd6c7d144bfa4
  $ find .git/objects/ -type f
  $ git cat-file -p 9daeafb9864cf43055ae93beb0afd6c7d144bfa4
  $ echo "version" > tag.txt
  $ git hash-object -w tag.txt
  $ git cat-file -t 088eda41aa61dc62fefef5d183a1f703bb01bfa6
  blob
  $ git cat-file -p master^{tree}      
  # logs
  $ git log | git log --pretty=oneline
  #info
  # refs
  ```

  - 命令注释

  ```shell
  fetch 抓取远程仓库
  pull == fetch + merge  抓取且合并
  merge 合并分支
  git branch -a 列出全部分支包括远程分支
  演示合并冲突
  rebase 变基 //将分支直接移植到主分支，危险
  # 需要创建哪些分支
  production 分支
  develop分支
  feature 分支
  release 分支
  hotfix 分支
  
  # 分支命名规范
  开发阶段
  dev-mmdd-需求名称
  测试阶段
  test-mmdd
  发布阶段
  tag-mmdd
  
  # 回退操作
  # 查看尚未缓存的文件
  $ git diff
  # 显示工作版本和HEAD的差异
  $ git diff HEAD
  # 比较两个历史版本之间的差异
  $ git diff SHA1 SHA2
  # 查看已经暂存起来的文件（staged）和上次提交时的快照之间（HEAD）的差异
  $ git diff --cached
  # 撤销修改（暂存区 git reset HEAD file 参数 soft mixed hard）
  git reset --hard branch名
  本质是移动HEAD 以及它所指向的branch
  reset --hard & reset --soft & reset 不加参数
  --soft:重置位置的同时，保留工作目录和暂存区的内容，并把重置HEAD的位置所导致的新的文件差异放到暂存区
  --mixed:(默认) 重置位置的同时，保留工作目录的内容，并清空暂存区
  --hard:重置位置的同时，清空工作目录的所有改动
  git checkout branch名的本质其实是把HEAD 指向指定的branch，然后迁出这个branch所对应的commit的工作目录
  # git reset 和 git checkout差异
  checkout 和 reset 都可以切换HEAD的位置，他们除了有许多细节的差异外，最大的区别在于 reset在移动HEAD时会带着它所指向的branch 一起移动，而checkout不会
  #暂存修改（stash) 临时存放工作目录的改动
  git stash
  git stash pop
  git stash -u
  # 查看暂存 （stash list)
  # 使用暂存（git stash apply stash@(1)/stash pop）
  # 清空暂存 （stash clear）
  # git reset [HEAD]
  $ git reset 3d845524c1db2fe15a99f0f1227a0bbfa385c52b [HEAD]
  #revert
  $ git revert HEAD^
  checkout
  
  解决冲突 
  先手工合并差异文件
  $ git add 差异文件
  $ git commit -m "tag 注释"
  $ git push origin 分支名
  ```

  