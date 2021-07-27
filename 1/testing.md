# go语言中的测试

- 错误测试
- 基准测试
- 例子输出
- main测试
- 子测试
- 跳过测试
- 文件系统测试
- io测试
- 黑盒测试
---
- http测试
- **性能分析：cpu，g等的使用情况**
- http请求跟踪测试


测试文件的命名是有一套规则的，通常是某个文件相对应的测试文件，比如`app.go`的测试文件就是`app_test.go`

## 错误测试
错误测试，也是测试中最基础的一种,test首字母要大写，后面的函数（测试谁写谁）首字母也要大写。使用`go test`命令进行启动。
```go
func TestXxx(t *testing.T){
    if xxx {
        t.Errorf("xxx")
    }
}
```
## 基准测试
所谓基准测试，指的是go语言提供的某个算法或者程序执行一定的次数，然后输出平均的执行时间这个就叫做基准测试

跟test一样`B`大写，`Benchmark` 后面的函数首字母也要大写。
```go
func BenchmarkXxx(*testing.B){
    // 这里的b.N go会自动提供，次数不一定。
     for i := 0; i < b.N; i++ {
        // 这里就是要测试的内容
        rand.Int()
    }
}
```
如果想在多线程的环境中测试，go给出了一个例子：

```go
func BenchmarkTemplateParallel(b *testing.B) {
    templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
    // 这里的 RunParallel函数是关键。
    b.RunParallel(func(pb *testing.PB) {
        var buf bytes.Buffer
        for pb.Next() {
            buf.Reset()
            templ.Execute(&buf, "World")
        }
    })
}
```
## 范例测试
范例测试的意思就是说，运行的结果要跟你提供的例子保持一致
```go
func ExampleHello() {
    fmt.Println("hello")
    // Output: hello
}

func ExampleSalutations() {
    fmt.Println("hello, and")
    fmt.Println("goodbye")
    // Output:
    // hello, and
    // goodbye
}
```
这里是有固定形态的，`//Output:`是固定的用法，后面跟例子输出的结果。
## main测试
测试来控制哪些代码在主线程上运行。

```go
func TestMain(m *testing.M){
os.Exit(m.Run())
}
```
## 子测试
```go
使用`t.Run()`可以进行子测试。
func TestTeardownParallel(t *testing.T) {
    // This Run will not return until the parallel tests finish.
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}

```
关于子测试，命令行里的命令是不一样的：

```bash
go test -run ''      # Run all tests.

go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".

go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".

go test -run /A=1    # For all top-level tests, run subtests matching "A=1".

```
## 跳过测试
如果想跳过某些条件，可以使用`t`或者`b`的`.Skip()`方法。
```go
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    ...
}
```

## 文件系统测试
这个包是啥意思呢？其实就是帮你模拟了一个文件系统，因为你比如要打开xx文件吧，你不需要单独真的去新建一个，使用这个文件系统的测试，就可以达到这个目的，这个包是`testing/fstest`

```go
	// 声明一个fstest.MapFs 对象，因为这个对象实现了fs.Fs接口
	//【func (fsys MapFS) Open(name string) (fs.File, error)】
	var ms fstest.MapFS
	// 如果这里不声明ms是这个对象，而是直接就初始化底层类型给ms，
	//下面函数要使用的ms就不是fstest.Mapfs对象，而是map[string]*fstest.MapFile对象
	// 当然你也可以直接不初始化，下面使用的时候显示转化一下即可
	// fstest.MapFs(ms) 即可。


	ms = make(map[string]*fstest.MapFile)
	mf1 := &fstest.MapFile{
		Data: []byte("test"),
		Mode: 30,
		ModTime: time.Now(),
		Sys: "12",
	}
	mf2 := &fstest.MapFile{
		Data: []byte("test1"),
	}
    // 前面是路径，后面是文件。这是一个模拟。
	ms["a/1"] = mf1
	ms["a/2"] = mf2
	fmt.Println(fstest.TestFS(ms,"a/1","a/2","a/3"))

```
这里多说一点，go里面的显示类型转化，刚才讲的一般使用的时候，type A int，A类型并不是int，只是它的底层是int，虽然它的一切操作都可以按照int来做，比如
```go	
    type A int

     var a A
	//a+1 就等于1，
    
	// 但是它仍然是A类型不是int，要注意类型转化，只有一个地方go会自动的语法糖，就是return的时候
	func fast()A{
			return 1
	}
	// 这里 return的时候进行自动类型判断了，没有把1判断为int，而是判断为A类型了。这属于语法糖。
```  
## io测试
io测试包`testing/iotest`主要是实现了readers和writers，具体我们可以理解为，它实现了很多读取和写入。

