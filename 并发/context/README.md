<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2022-11-17 20:40:42
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-01-04 19:06:48
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

## withValue
WithValue 基于 parent Context 生成一个新的 Context，保存了一个 key-value 键值
对。它常常用来传递上下文。

context 在查询 key 值的时候还支持链式查找，如果没有发现数据就往 parent context 中查询

```go
ctx = context.TODO()
ctx = context.WithValue(ctx, "key1", "0001")
ctx = context.WithValue(ctx, "key2", "0001")
ctx = context.WithValue(ctx, "key3", "0001")
ctx = context.WithValue(ctx, "key4", "0004")
fmt.Println(ctx.Value("key1"))
```

### WithCancel
withCancel 返回父 context 中的 ctx 实例副本，它相当于父 context 的子 context，并且在父 context 被取消时，子 context 也会被取消。

```go
func withCancel(parent Context) *cancelCtx {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	c := &cancelCtx{}
  // 向上寻找
	c.propagateCancel(parent, c)
	return c
}
```
propagateCancel 部分代码：
```go
func (c *cancelCtx) propagateCancel(parent Context, child canceler) {
	c.Context = parent

	done := parent.Done()
	if done == nil {
		return // parent is never canceled
	}

	select {
	case <-done:
		// 如果父done了，那么子ctx一定也会出发cancel
		child.cancel(false, parent.Err(), Cause(parent))
		return
	default:
	}

	if p, ok := parentCancelCtx(parent); ok {
		// parent is a *cancelCtx, or derives from one.
		p.mu.Lock()
		if p.err != nil {
			// 如果父发生了cancel，那么子ctx也要触发cancel
			child.cancel(false, p.err, p.cause)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
      // 将子ctx添加到父ctx中
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
		return
	}
  ...
  go func() {
		select {
      // 父 done被触发，那么子ctx就会被触发cancel操作
		case <- parent.Done():
			child.cancel(false, parent.Err(), Cause(parent))
		case <-child.Done():
		}
	}()
}

type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}
```
propagateCancel 将 c 向上传播，顺着 parent 的路径一直向上查找，直到找到 parentCancelCtx，如果不为空，就把自己加入到这个 parentCancelCtx 的 children 切片中，然后就可以在父 ctx 取消的时候，通知自己也被取消

当这个 cancelCtx 的 cancel 函数被调用的时候，parent 的 Done 被 close 的时候，或者父 ctx 触发了 cancel 的时候，这个子 ctx 会被触发 cancel 动作

cancel 是向下传递的，如果一个 WithCancel 生成的 Context 被 cancel 时，如果它的子 Context (也有可能是孙，或者更低)，就会被 cancel，但是不会向上传递。parent Context 不会因为子 Context 被 cancel 而 cancel。

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

综上所述，done 这个 channel 被 close 有三个原因：
- 截止时间到了
- cancel 函数被调用了
- parent context 的 done close 了，然后子 ctx 也要触发 cancel 方法
- parent context cancel 了触发子 ctx cancel 方法

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
Go 语言中 context 实现并发安全的主要手段是通过原子操作和 Mutex 来保证状态的原子性

根据底层代码可知，当不同的 goroutine 获取 ctx 的时候，每次操作都会加上互斥锁，来保证数据的非竞争性，线程的安全，例如

```go
if p, ok := parentCancelCtx(parent); ok {
		// parent is a *cancelCtx, or derives from one.
		p.mu.Lock()
		if p.err != nil {
			// 如果父发生了cancel，那么子ctx也要触发cancel
			child.cancel(false, p.err, p.cause)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
      // 将子ctx添加到父ctx中
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
		return
	}
```
