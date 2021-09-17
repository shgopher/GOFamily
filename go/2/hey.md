# 压力测试

压力测试，通常来说是对于某个运行中的系统进行加压，比如说模拟很多请求来对这个系统的一些功能进行测试，这种方法就是压力测试，我们今天介绍的方法，主要是使用`hey`这个软件。

```bash
-n 要运行的请求数。默认是200。
 
-c 并发运行的请求数。请求的总数不能小于并发级别。默认是50。
 
-q 速率限制，以每秒查询(QPS)为单位。默认没有限制。
 
-z 发送请求的应用程序配置。当时间到了，应用程序停止并退出。如果指定持续时间，则忽略n。
    例子:- z 10s - z 3m。

-o 输出类型。如果没有提供，则打印摘要。“csv”是唯一受支持的替代方案。转储文件的响应以逗号分隔值格式的度量。
-m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
-H 自定义HTTP头。您可以通过重复标记指定所需的数量 
    For example, -H "Accept: text/html" -H "Content-Type: application/xml" 
-t 每个请求的超时时间(以秒为单位)。默认值是20，使用0表示无穷大。
-A  HTTP Accept header.
-d  HTTP request body.
-D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
-T  Content-type, defaults to "text/html".
-a  Basic authentication, username:password.
-x  HTTP Proxy address as host:port.
-h2 Enable HTTP/2.
 
-host   HTTP Host header.
 
-disable-compression  禁用压缩。
-disable-keepalive    禁用keep-alive，防止重用TCP不同HTTP请求之间的连接。
-disable-redirects   禁用HTTP重定向的后续操作
-cpus        使用的cpu核数。(当前机器默认为48核)    

```

举个例子：

```go
hey -n 1000 -c 100 -m POST -T "application/x-www-form-urlencoded" -d 'username=1&message=hey' http://your-url/resource
```

输出的结果，默认如下：

```go
Summary:
  Total:	2.2266 secs
  Slowest:	1.2503 secs
  Fastest:	0.0667 secs
  Average:	0.2296 secs
  Requests/sec:	449.1198
  

Response time histogram:
  0.067 [1]	|
  0.185 [553]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.303 [234]	|■■■■■■■■■■■■■■■■■
  0.422 [101]	|■■■■■■■
  0.540 [80]	|■■■■■■
  0.659 [2]	|
  0.777 [12]	|■
  0.895 [8]	|■
  1.014 [1]	|
  1.132 [1]	|
  1.250 [7]	|■


Latency distribution:
  10% in 0.1180 secs
  25% in 0.1363 secs
  50% in 0.1682 secs
  75% in 0.2503 secs
  90% in 0.4301 secs
  95% in 0.4475 secs
  99% in 0.7996 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0457 secs, 0.0667 secs, 1.2503 secs
  DNS-lookup:	0.0010 secs, 0.0000 secs, 0.0075 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0027 secs
  resp wait:	0.1703 secs, 0.0605 secs, 1.0125 secs
  resp read:	0.0127 secs, 0.0002 secs, 0.7614 secs

Status code distribution:
  [404]	1000 responses

```