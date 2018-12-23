
这次作业很简单，没什么坑好踩，按照博客的步骤就可以完成golang的配置，官方文档也很友善，由浅入深的教程很适合新手学习

>博客地址：[安装 go 语言开发环境][1]

那么就简单记录下怎么完成作业和一些笔记
**（以下皆为在centOs7下的操作）**


@[toc]


## 1、安装vscode

>Visual Studio Code 是一个轻量级但功能强大的源代码编辑器，可在 Windows，macOS 和 Linux 桌面上运行。它内置了对JavaScript，TypeScript和Node.js的支持，并为其他语言（如C ++，C＃，Java，Python，PHP，Go）和运行时（如.NET和Unity）提供了丰富的扩展生态系统。

嘛，vscode是个好东西，vim和emacs自然是更好的东西。不过也没必要死磕vim和emacs，集成的工具就是让人更有效率地去开发，尽管vim和emacs经过个人调教后可以成为强大的个人定制的IDE，但也需要丰富的理论知识和大量的实操支持，人生嘛，就是把有限的时间用在更想做的事情上，所以心安理得地选择vscode吧。

安装操作：

```C++
$ sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc

$ sudo sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/vscode.repo'

$ yum check-update
$ sudo yum install code
```
这几条命令的意思是导入vscode的安装包和设置code安装和使用的一些基本配置（安装密钥和版本库）

## 2、安装golang

基于一些大家都懂的原因，我们和golang的官方网站失去了联系
所幸世上存在一些善良的人们，为我们不远万里搬运来了中文goooo

>中国项目组：https://go-zh.org/

安装操作
```
$ sudo yum install golang //安装
$ rpm -ql golang |more //目录
$ go version //版本信息
```
不一定是最新版本，不过不影响使用

## 3、配置工作空间

工作空间是golang的一个重要模型
>go 工具为公共代码仓库中维护的开源代码而设计。 无论你会不会公布代码，该模型设置工作环境的方法都是相同的。
Go代码必须放在工作空间内。它其实就是一个目录，其中包含三个子目录：
1. src 目录包含Go的源文件，它们被组织成包（每个目录都对应一个包）
2. pkg 目录包含包对象
3. bin 目录包含可执行命令
go 工具用于构建源码包，并将其生成的二进制文件安装到 pkg 和 bin 目录中。
src 子目录通常包会含多种版本控制的代码仓库（例如Git或Mercurial）， 以此来跟踪一个或多个源码包的开发。

基本上我们需要维护的就是src这个子目录，里面包含了我们开发的源代码，而bin和pkg是通过使用go命令自动维护的目录，可以说一目了然，十分方便了。

那么**创建工作空间**，可以在你喜欢的任意目录

    $ mkdir $HOME/gowork

**配置环境变量**，对于 centos 在 ~/.profile 文件中添加:

    export GOPATH=$HOME/gowork
export PATH=$PATH:$GOPATH/bin

第一行使你可以在任意目录下执行go命令
第二行使你可以在任意目录下输入文件名执行bin中二进制文件

不要忘记**更新配置**

    $ source $HOME/.profile

**检查配置**

    $ go env
    

## 4、写一个helloworld

创建源代码目录：

    $ mkdir $GOPATH/src/github.com/github-user/hello -p

加上-p可以直接创建每一级目录，不用一级级创建

创建github.com目录有利于直接从github上载入golang的包和工具，而不需要调整目录，github-user是你自己的github用户名

总的来说，这个目录格式最适用于golang的开发

使用 vs code 创建 hello.go

    code $GOPATH/src/github.com/github-user/hello/hello.go
    
