# go标准库的简单介绍
https://pkg.go.dev/std
## archive
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
### compress/bzip2
### compress/flate
### compress/gzip
### compress/lzw
### compress/zlib
## container
### container/heap
### container/list
### container/ring
## context
## crypto
### crypto/aes
### crypto/cipher
### crypto/des
### crypto/dsa
### crypto/ecdh
### crypto/ecdsa
### crypto/ed25519
### crypto/elliptic
### crypto/hmac
### crypto/md5
### crypto/rand
### crypto/rc4
### crypto/rsa
### crypto/sha1
### crypto/sha256
### crypto/sha512
### crypto/subtle
### crypto/tls
### crypto/x509
#### crypto/x509/pkix
## database
### database/sql
### database/sql/driver
## debug
### debug/buildinfo
### debug/dwarf
### debug/elf
### debug/gosym
### debug/macho
### debug/pe
### debug/plan9obj
## embed
## encoding
### encoding/ascii85
### encoding/asn1
### encoding/base32
### encoding/base54
### encoding/binary
### encoding/csv
### encoding/gob
### encoding/hex
### encoding/json
### encoding/pem
### encoding/xml
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
