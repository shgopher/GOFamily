# 子连接

```bash

SELECT username FROM db
WHERE year >
(SELECT year FROM db WHERE stuentName=?)
```

()这就是子连接，然后自连接的结果给year > xxx
