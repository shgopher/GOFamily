# go标准库的简单介绍
https://pkg.go.dev/std
## archive
https://pkg.go.dev/archive

目前不存在 archive 包本身，存在其子包，这个包，包含的内容是对于档案文件的处理，比如 tar，zip 这种文档文件。
### archive/tar
https://pkg.go.dev/archive/tar
```go
import "archive/tar"
```
`archive/tar` 包，实现了对 tar 文件的处理
- tar 文件头的设置
- tar 文件的读写
### archive/zip
https://pkg.go.dev/archive/zip
```go
import "archive/zip"
```
`archive/zip` 包，实现了对 zip 文件的处理
- zip 文件头的设置
- zip 文件的读写
## bufio
https://pkg.go.dev/bufio
```go
import "bufio"
```
`bufio` 包， 提供了有缓冲的 i/o，比 io 包封装程度更高。使用缓冲区来一次读取多个字节，从而减少系统调用的次数。
- 提供了基本的读写功能
- 提供了逐行读取的功能
## builtin 
https://pkg.go.dev/builtin

注意，此包无法使用 import 的方式引入，因为它只是内部包和内部类型的一个存放位置而已

比如有 append() clear() max() min() int float64 等等
## bytes
https://pkg.go.dev/bytes
```go
import "bytes"
```
bytes 包，包含了很多处理 bytes 类型的操作。

跟字符串的基本操作类型
## cmp
https://pkg.go.dev/cmp

```go
import "cmp"
```
cmp 包提供了有关比较的相关内容。
- cmp.Ordered 表示可比较的类型约束
- 提供了比较函数
## compress
https://pkg.go.dev/compress

不直接提供 compress 包，提供了众多子包，所有的内容都是关于压缩算法
### compress/bzip2
https://pkg.go.dev/compress/bzip2
```go
import "compress/bzip2"
```
提供了 bzip2 的解压功能，然而并没有提供压缩的功能
### compress/flate
https://pkg.go.dev/compress/flate
```go
import "compress/flate"
```
flate 实现了 RFC1951 中描述的 DEFLATE 压缩数据格式。
### compress/gzip
https://pkg.go.dev/compress/gzip
```go
import "compress/gzip"
```
gzip 实现了对 gzip 格式压缩文件的解压（读取）和压缩（写入）
### compress/lzw
https://pkg.go.dev/compress/lzw
```go
import "compress/lzw"
```
lzw 实现了 lzw 文件（ Lempel-Ziv-Welch ）的解压缩操作。
### compress/zlib
https://pkg.go.dev/compress/zlib
```go
import "compress/zlib"
```
zlib 实现了对 zlib 格式文件的解压和压缩操作。
## container
https://pkg.go.dev/container

container 没有提供本包，提供了众多子包，这个目录下的内容都是关于容器的，这也是 go 提供的内置容器
### container/heap
https://pkg.go.dev/container/heap
```go
import "container/heap"
```
go 内置的堆，值得注意的是，go语言仅提供了接口以及接口的相关函数，并没有具体的实现，使用时还需要自行实现接口。
### container/list
https://pkg.go.dev/container/list
```go
import "container/list
```
go 内置的双向链表，这里不是接口了，是已经实现好了的双向链表
### container/ring
https://pkg.go.dev/container/ring
```go
import "container/ring"
```
go 内置的循环链表，非接口，已经实现好了
## [context](../并发/context)
https://pkg.go.dev/context
```go
import "context"
```
context 提供了在“多线程”的场景下的线程控制功能，简单的说就是 context 这个上下文可以统一取消所有的上下文环境中的 goroutine
```go
func main() {
  // cal的调用，以及计时器的到达均可调用 ctx.Done() 的发生。
	ctx, cal := context.WithTimeout(context.Background(), time.Second*1)
	defer cal()
	go func() {
		select {
		case <-time.After(time.Second * 2):
			fmt.Print(1)
		case <-ctx.Done():
			fmt.Print(2)
		}
	}()
	time.Sleep(time.Second * 3)
}
```
## crypto
https://pkg.go.dev/crypto
```go
import "crypto"
```
crypto 包，包含了很多加密算法
### crypto/aes
https://pkg.go.dev/crypto/aes
```go
import "crypto/aes"
```
提供了 aes 加密算法的加密过程。此算法为对称加密算法
- 创建一个密钥
### crypto/cipher
https://pkg.go.dev/crypto/cipher
```go
import "crypto/cipher"
```

