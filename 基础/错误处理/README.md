<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-11-29 11:49:04
 * @FilePath: /GOFamily/工程/错误处理/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 错误处理
## 错误处理的基本认识
在 go 语言中，没有传统编程语言的 `try - catch` 操作，go 语言中一切错误都需要显式处理：
```go
if err := readFile("./x"); err != nil {
	return err
}
```
通常，我们规定函数返回的最后一个数据是错误接口：
```go
func age(v int)(int, error){
	if v > 10 {
		return 10, nil
	}
	return -1, fmt.Errorf("错误x")
}
```
我们直接返回一个 err 是一种简单的做法，如果错误比较初级也可以这么做，但是如果想要带有更精确的提示信息，可以在返回的时候 wrap 一层信息：

就以上文的读取数据为例
```go
if err := readFile("./"); err != nil {
	return fmt.Errorf("在读取数据的时候发生了错误，错误信息是：%w", err)
}
```
wrap 一层信息，对于错误的定位更加高效

## Error 的本质是什么？
错误处理的核心就是下面这一个 error 的接口
```go
type error interface{
	Error()string
}
```
所以只要我们的自建类型实现了这个接口，就可以使用很多的错误处理的方法。
### 自定义 error
我们使用 errors.New() 的时候其实就是返回了一个 go 自建的，叫做 errorString 的实现了 error 接口的结构体。
这就是自建 error 的方法
```go
func main() {
	e := errors.New("a")
	println(e)
}
// (0x47fd48,0xc00003e730)
```
为了防止在比较错误的时候发生错误一致的情况，所以自建 error，返回的实际上是一个指针。
> 下文会提用什么方法进行比较 err，实际上就是 “两个接口类型是否相等 --- 类型一致，值一致”，如果返回的值是指针，那么值肯定就不可能一样了。

```go
// go 源代码

func New(s string) error {
	return &errorString{s}
}
```
当我们使用 fmt.Errorf() 的时候，其实也是使用的上述方法。

不过，如果我们使用了占位符 `%w` 时，将不会使用上述方法，而是使用了 wrapError：

```go
type wrapError struct {
	msg string
	err error
}
func(e *wrapError) Error()string{
	return e.msg
}
func(e *wrapError) Unwrap() error{
	return e.err
}
```
使用这种方式主要是为了错误链，那就让我们看一下如何使用错误链的相关操作。

### errors.Is()
上文我们说到，错误可以使用 wrap 的方式进行封装，那么如果我们想判断封装的这么多层的错误中，有没有哪一层错误等于我们要的值，可以使用这个函数进行判断：

```go
func main(){
	err := errors.New("err is now")
	err1 := fmt.Errorf("err1:%w",err)
	err2 := fmt.Errorf("err1:%w",err1)
	if errors.Is(err2,err) {
		fmt.Println("是一个错误")
	}
}
```
### errors.As()
这个方法跟上文的 Is() 类似，只不过它判断的是类型是否一致。
```go
type errS struct {
	a string
}

func (t *errS) Error() string {
return t.a
}

err := &errS{"typical error"}
err1 := fmt.Errorf("wrap err: %w", err)
err2 := fmt.Errorf("wrap err1: %w", err1)

var e *errS
// 这里的 target 必须使用指针
if !errors.As(err2, &e) {
	panic("TypicalErr is not on the chain of err2")
}
	println("TypicalErr is on the chain of err2")
	println(err == e)
```
### errors.Join()
这个方法是将几个错误结合在一起的方法：
```go
	err1 := fmt.Errorf("err1:%w",err)
	err2 := fmt.Errorf("err1:%w",err1)
	err := errors.Join(err1,err2)
```
## 当错误处理遇到了 defer
```go
func age() (int, error){
	if xx {
		return 0 ,err
	}
	defer f.close
}
```
这段伪代码的意思是说，当有条件之后，返回一个错误，但是 defer 的内容发生在这个 err 被固定之后，所以 defer 中如果再有错误将不会被处理。

那么我们该怎么更改呢？

我想你一定想到了前文我们说过，使用带有变量的返回值是可以将 defer 的值进行返回的：

