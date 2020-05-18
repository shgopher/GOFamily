# rpc

go在net/rpc包中实现了基本的rpc协议的东西，我们通过一个例子来快速实现一个rpc的例子

服务端：
```go
package main

import (
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func main() {
	rpc.Register(new(Server))// 注册structure
	rpc.Register(new(Server2))
	rpc.HandleHTTP()// 调用了http协议
	l, _ := net.Listen("tcp", ":442")
	http.Serve(l, nil)
}

type Server struct {
	value string
}
type Replay string
type Arg int

func (s *Server) Hello(arg Arg, replay *Replay) error {// 此函数是注册到regist中的数据带有的方法，首先必须是指针方法，然后第一个数据是非引用，第二个是引用类型的返回值，第一个是参数
	a := strconv.FormatInt(int64(arg), 10)
	*replay = Replay(a +"--->>>>-----" +a)
	return nil
}

type Server2 struct {
	Vaue int
}
type Replay2 string
type Are2 int

func (s *Server2) Hello(arg Arg, replay *Replay) error {
	a := strconv.FormatInt(int64(arg), 10)
	*replay = Replay(a +":::::::::::" +a)
	return nil
}

```

客户端：

```go
package main

import (
	"fmt"
	"net/rpc"
)

func main(){
	t := ""
	c,_ := rpc.DialHTTP("tcp",":442") // 初始化这个rpc
	c.Call("Server.Hello",12,&t)// 直接调用方法注意第一个参数的名字是struct.Method
	defer c.Close()
	fmt.Println(t)
	c.Call("Server2.Hello",90,&t)
	fmt.Println(t)
}
```