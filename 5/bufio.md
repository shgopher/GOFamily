# bufio

bufio == buffer  io 意思就是封装了io，是拥有缓存的io操作

bufio跟io一样，主要是实现了三种方法
- read：读
- write：写
- scan：扫描

`func NewReader(rd io.Reader) *Reader` 实现了一个带有缓存的io.reader接口，这里返回的是bufio.Reader ,这个struct实现了io.Reader接口。

`func NewWriter(w io.Writer) *Writer`实现了一个带有缓存的io.Writer接口，其中bufio.Writer是一个实现了io.Writer的struct

`func NewScanner(r io.Reader) *Scanner`提供了一些扫描相关的method