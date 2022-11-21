# 代码检查工具
代码检查通常会被配置再cicd中，用于自动检查代码的质量，本次我们介绍三个用于代码检查的工具
- golangci-lint
- golint
- govulncheck
## golangci-lint
## golint
## govulncheck
go官方维护了一个 https://vuln.go.dev/ 的漏洞库，我们可以使用 `go install golang.org/x/vuln/cmd/govulncheck@latest` 的方式，下载目前（go 1.19）还在测试阶段的这一功能，`govulncheck` 将会是一个独立的工具，并且go在 https://pkg.go.dev/golang.org/x/vuln/vulncheck 还提供了相关功能的API，可以更灵活的去使用这个功能，目前已知的即将推出的功能分别是提供vscode插件，以及将此功能集成在pkg.go.dev这个go包的集合地，也就是说只要被收录在这个网站的go包都将自动接受漏洞检查，另外，go以后可能还会将这个功能直接集成在例如 `go build` 这种常用命令上。

- 下载govulncheck `go install golang.org/x/vuln/cmd/govulncheck@latest`
- 在一个拥有 go.mod 的目录下，使用 govulncheck 跟上一个有go文件的路径，例如：`govulncheck ./pkg/watcher`

只需要这样简单的设置就可以去检查代码中存在的风险和漏洞，govulncheck 就会打印出这样的信息：

```go
Vulnerability #2: GO-2022-0493
  When called with a non-zero flags parameter, the Faccessat
  function can incorrectly report that a file is accessible.
  Found in: golang.org/x/sys/unix@v0.0.0-20211020064051-0ec99a608a1b
  Fixed in: golang.org/x/sys/unix@v0.0.0-20220412211240-33da011f77ad
  More info: https://pkg.go.dev/vuln/GO-2022-0493
```
信息中包括了你引用的某些包出现的一些漏洞，在fix中有修复的信息，可以把你引用的包进行一个升级。
## 参考资料
- - https://go.dev/blog/vuln