```go
func age() ( i int, e error){
	if xx {
			e = xx
			i = xx
		return 
	}
	defer f.close
}
```
那么这种写法，defer 中如果发生了错误就会覆盖掉了程序执行中的 err，所以这种方法也是不行的，即使它能照顾到了 defer 中的错误处理。

我们可以将错误处理都放在 defer 中处理就可以了
```go
func age() ( i int, e error){
	if xx {
		i = xx
		e = xx
		return 
	}
	defer func(){
		e1 := f.close()
		if e != nil {
				if  e1 != nil {
					log(e1) 
				}
				return 
		}
		err = e1
	}()
}
```
这样，两种错误都能处理到了

## 错误处理实战的五种方式
### 经典的错误处理方式
每一个步骤分别直接处理错误
```go
type age interface{
	getAge()error
	putAge()error
	allAge()error
}
func D(ag age)error{
	if err != ag.getAge();err != nil {
		return fmt.Errorf("%w",err)
	}
	if err != ag.putAge();err != nil {
		return fmt.Errorf("%w",err)
	}
	if err != ag.allAge();err != nil {
	  return fmt.Errorf("%w",err)
	}
}
```
### 屏蔽过程中的错误处理
将错误放置在对象的内部进行处理：
```go
type FileCopier struct{
	w *os.File
	r *os.File
	err error
}
func (f *FileCopier)open(path string)( *os.File,error){
	if f.err != nil {
		return nil, f.err
	}

	h, err := os.Open(path)
	if err!= nil {
		f.err = err
		return nil, err
	}
	return h,nil
}

func(f *FileCopier)OpenSrc(path string) {
	if f.err != nil {
		return
	}
	f.r,f.err = os.Open(path)
	return 
}
func(f *FileCopier) CopyFile(src, dst string) error {
	if f.err != nil {
		return f.err
	}
	defer func(){
		if f.r != nil {
			f.r.Close()
		}
		if f.w!= nil {
			f.w.Close()
		}
		if f.err != nil {
			if f.w != nil {
				os.Remove(dst)
			}
		}
	
	f.opensrc(src)
	f.createDst(dst)
	return f.err
		}
	}()
}
```
这段代码并不是特别完整，但是从中我们还是可以理解这种将错误放在对象中的写法的技巧。

首先，错误直接放置在对象自身，在方法中首先去调用这个字段来看是否拥有错误，如果有，直接退出即可

如果没有错误继续往下走，如果本次方法发生错误就继续将这个错误赋值给这个字段，

当最后处理的方法时，这里也就是 copyfile 方法，我们在 defer 中要对于各个子方法进行判断，到底是哪个方法有错误，然后逐一进行判定。相当于处理错误的逻辑集中放置到了最后一个函数进行执行了。

也就是说，将错误放置在对象本身的时候，通常应该为顺序调用的方法，一旦前者出现错误，后者即可退出

如果不是顺序的执行过程，那么有些的错误就可能被湮没，导致错误无法被感知。

### 分层架构中的错误处理方法
常见的分层架构
- controller 控制层
- service 服务层
- dao 数据访问层

dao 层生产错误
```go
if err != nil {
	return fmt.Errorf("%w",err)
}
```
service 追加错误
```go
err := a.Dao.getName()
	if err != nil {
	return fmt.Errorf("getname err: %w",err)
	}
}
```
controller 打印错误
```go
if err!= nil {
	log(err)
}
```
### pkg/errors

如果感觉标准库提供的错误处理不够丰富，也可以使用 github.com/pkg/errors 来处理错误

此包常用的方法有
```go
// 生成新的错误，同样会附加堆栈信息
func New()error
// 只附加新的消息
func WithMessage(err error,message string) error
// 只附加堆栈信息
func WithStack(err error)error
// 附加信息 + 堆栈信息(就是一大堆的各种文件的堆栈调用过程的详细信息)
func Wrapf(err error,format string,args...interface{}) error
// 获取最根本的错误（错误链的最底层）
func Cause(err error) error
```
例如：
```go
func main(){
	err := age()
	if err != nil {
		// %+v 是 pkg/errors 包提供的格式化输出格式，输出错误堆栈
		fmt.Printf("%+v",err)
	}
}

func age()error {
	return errors.Wrap(err,"open error")
}
```

