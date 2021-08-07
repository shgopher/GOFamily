# go语言利用goland进行动态调试
> [jetbrains](https://blog.jetbrains.com/go/2019/02/06/debugging-with-goland-getting-started)

> 这里谈及的"服务"二字值得就是一个正在运行的系统，比如一个监听了80接口的正在运行的一个"服务"

goland 提供了调试的功能，这里谈及的动态调试，说的是调试一个服务，当然这里也会讲解一下关于goland debug的全部内容。

debug最开始要去设置debug的内容 -->  edit configurations -->  进入debug的设置界面，这里有 go run 某个文件，也有 go build 这个项目，其中有一个参数叫做 设置参数，这个意思就是说在debug的时候，如果有参入要填入，就是在这个地方填入的。设置完毕以后就可以debug了，下面我们将debug的类型分为四种类型来分别阐述其方法。

## debug一个应用程序
## debug一个测试
## 在本地机器上debug一个正在运行的“服务”
## 往远程机器上debug一个正在运行的“服务”


