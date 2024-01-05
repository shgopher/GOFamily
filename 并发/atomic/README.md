<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-01-05 22:33:23
 * @FilePath: /GOFamily/并发/atomic/README.md
 * @Description: 
 * 
 * Copyright (c) 2024 by shgopher, All Rights Reserved. 
-->
# atomic
原子包是比互斥锁更底层的包，如果在简单的场景下，使用 sync.Mutex 可能会比较复杂，并且耗费资源，那么使用更加底层的 atomic 就更加划算了

所谓原子操作，就是当某 goroutine 去执行原子操作时，其它 goroutine 只能看着，这个操作要么成功，要么失败，不会有第三个状态

原子包操作对象的时候，都是操作的地址，所以谨记不要使用值操作而是要指针操作

## 介绍一下 atomic 的内容

- Add：例如 `func AddInt32(addr *int32,delta int32)(new int32)` 给第一个参数地址指向的数据值增加一个 delta 并返回新的数据
- CompareAndSwap：例如 `func CompareAndSwapInt32(addr *int32,old,new int32)(swapped bool)` 比较 addr 指向的数据是否等于 old，如果不等于返回 false，如果等于就将此地址的值切换为 new 值，并且返回 true
- Load：例如 `func LoadInt32(addr *int32)(val int32)` 读取 addr 指向的值并返回
- Store：例如 `func StoreInt32(addr *int32,val int32)` 将 val 值写入到 addr 指向的内存空间中
- Swap：例如 `func SwapInt32(addr *int32,new int32)(old int32)` 将 addr 指向的值切换为 new 值，并返回旧值
- Value：`type Value` `func(*Value) Load` `func(*Value) Store` 原子的存取数据


目前 atomic 还没有部署泛型，所以里面到处充斥者 `LoadInt32` `LoadInt64` 这种类型的函数，以后等泛型部署到原子包后就不会这么繁琐了

## 基于原子库的第三方扩展
- uber-go/atomic 定义扩展了几种常见类型的原子操作，例如 bool error string 等

  ```go
  var atom atomic.Uint32
  atom.Store(42)
  atom.Sub(2)
  atom.CompareAndSwap(40, 11)
  ```
  看起来的确比官方提供的原子包更加简洁一些

## issues
### 对一个地址的赋值是原子操作吗？
如果对于单核处理器的机器来说，地址的赋值是原子操作

在现在的系统中，write 的地址基本上都是对齐的

对齐地址的写，不会导致其他人看到只写了一半的数据，因为它通过一个指令就可以实现对地址的操作，如果地址不是对齐的话，那么，处理器就需要分成两个指令去处理，如果执行了一个指令，其它人就会看到更新了一半的错误的数据，这被称做撕裂写

所以，你可以认为赋值操作是一个原子操作

但是，对于现代的多处理多核的系统来说，由于 cache、指令重排，可见性等问题，我们
对原子操作的意义有了更多的追求。

***在多核系统中，一个核对地址的值的更改，在更新到主内存中之前，是在多级缓存中存放的。这时，多个核看到的数据可能是不一样的***

多处理器多核心系统为了处理这类问题，使用了一种叫做内存屏障 (memory fence 或
memory barrier) 的方式。一个写内存屏障会告诉处理器，必须要等到它管道中的未完成
的操作 (特别是写操作) 都被刷新到内存中，再进行操作。

atomic 包提供的方法会提供了一些的功能，不仅仅可以保证赋值的数据完整性，还能保证数据的可见性，一旦一个核更新了该地址的值，其它处理器总是能读取到它的最新值。

atomic 包主要利用了以下几点技术：

- 编译器插入内存屏障 (Memory Barrier) 指令
编译器会在 atomic 操作前后插入内存屏障指令，来限制 CPU 的乱序执行，保证在该操作前的读写操作都完成，之后的读写都待其完成后再执行。

- 硬件支持的原子 CPU 指令
如 X86 的 LOCK 指令可以将某些指令变为原子指令。atomic 会利用 CPU 提供的这些原子指令实现加锁。

- 缓存一致性硬件协议
如 Intel 的 MESI 协议可以在多核间保证缓存的一致性。atomic 利用缓存一致性，使得多个核心缓存中的数据版本是一致的。

- 核心间互斥
atomic 中的原子操作会在多核间加锁，保证同时只有一个核心可以操作共享变量。

需要注意的是，因为需要处理器之间保证数据的一致性，atomic 的操作是会降低性能的。

综上所述，***对于单核机器来说，普通的地址赋值就是原子操作，但是对于多核机器来说，不属于原子操作， 原子包去进行的赋值一定是原子操作***
## 参考资料
- go.dev
- https://time.geekbang.org/column/intro/100061801
- 《go 进阶训练营》
- 《go 语言精进之路》