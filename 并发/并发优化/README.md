<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2024-02-26 00:23:06
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-02-27 23:21:55
 * @FilePath: /GOFamily/并发/并发优化/README.md
 * @Description: 
 * 
 * Copyright (c) 2024 by shgopher, All Rights Reserved. 
-->
# go 语言实战项目并发优化
上文我们已经讲解了 goroutine，channel，并发原语，atomic，context，以及 go 的内存模型，内容还是比较多的，我们需要通过实战的项目的优化演进过程来更好的理解并发的最佳实践
## 能不并发就不并发
并发是一个双刃剑，一方面它可以加速程序，另一方面它也增加了程序的复杂度，所以需要在并发和不并发之间做取舍，如果你发现你的程序在不使用并发的时候也能满足你的需求，答应我，不要使用并发，***累活脏活自己干，不要委派给另一个 goroutine 去做所谓的并发工作。***


```go
// bad
func main() {
  http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request,)) {
    fmt.Fprintln(w,"hello wrold!")
  }
  go func(){
    if err := http.ListenAndServe(":8080",nil);err != nil {
      log.Fatal(err)
    }
  }()

  select{}
}
```
在这段代码中，委派一个 goroutine 去启动一个监听服务，又使用 select {} 去阻塞主 goroutine 的运行

确实，这可以满足需求

但是，委派一个后台 goroutine 去执行监听服务没有带来任何有利的收益，反而增加了代码的复杂度，所以我们应该取消委派，主 goroutine 去执行监听即可

```go
// better

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request,)) {
    fmt.Fprintln(w,"hello wrold!")
  }
  
  if err := http.ListenAndServe(":8080",nil);err != nil {
    log.Fatal(err)
  }
}
```


## 优先使用 channel + context 的方法去优雅关闭
## 使用方去决定是否并发