## 黑盒测试
黑盒测试，使用的包是`testing/quick`

比如一个可以快速得到是否正确的函数

```go
func main(){
	f := func() bool{
		return 1 == 2
	}
	quick.Check(f,nil)
}

```
首先，对于我们来说，qucik.Check()是这个盒子的外壳，我们只能看到它，f我们是看不到的，所以我们可以通过check得到f的返回bool结果。

下面这个函数就是可以黑盒看f f1 是否是一致的。

```go
func main(){
	f := func() bool{
		return true
	}
	
	f1 := func()bool {
		return false
	}
	
	quick.CheckEqual(f,f1,nil)
}
```

## http测试
http测试，意思就是当你需要一个服务的时候，不需要自己再写一个http服务，你只需要`net/http/httptest`包即可。

这个包大致可以分为三个内容
1. request，请求
注意此包并不是客户端的请求，这是服务端的请求。
```go
func NewRequest(method, target string, body io.Reader) *http.Request
```

目标是RFC 7230“请求目标”：它可以是路径或绝对URL。如果目标是绝对URL，则使用URL中的主机名。否则，将使用“ example.com”。

method 空是get
2. response，响应
这个包就是生成一个响应。
```go
func NewRecorder() *ResponseRecorder
```
3. server，服务
服务器是侦听本地接口上系统选择的端口的HTTP服务器，用于端到端HTTP测试。

```go
// 在这段代码中，第一段是一个server，下面是一个客户端get请求。所以上面哪个server监听了本地的请求。
func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

```
## 性能分析
`net/http/pprof` 包提供了例如gc，内存，cpu等数据的性能分析包。

如果要使用这个功能，需要写入这个`import _ "net/http/pprof"`

以及将下面代码加入

```go
go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()

```
新开一个go的groutine，然后来进行性能分析。

> https://golang.org/pkg/net/http/pprof/

```go 
go tool pprof -http=:6062 http://localhost:6060/debug/pprof/block

go tool pprof -http=:6062 http://localhost:6060/debug/pprof/goroutine

go tool pprof -http=:6062 http://localhost:6060/debug/pprof/cpus


```
使用这个命令，可以把数据的分析，使用浏览器打开，http跟的端口，是自己设定的，后面的是分析的具体参数，比如`/block` `/heap `等。下面有个列表，最前面就是这些命令。

> runtime/pprof
pprof的具体实现，所有类型的代码都可以使用。如果不是Web应用程序，建议使用该包。

|类型|描述|备注|
|:---:|:---:|:---:|
|allocs|	内存分配情况的采样信息|	可以用浏览器打开，但可读性不高|
|blocks|	阻塞操作情况的采样信息	|可以用浏览器打开，但可读性不高|
|cmdline|	显示程序启动命令及参数|	可以用浏览器打开，但可读性不高|
|goroutine|	当前所有协程的堆栈信息	|可以用浏览器打开，但可读性不高|
|heap|	堆上内存使用情况的采样信息	|可以用浏览器打开，但可读性不高|
|mutex|	锁争用情况的采样信息	|可以用浏览器打开，但可读性不高|
|profile	|CPU 占用情况的采样信息	|浏览器打开会下载文件|
|threadcreate	|系统线程创建情况的采样信息	|可以用浏览器打开，但可读性不高|
|trace|	程序运行跟踪信息|	浏览器打开会下载文件|
## http请求跟踪测试
`net/http/trace`包提供了监听http请求的各个过程的功能，我们来看一个例子

```go
func main() {
    // 这里有一个新的request
	req, _ := http.NewRequest("GET", "http://example.com", nil)
    // 这里，有两个参数被监听
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
	}
    // 将钩子放入这个http的请求之内，实现监听的效果。
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
```