在使用这个 pkg/errors 包的时候要注意一件事，因为 wrap 可以包装堆栈向上输出，如果你调用的第三方库使用了 wrap，你再次使用 wrap，那么就会出现两堆相同的堆栈信息，这造成了极大的冗余。

所以，

- 在提供多人使用的三方库的时候不要使用 wrap，只有逻辑代码的时候使用 wrap 的功能
- 遇到一个错误不打算处理，那么要带上足够多的信息再向上抛出
- 一旦错误处理完成之后就没有错误了，不再需要把错误继续网上抛，返回 nil 即可

所以我们使用 pkg/errors 包将上面的分层写法做一个更完善的改进：

dao 层生产错误
```go
// a
func getName()error {
	if err != nil {
		// 返回此处的错误堆栈
		return errors.Wrap(err,"error:")
	}
	return nil
}

```
service 追加错误
```go
err := a.getName()
	if err != nil {
		// 不返回堆栈了，仅仅添加错误
		return errors.WithMessage(err,"getName error")
	}
}
```
controller 打印错误
```go
if err!= nil {
	// 添加日志
	log(err)
}
```
### errgroup 的使用技巧
errgroup 的使用方法是 golang.org/x/sync/errgroup
```go
package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)


func main() {
	g, ctx := errgroup.WithContext(context.Background())
	// 启动一个 goroutine去处理错误
	g.Go(func() error {
		return fmt.Errorf("error1")
	})
	g.Go(func() error {
		return fmt.Errorf("error2")
	})
	// 类似 waitgroup 的 wait 方法
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
```

## 错误处理相关技巧
这里会介绍在实战过程中用到的诸多技巧
### 使用 errors.New() 时要写清楚包名
```go
package age

ErrMyAge := errors.New("age: ErrMyAge is error")
ErrMyAddress := errors.New("age: ErrMyAddress is error")
```
### 使用 error 处理一般错误，使用 panic 处理严重错误 (异常)
使用这种模型就避免了类似 Java 那种所有错误都一样的行为，Java 的使用 try-catch 的方式导致任何错误都是一个方式去处理，非常有可能让程序员忽略错误的处理

然而 go 不同，**错误**使用 error，**异常**使用 panic 的方式去处理。
- 错误：error
- 异常：panic

假设我们在代码中使用了 panic，通常来说，为了代码的健壮性还是会使用 defer 函数去运行一个 recover() 的，程序的存活比啥都重要。
### 基础库，应该避免使用 error types
因为这种写法容易造成代码的耦合，尤其是在我们写的基础库中，非常容易造成改动代码来引入的不健壮性。

使用自定义的 error type
```go
package main

import (
	"fmt"
)

// 为了额外增加更多的错误信息，字段需大写
type ErrMyAge struct {
	EV string
	MErr int
}

func (e *ErrMyAge) Error() string {
	return fmt.Sprintf("age: %s", e.EV)
}

func main() {
	err := &ErrMyAge{"err age is hi"}
	fmt.Println(err)
}
```

使用 errors.New() 哨兵错误模式：

```go
// 我方代码
ErrAge := errors.New("age: ErrAge is error")
ErrAddress := errors.New("age: ErrAddress is error")

---------
// 使用者

import(
	"github.com/shgopher/age"
)

func age(){
		// 带来了耦合
		if errors.Is(err, age.ErrAge){
		// 处理
		}
	}
}
```

实际上他们都是 error types，只不过前者比后者增加了很多额外的信息，但是相同点是，他们都造成了耦合

如果别人使用了这个基础库，那么势必这些错误就会跟使用者的代码耦合，我们改动了代码，第三方的代码就会因此受到影响。

因此，在对外暴露的基础包中，我们应尽量减少定义哨兵错误 (上述定义方法被称之为哨兵模式的错误)

上述提供的哨兵模式是透明的错误导出机制，所以容易造成耦合

我们可以提供不透明的机制，不导出透明的错误类型，仅仅让用户判断是否等于 nil，就可以防止耦合的存在

```go
import ("github.com/shgopher/age")

func age(){

	if err:= age.Bar(); err == nil {
		return
	}
	
}
```
这也是大多数程序应该提供的模式，不对外暴露，避免了耦合

