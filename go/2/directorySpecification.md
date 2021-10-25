# 目录规范
## go包目录规范
直接所有文件放在root即可。比如[glog](https://github.com/golang/glog)，而且一般来说，不太复杂的go的包采用的这种方式。
## go项目目录规范
go项目可以包括的内容：

- 项目介绍：README.md。
- 客户端：xxxctl。
- API 文档。
- 构建配置文件，CICD 配置文件。
- CHANGELOG。
- 项目配置文件。
- kubernetes 部署定义文件（未来容器化是趋势，甚至会成为服务部署的事实标准，所以目录结构中需要有存放 kubernetes 定义文件的目录）。
- Dockerfile 文件。
- systemd/init 部署配置文件（物理机/虚拟机部署方式需要）。
- 项目文档。
- commit message 格式检查或者其他 githook。
- 请求参数校验。
- 命令行 flag。
- 共享包：
    - 外部项目可导入。
    - 只有子项目可导入。
- storage 接口。
- 项目管理：Makefile，完成代码检查、构建、打包、测试、部署等。
- 版权声明。
- _output 目录（编译、构建产物）。
- 引用的第三方包。
- 脚本文件（可能会借助脚本，实现一些源码管理、构建、生成等功能）。
- 测试文件。

下面看一下具体的目录结构图：

![](https://gitee.com/shgopher/img/raw/master/go-subject-complex-menu0.png)
<br>
![](https://gitee.com/shgopher/img/raw/master/go-subject-complex-menu1.png)
<br>
![](https://gitee.com/shgopher/img/raw/master/go-subject-complex-menu2.png)

解释一下基本的目录意义：
- /cmd 将不同的组建的`main`函数存放在这里
- /internal 存放私有应有和代码，别人无法引用里面的代码
- /internal/pkg 私有的，内部共享的包
- /pkg 可以被外部引用的包
- /third_party fork的第三方包，并且做了更改，那么就放这里
- /test 测试用的示例，如果想被go忽略可以使用_test或者.test
- /configs 项目的配置目录
- /deployments 存放容器编排部署文件，比如k8s
- /init 初始化系统和进程管理配置文件
- /Makefile 项目管理，通常是部署一些自动化的东西
    - 静态代码检查，推荐 golangci-lint
    - 单元测试，运行go test
    - 编译，编译源码
    - 镜像的打包和发布，通常就是部署的docker和k8s的命令
    - 清理
    - 代码生成
    - 部署，一键部署功能
    - 发布，比如发布到docker hub ，github等
    - 帮助
    - 版权声明
    - api文档，通过 swagger生成api文档
- /scripts 存放脚本文件
    - /scripts/make-rules 通常来说makefile中用的脚本就是放在这里的。
    - /scripts/lib shell库存放自动化的脚本代码
    /scripts/install 安装脚本
- /build 存放安装包和持续集成相关的文件
    - /build/package 存放容器，系统的脚本和配置
    - /build/ci 存放ci(travis circle drone)配置文件和脚本
    - /build/docker 存放docker的dockerfile文件
- /tools 存放工具
- /githooks git 钩子例如说commit-msq
- /assets 项目使用的其它资源，比如图片，css js
- /website 存放项目网站相关的数据
- /docs 存放文档，上文详细说明了
- README.md  
- /CONTRIBUTING.md  说明如何贡献代码的说明
- /api 存放当前项目对外提供的不同服务
- /license 协议
- /CHANGELOG 生成change log
- /examples 示例文件，通常来说越多越好，越详细越好

不建议的目录⚠️：
- /src
- /model

总结一下：

- /cmd
- /internal
- pkg
- README.md

为必须拥有的目录结构（项目工程目录）其它的可以后期添加。如果你想添加文件夹但是里面没有内容，可以在文件夹里添加一个`.keep`的空文件即可。