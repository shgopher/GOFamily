# atomic
`import "sync/atomic"`

此包为原子操作，意思就是同一时间，此包的操作，在cpu层面就是单一存在的，不存在各线程的竞争，是原子性质的操作。在cpu来看属于绝对的单位时间内唯一的操作。

鉴于目前go还没有实现泛型，所以实现的方法稍微优点冗余，实际上，atomic包大概就是这么几种方法：

- load
    ```go
    func LoadInt32(addr *int32) (val int32)`
    ```
    仅列出一个例子，其它的类似，这个函数就是原子的获取addr指针上的值
    
- store 

    ```go
    func StoreInt32(addr *int32, val int32)
    ```
    将val的值保存到addr这个指针类型里
- add 
    ```go
    func AddInt32(addr *int32, delta int32) (new int32)
    ```
    原子性的将delta的值，添加到addr的值上，并返回一个新的值
- swap 
    ```go
    func SwapInt32(addr *int32, new int32) (old int32)
    ```
    SwapInt32原子地将新存储到*addr中，并返回之前的 old *addr值。
- compare and swap
    ```go
    func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
    ```
    CompareAndSwap原子性的比较*addr和old，如果相同则将new赋值给*addr并返回一个bool值

