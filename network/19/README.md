# HTTP状态码

- 1XX	Informational（信息性状态码）	接收的请求正在处理
- 2XX	Success（成功状态码）	请求正常处理完毕
  - 200 请求0k
- 3XX	Redirection（重定向状态码）	需要进行附加操作以完成请求
  - 301 301是永久的重定向，如果在暂时跳转到比如说其它的一个登陆界面再跳转过来的时候一定不能用301
  - 302 暂时的重定向
  - 303 同302 + 必须用GET
  - 304 如果请求报文首部包含一些条件，例如：If-Match，If-Modified-Since，If-None-Match，If-Range，If-Unmodified-Since，如果不满足条件，则服务器会返回 304 状态码。
  - 307 同 302 +  不会把请求的POST改成GET
- 4XX	Client Error（客户端错误状态码）	服务器无法处理请求
  - 400 请求报文语法错误
  - 401 用户认证失败
  - 403 请求被拒绝
  - 404 就是没有发现资源
- 5XX	Server Error（服务器错误状态码）
  - 500 服务器出错
  - 503 服务器正在调试，或者超负荷了

  ## HTTP的不同METHOD之前的区别

  ![p](https://github.com/imgoogege/PHP-Interview-QA/blob/master/docs/assets/network-http-method.png)