hello.go
```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```
在终端运行!(在hello目录下）
```
$ go run hello.go
hello,world
```

也可以使用
```
$ go install
$ hello
```
在bin中生成二进制文件，并执行

## 5、安装go的一些工具

出于某些大家都懂得的原因，我们无法从https://golang.org/x/tools/上安装东西

所幸我们还有github
https://github.com/golang/tools 是 golang.org/x/tools的一个镜像，代码是一样的

所以我们可以直接把代码下载到本地，直接从本地中链接工具使用，而不是通过网络

```
# 创建文件夹
mkdir $GOPATH/src/golang.org/x/
# 下载源码
go get -d github.com/golang/tools
# copy 
cp $GOPATH/src/github.com/golang/tools $GOPATH/src/golang.org/x/ -rf
# 安装工具包
$ go install golang.org/x/tools/go/buildutil
```



顺便了解一下几个命令
[Go 命令教程][2]

### go get

    $ go get github.com/golang/tools
    
>命令go get可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装

所以上面这条命令就是从https://github.com/golang/tools上把包克隆了下来，相当于git clone，放在了

    $GOPATH/src/github.com/golang/tools
目录下
并且编译和安装了目录下的.go文件，生成二进制文件放在了bin中

### go build

>go build命令用于编译我们指定的源码文件或代码包以及它们的依赖包。

### go install

>命令go install用于编译并安装指定的代码包及它们的依赖包。当指定的代码包的依赖包还没有被编译和安装时，该命令会先去处理依赖包。与go build命令一样，传给go install命令的代码包参数应该以导入路径的形式提供。并且，go build命令的绝大多数标记也都可以用于go install命令。实际上，go install命令只比go build命令多做了一件事，即：安装编译后的结果文件到指定目录。

[go build 和 go install 的区别][3]

### go run

>go run命令可以编译并运行命令源码文件。由于它其中包含了编译动作，因此它也可以接受所有可用于go build命令的标记。除了标记之外，go run命令只接受Go源码文件作为参数，而不接受代码包。与go build命令和go install命令一样，go run命令也不允许多个命令源码文件作为参数，即使它们在同一个代码包中也是如此。而原因也是一致的，多个命令源码文件会都有main函数声明。

## 6、把包推向远程仓库

假设你已经按步骤实现了[如何使用Go编程][4]的两个包hello和stringutil，一个main程序hello.go，一个库文件reverse.go，和一个测试文件reverse_test.go

这部分很简单，只要建立好清晰的工作空间目录，在目录下编程是件很简单的事情。

那么我们是时候考虑把我们编写好的包推向github上的远程仓库共享给别的开发者，让他们使用我们~~精心编写的强大~~的库

这是一个实用的教程
[Git教程][5]

**首先安装git客户端**

    $ sudo yum install git

**初始化git仓库**

在

    $GOPATH/src/github.com/github-user/

目录下

    $ git init

我们会发现目录下多了一个.git目录，这个目录是Git来跟踪管理版本库的，我们不需要操作它

**把目录文件添加到git仓库**

    $ git add .
    
.的意思是把当前目录下**所有新添加的目录及目录下文**件添加到git仓库（准确地说是暂存区（stage））

**然后提交修改**

    $ git commit -m "创建了hello和stringutil包"
    
-m是注释，这是必须的，不然你不知道你这次修改干了什么，而git本来就是用来管理历史版本的工具，如果你什么都不写，就失去了使用git的意义

现在我们本地的分支已经实现好了，它的默认名字是master，接下来我们就需要把master这个分支所管理的代码push到github的远程仓库上

**那么首先，新建一个远程仓库**
我们假设你已经有了一个github账号，找到新建仓库这个页面
https://github.com/new

![新建仓库](https://img-blog.csdn.net/20180929220440214?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

输入你喜欢的仓库名，并按照默认设置新建就行了

**那么接下来就是连接本地仓库和远程仓库**

    $git remote add origin https://github.com/github-user/go
    
github-user是你的用户名，go是你的仓库名，origin是默认的远程仓库名，也可以改成别的，然后按照提示输入用户名和密码就成功链接到远程仓库了

**把本地库的内容推向远程仓库**

    $ git push -u origin master

把本地库的内容推送到远程，用git push命令，实际上是把当前分支master推送到远程。这样就能在https://github.com/github-user/go上看到hello和stringutil包了
![远程仓库](https://img-blog.csdn.net/20180929220518287?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
由于远程库是空的，我们第一次推送master分支时，加上了-u参数，Git不但会把本地的master分支内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来，在以后的推送或者拉取时就可以简化命令。

第一次使用git的时候还可能遇到各种错误，但是git的错误提示和引导都十分清晰，只要认真阅读，按照提示键入命令即可，当然更好的方式是去了解git的工作原理，例如阅读最开始那个教程


  [1]: https://pmlpml.github.io/ServiceComputingOnCloud/ex-install-go
  [2]: http://wiki.jikexueyuan.com/project/go-command-tutorial/0.1.html
  [3]: https://www.cnblogs.com/ghj1976/archive/2013/04/23/3038347.html
  [4]: https://go-zh.org/doc/code.html
  [5]: https://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000
