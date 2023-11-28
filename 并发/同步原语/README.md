<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-11-28 20:58:22
 * @FilePath: /GOFamily/并发/同步原语/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# 同步原语
同步原语是 go 提供的相对底层的同步机制，它更加灵活，但是同时也更加复杂，如果可能的话，我们应该尽量使用 csp 的并发模型，使用 channel 去代替并发原语。
## sync.Mutex
下面介绍的众多并发原语，甚至下一章的 channel，都使用了这个核心内容，它就是 `sync.Mutex`，它是所有同步原语包含 channel 的底层核心。
### sync.Mutex 基础操作
我们现看一下它的基础使用功能：
```go
func age(){
  var mu sync.Mutex
  for i := 0; i <10; i++ {
    go func(){
      mu.Lock()
      defer mu.Unlock()
    }()
  }
}

// 跨goroutine 加解锁
func age1() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	go func() {
		defer mu.Unlock()
		time.Sleep(1000)
		fmt.Println("hi there")
	}()
	mu.Lock()
}
```
sync.Mutex 即为互斥锁，规则是：

- 锁的加锁和解锁可以跨 goroutine 使用，比如 a goroutine 加锁，在 b goroutine 将锁解开。
- 只有现解锁才能继续上锁，happens-before 就是：***n 次解锁一定 happens-before n+1 次加锁***
### sync.Mutex 互斥锁的实现原理
go 语言互斥锁的实现非常简单，只有这一个结构体就是核心：
```go
type Mutex struct {
  state int32
  sema uint32
}
```
## sync.RWMutex
## sync.Locker
## sync.WaitGroup
## sync.Cond
## sync.Once
## 讨论 map 在多线程中的场景
## Pool
## errgroup
## semaphore
## singleflight
## syncmap
## 参考资料
- https://mp.weixin.qq.com/s/iPpWd8vjyaN2sJFwxzN9Bg
- https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/
- https://time.geekbang.org/column/intro/100061801
- 《go 语言精进之路》