那么这种方法的缺陷也很明显了，就是无法获取更多的错误信息，理论上来说，我们也没必要获取那么多错误信息，但是如果真的要获取错误信息了，该如何去做呢？

我们可以向外暴露一些动作，只暴露行为：

```go
type mage interface {
	age() bool
}

// 通过一个对外暴露的函数可以对外输出行为
// 并且还不用暴露出具体的对象
func IsMage(err error)bool {
	 a, ok := err.(mage)
	 return ok && a.age()
}
```
### 优化不必要的代码让程序变得更简洁
方法一将大函数变小函数，通过封装函数的方法从视觉上降低 if err 的影响。

```go
// 改造之前

func OpenFile(src ,dst string) error {
	if r,err := os.Open(src); err != nil {
		return err
	}
	if w,err := os.Create(dst); err != nil {
		r.Cloase()
		return err
	}
	if _,err := io.Copy(w,r); err!= nil {

		return err
	}
	return nil
}

// 将前面两个操作封装成一个函数

func OpenD(src dst string) (*os.File,*os.File,error) {
	var r ,w *os.File 
	var err error

	if r,err = os.Open(src); err!= nil {
		return nil,nil,err
	}
	if w,err = os.Create(dst); err!= nil {
		r.Close()
		return nil,nil,err
	}

	return r,w,err
}
// 主函数就只有一个 if err 了
func OpenFile(src,dst string)error {
	var err error
	var r, w *os.File
	if r,w,err = OpenD(src,dst); err!= nil {
		return err
	}
	defer func(){
		r.Close()
		w.Close()
	}()
	if _,err = io.Copy(w,r); err!= nil {
		return err
	}
	return nil
}
```
方法二将代码中不必要的成分删除，来保证代码的简洁
```go
// 改造前
func AuthticateRequest(r *Request)error{
	err := authticate(r.User)
	if err != nil {
		return err
	}
	return nil
}
// 实际上 authticate 只返回一个 error 类型的接口，根本不需要判断
func AuthticateRequest(r *Request)error{
	return authticate(r.User)
}
```
方法三使用更加合适的方法
```go
// 我们要逐行去读取数据

// 改造前
func Countlines(r io.Reader)(int,error){
	var(
		br = bufio.NewReader(r)
		line int
		err error
	)

	for {
		_,err = br.readline('\n')
		line ++ 
		if err != nil {
			break
		}
	}
		if err != io.EOF {
			return 0,err
		}
		return lines,nil
		
}

// 代码看起来也是很合理的样子，也很简洁，但是，我们其实用的函数不是特别的合适

//其实这里使用 scan 更加合适，代码量更加精简，并且结构异常舒服
func Countlines(r io.Reader) (int,error){	
	sc := bufio.NewScanner(r)
	lines := 0

	for sc.Scan() {
		lines++
	}
	
	return lines, sc.Err()
}

``` 

### 错误应该只处理一次

我们有日志系统，有些时候我们发现一个错误，然后打了一个日志，然后又把错误给 return 了，实际上这与 go 语言哲学中说的错误只处理一次相违背
```go
// ❌
func age()error{
	 _, err := os.Open("./a")
	if  err!= nil {
		log.Prinln(err)
		return err
	}
}
```
```go
// ✅
func age()error{
	 _, err := os.Open("./a")
	if  err!= nil {
		return errors.Wrap(err,"open error")
	}
}
```
正确的处理方法是错误只处理一次，那么在什么时候处理呢？

上文提到了多层的架构设计，我们在底层 (带有堆栈的错误向上抛出) 和中层 (仅仅附加信息再次向上抛出) 仅仅是向上抛出，并不需要将错误记录在日志中，在应用层才需要去使用日志记录错误，日志记录完错误以后，也不需要再向上抛出错误了 (最顶端了) 完全满足 “只处理一次错误” 的要求。


## 业务 code 码的设置
常见的 http 错误码数量较少，比如常见的只有例如 404 301 302 200 503 等，绝对数量还是较少，无法去表达业务上的错误，因此我们需要设置一套能表达具体生产业务的 code 码。

