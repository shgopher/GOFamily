# go 包管理工具
- go 导包的过程
- go 包的版本管理
- go 包的最小版本原则
- go module 命令的使用
	- go get / go install
	- GOPROXY
	- GOSUMDB
	- go 使用私有服务器
	- workspace
## go 语言的包导入过程
- go 编译快的原因
- go 程序的构建过程

go 程序编译快的原因。1 是因为 go 要求每一个包都必须显式的标记处引入的包。
```go
package hi

import (
  "fmt"
  "io"

  "github.com/shgopher/hui"
)
```
所以编译器无需去查找这个文件中具体引入了哪些包就可以在头部检索出全部的导包名单

2 是因为 go 语言不允许循环引用，比如说 a 引用了 b，b 又引入了 a，所以每一个包都是独立的存在，处于 “有向无环图 “这种格式下，进而可以并行去编译不同的包，提高编译效率。

3 包编译的结果里 (`xx.o`，`xx.a`) 中不仅存储了此包的导出信息，还存储了它引入的包的导出信息，这样编译器只需要读取一个目标文件就可以获取全部的包信息，不需要读取全部的文件。比如 p 包要引入 q 包，那么 q 包的。a 文件已经包括了它引入的全部包的导出信息，这样编译器可以快速统计某一个包被引用的情况，比如包引用了 a b，a b 分别引用了 c 但是版本不同，那么这种情况下因为主包已经了解了他们的情况，所以就可以用最小版本原则去选择一个版本的包进行编译即可。有点类似分封制，我帝王 p 包，我的直接负责人是诸侯 q，q 掌握了它下级 s 的所有数据，假设 s 还有下级，s 掌握了 w 的所有数据，然后层层回报，当帝王 p 最终编译的时候，它已经完全拥有了所有的数据，可以精准调控。再结合每一个包都可以独立的并行编译

所以这一整套下来 go 的编译速度就会很快。并且在 go.mod 文件中也会详细记录使用过的包及其版本，那么并行下载，快速编译就不是梦了。

## go 程序的构建过程。

go 程序的建立是由编译和链接两个阶段组成的，举个例子
有一个项目拥有一个 main 包，一个 lib 包，lib 包会被引入到 main 包中，那么编译的过程是这样的，首先创建一个临时工作区域，编译 lib 包为 lib.a，编译 main 包为 main.a，链接器将 lib.a 和 main.a 链接成 name.out，然后改名为 name (unix-like)，**使用第三方包的意思就是链接了该包的源代码编译的最新的。a 文件而已**，并且**每次编译**都会重新编译最新的。a 文件 (所以只保留。a 文件，删除源码是不可行的)，但是标准库除外，并不会每次编译都会重新编译标准库，所以说如果你修改了源码一定得把标准库。a 文件删除，并且 build 的时候使用 `build -a` 的方式才可以使用自己更新的标准库。

一个小 tip：go 导包的时候引入的是路径名称，路径名称**通常**最后一位路径名称跟包名保持一致，当然也可以不一致。例如 `github.com/shgopher/go-hui` 但是实际上包名称是这个 `package hui`，所以导入的时候按照 `github.com/shgopher/go-hui`，使用的时候用 `hui.xxx()`。

因为包的名称很容易发生冲突，所以 go 接受包名称的重命名

```go
import (

  app  "github.com/shgopher/app"
  app1  "github.com/googege/app"
)
```
除此之外还有 `_` 方式，它会**计算包级变量的初始化表达式**和**执行导入包的 init 初始化函数**，意思就是跟 init 函数有关的内容都会被计算，并且导入进去。

```go
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

```
比如这个例子，image.Decode 会查询注册表，看看注册表里都有谁，这个时候我们引入的 _ iumage/png 的 init 函数就是将 png 中实现了接口的具体数据导入到了注册表中，所以说这里只需要导入这个 init 函数即可。

