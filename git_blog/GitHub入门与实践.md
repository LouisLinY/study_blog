# GitHub入门与实践

## 第 1 章

### 1.1  什么是github

​       Github为开发者提供Git仓库的托管服务。这是一个让开发者与世界上的其他人共享代码的完美场所。



- **github公司与octocat**

  ​        github总部位于美国旧金山，吉祥物octocat章鱼猫

- **并不只是git仓库的托管服务**

  ​        github除提供git仓库的托管服务外，还为开发者或团队提供了一系列功能，帮助其高效率、高品质地进行代码编写。

  github的创始人之一Chris Wanstrath

- **Github的使用情况**

  ​        在git中，开发者将源代码存入名叫"git仓库"的资料库中兵加以使用。而github则是在网络上提供git仓库的一项服务。

  ​        github上公开的软件源代码全部由git进行管理。理解git，是熟练运用github的关键所在。

  

  ### 1.2  使用github会带来哪些变化

- **协作形式变化**

  ​        以程序员为代表的软件开发者之间，一直都没有一个用来辅助多人协同编程的关键性软件。因此软件开发者们往往要将版本管理系统、BUG跟踪系统、代码审查工具、邮件列表、IRC等众多工具组合在一起，以实现多人协作。

- **在开发者之间引发化学反应的Pull Request**

  Pull Request是指开发者在本地对源代码进行更改后，向github中托管的git仓库请求合并的功能。开发者可以在Pull Request上通过评论交流。通过这个功能，开发者可以轻松更改源代码，并公开更改的细节，然后向仓库提交合并请求。同时也可以选择拒绝合并。

  github的Pull Request不但能轻松查看源代码的前后差别，还可以对指定的一行代码进行评论。通过这一功能，开发者们可以针对具体的代码进行讨论，是代码审查的工作变得前所未有地惬意。

- **对特定的用户进行评论**

  任务管理和BUG报告可以通过Issue进行交互。如果想让特定的用户看到，只要用“@用户名”的格式书写，对方便会接收到通知（Notifications），查看Issue。由于也提供了Wiki功能，开发者可以轻松创建文档，进行公开、共享。wiki更新的历史记录也在git中管理，可以让用户轻松更改。

- **github flavored Markdown**

  用户所有用文字输入的功能都可以用Github Flavored Markdown（GFM）语法进行描述

  输入“@组织名”，可以让属于该组织的所有成员收到通知。“@组织名/团队”，可以让该团队的所有成员收到通知

  输入“#编号”，会连接到该仓库所对应的Issue编号，输入”用户名/仓库名#编号“，则可以连接到指定仓库所对应的Issue编号。‘

  按照上述这类特定格式书写便会自动创建链接。

- **能看到更多其他团队的软件。**

  只要将感兴趣的仓库添加至Watch中，就可以在News Feed查看该仓库的相关信息。便能在第一时间掌握最新版本的新功能或BUG修正的信息

- **与开源软件相同的开发模式**



### 1.5  GitHub提供的主要功能

- Git 仓库

- Organization

- Issue

  是将一个任务或者问题分配给一个Issue进行跟踪和管理的功能，可以像BUG管理系统或者TiDD的Ticket一样使用，每当进行Pull Request，都会同时创建一个Issue

  每一个功能更改或者修正都对应一个Issue，讨论或修正都以这个Issue为中心进行。只要查看Issue，就能知道和这个更改相关的一切信息，并以此进行管理。

  在git的提交信息中写上Issue的ID（“#7”）,github就会自动生成从Issue到对应提交的链接。另外，只要按照特定的格式描述提交信息，还可以关闭Issue。

- Wiki

  通过wiki功能，任何人都能随时对一篇文章进行更改并保存，以此可以多人共同完成一篇文章。该功能常用在开发文档或手册的编写中。

  wiki页也是作为git仓库进行管理的，改版的历史记录会被切实保存下来，使用者可以放心改写。

- Pull Request

  开发者向github的仓库推送更改或功能添加后，可以通过Pull Request功能向别人的仓库提出申请，请求对方合并。

  pull request发送后，目标仓库的管理者等人将能够查看pull request的内容及其中包含的代码更改。

  提供了对pull request和源代码前后差别进行讨论的功能。通过此功能，可以以行为单位对源代码添加评论。



## 第 2章 Git的导入

- 什么是版本管理

  版本管理就是管理更新的历史记录。（记录一款软件添加或更改源代码的过程，回滚到特定阶段，恢复误删除的文件等。）

- 集中式与分布式

  subversion将仓库集中存放在服务器中，所以只存在一个仓库。将所有的数据集中存放在服务器中，有便于管理的优点，但是一旦开发者所处的环境不能连接到服务器，就无法获取最新的源代码，开发也就几乎无法进行。服务器宕机时同样的道理，而且万一服务器故障导致数据消失，恐怕开发者就再也见不到最新的源代码

  github将仓库fork给了每一个用户。fork就是将github的某个特定仓库复制到之自己的账户下。fork出的仓库与原仓库是两个不同的仓库，开发者可以随意编辑

- 初始设置

  提高代码的可读性

  ```shell
  git config --globa color.ui auto    #让命令的输出拥有更高的可读性
  ```

- 设置SSH Key

  ```shell
  $ ssh-keygen -t rsa -C "louislin15@yahoo.com"
  Generating public/private rsa key pair.
  Enter file in which to save the key (/c/Users/Administrator/.ssh/id_rsa):
  Created directory '/c/Users/Administrator/.ssh'.
  Enter passphrase (empty for no passphrase):
  Enter same passphrase again:
  Your identification has been saved in /c/Users/Administrator/.ssh/id_rsa.
  Your public key has been saved in /c/Users/Administrator/.ssh/id_rsa.pub.
  The key fingerprint is:
  SHA256:iBw10OLmrx0OM47thcLCEv7zmVZY1FfnfeqvUjsT7fI louislin15@yahoo.com
  The key's randomart image is:
  +---[RSA 3072]----+
  |    .oo .   .. . |
  |    ...o . .  o .|
  |   ....   .    .o|
  |   .oo o       ..|
  |.  oo + S     .. |
  |oo  .o .     .o .|
  |ooo =.+      ..+ |
  |...* O+.    . =..|
  |  .oO=o      ..*E|
  +----[SHA256]-----+
  
  ```

- 与github进行认证和通信

  ```
  ssh -T git@github.com
  ```

  



