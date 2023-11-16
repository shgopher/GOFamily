# 命令
go 拥有众多命令操作，这里将讲述关于这些命令的使用方法

介绍一下最常见的命令
- go help 显示一个命令基本的用法，例如：`go help fmt`
- go doc 显示一个命令全部的用法，例如：`go doc cmd/gofmt`

使用 go help 可以显示全部的形如 `go fmt` `go build` 这种挂靠在 go 后面的命令，然后 help 加具体的命令，就可以显示基本用法，然后在 help 提示的内容中，通常会有提示你，如果使用 go doc 命令去寻找更加详细的内容，比如下文要写到的，使用 `go help fmt` 就会显示去寻找 `go doc cmd/gofmt`

## gofmt
> go fmt 命令简单封装了 gofmt 命令

gofmt 的目的是标准化 go 语言的代码，增加代码的亲切感，消除不同人员写的代码的之间的隔阂感

介绍几个常见的使用方法，详细内容可以使用 `go doc cmd/gofmt `

- `gofmt -s` 简化代码

  ```go
  v := []int{1,2,3}
    // 复杂
    for _ = range v {

    }

    // 使用 -s 后

    for range v {

    }
  ```
  不过，这个命令虽然会显示出来要优化的简单写法，但是，并不会更改用户的代码，需要自己去更改。

- `gofmt -r` 代码重构 replace 能力

  例子：`gofmt -r 'a -> Student'` 意思可不是 a 字符改变成 Student，这里是采用的通配符，意思就是所有的英文字符都要改成 student。只要是小写字母都会被视为通配符。再举一个例子，`gofmt -r 'a[b:len(a)] -> a[b:]'` 这里的 a 代表所有的英文字符串，b 就会代表整数类型

- `gofmt -l` 输出不满足 gofmt 要求的文件

  比如 `gofmt -l $GOROOT` 就会输出这个路径下不满足的文件列表，可以看出 go 的标准库不满足标准的也不少，😂

## goimports

安装方法 `go get golang.org/x/tools/cmd/goimports`，一般的 IDE 都会内置这个工具，比如 goland

- 对于代码中使用了，但是没有 import 的包
- 对于代码中没有使用，但是 import 了的包

这个工具都会一一管理，少了加上，多了取消掉

## go build
## go install
## go get
## go clean
## go doc godoc
## go run
## go test
## go list
## go fix go tool fix
## go vet / go tool vet
## go tool pprof
## go tool cgo
## go env