这种用法还是比较重要的，我们来自己编写一段代码，来看看实现这种注册器的基本原理：
```go
// 这里是实现这个注册器
package go2
// 在某个包实现go2的时候，这里是go3，这个drivers就被初始化了。
var drivers = make(map[string]Driver)

type Driver interface {
	Open()
}

func Register(name string, driver Driver) {
	drivers[name] = driver
}

func DDD(name string) {
	var d Driver
	if d1, ok := drivers[name]; !ok {
		panic("no real driver registered")
	} else {
		d = d1
	}
	d.Open()
}

```
```go
// 这里是实际的实现我们注册器那个包实现的抽象方法
package go3

import (
	"fmt"

	"shgopher.com/go2"
)

type ddd struct{}

func (d *ddd) Open() {
	fmt.Println("this is go3")
}

func init() {
	go2.Register("go3", &ddd{})

}

```
```go

// 最后我们来使用一下：

package main

import (
	"shgopher.com/go2"
	_ "shgopher.com/go3"
)

func main() {

	go2.DDD("go3")
}

```

另外还有一种导入包的奇妙用法，就是

```go
	. "fmt"

	func main(){
		println("hi")
	}
```
它的意思是省略包名，no！，请不要这么用。

## go module 构建工具
- go111MODULE
- GOPROXY
- GOSUMDB
- module 版本的升级
- workspace

go 自带包管理工具 go module，本文章写作时 go 的版本是 1.19，所以过去的 go path，go vendor，go dep 均不再提及，读者也不用去 care，鉴于干谈 go module 非常的枯燥，我设立一个例子来讲述 go module。

在桌面，我们使用 go module 的创建命令

```bash
cd ~/Desktop

mkdir hello

go mod init github.com/shgopher/hello 

```

创立了一个以 hello 作为包名称的包。这里我们使用 `github.com/shgopher/hello` 是证明这个包使用 GitHub 进行存储，因为 go 使用 git 来管理版本，在项目内部，这个项目的包名称就是最后一个名称，hello 才是这个包的正式名称，此处是惯用，实际上真实管理包正式名称的是每一个文件上的 package xx 管理的，整个就是路径而已，不过习惯于最后一个路径名称为包的名称，即：最后一个名称 == package 后面的 xxx。

然后我们在这个项目的 root 路径下会发现 go 帮我们创建了一个 go.mod 的文件，它就是 go module 的版本文件，我们拿一个 go.mod 来作为举例

```go
// 一个 go.mod 目录：

module github.com/shgopher/hello

go 1.19


require (
  github.com/apache/thrift v0.13.0
  github.com/bytedance/gopkg v0.0.0-20220531084716-665b4f21126f
  github.com/appleboy/gin-jwt/v2 v2.6.4
  gopkg.in/yaml.v3 v3.0.1
)

require (

	github.com/beorn7/perks v1.0.1 // indirect
)

replace (
	// 后面指向的要么是相对路径例如 ../shgopher/.com/net
	// 要么一定要在后面加上版本，并且是可以获取到的包
	golang.org/x/net v1.1.2 => shgopher.com/net v1.4.5
)

```

我们一个一个解释，首先最开头的是这个包的路径名称，实际上 go 会使用 `git clone https://github.com/shgopher/hello.git ` 的方式下载包，当然，假设你不使用 GitHub 的 git 作为项目的存储位置，使用自建的也是 OK 的，比如 `module shgopher.com/hello` 只要同样配置了 git 服务器都可以，因为底层都是使用的 `git clone https://xxx.com/xx.git` 模式。

接下来的 `go 1.19` 是这个项目使用的 go 的大版本，比如你使用的是 `1.19.1`，`1.19.2`，上面写的都是 `go 1.19`，比如你使用 `go mod edit -go=1.19 ` 来更新此项目使用的 go 版本的时候，只能写到 minor，而不能加上 patch (minor 和 patch 下文有说)。

接下里有两个 require，其中第一个 require 指的是直接引入的包，后面的 require 是间接引用的包。意思就是你引用的包，它引入的包。

第一个 require 中，我给出了四种常见的版本用法

