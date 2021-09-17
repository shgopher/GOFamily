# go语言代码的性能优化
概括：
- 减少堆内存的使用
- 进行内存对齐
- 减少反射的使用
- strconv.Iota 效率比 fmt.Sprintf` 高
- 减少string转化为[]byte
- slice append的过程尽量人工扩容
- 多用sync.Pool复用对象
- 用bufio代替io包
- 使用protobuf来代替json
- map中key最好用整型
---
1. 尽量减少堆内存的使用，尽量减少大片内存的申请【减少使用页堆】

2. 要进行内存对齐

    内存字节对齐主要是用在结构体中，一个较大的结构体，如果能内存对齐的话，就会节省不少的内存使用空间，

    这里要提一下，计算机中最小的单位是字节，一个字节等于8位，但是，1位仍然是一个字节，因为最小的单位是字节，不存在0.125个字节这种说法。
    
    我们使用`unsafe.Sizeof()`来获取类型的大小

    所以说bool类型，虽然就一位，但是它是1个字节。int类型在64位的计算机中，它就是`64 / 8 = 8个字节`,`string` 是16个字节，主要包括了，一个指针类型8个字节（`unptr`，其实这个家伙就是`uint8`）一个`byte`类型 8个字节（还是`uint8`）所以一共是16个字节，`slice`类型，一个指针和一个一个`length`【8个字节的`uint8`】一个`cap`【8个字节】所以`slice`类型是24个字节【这仅仅是`slice`类型，还没有装数据呢】，`map`是8个字节，就一个指针【说明底层结构体中没有len和cap】

    这里我们要知道一个不占内存的数据，就是`struct{}` 它是0字节，所以说一个空的结构体是没有大小的，小tips，go里面没有set，可以用空结构体作为map的value，就可以充当一个`set`类型了。

    下面演示一下内存对齐：

    这是没有对齐的时候
    ```go
    type a struct {
        a int8
        b int32
        c int16
    }
    ```
    下面是对齐的时候
    ```go
    type a struct {
        a int8
        c int16
        b int32
    }
    ```

3. **尽量减少反射的使用**，例如不使用标准库的json包，使用fastjson等包，不使用反射来处理多个chan，使用递归的方法来处理chan，递归如果可以尾递归或者不用递归，更好。

4. 数字转化为字符串的时候，使用 `strconv.Iota` 效率比 `fmt.Sprintf` 高
    
    ```go
    func BenchmarkT(b *testing.B) {
        for i := 0; i < b.N; i++ {
            FmtT()
        }
    }
    func BenchmarkT2(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Tota()
        }
    }
    func FmtT(){
        a := 1234567
        fmt.Sprintf("%d",a)
    }
    func Tota(){
        a:= 1234567
        strconv.Itoa(a)
    }
    ```

    ```bash
    goos: darwin
    goarch: arm64
    pkg: a
    BenchmarkT
    BenchmarkT-8    	20879263	        54.95 ns/op
    BenchmarkT2
    BenchmarkT2-8   	64085587	        18.53 ns/op
    PASS
    ```
    可以看出，使用strconv.Itoa的效率至少要比fmt.Sprintf高了一倍以上。
5. **尽量减少从string转化为[]byte的过程**，因为string是只读的，改变string只能先转化为一个slice ，因为string是只读，所以说它转化为[]byte必须复制数据，以及，当字符串的量过大的时候，分配的slice就很有可能分配到堆上，所以说字符串的添加或者是各种处理都是一件非常耗费资源的事情，要尽量的减少字符串的操作。如果真的要拼接字符串，**不要使用 `+`,尽量使用bytes.Buffer**  

    ```go
    var b bytes.Buffer // A Buffer needs no initialization.
    b.WriteString("a")
    ```
    也就是说，我们在生成string的过程中，要保持[]byte的状态，中间的各种处理都是[]byte，最后转化为string(x) 即可

    在go的strings包中，处理strings的过程就是使用了bytes.buffer
    ```go
    var b Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
    ```
6. slice的append过程，如果cap不够，尽量自己计算好，人工分配新的cap，也不要让系统分配，因为系统很有可能直接扩展到两倍的容量，会造成内存的浪费。

7. 不要在热代码中进行内存分配，这样会让 gc 很高。要使用 sync.Pool 这种 `池` 类技术来复用。

8. 尽量使用bufio包的io操作来替代io包的无buffer的操作，因为io的过程很慢

9. 对于在 for 循环中的的固定正则表达式，要使用 regexp.Compile 编译正则表达式

10. 如果要获得更好的编码速度，就不要使用go的json包，因为这个包内使用了反射，可以使用protonbuf来代替，protobuf是google出品，go完美支持。`google.golang.org/protobuf` 引入此包即可。

11. map中的key尽管可以使用任何的可以[比较](../1/int.md#关于go里面的比较问题)的数据作为key，但是平心而论，使用整型例如int类型是最佳选择，因为比较的时候更容易，例如string也好，其他类型也好都没有整型快
