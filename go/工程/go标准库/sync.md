# sync
## sync.Map
### sync.Map 的底层数据

sync.Map 适合在读多写少的场景下使用，sync.Map 的核心思想是读写分离，以及用空间换时间。

有两个数据结构：

`dirty`数据结构

```go
type Map struct {
  // 互斥锁
    mu Mutex
  // 原子操作
    read atomic.Value
    // dirty数据 
    dirty map[interface{}]*entry
    // 标记tag 表示数据每次从diry中往read中迁移一个就+1
    misses int
}
```

`read` 数据结构

```go
type readOnly struct {
    m       map[interface{}]*entry
    // 标记：false 证明数据不只是read中有，dirty也有
    amended bool 
```

读数据是优先从read中读（注意，read是只读的数据结构，所以不加锁），如果read不包含这个数据，会从dirty中读取（注意，从dirty查找的时候会加锁），并misses+1，直到 misses 大于dirty，开始迁移数据，数据一次性从dirty传到read中。

因为read没有锁，所以才说，sync.Map 读的时候效率高。

当新增的时候，从read和dirty中查找是否用数据，在read有数据且没有被标记为清除就无锁覆盖原来的值。这个时候开始加锁了，如果在read中虽然有数据，但是被标记为清除了，并且dirty没有数据，就加入到diry中，如果这个时候read有数据，虽然被被标记为删除了，dirty也有数据，那么直接原子更新数据即可。如果数据不在read，在dirty中，直接在dirty中更新即可，如果数据不在read中也不再dirty中，如果这个时候read.amended 是false（意思就是dirty空了），**那么就必须将数据一点一点从read中复制到dirty中**，如果是true就不用复制了，在完成这个工序后，再把数据放置在dirty中。

删除的时候，就先从read中查找，有了就原子性质的将其标记为nil，并不是真的删，没有就加锁从dirty中找，有了就删。这个时候是真删，如果都没有就返回。

在写多读少的时候，涉及到read中无数据就会频繁的加锁去dirty中查找、read和dirty交换等开销，比常规的map加锁性能更差。因为普通的只是加锁，sync.Map 不仅仅是加锁，还得复制数据，这花销就更大了。

总结：**高并发，读多写少 可以使用sync.Map ,如果 读少写多，不要使用sync.Map,比常规锁+map更慢**

## 参考资料
- https://zhuanlan.zhihu.com/p/353440086