为了保证服务端的安全，我们设置的 code 码应该设置两套数据，一套显示给客户端，一套自用，以此来保证服务端的绝对安全。

有三种设计业务 code 码的方式：

### 一律返回 http status 200，具体 code 码单独设置

例如
```json
{
  "error": {
    "message": "Syntax error \"Field picture specified more than once. This is only possible before version 2.1\" at character 23: id,name,picture,picture",
    "type": "OAuthException",
    "code": 2500,
    "fbtrace_id": "xxxxxxxxxxx"
  }
}
```
- http status code 通通 200
- code 2500，才是真实的面向客户端的 code 码

使用这种方法的一大缺陷就是必须解析 body 内容才能发现具体的错误业务码，很多场景我们仅仅需要知道返回的是成功或者错误，并不需要知晓具体的业务码，这是这种方式的一大弊端。

### 使用合适的 http status code + 简单的信息以及业务错误代码
```bash
HTTP/1.1 400 Bad Request
x-connection-hash: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
set-cookie: guest_id=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
Date: Thu, 01 Jun 2017 03:04:23 GMT
Content-Length: 62
x-response-time: 5
strict-transport-security: max-age=631138519
Connection: keep-alive
Content-Type: application/json; charset=utf-8
Server: tsa_b

# 注意这里：仅仅返回简单的错误信息
{"errors":[{"code":215,"message":"Bad Authentication data."}]}
```
这种方案也是大多数公司采纳的方案，使用具体的 http status code 可以知晓大概的业务类型，是错误还是正常运行，然后使用简单的错误信息和业务错误代码去定位具体的错误

如果业务不是特别复杂，使用这种方式即可
### 使用合适的 http status code + 非常详细的业务错误代码以及信息
```bash
HTTP/1.1 400
Date: Thu, 01 Jun 2017 03:40:55 GMT
Content-Length: 276
Connection: keep-alive
Content-Type: application/json; charset=utf-8
Server: Microsoft-IIS/10.0
X-Content-Type-Options: nosniff

{"SearchResponse":{"Version":"2.2","Query":{"SearchTerms":"api error codes"},"Errors":[{"Code":1001,"Message":"Required parameter is missing.","Parameter":"SearchRequest.AppId","HelpUrl":"http\u003a\u002f\u002fmsdn.microsoft.com\u002fen-us\u002flibrary\u002fdd251042.aspx"}]}}
```
当业务逻辑稍微复杂一些，并且需要极其精准和快速的定位错误时，就需要在返回的 body 中去设置非常详细的错误信息

**综上所述：**

- 使用正确的 http status code 让业务的第一步变得更加直观
- 区别于 http status code，具体业务的 code 码会更加丰富
- 返回尽可能详细的错误，有助于复杂逻辑的快速错误定位
- 直接返回给客户的错误代码不应该包括敏感信息，敏感信息的 code 码仅供内部使用
- 错误信息要求规范，简洁以及有用

### 业务 code 码的具体设计
引入业务 code 码的核心原因就是 http status code 太少，以及他们并不能跟具体业务挂钩。

当我们设置好良好又详细的 code 码时，我们就可以快速定位业务代码，以及可以快速知晓发生错误的等级模块，具体信息等

下面给出具体的设计思路：**纯数字表达，不同的数字段表达不同的模块不同的业务**

例如 100101
- 10：表示某个服务
- 01：表示某个服务下的模块
- 01：模块下的错误码

10 服务 01 模块 01 错误，--- 服务 10 数据库模块未找到记录错误

一共最多有 100 个服务，每个服务最多有 100 个模块，每个模块最多有 100 个错误，如果某些模块 100 个都不够用，那怎么这个模块有必要去拆分一下了。

### 如何设置 http status code

- `1xx`：请求已接收，继续处理
- `2xx`：成功处理了请求
- `3xx`：请求被重定向
- `4xx`：请求错误
- `5xx`：服务器错误

由于 http status code 相对数量也不算太少，如果每一个都利用上，难免会增加复杂度，建议仅使用基本的几个即可

- 200 - 表示请求成功执行。
- 400 - 表示客户端出问题。
- 500 - 表示服务端出问题。

如果上述的感觉太少，再增加下面几个也可以