### crypto/des
https://pkg.go.dev/crypto/des
```go
import "crypto/des"
```
### crypto/dsa
https://pkg.go.dev/crypto/dsa
```go
import "crypto/dsa"
```
### crypto/ecdh
https://pkg.go.dev/crypto/ecdh
```go
import "crypto/ecdh"
```
### crypto/ecdsa
https://pkg.go.dev/crypto/ecdsa
```go
import "crypto/ecdsa"
```
### crypto/ed25519
https://pkg.go.dev/crypto/ed25519
```go
import "crypto/ed25519"
```
### crypto/elliptic
https://pkg.go.dev/crypto/elliptic
```go
import "crypto/elliptic"
```
### crypto/hmac
https://pkg.go.dev/crypto/hmac
```go
import "crypto/hmac"
```
### crypto/md5
https://pkg.go.dev/crypto/md5
```go
import "crypto/md5"
```
### crypto/rand
https://pkg.go.dev/crypto/rand
```go
import "crypto/rand"
```
### crypto/rc4
https://pkg.go.dev/crypto/rc4
```go
import "crypto/rc4"
```
### crypto/rsa
https://pkg.go.dev/crypto/rsa
```go
import "crypto/rsa"
```
### crypto/sha1
https://pkg.go.dev/crypto/sha1
```go
import "crypto/sha1"
```
### crypto/sha256
https://pkg.go.dev/crypto/sha256
```go
import "crypto/sha256"
```
### crypto/sha512
https://pkg.go.dev/crypto/sha512
```go
import "crypto/sha512"
```
### crypto/subtle
https://pkg.go.dev/crypto/subtle
```go
import "crypto/subtle"
```
### crypto/tls
https://pkg.go.dev/crypto/tls
```go
import "crypto/tls"
```
### crypto/x509
https://pkg.go.dev/crypto/x509
```go
import "crypto/x509"
```
#### crypto/x509/pkix
https://pkg.go.dev/crypto/x509/pkix
```go
import "crypto/x509/pkix"
```
## database
https://pkg.go.dev/database
此包不可直接使用，它包含了子包，这个子包是处理sql的统一接口，并不提供实际的实现
### database/sql
https://pkg.go.dev/database/sql

```go
import _ "database/sql"
```
当我们使用 MySQL，Redis 等数据库的时候，通常要使用上述的方式引入这个包，此包提供了 SQL 操作的接口。

使用此包必须跟第三方的数据库驱动结合，比如下面这种操作：

```go
import(
_ "database/sql"
  "xx/mysql"
) 
```
### database/sql/driver
https://pkg.go.dev/database/sql/driver
```go
import "database/sql/driver"
```
driver 包，包含了数据库 driver 的接口，要想实现某个数据库的驱动（driver）就必须引入此包，实现此包定义接口的具体内容。

在被用户使用的时候，引入第三方数据库驱动，加上引入 database/sql 这个 SQL 操作包，就可以实现正常的 SQL 操作。
## debug
https://pkg.go.dev/debug

不提供 debug 包本身，debug 包含了众多子包，都是跟 调试相关。
### debug/buildinfo
https://pkg.go.dev/debug/buildinfo
不直接使用此包，此包提供二进制的功能，由 runtime/debug 来调用。
### debug/dwarf
https://pkg.go.dev/debug/dwarf
```go
import "debug/dwarf"
```
用于解析 DWARF 调试信息，DWARF 调试信息包含了程序的源代码、变量、函数、类型等相关信息，可以帮助调试器进行源代码级别的调试。
### debug/elf
https://pkg.go.dev/debug/elf
```go
import "debug/elf"
```
用于解析 ELF 可执行文件格式，
### debug/gosym
https://pkg.go.dev/debug/gosym
```go
import "debug/gosym"
```
debug/gosym 用于解析 Go 语言程序的符号表信息。
### debug/macho
https://pkg.go.dev/debug/macho
```go
import "debug/macho"
```
debug/macho 用于解析 Mach-O 格式的可执行文件
### debug/pe
https://pkg.go.dev/debug/pe
```go
import "debug/pe"
```
debug/pe 用于解析 PE 格式的可执行文件
### debug/plan9obj
https://pkg.go.dev/debug/plan9obj
```go
import "debug/plan9obj"
```
debug/plan9obj 用于解析 Plan 9 object 文件格式。
## embed
https://pkg.go.dev/embed
```go
import _ "embed"
```
```go
//go:embed hello.text
```
mbed 是 Go 语言自 1.16 版本引入的一个标准库，用于将静态文件嵌入到 Go 代码中。

通过 embed 包，我们可以将静态文件（如文本文件、JSON 文件、HTML 文件、图像文件等）直接嵌入到 Go 代码中，而无需将文件作为独立的资源文件放在磁盘上。这样做的好处是，可以将所有的资源文件打包到可执行文件中，方便分发和部署。

将一个文件嵌入到一个string中
```go
package main

import _ "embed"

// 这里的注释开头没有空格，这是 go 的自有命令注释；//go:xx 
//
//go:embed hello.txt
var s string

func main() {
	print(s)
}
```
将一个文件嵌入到文件系统中
```go
package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var helloFile embed.FS

func main() {
	content, err := helloFile.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}

	fmt.Println(string(content))
}
```
## encoding
https://pkg.go.dev/encoding
使用时，不直接使用 encoding，除非你想亲自实现某个编码的加解码

encoding 是 Go 语言标准库中的一个包，用于处理数据的编码和解码。

encoding 包提供了许多子包，每个子包都专门用于处理特定的数据编码格式

