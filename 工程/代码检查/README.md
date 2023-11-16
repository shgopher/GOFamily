# 代码检查工具
代码检查通常会被配置再 CI 中，用于自动检查代码的质量，本次我们介绍三个用于代码检查的工具
- go vet / go tool vet
- golangci-lint
- govulncheck
## go vet / go tool vet
go vet 命令是 go tool vet 的简单封装，go vet 实际上还是需要调用 go tool vet 才能完成工作，这俩命令的主要目的就是为了基础的代码检查。不过这个命令只能做简单的检查，下面我们介绍一下更常用的工具。

## golangci-lint
首先在介绍 golangci-lint 之前我们先下载它，它是一个 go 语言写的可执行文件，使用
```go
go install  github.com/golangci/golangci-lint/cmd/golangci-lint@latest 
``` 
即可下载到本地的～/go/bin/ 目录，这里专门存储使用 go install 下载的使用 go 写的可执行文件，记得将这个路径加入 PATH。

通过在要检查的项目中设置配置文件，用来配置 lint 工具的选项，使用 `.golangci.yaml` 即可

例如：
```bash

run:
  skip-dirs: # 设置要忽略的目录
    - util
    - .*~
    - api/swagger/docs
  skip-files: # 设置不需要检查的go源码文件，支持正则匹配，这里建议包括：_test.go
    - ".*\\.my\\.go$"
    - _test.go
linters-settings:
  errcheck:
    check-type-assertions: true # 这里建议设置为true，如果确实不需要检查，可以写成`num, _ := strconv.Atoi(numStr)`
    check-blank: false
  gci:
    # 将以`github.com/marmotedu/iam`开头的包放在第三方包后面
    local-prefixes: github.com/marmotedu/iam
  godox:
    keywords: # 建议设置为BUG、FIXME、OPTIMIZE、HACK
      - BUG
      - FIXME
      - OPTIMIZE
      - HACK
  goimports:
    # 设置哪些包放在第三方包后面，可以设置多个包，逗号隔开
    local-prefixes: github.com/marmotedu/iam
  gomoddirectives: # 设置允许在go.mod中replace的包
    replace-local: true
    replace-allow-list:
      - github.com/coreos/etcd
      - google.golang.org/grpc
      - github.com/marmotedu/api
      - github.com/marmotedu/component-base
      - github.com/marmotedu/marmotedu-sdk-go
  gomodguard: # 下面是根据需要选择可以使用的包和版本，建议设置
    allowed:
      modules:
        - gorm.io/gorm
        - gorm.io/driver/mysql
        - k8s.io/klog
      domains: # List of allowed module domains
        - google.golang.org
        - gopkg.in
        - golang.org
        - github.com
        - go.uber.org
    blocked:
      modules:
        - github.com/pkg/errors:
            recommendations:
              - github.com/marmotedu/errors
            reason: "`github.com/marmotedu/errors` is the log package used by marmotedu projects."
      versions:
        - github.com/MakeNowJust/heredoc:
            version: "> 2.0.9"
            reason: "use the latest version"
      local_replace_directives: false
  lll:
    line-length: 240 # 这里可以设置为240，240一般是够用的
  importas: # 设置包的alias，根据需要设置
    jwt: github.com/appleboy/gin-jwt/v2         
    metav1: github.com/marmotedu/component-base/pkg/meta/v1

```
使用 golangci-lint run (等于 golangci-lint run ./...  意思就是把所有的包，子包，遍历完全) 你会得到一个类似：

```bash

collie.go:171:41: composites: image/jpeg.Options struct literal uses unkeyed fields (govet)
				if err := jpeg.Encode(file, i.img, &jpeg.Options{q}); err != nil {
				                                    ^
collie.go:241:2: printf: `fmt.Println` arg list ends with redundant newline (govet)
	fmt.Println("声明：本程序来自GitHub：shgopher,欢迎关注公众号：科科人神；\n免费软件，如果使用期间出现任何后果，本软件不承担任何责任谢谢\n")
	^
collie.go:244:2: printf: `fmt.Println` arg list ends with redundant newline (govet)
	fmt.Println("运行结束 ☕️ ☕ ☕\n")
	^
collie.go:162:5: SA9001: defers in this range loop won't run unless the channel gets closed (staticcheck)
				defer file.Close()
				^
```
这样的结果，这样你就会发现是哪个配置的 linter 发出的警告，以及是什么样子的警告。

## govulncheck
go 官方维护了一个 https://vuln.go.dev/ 的漏洞库，我们可以使用 `go install golang.org/x/vuln/cmd/govulncheck@latest` 的方式，下载目前 (go 1.19) 还在测试阶段的这一功能，`govulncheck` 将会是一个独立的工具，并且 go 在 https://pkg.go.dev/golang.org/x/vuln/vulncheck 还提供了相关功能的 API，可以更灵活的去使用这个功能，目前已知的即将推出的功能分别是提供 vscode 插件，以及将此功能集成在 pkg.go.dev 这个 go 包的集合地，也就是说只要被收录在这个网站的 go 包都将自动接受漏洞检查，另外，go 以后可能还会将这个功能直接集成在例如 `go build` 这种常用命令上。

- 下载 govulncheck `go install golang.org/x/vuln/cmd/govulncheck@latest`
- 在一个拥有 go.mod 的目录下，使用 govulncheck 跟上一个有 go 文件的路径，例如：`govulncheck ./pkg/watcher`

只需要这样简单的设置就可以去检查代码中存在的风险和漏洞，govulncheck 就会打印出这样的信息：

```go
Vulnerability #2: GO-2022-0493
  When called with a non-zero flags parameter, the Faccessat
  function can incorrectly report that a file is accessible.
  Found in: golang.org/x/sys/unix@v0.0.0-20211020064051-0ec99a608a1b
  Fixed in: golang.org/x/sys/unix@v0.0.0-20220412211240-33da011f77ad
  More info: https://pkg.go.dev/vuln/GO-2022-0493
```
信息中包括了你引用的某些包出现的一些漏洞，在 fix 中有修复的信息，可以把你引用的包进行一个升级。
## x/tools 工具系列
> https://pkg.go.dev/golang.org/x/tools#section-readme 

比如检测变量 shadow 的工具

```bash
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
```


## 参考资料
-  https://go.dev/blog/vuln
- https://time.geekbang.org/column/article/390401