1. 使用 git version 命名的版本 `v0.13.0 `
2. 项目没有使用 version 命名，go 官方使用最新的文件，并且给予了它一个临时的版本号 `v0.0.0-20220531084716-665b4f21126f`
3. 当项目的版本超过 1.x 的时候，go 推荐使用再加上一个/v2 的方式进行命名，但是实际上它的包名称仍然是 gin-jwt，并且 1.x 和 2.x 的包可以同时引入项目中，因为他们算两个包，只需要重命名即可。版本 2 和版本 1 的 module 后面写的也不一样，比如一个是 `module github.com/shgopher/collie` 版本 2 就是 `module github.com/shgopher/collie/v2` 不过虽然最后结尾的是 v2 但是 package 后面写的仍然是 collie
4. 也可以使用 yaml.v3 的方式进行命名，因为路径名称并不是包的名称，所以这种方式也有不少项目使用

第二个 require，全部都是间接引用的包名称，go 也会一并下载到本地缓存，go 会使用 `～/go/pkg/mod/包名称@版本号` 的地方去保存下载下来的包。

使用 `go list -m -json all` 命令可以查看当前项目的所有依赖的包名称，例如

```json
{
	"Path": "github.com/marmotedu/iam",
	"Main": true,
	"Dir": "/Users/shgopher/Desktop/github-projects/iam",
	"GoMod": "/Users/shgopher/Desktop/github-projects/iam/go.mod",
	"GoVersion": "1.18"
}
{
	"Path": "cloud.google.com/go",
	"Version": "v0.93.3",
	"Time": "2021-08-17T22:38:11Z",
	"Indirect": true,
	"GoMod": "/Users/shgopher/go/pkg/mod/cache/download/cloud.google.com/go/@v/v0.93.3.mod",
	"GoVersion": "1.11"
}
{
	"Path": "github.com/marmotedu/component-base",
	"Version": "v1.6.2",
	"Time": "2021-12-21T06:47:41Z",
	"GoMod": "/Users/shgopher/go/pkg/mod/cache/download/github.com/marmotedu/component-base/@v/v1.6.2.mod",
	"GoVersion": "1.17"
}
```

我们重点关注几个字段，首先第一个对象中，`"Main":true`，表示这个对象描述的包属于主包，即，这个项目的 root 路径下的 `go.mod`；第二个对象中，我们看到 `	"Indirect": true,` 意思就是指这个包是间接引入的包，第三个对象里，并没有出现 indirect 的字段，就证明这个包是直接引入的。因为这个列举的主要是 go.mod 的目录，又因为拥有一个 go.mod 的包拥有同样的版本号，所以包和包的子包是公用一个 go.mod 的，只会出现一次。

### 如何控制 go 引入的包的版本号。

当我们更新一个旧的项目中的依赖时，我们可以使用 `go clean -modcache` 的命令去删除 go 缓存的 go 包。然后显式的为包设置版本，第一种方法是直接在 go.mod 写入要引入的包的版本，第二种方法是使用 `go mod -require=github.com/shgopher/hello@v0.1.1` 的方式写入你想要的包的版本，除了第二种这种比较精确的方式，go mod 还支持 query 的方式去指定范围，例如说 `go mod -require=github.com/shgopher/hello@>=v0.1.1`。

go 包采用 “最小版本” 选择的理念。举个例子，我的项目 hello，直接引入了 c 包和 d 包，然后 c 和 d 又分别引入了 e 包，那么最后本项目使用的 c,d,e 包采用的是哪个版本呢？

我们详细说明一下，假设我们使用的 cde 都是@latest，那么每次 build 的时候，都会去下载最新的包；如果我们使用了具体的版本，假设我们在 go.mod 中给定的 c d 分别是 v0.1.0 v1.0.1 但是呢，c 存在多个版本，比如现在有 v0.1.1 v0.2.1 那么根据最小版本的选择问题，go 会选择一个符合 v0.1.0 的最小版本，即：>=v0.1.0，所以此处会选用 v0.1.1，也就是说 go 的版本的隐藏含义是大于等于选最小，很多别的语言都是大于等于选最大，例如 rust，但是强调一下，**这个大于等于选最大或者最小说的是一个版本的情况下，超过版本号的时候，算俩包 (例如 v1.1.0 和 v2.1.0 就算俩大版本了，所以说他们是俩包都不为过)**，而且我们在改变版本的时候应该先把缓存给清除了 `go clean -modcache`，接下来我们的 cd 包分别引入了 e 包，假如 e 存在 v1.1.0 v1.2.0 v1.3.0 v1.4.0，c d 分别引入的是 v1.1.0 v1.3.0 那么 go 会**合并需求 (只引入包众多版本的一个版本)**，并且依然采用最小版本的方式即：最终只选用 e 包的 v1.3.0 这一个版本的包引入，总结一下：单个包，使用 >= 的最小值引入，多个包引入同一个包的情况，使用满足他们共同条件下的最小值。

