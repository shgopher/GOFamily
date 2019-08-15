# ETage验证原理

可以将缓存资源的 ETag 值放入 If-None-Match 首部，服务器收到请求后，判断缓存资源的 ETag 值和资源的最新 ETag 值是否一致，如果一致则表示缓存资源有效，返回 304 Not Modified。

304的话：

> If-Match，If-Modified-Since，If-None-Match，If-Range，If-Unmodified-Sinc

这些的字段如果不符合这些的规定，出错的时候就会出现304的错误代码。
