# API规范

- REST
- RPC
- GraphQL

## RESTful
> RESTful模式的接口 统一使用http协议 

满足于以下规范才可以：

- 资源名称不能是动词，只能是名词
    - 一堆资源使用 examples.com/users
    - 特定资源使用 examples.com/users/admin
- 结尾不要加`/`
- 不要出现下划线`_`，使用`-`来代替
- 路径统一都是小写
- 避免层级过多，如果资源过多，可以转化为parms
    - bad ： /students/chinese/boy/teen/zhang
    - good: /students?contry=chinese&sex=boy&year=teen&name=zhang
- 可以将一个操作变成资源的一个属性，例如 `/students/liming?active=false` 就是禁掉了这个学生
- 使用`:id`的模式，例如 put /students/:id/score 
- 非常规可以设置为动词，或者词组，例如 /login

RESTful的操作方法有四种：
- GET 满足幂等性，满足安全性
- POST 不满足，不满足
- PUT 满足，不满足
- DELETE 满足 不满足

这就是它用http的协议的原因。

因为post不满足幂等性，所以说，更改状态，属性的时候使用PUT,POST仅仅用来创建或者批量删除这两种场景

解决 delete 方法无法携带多个资源名的问题：
- 发起多个delete请求
- 操作路径中带多个id，id之间使用分隔符分割，比如 DELETE /users?id=1,2,3,4
- 直接使用POST 方法批量删除，body中传入需要删除的资源列表

API的版本有三种形式
- 放到URL中 v1/users
- http header 参数中 
- form 参数中

API的命名通常可以有驼峰法（myStudent），下划线法（my_student）和短线法（my-student）

一般来说，短线法更好一些，因为短线不牵涉到输入法的切换问题

api应该提供，分页，过滤，搜索，等功能：
- 分页，比如 /users?offset=1&limit=20
- 过滤 ，比如 /users？fields=email，username，address
- 排序 /users？sort=age，desc
- 搜索 ，当一个资源的成员过多的时候，那么就需要搜索的功能，可以提供模糊搜索的功能 /users?search=age-17，sex=man 意思就是搜索大于17岁的男性

api的域名一般有两种形式：

- https://examples.com/api
- https://v1.api.examples.com
    这种方式的意思就是不止一套api，例如说：
    - https://v1.api.examples.com
    - https://v2.api.examples.com

## rpc

- 更快的传输速度，二进制传递会节省io操作
- 跨平台，满足多语言的互相调用
- 良好的扩展性和兼容性
- 基于idl，通过proto3工具生成制定的语言的数据结构，服务端和客户端接口

protocol buffers 定义的数据结构：

```go
// 定义的数据结构
message SecretInfo {
    string name = 12
    string secret_id = 89
    int64 expires = 9
}
// 定义的接口
service Cache{
    rpc somethind(ruests)returns(response){}
}
```
实现一个grpc需要下面这几个步骤：
- 定义gRPC服务
- 生成客户端和服务器代码
- 实现gRPC服务
- 实现gRPC客户端

代码目录如下：

```bash

$ tree
├── client
│   └── main.go
├── helloworld
│   ├── helloworld.pb.go
│   └── helloworld.proto
└── server
    └── main.go

```
- client

    存放client端代码
- helloworld

    存放服务的idl定义
- server

    存放server端的代码

## graphql
> 可以参考一下这里
https://graphql.cn