使用 go mod tidy 可以对这个项目的包进行梳理，比如使用 latest 的包会重新比对，然后下载最新版本的包，比如在 require 中明确引入的包，但是在实际上线前发现没有继续使用这个包了，那么使用 go mod tidy 也可以删除这个包，这个命令类似于 “刷新” 这个概念。

### 如何优雅的升级和降级引入的包版本？

我们会用到下面两个命令
- go list
- go get

首先，我们使用 `go list -m -versions github.com/shgopher/hello` 的命令查找这个包的众多版本 (没有使用 git 给定版本的就没办法了，给定@latest，直接用最新的了)，例如说会有 v1.1.0 v1.2.0，假设我们本来用的是 v1.1.0 想升级一下，我们使用 `go get github.com/shgopher/hello@v1.2.0` 即可优雅的升级这个包的版本，当然降级也是一样的，反正指定一个包即可，要注意的是，go 在升级或者降级的时候，会自动将间接使用的包也做出相应的版本调整。假如我们现在想把所有的依赖包升级为这个版本下的最新包，使用 `go get -u` 就可以更新所有的直接依赖和间接依赖的包为最新包。如果只想升级 patch 而不是 minor (v1.2.0；1 是 version，2 是 minor，3 是 patch) 使用 `go get -u=patch`，如果只想升级某一个包 (及其依赖包) 到最新版本，在后面指出来具体的包名称即可，`go get -u github.com/shgopher/hello`。

go install 和 go get 是两个比较像的命令，其中 go get 仅仅用在 go module 中调整版本的时候，比如升级，降级，go install 用于下载包，两者都是一样的，如果包的后面带上版本那么就是指定版本，如果不带就是最新版本。

```go
go install github.com/shgopher/hello 

go install github.com/shgopher/hello@v1.1.2
```

```go
go get -u github.com/shgopher/hello
go get github.com/shgopher/hello@v1.3.4
```
在使用 go install 的时候其实是忽略 go.mod 的，所以 go install 跟 go.mod 没有任何关系，也不会记录 go.mod 中，它的作用就是下载包，go get 跟 go.mod 紧密相关，它需要 go.mod，并且每次更改都会记录在 go.mod 中，当你在一个没有 go.mod 的路径下 (指的是此路径，此路径的父路径也没有) 使用 go get 会提示如下错误
```go
go: go.mod file not found in current directory or any parent directory.
	'go get' is no longer supported outside a module.
	To build and install a command, use 'go install' with a version,
	like 'go install example.com/cmd@latest'
	For more information, see https://golang.org/doc/go-get-install-deprecation
	or run 'go help get' or 'go help install'.
```

所以这个时候应该使用的是 go install。

### go 包代理

GOPROXY 是 go 的代理服务器。go 之前使用 GitHub，gitlab 等托管平台，goproxy 命令可以设置一个集中式的代理服务，比如

```bash
export GOPROXY = https://goproxy.cn,direct 

export GOPROXY = https://proxy.golang.org,direct

```

其中前者是国内常用的代理服务器，后者是 go 官方的代理服务器；`direct` 的意思是，**直接**使用代理服务器的内容，`,` 的意思是前面的服务器**只有**出现 404 和 410 错误的时候才会去选择逗号后面的服务，如果想设置只要发生错误就使用后者的命令，那么可以使用 `|`，例如使用

```bash
go env -w GOPROXY=https://proxy.golang.org|https://goproxy.cn|direct`
```

这里有个小知识，因为 `｜` 在 unix-like 操作系统中通常还表示通道的含义，就是前面的输出等于后面的输入，所以我们需要将这个符号进行转义才能正常使用：

```bash
go env -w GOPROXY=https://proxy.golang.org\|https://goproxy.cn\|direct