- 401 - 表示认证失败。
- 403 - 表示授权失败。
- 404 - 表示资源找不到，这里的资源可以是 URL 或者 RESTful 资源。

将 http status code 控制在**个位数**，有利于后端的逻辑代码简洁性，比如 301 302 确实是代表不同的含义，**前端或许可以设置丰富的 http status code，因为浏览器会进行相关的具体操作，但是后端返回给前端的 http status code 并没有任何的操作，使用过多只会增加复杂度。**

## 设计一个生产级的错误包
### 生产级的错误包需要的功能

1. 支持错误堆栈

```go
2021/07/02 14:17:03 call func got failed: func called error

main.funcB /home/colin/workspace/golang/src/github.com/marmotedu/gopractise-demo/errors/good.go:27
main.funcA /home/colin/workspace/golang/src/github.com/marmotedu/gopractise-demo/errors/good.go:19
main.main /home/colin/workspace/golang/src/github.com/marmotedu/gopractise-demo/errors/good.go:10
runtime.main /home/colin/go/go1.16.2/src/runtime/proc.go:225runtime.goexit /home/colin/go/go1.16.2/src/runtime/asm_amd64.s:1371
exit status 1
```
拥有错误的堆栈，我们才能定位错误的根源。

2. 支持打印不同的格式比如 %s %w %d

3. 支持 Wrap() Unwrap() 的功能，就是错误的嵌套和反嵌套

4. 错误包要支持 Is() 和 As() 方法，这主要是因为有错误的嵌套，所以无法再使用接口相比较的方式进行判断接口类型是否相等 (类型相同，值相同)

5. 要支持格式化和非格式化的创建方式

```go
errors.New("err")
fmt.Errorf("%w",err)
```

### 具体实现
从 github.com/pkg/errors 包中改造即可。

增加以下字段的结构体就可以满足上面的需求
```go
type withCode struct {
	err   error
	code  int
	cause error
	*stack
}
```

## issues
`问题一：` **请说出下列代码的执行输出***

```go
package main

import "fmt"

func main() {
	defer func() {
		println("3")
	}()
	defer func() {
		println("2")
	}()
	defer func() {
		println("1")
	}()
	fmt.Println("hi1")
	panic("oops")
	// 这里的defer将不会进栈，所以也就不会执行了。
	defer func() {
		println("x")
	}()
	fmt.Println("hi2")
}

```
答案是
```go
hi1
1
2
3
oops
panic: oops

goroutine 1 [running]:
main.main()
	/tmp/sandbox1932632082/prog.go:18 +0xa7

Program exited.
```
解释：Panic，意味着恐慌，意思等于 return，所以 panic 下面的数据是无法执行的，defer 不同，他们是顺序的将这些 defer 函数装入函数内置的 defer 栈的，所以在 return 之后，defer 栈会执行，所以这里的 defer 1 2 3 可以执行，Panic 前面的 hi1 可以执行，但是 Panic 之后，相当于 return 后面的 hi2 就无法执行了。

`问题二：` **看一段代码，分析答案**

```go
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("starting f")
	g(2)
	fmt.Println("behind  g") //会终止执行
}

func g(i int) {
	defer fmt.Println("Defer in g", i)
	fmt.Println("Panicking!")
	panic(fmt.Sprintf("%v", i))
	fmt.Println("Printing in g", i) //终止执行
}
```
答案是
```js
starting f
panicking
Defer in g 2
recoverd in f
```

解释一下，首先执行的是 f 函数的代码，然后开始执行 g，在 g 中遇到了 Panic，所以 panic 后面的 parinting in g 就无法执行了，所以执行了 defer in g
这个时候 f 中的 g(2) 后面的数据也无法执行了，因为整个 f 也陷入了恐慌，所以它只能 return 进入 defer 了，defer 中刚好有 recover，所以执行了 recover 信息后，就退出了函数。

## 参考资料
- https://mp.weixin.qq.com/s/EvkMQCPwg-B0ffZonpwXodg
- https://mp.weixin.qq.com/s/D_CVrPzjP3O81EpFqLoynQ
- https://time.geekbang.org/column/article/391895
- 极客时间《go 进阶训练营》