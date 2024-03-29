<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2024-02-26 00:23:06
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-03-02 17:50:06
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
核心就是把众多启动的 goroutine 改成 worker pool + context 的模型

在 channel 中我们讲过 worker pool 的实现，为什么一定要优先使用 worker pool？因为很多时候你启动了很多的 goroutine，不知不觉就会造成混乱，以及 goroutine 的泄露问题，我们使用 worker pool 的方式，可以很好的控制 goroutine 的并发数量，以及优雅的关闭 goroutine

看一个优秀的例子：
```go

type Tracker struct {
	ch   chan string // 作为工作池子让worker消费
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch:   make(chan string, 10),
		stop: make(chan struct{}, 2),
	}
}
func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (t *Tracker) Run() {
	for data := range t.ch {
		// 模拟消费
		time.Sleep(time.Second * 5)
		fmt.Println(data)
	}
	// run 数据结束之后就可以发送信号了
	t.stop <- struct{}{}
}

// shutDown 通过关闭ch 让 run中的range结束，进而再发送stop信号，让shutdown函数退出
func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():  
	}
}
```
```go
func main() {
	tr := NewTracker()
	// 开启两个消费者
	go tr.Run()
	go tr.Run()
	
  tr.Event(context.Background(), "test")
	tr.Event(context.Background(), "test2")
  tr.Event(context.Background(), "test3")
	tr.Event(context.Background(), "test4")
  tr.Event(context.Background(), "tes5")
	tr.Event(context.Background(), "test6")

	time.Sleep(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}
```
## 使用方去决定是否并发
```go
func ListDirectory(dir string) chan string
```
这是一个返回目录下所有文件路径的函数，它返回的是一个 channel，很明显，这跟题目所说的让使用方去决定是否并发是相违背的

因为在这个函数体内部，已经完成了一个 goroutine 的创建，但是作为使用方，我们并不能说一定要使用并发，而且你也无法控制发送的停止，假设你在这个函数内部使用一个 close 作为关闭的信号，也是有问题的，比如我想恢复读取，那么你已经 close 了，该如何继续开启呢？

还有一个问题，这个函数返回的是一个 channel，但是并没有返回 error，我们如果去判断 channel 的状态，并不能分辨是读取完毕 channel 被 close 还是出现 error，然后 channel 被关闭了，这就是二象性的问题，你不能设置二象性这种状态的函数，就跟刚才说的那样，你无法获悉真实的最精准的状态。

也就是说你的发送过程不够透明，调用方无法决定暂停，恢复读取这些行为

再提一个需求，我只需要读取的数据中的某些符合规定的路径，那么你如果作为一个黑箱的函数，完全无法做到定制化对不对

那么为什么不能设计成一个普通函数，然后在调用方再决定是否并发这个行为模式呢？

我们看一下 go 语言标准库中的用法
```go
filepath.Walk(root string,fn filepath.WalkFunc) error
```
很明显这是一个实现了功能的普通函数，那么让我们分别看看它的普通模式和并发模式
```go
// 普通模式

err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name()) // 可以看到我们返回的err是不同的err
			return filepath.SkipDir
		}
    // 定制需求
    if path != "xxx"{
      fmt.Printf("visited file or dir: %q\n", path)
    }
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
```
那么再让我们看一下并发模式
```go
// 并发模式
go func() {
		defer close(value)
		err <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// if the file is noe regular, it mean the file is done,you should return
			if !info.Mode().IsRegular() {
				return nil
			}
			value <- path
			return nil
		})
	}()
```
解释一下这个 Walk 的函数本质，它传入一个路径和一个函数，这个函数拥有的参数是一个要输出的路径，一个文件信息和一个错误类型，返回一个错误类型，参数函数被执行的时候 walk 底层会将真实 path 值作为实际参数传入这个函数里，所以我们可以在 path 中，指定一个 channel 去作为流的方式导出数据。

```go
// walk 的底层源码
for _, name := range names {
		filename := Join(path, name)
		fileInfo, err := lstat(filename)
		if err != nil {

      // 这里将读取到的path值传入了 walkFn 函数里，来完成实际参数的赋值行为
			if err := walkFn(filename, fileInfo, err); err != nil && err != SkipDir {
				return err
			}
		} else {
			err = walk(filename, fileInfo, walkFn)
			if err != nil {
				if !fileInfo.IsDir() || err != SkipDir {
					return err
				}
			}
		}
}
```

可以看到这个函数的实现完成了基本的功能，既可以并发，又可以不并发
## 必须让发送方决定 channel 的关闭
如果不能让发送方决定 channel 的关闭，而是不控制或者是接收方去控制，那么数据的丢失就不可避免了

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    ch := make(chan int)
    wg := sync.WaitGroup{}

    // 发送者
    wg.Add(1) 
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch) // 发送方关闭channel
    }()

    // 接收者
    wg.Add(1)
    go func() {
        defer wg.Done() 
        for {
            if data, ok := <-ch; ok {
                fmt.Println(data)
            } else {
                break // channel已关闭
            }
        }
    }()
    
    wg.Wait() 
}
```
那么让我看看一下几个反例

```go
// 反例1:接收方关闭channel
func main() {
  ch := make(chan int)

  go func() {
    for i := 0; i < 5; i++ {
      ch <- i 
    }
  }()

  for {
    fmt.Println(<-ch)
      // 当你接收方去关闭channel的时候
      // 发送方如果不知道的话，向一个close的channel发送数据会panic的
      if xx {
         close(ch) // 接收方关闭channel
         break
      }
  }
}
```
```go
// 反例2:没有关闭channel
func main() {
  ch := make(chan int)
  
  go func() {
    for i := 0; i < 5; i++ {
      ch <- i
    }
  }()

  for {
    // 这里 ok 永远都不会false，所以 这里会死循环
    if data, ok := <-ch; ok {
      fmt.Println(data)
    } else {
      break
    }
  } 
}
```