```
go module 还有 go.sum 这个文件，它存储的是包的一些基础信息，最重要的是对于一个包求 hash 值，记录这个 hash，在每次 build 的时候对于缓存的包和 go.sum 中的 hash 值做对比，来规避恶意更改，go.sum 会在项目的更新换代过程中保存多个版本的包信息。

GOSUMDB 命令指向的服务就是保存公有包的校验和的数据库。一个新的包，在一切运行正确的情况下，go 会通过 GOSUMDB 配置的数据库去查询这个包的校验和，查询出结果后和下载的包进行比对，正确的情况下存入 go.sum；如果一个已经缓存的包，每次 run build 的时候都会将缓存的包文件校验跟 go.sum 进行比对来保证正确性。

当然，如果你不想使用 GOSUMDB，使用 `go env -w GOSUMDB=off` 即可。这样就无法对比包和数据库中的校验和，只能做本地校验了。

### 如何使用一个私有包

我们讲解了如果配置公有的代理服务器 GOPROXY，文件校验和数据库 GOSUMDB，接下来我们谈一下如果我们想使用一个私有的包，比如一个 GitHub 上的私有包，一个本地 git 服务器上的包，我们使用 GOPRIVATE，它的目的就是绕过 GOPROXY 和 GOSUMDB，因为是私有的所以在代理服务器和文件校验和数据库都不会有它的记录，我们可以这么设置
```bash
# 意思是这个路径下的所有包都不会经过代理服务器了。
# 这个命令支持多个路径使用逗号分隔
export GOPRIVATE = shgopher.com,shgopher.io,*.api.shgopher.com
```
除了设置这个命令之外，还需要设置一个密钥用来 ssh 的方式去访问 GitHub 上的私有仓库，或者是 GitHub --- personal access tokens

使用 ssh，将主机公钥 (～/.ssh/id_rsa.pub) 添加到 github.com 的 ssh keys 中。

```bash
	#我们谈一下如果生成公钥：
	## 在～/.ssh/ 路径下
	ssh-keygen -t rsa -C "个人邮箱"
	## 将这个id_rsa.pub中的公钥 添加到GitHub中的ssh keys 中
```
如果使用 ssh 的方式获取代码，那么在~/.gitconfig 中添加 (这一步其实就是一个映射：保持你的日常习惯的情况下，使用了 ssh)
```bash

[url "ssh://git@github.com"]
	insteadOf = https://github.com
```
如果是本地服务器那么就是
```bash
[url "ssh://git@local.com"]
	insteadOf = https://git.local.com
```
不过要注意一下，如果使用 ssh 那么远程的服务就得变更名称，因为通常我们的 GitHub 给我们的都是 https 的方式，使用 ssh 的话就是：
```bash
 git remote set-url origin git@github.com:USERNAME/REPOSITORY.git
```
从 ssh 更改为 https 就是

```go
git remote set-url origin https://github.com/USERNAME/REPOSITORY.git
```

使用 GitHub personal access token 的方式
```bash
# 在GitHub personal access tokens中申请即可，然后配置在～/.netrc

machine github.com login shgopher password 你的 personal access tokens
```
在 Linux 中要配置~/.netrc，但在 macOS 中，git 输入的 username 和令牌会自动的缓存，不用设置这个配置文件。

更多关于 GitHub 访问的信息可以访问[这里](https://docs.github.com/en/get-started/getting-started-with-git/managing-remote-repositories#switching-remote-urls-from-ssh-to-https)

我个人建议直接使用 GitHub 的 personal access tokens (仅支持 https) 这种令牌的方式代替密码，并且使用 https 即可，简单，安全，好用，
另外如果你的令牌更新了，假设是 Macos 的情况下，可以去钥匙串访问的互联网密码种类中去更新令牌

如果我们引入的包是不支持 https 协议的，那么我们可以设置 `GOINSECURE = private.res.com` 来使用这种私有库。

go 仅仅支持 https 和 http 标准端口的带有域名的包，比如我们使用 ip，或者端口不是 80 和 443，那么设置 GOPROXY 和 GOPRIVATE GOINSECURE 都没用了。

比如类似这种：

```go
// 这种无法使用
import 192.168.1.1:9090/test/t
```
这种直接使用 ip 的方式不能写到 go 的代码中，所以我们可以使用 git 的 insteadof 功能来改变一下，使用一个正常的 url 去改变掉这个 ip

```go
# ~/.gitconfig
[url "132.148.1.1:9090/test/t"]
	insteadof="shgopher.com/test/t"
