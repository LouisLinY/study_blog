

git在执行git add命令时出现以下警告：

```shell
Administrator@wlkj-0010 MINGW64 ~/Desktop/git (master)
$ git add first.txt
warning: LF will be replaced by CRLF in first.txt.
The file will have its original line endings in your working directory
```

因为Git的换行符检查功能。LF是linux下的换行符，而CRLF是enter + 换行。



Git提供了一个换行符检查功能（core.safecrlf），可以在提交时检查文件是否混用了不同风格的换行符。这个功能的选项如下：

- false - 不做任何检查
- warn - 在提交时检查并警告
- true - 在提交时检查，如果发现混用则拒绝提交

建议使用最严格的 true 选项。



假如你正在Windows上写程序，又或者你正在和其他人合作，他们在Windows上编程，而你却在其他系统上，在这些情况下，你可能会遇到行尾结束符问题。这是因为Windows使用回车和换行两个字符来结束一行，而Mac和[Linux](http://lib.csdn.net/base/linux)只使用换行一个字符。虽然这是小问题，但它会极大地扰乱跨平台协作。

Git可以在你提交时自动地把行结束符CRLF转换成LF，而在签出代码时把LF转换成CRLF。用core.autocrlf来打开此项功能，如果是在Windows系统上，把它设置成true，这样当签出代码时，LF会被转换成CRLF：

```shell
Administrator@wlkj-0010 MINGW64 ~/Desktop/git (master)
$ git config --global core.autocrlf true
```



Linux或Mac系统使用LF作为行结束符，因此你不想 Git 在签出文件时进行自动的转换；当一个以CRLF为行结束符的文件不小心被引入时你肯定想进行修正，把core.autocrlf设置成input来告诉 Git 在提交时把CRLF转换成LF，签出时不转换：

```shell
git config --global core.autocrlf input
```

这样会在Windows系统上的签出文件中保留CRLF，会在Mac和Linux系统上，包括仓库中保留LF。

如果你是Windows程序员，且正在开发仅运行在Windows上的项目，可以设置false取消此功能，把回车符记录在库中：

```shell
git config --global core.autocrlf false
```





配置邮箱和用户名：

```shell
Administrator@wlkj-0010 MINGW64 ~/Desktop/git (master)
$ git config --global user.email "louislin15@yahoo.com"

Administrator@wlkj-0010 MINGW64 ~/Desktop/git (master)
$ git config --global user.name "LouisLinY"

```



版本库（repository）的回退操作

- git log的命令拓展  --pretty=oneline

- git reset 还原版本命令

- git reflog查看命令历史

  



git log

```shell
$ git log
commit 874ab35471f41a4468303356af83dbdc23fd962d (HEAD -> master)
Author: LouisLinY <louislin15@yahoo.com>
Date:   Sun Dec 15 16:55:57 2019 +0800

    three commit, add file

commit b2bb7eaa1079e9827f9608c3744bc0e8c24cd468
Author: LouisLinY <louislin15@yahoo.com>
Date:   Sun Dec 15 16:54:18 2019 +0800

    two commit, add file

commit e819e7cbf825b22afaaaec41822c7c3f8650771d
Author: LouisLinY <louislin15@yahoo.com>
Date:   Sun Dec 15 16:51:29 2019 +0800

    first commit, new file

Administrator@wlkj-0010 MINGW64 ~/Desktop/git (master)
```



git log --pretty=oneline

```shell
$ git log --pretty=oneline
874ab35471f41a4468303356af83dbdc23fd962d (HEAD -> master) three commit, add file
b2bb7eaa1079e9827f9608c3744bc0e8c24cd468 two commit, add file
e819e7cbf825b22afaaaec41822c7c3f8650771d first commit, new file
```



git reset  --hard HEAD^

```SHELL
$ git reset --hard HEAD^
HEAD is now at b2bb7ea two commit, add file

```



git reflog

```shell
$ git reflog
b2bb7ea (HEAD -> master) HEAD@{0}: reset: moving to HEAD^
874ab35 HEAD@{1}: commit: three commit, add file
b2bb7ea (HEAD -> master) HEAD@{2}: commit: two commit, add file
e819e7c HEAD@{3}: commit (initial): first commit, new file

```



git reset

```shell
$ git reset --hard 874ab35
HEAD is now at 874ab35 three commit, add file

```



版本撤销及删除操作

- git checkout -- file   #签出
- git reset HEAD file （撤销到最近版本缓存区清空）
- git rm file



在add之前如何撤销

git checkout -- file

在add之后如何撤销

git reset HEAD file

git checkout -- file



删除文件

git rm file

如何撤销刚才的删除

git reset HEAD file

撤销清单 回到上一个版本

git checkout --- filename



删除之后直接放入版本库



git与github链接

创建密钥：ssh-keygen -t rsa -C  "louislin15@yahoo.com"











