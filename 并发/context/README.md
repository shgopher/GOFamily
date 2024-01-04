<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-01-04 17:32:45
 * @FilePath: /GOFamily/并发/context/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# context
context 应该被翻译为上下文，但是 go 语言中的 context 更多的作用是**取消子 goroutine**，传递信息这个上下文本来的含义，在 go 的 context 反而不是重点

context 本身是一个接口类型，它拥有四个方法：
- Deadline()(deadline time.Time,ok bool)：返回一个代表此上下文完成工作的时间，如果没有截止时间，返回 false
- Done() <-chan struct {}：代表了完结
- Err() error：如果 context 已经 done 了返回一个 error 错误，错误值分别为：Canceled，DeadlineExceeded，如果没有 done，error 返回 nil
- Value(key any) any：context 中存储的键值对

context 作为上下文，需要一个最顶层的 context 接口类型，你可以使用 `context.Background()` 或者 `context.TODO()` 去充当这个顶端，这两者是一个意思，用哪个都可以，这是 go 提供的已经实现了 context 接口类型的对象，它的底层是一个结构体

context 有一些编程范式：
- 将 cotext 设置为参数的第一个，例如 `func age(ctx context.Context,a string)`
- 不要使用 nil 作为上下文参数，如果想要空的顶端上下文，使用 `context.Background`，虽然 background 底层实现接口的时候，也是内容为空，但是它的确是实现了接口
- context 只能作为函数的临时传递对象，不能持久化它，使用数据库保存，等持久化方式都是不可取的
- 使用 withValue 方法的时候，key 值不要使用 string，如果起冲突，使用自建的类型，例如 `type A struct{}`
- 尽量不要定义输出的 key 值

## 使用 context
标准库中提供了多个 context 接口类型实例：
- WithCancel
- WithCancelCause
- WithDeadline
- WithDeadlineCause
- WithTimeout
- WithTimeoutCause

其中，带有 Cause 的函数跟不带的函数基本意思相同，但是多了一个 cause 的内容，它是指的是取消的原因

### WithCancel
当你的任务执行完毕了，执行返回值中的 cancel 函数，cancel 函数会 close 这个 done channel，这样就可以释放资源

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
			}
		}
	}
  time.Sleep(time.Second * 10)
}
```
WithCancel 返回 parent context 的一个副本，它自然就是子 context，当父 context 被 cancel 的时候，子 context 也会被 cancel

### withTimeout withDeadline

这两个只是添加了到期时间，一个是超时时间，一个是截止时间，一旦超过时间后，自动 close 这个 done 这个 channel

done 这个 channel 被 close 有三个原因：
- 截止时间到了
- cancel 函数被调用了
- parent context 直接控制子 context 的 done close

关于第三条，解释一下：
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个父context,设置deadline为3秒
	parentCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 创建一个子context,deadline继承父context
	childCtx, cancel := context.WithCancel(parentCtx)
	defer cancel() // 注意这里需要调用cancel

	go doWork(childCtx)

	time.Sleep(14 * time.Second)
}

func doWork(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("over")
			return
		default:
			fmt.Println("11")
		}
	}
}

```

## 带有 cause 的函数

我们看一个例子
```go
package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	var myError = errors.New("myError")
	ctx, cancel := context.WithCancelCause(context.TODO())
	cancel(myError)
	ctx.Err()                       
	fmt.Println(context.Cause(ctx)) // returns myError
}
```
但我们调用 cancel 函数的时候，内部参数是一个 error 类型，调用 context.Cause(ctx) 返回的就是它的取消原因，那么这里的话就是 MyError



## issues
### contex.Contex 如何实现并发安全的？
