# Inode 和Block

之前说过Inode存放的是block的编号block是存放的内容有点像 Inode是文件头block是实际的文件内容这种感觉。

一个block只能被一个文件所使用，没有被使用到的部分就废了，所以要选择一个合适自己的block大小。

inode 具有以下特点：

- 每个 inode 大小均固定为 128 bytes。
- 每个文件都**仅**会占用一个 inode。

每个inode都有很多的block，为了管理这些block，inode使用了 “间接、双间接、三间接引用” 这种模式。

inode 具体包含以下信息：

```go
- 权限 (read/write/excute)；
- 拥有者与群组 (owner/group)；
- 容量；
- 建立或状态改变的时间 (ctime)；
- 最近一次的读取时间 (atime)；
- 最近修改的时间 (mtime)；
- 定义文件特性的旗标 (flag)，如 SetUID...；
- 该文件真正内容的指向 (pointer)。
```
