# flag

实现了命令行参数的解析，意思即是这个包主要的作用就是跟系统中的那个脚本命令行有着很大的关系

例：

```go
var outPutYes
flag.IntVar(&outPutYes, "n", 0, "")
	flag.Parse()
```
