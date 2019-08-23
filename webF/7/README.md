# Fetch

Fetch是github公司开发的一个window对象，目的是代替传统的Ajax的使用方法。Fetch提供的是基于Promise的
串行处理方法。

```js
fetch(url).
then(res =>res.json()).

then( (json) => {
  console.log(json)

}).catch((e) => {

  console.log(e)
})


```

跨域问题在fetch上设置没有意义，跨域问题要在服务器上进行设置，`Access-Control-Allow-Origin`。设置http
协议中的这个响应头的头信仰即可，如果设置为一个地址，那么就是只允许这个地址请求，如果设置为*就是全部都可以请求。
可以[参考](http://www.ruanyifeng.com/blog/2016/04/cors.html)
