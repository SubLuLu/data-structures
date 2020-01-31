# 数据结构与算法分析(In Golang)

### 简介

这是用gitbook编写的书，代码实现使用的编程语言是`Golang`，这样做的目的仅仅是为了熟练一下`Golang`，同时能够很方便的将文档导出为`PDF`、`MOBI`、`EPUB`等电子书格式，以便后期阅读参考。

### 分支说明

本书采用多分支原则编辑书籍，master为所有分支的合并，book为文档分支，其它分支均为代码实现，每一章对应一个分支。

| branch | comment                        |
| ------ | ------------------------------ |
| master | 主分支，代码和文档集合         |
| book   | book分支，文档，用于导出电子书 |
| list   | list分支，链表实现相关代码     |
| queue  | queue分支，队列实现相关代码    |
| stack  | stack分支，栈实现相关代码      |
| hash   | hash分支，散列实现相关代码     |
| tree   | tree分支，树实现相关代码       |
| heap   | heap分支，堆实现相关代码       |
| sort   | sort分支，排序实现相关代码     |

### 如何导出电子书

* 安装必要软件

gitbook已经不再提供客户端，只能在线编辑，而且不能导出PDF，造成了大大的不便，所以不得不采取某些措施去让`gitbook-editor`涅槃重生。

1. node 5.6.0
2. gitbook-cli 2.3.2
3. gitbook 2.6.9
4. calibre 3.46.0

`node`是gitbook依赖的环境，也就是说老版本的`gitbook-editor`是基于node实现的，而且界面使用的是nw，需要安装5.6.0版本，若多个版本共存，使用gitbook命令时，必须切换到5.6.0。

`gitbook-cli`是gitbook的命令行接口工具，最新的版本是2.3.2，已经有两年多没有更新了，里面的依赖模块和语法都是两年前的，这也就是`node`必须使用5.6.0版本的原因。

`gitbook`是用来美化文档生成书籍的命令行工具，需要配合`gitbook-cli`使用，而且需要保持大版本一致，所以`gitbook`只能使用2.x.x的版本，在2.x.x的版本中，2.6.9是最新最稳定的，所以使用该版本。

`calibre`是一款电子书转换工具，它提供了命令行接口，gitbook导出常见格式的电子书正是调用该软件提供的接口，对应当时gitbook可以调用的版本为3.4x.x，所以推荐安装版本为3.46.0，[官方下载链接](https://download.calibre-ebook.com/3.46.0/)

**Mac用户可以参照[gitbook重生](https://github.com/SubLuLu/awesome-tools/blob/master/gitbook-editor-reborn.md)进行安装，其他系统的用户，步骤类似，具体问题还需具体分析。**

* 使用步骤

安装好上述软件之后，切换到`book`分支

```bash
git checkout book
```

安装依赖以及`book.json`中指定的插件

```bash
gitbook install
```

利用命令导出电子书

```bash
# pdf为导出格式，可以换为epub mobi
# . 指定文档所在目录
# ds.pdf 指定导出的文件名
gitbook pdf . ds.pdf
```
