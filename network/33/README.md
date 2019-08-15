# keep-alive

在HTTP1.1之间，模式使用的是短链接，所以如果想使用长连接就要使用本字段

connection：keep-alive 就可以实现长连接了。

那么在http1.1之后默认长连接了，该如何取消呢，只需要客户端或者是服务器调用 connection:close 即可。
