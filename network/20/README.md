# cookie Session

cookie和session是一般的保持会话的一种方法

具体实用的时候是这样的 在本次保存cookie，cookie中是key-value结构，一般一个value中保存的是一个session的sessionID
然后这个sessionID对应了一个Session 当然也是KY  实用一个散列表将每一个ID跟Session保存在散列表中。然后每一次请求都
带着cookie然后带着这个ID，session临时对话就保存在服务器的内存或者内存数据库中（目的是为了快）然后每次使用sessionID来标记
一个用户，使用这个id来查找出session中的很多信息，就不需要每次再去请求数据库了。 这就是会话机制。

cookie可以使用Web storage API（本地存储和会话存储）或 IndexedDB。代替

cookie可是大概的分为两种一种是可持续的cookie一种是会话cookie也就是说关了浏览器就gg了。
cookie有很多的参数比如：

- HTTPOnly 是否可以被js调用
- secure 只能通过https传输。
- maxage 最大失效时间段

等。