```
这个时候你引入的时候写 shgopher.com/test/t 就可以了，它会自己找 192.168.1.1:9090/test/t 真实值去代理，并且 shgopher.com/test/t 也要加入到 GOPRIVATE 中才可以。

这个时候有个问题，就是如果 shgopher.com 没有配置版本管理软件，例如 git 这种，go 是无法获取数据的，go 一般会存储一些网站，例如 GitHub，gitlab，所以它看到类似 github.com/xxx/的地址就去找使用 git，或者你手动写上 git，比如 shgopher/d/x.git go 也会明白你使用的 git，go 还能发送请求的方式去获取你使用的版本软件，比如 golang.org/x/net go 就会请求 https://golang.org/x/net?go-get=1 这个服务提供了一个 html：

```html
<head>
<meta name="go-import" content="golang/x/net git https://go.googlesource.com/net">
<meta name="go-source" content="golang.org/x/net https://github.com/golang/net">
<head>
```
这表示这个 golang.org/x/net 的包实际是通过 git 的方式从 https://go.googlesource.com/net 获取的。

我们可以让 `GET shgopher.com?go-get=1 ` 返回一个 html，在里面设置为 content=“shgopher.com/test/t git 132.148.1.1:9090/test/t”。

其实，老老实实的使用 go 推荐的导包方式挺好的，即便是私有包也好好的安排一个 git 版本服务器，仅限内部使用的话，配置好 GOPRIVATE 并且使用 ssh 公钥和私钥或者 GPG 的方式加密获取包，这种才是王道。

配置私有的 GOPROXY，你不能总想着靠别人，如果想搭建一个属于自己的 goproxy 服务器，那么可以使用 https://github.com/goproxy/goproxy 这个项目

### 本地开发使用的配置文件 go.work
go 推出了仅用于本地开发的 workspace，我们来介绍一下这个功能，比如说我们要在一个项目中你引入一个还未公开到公共仓库的包 github.com/shgopher/hui，那么在这个时候就需要 workspace 了。
```go
package main

import(
	"github.com/shgopher/hui"
)
```
很显然这个包还未发布，所以不可能引入，在之前通常使用 replace 的方式，但是有了 workspace 以后，不到特定的场景就不需要使用 replace，比如尚未发布的包这种场景用 workspace 最好用 (包的子包不需要任何设置，就可以主包直接引入这个子包了，这一点要搞清楚)。

使用 go work init 来创建一个工作区，go.work 形如：
```go
//go.work
go1.19

use(
	. //go.work当前路径是可以省略的
	/Users/ddd/src/go/hui
)

```
go.work 可以设置在需要工作区的路径的父路径 (如果父路径没有就会一直忘外寻找直到根路径)，工作区中的命令会向外部去寻找 go.work，所以我们通常可以在需要 workspace 的地方的父路行下设置 go.work，并且设置为绝对路径这样简单高效，记住，路径是不包含子路径的，比如本来是/workspace/go1 但是你设置的是/workspace 那么就是错的，它会寻找/workspace 中有没有 go.mod 它以 go.mod 作为寻找对象。当然了一个包的子包 (例如/workspace/go1/a) 无需再写进去 go.work 中，它跟外部的包是一个包。

```go
use(
	/workspace/a
	/workspace/b
	/workspace/c
)
```

在工作区同样可以设置 replace，但是级别没有 go.mod 中的高，会被 go.mod 覆盖。

go 的包和包的管理工具 go module 基本上已经讲解完毕了，以后有了新的见解再更新这篇文章。

## 参考资料
- https://go.dev/blog/get-familiar-with-workspaces
- https://book.douban.com/subject/35720728/  图书下册 321 页 - 349 页；图书上册 120 页 - 131 页
