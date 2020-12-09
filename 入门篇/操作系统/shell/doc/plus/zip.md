# 压缩文件
## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
## 列举几个命令
- gzip – 压缩或者展开文件

- bzip2 – 块排序文件压缩器

- tar – 磁带打包工具

- zip – 打包和压缩文件

- rsync – 同步远端文件和目录

## 压缩是什么？
压缩其实就是去除冗余的部分只保留关键部分的技术。
## 压缩和解压缩
- gzip 压缩文件。
- gunzip 解压缩文件。

压缩的时候并不是新造而是替换，也就是源文件gg了，只剩下了新的压缩文件。

### gzip 选项

|选项|	说明|
|-|-|
|-c|	把输出写入到标准输出，并且保留原始文件。也有可能用--stdout 和--to-stdout 选项来指定。
|-d|	解压缩。正如 gunzip 命令一样。也可以用--decompress 或者--uncompress 选项来指定.
|-f|	强制压缩，即使原始文件的压缩文件已经存在了，也要执行。也可以用--force 选项来指定。
|-h|	显示用法信息。也可用--help 选项来指定。
|-l|	列出每个被压缩文件的压缩数据。也可用--list 选项。
|-r|	若命令的一个或多个参数是目录，则递归地压缩目录中的文件。也可用--recursive 选项来指定。
|-t|	测试压缩文件的完整性。也可用--test 选项来指定。
|-v|	显示压缩过程中的信息。也可用--verbose 选项来指定。
|-n|umber	设置压缩指数。number 是一个在1（最快，最小压缩）到9（最慢，最大压缩）之间的整数。 数值1和9也可以各自用--fast 和--best 选项来表示。默认值是整数6。

### gunzip

如果我们的目标只是为了浏览一下压缩文本文件的内容，我们可以这样做：

```bash
gunzip -c foo.txt | less

```
## bzip2
bzip2 跟gzip不同它舍弃了高速的压缩速度，反而追求的是压缩比例

> 来压缩一个已经被压缩过的文件,只是在浪费时间和空间！如果你再次压缩已经压缩过的文件，实际上你 会得到一个更大的文件。这是因为所有的压缩技术都会涉及一些开销，文件中会被添加描述 此次压缩过程的信息。如果你试图压缩一个已经不包含多余信息的文件，那么再次压缩不会节省 空间，以抵消额外的花费。
## tar

中文名叫做 “归档”

归档就是收集许多文件，并把它们 捆绑成一个大文件的过程

经常看到扩展名为 `.tar` 或者 `.tgz` 的文件，它们各自表示“普通” 的 tar 包和被 gzip 程序压缩过的 tar 包。
###使用方式

```bash
tar 模式 要命名的tar包 源文件
```
其中mode的模式有以下：

|c|	为文件和／或目录列表创建归档文件。
|-|
|x|	抽取归档文件。
|r|	追加具体的路径到归档文件的末尾。
|t|	列出归档文件的内容。

当抽取一个归档文件时，有可能限制从归档文件中抽取什么内容。例如，如果我们想要抽取单个文件， 可以这样实现：

tar xf archive.tar pathname
通过给命令添加末尾的路径名，tar 命令就只会恢复指定的文件。

当指定路径名的时候， 通常不支持通配符；然而，GNU 版本的 tar 命令（在 Linux 发行版中最常出现）通过 --wildcards 选项来 支持通配符
```bash
tar xf ../playground2.tar --wildcards 'home/me/playground/dir-\*/file-A'
```

## zip

zip即是压缩工具又是打包工具

zip -r playground.zip playground

除非我们包含-r 选项，要不然只有 playground 目录（没有任何它的内容）被存储。虽然会自动添加 .zip 扩展名，但为了清晰起见，我们还是包含文件扩展名。
## unzip ../playground.zip
## rsync命令

rsync命令是一个远程数据同步工具，可通过LAN/WAN快速同步多台主机间的文件。rsync使用所谓的“rsync算法”来使本地和远程两个主机之间的文件达到同步，这个算法只传送两个文件的不同部分，而不是每次都整份传送，因此速度相当快。

- rsync [OPTION]... SRC DEST
- rsync [OPTION]... SRC [USER@]host:DEST
- rsync [OPTION]... [USER@]HOST:SRC DEST
- rsync [OPTION]... [USER@]HOST::SRC DEST
- rsync [OPTION]... SRC [USER@]HOST::DEST
- rsync [OPTION]...
```bash
 rsync://[USER@]HOST[:PORT]/SRC [DEST] 对应于以上六种命令格式

 rsync有六种不同的工作模式：

 拷贝本地文件。当SRC和DES路径信息都不包含有单个冒号":"分隔符时就启动这种工作模式。如：rsync -a /data /backup 使用一个远程shell程序(如rsh、ssh)来实现将本地机器的内容拷贝到远程机器。

 当DST路径地址包含单个冒号":"分隔符时启动该模式。如：rsync -avz  \*.c foo:src 使用一个远程shell程序(如rsh、ssh)来实现将远程机器的内容拷贝到本地机器。

 当SRC地址路径包含单个冒号":"分隔符时启动该模式。如：rsync -avz foo:src/bar /data 从远程rsync服务器中拷贝文件到本地机。

 当SRC路径信息包含"::"分隔符时启动该模式。如：rsync -av root@192.168.78.192::www /databack 从本地机器拷贝文件到远程rsync服务器中。

 当DST路径信息包含"::"分隔符时启动该模式。如：rsync -av /databack root@192.168.78.192::www 列远程机的文件列表。这类似于rsync传输，不过只要在命令中省略掉本地机信息即可。如：rsync -v rsync://192.168.78.192/www
```
