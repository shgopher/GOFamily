# go语言官方提供的漏洞检查工具 govulncheck

go官方维护了一个 https://vuln.go.dev/ 的漏洞库，我们可以使用 `go install golang.org/x/vuln/cmd/govulncheck@latest` 的方式，下载目前（go 1.19）还在测试阶段的这一功能，`govulncheck` 将会是一个独立的工具，并且go在 https://pkg.go.dev/golang.org/x/vuln/vulncheck 还提供了相关功能的API，可以更灵活的去使用这个功能，目前已知的即将推出的功能分别是提供vscode插件，以及将此功能集成在pkg.go.dev这个go包的集合地，也就是说只要被收录在这个网站的go包都将自动接受漏洞检查，另外，go以后可能还会将这个功能直接集成在例如 `go build` 这种功能上，或许使用参数打开，拭目以待。



## 参考资料
- https://go.dev/blog/vuln