- encoding/json：用于处理 JSON 数据的编码和解码。
- encoding/xml：用于处理 XML 数据的编码和解码。
- encoding/csv：用于处理 CSV（逗号分隔值）数据的编码和解码。
- encoding/base64：用于进行 base64 编码和解码。
- encoding/hex：用于进行十六进制编码和解码。
- encoding/gob：用于进行 Go 对象的编码和解码。
### encoding/ascii85
https://pkg.go.dev/encoding/ascii85
```go
import "encoding/ascii85"
```
encoding/ascii85 用于进行 ASCII85 编码和解码
### encoding/asn1
https://pkg.go.dev/encoding/asn1
```go
import "encoding/asn1"
```
encoding/asn1 用于进行 ASN.1（Abstract Syntax Notation One）编码和解码。
### encoding/base32
https://pkg.go.dev/encoding/base32
```go
import "encoding/base32"
```

### encoding/base64
https://pkg.go.dev/encoding/base64
```go
import "encoding/base64"
```
encoding/base32 用于进行 Base32 编码和解码
### encoding/binary
https://pkg.go.dev/encoding/binary
```go
import "encoding/binary"
```
encoding/base64 用于进行 Base64 编码和解码。
### encoding/csv
https://pkg.go.dev/encoding/csv
```go
import "encoding/csv"
```
encoding/csv 用于处理 CSV（逗号分隔值）格式的数据。

### encoding/gob
https://pkg.go.dev/encoding/gob
```go
import "encoding/gob"
```
encoding/gob 用于将 Go 的值编码为二进制格式，并进行序列化和反序列化。
### encoding/hex
https://pkg.go.dev/encoding/hex
```go
import "encoding/hex"
```
十六进制编码将二进制数据编码为十六进制字符串，每个字节对应两个十六进制字符。十六进制编码通常用于将二进制数据转换为可打印的 ASCII 字符串，例如在 URL 参数中传递二进制数据或在文本文件中嵌入二进制数据。

encoding/hex 包提供了一组函数和类型，用于进行十六进制编码和解码。主要有两个函数：EncodeToString 和 DecodeString。
### encoding/json
https://pkg.go.dev/encoding/json
```go
import "encoding/json"
```
encoding/json 用于对 JSON（JavaScript Object Notation）格式的数据进行编码和解码操作。

JSON 是一种常用的数据交换格式，用于在不同平台和编程语言之间传输和存储数据。JSON 数据由键值对组成，可以表示复杂的数据结构和层次关系。
### encoding/pem
https://pkg.go.dev/encoding/pem
```go
import "encoding/pem"
```
encoding/pem 用于对 PEM（Privacy-Enhanced Mail）格式的数据进行编码和解码操作。

PEM 是一种常用的文本格式，用于在非文本环境中传输和存储密钥、证书等数据。PEM 格式使用 ASCII 字符表示二进制数据，通常以 "-----BEGIN…" 和 "-----END…" 标记来标识不同类型的数据。
### encoding/xml
https://pkg.go.dev/encoding/xml
```go
import "encoding/xml"
```
encoding/xml 用于对 XML（eXtensible Markup Language）格式的数据进行编码和解码操作。

XML 是一种常用的文本格式，用于存储和传输结构化的数据。XML 数据由标签、属性和文本内容组成，可以表示复杂的数据结构和层次关系。
## errors

## expvar
## flag
## fmt
## go
### go/ast
### go/build
#### go/build/constraint
### go/constant
### go/doc
#### go/doc/comment
### go/format
### go/importer
### go/parser
### go/printer
### go/scanner
### go/token
### go/types
## hash
### hash/adler32
### hash/crc32
### hash/crc64
### hash/fnv
### hash/maphash
## html
### html/template
## image
### image/color
#### image/color/palette
### image/draw
### image/gif
### image/jpeg
### png
## index
### index/suffixarray
## io
### io/fs
### io/ioutil
## log
### log/slog
### log/syslog
## maps
## math
### math/big
### math/bits
### math/cmplx
### math/rand
## mime
### mime/multipart
### mime/quotedprintable
## net
### net/http
#### net/http/cgi
#### net/http/cookiejar
#### net/http/fcgi
#### net/http/httptest
#### net/http/httptrace
#### net/http/httputil
#### net/http/pprof
### net/mail
### net/netip
### net/rpc
#### net/rpc/jsonrpc
### net/smtp
### net/textproto
### net/url
## os
### os/exec
### os/signal
### os/user
## path
### path/filepath
## plugin
## reflect
## regexp
### regexp/syntax
## runtime
### runtime/cgo
### runtime/coverage
### runtime/debug
### runtime/metrics
### runtime/pprof
### runtime/race
### runtime/trace
## slices
## sort
## strconv
## strings
## sync
### sync/atomic
## syscall
### sync/js
## testing
### testing/fstest
### testing/iotest
### testing/quick
### testing/slogtest
## text
### text/scanner
### text/tabwriter
### text/template
#### template/parse
## time
### time/tzdata
## unicode
### unicode/utf16
### unicode/utf8
## unsafe
