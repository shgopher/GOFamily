# go语言常见的关键字

本文章会包含以下内容：
- for和range
- select
- defer
- panic和recover
- make和new

## for 和 range

### slice
```go
func main() {
	a:= []int{1,2,3}
	var ma []*int
	for _,v := range a{
		v+=1
		ma = append(ma,&v)
	}
	fmt.Println(a)
	for i := 0;i < 3;i++ {
		fmt.Println(*ma[i])
	}
}
```
```go
[1,2,3]

4,4,4
```
前面a为什么v+1了输出的还是1，2，3呢？

原因是因为这里的v只是切片数据的值复制，所以v的改变不会改变原切片的内容，如果要改变可以更改为 

```go
for i:= 0;i<len(a);i++ {
    a[i]++
}
```

而遇到这种同时遍历索引和元素的 range 循环时，Go 语言会额外创建一个新的 v2 变量存储切片中的元素，循环中使用的这个变量 v2 会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝。

因为在循环中获取返回变量的地址都完全相同，所以会发生神奇的指针一节中的现象。因此当我们想要访问数组中元素所在的地址时，不应该直接获取 range 返回的变量地址 &v2，而应该使用 &a[index] 这种形式。

```go
for k,v := range a{
		v+=1
		ma = append(ma,&a[k])
	}
```

那么这种方式为啥就是不一样了呢？

因为经过了复制，也就是说这里的&a[k]每次循环都不是一个&a[k],每次的a[k]都是一次运算，意思就是已经取到值了。


对于所有的 range 循环，Go 语言都会在编译期将原切片或者数组赋值给一个新变量 ha，在赋值的过程中就发生了拷贝。所以这里的range过程中的length就是老的length，新的切片跟这个循环次数一点关系都没有。

### map

在map中的遍历是随机的，原理就是寻找桶的时候是随机的，然后我们从某一个桶开始找，找到非零桶的时候把链表里的数据查询出来，一般这个时候是查询出来的数据的地址，当然在扩容期间，就是找到真的k-v，然后进而找到所有的正常桶，然后再寻找溢出桶。这个时候遍历就结束了。
### string
当遍历字符串的时候，跟遍历slice差不多，只是会把byte（uint8）类型转化为rune（int32）
### chan

range 可以变量chan中的数据，不过如果chan未关闭，但是没有数据了，range会Panic，所以如果没有数据了，chan需要关闭才行。如果不关闭就会阻塞，也因为不关闭，就意味着只有写，才能读，但是不写了，读也读不了，如果读一个关闭的chan，那么只会得到零值而已。

```go
func main() {
	a := make(chan int)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			a <- i
		}(i)

	}
	go func() {
		wg.Wait()
		close(a)
	}()
	//这里，当range读的通道被close后，就会【自动！】跳出循环。
	for v := range a {
		fmt.Println(v)
	}

}

```

## select
select跟io复用中的select比较相似，在select中的case中，只允许出现chan的读或者是写，当然还有default，形如
```go
ffunc main() {
	a := make(chan  int)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			a <- i

		}(i)

	}
	go func() {
		wg.Wait()
		close(a)
	}()
	L:
	for  {
		select {
            // 使用这种形式判断通道是否close
		case v,ok := <- a:
			if ok  {
				fmt.Println(v)
			}else {
				break L
			}
		}
	}

}

```
但是请注意，如果两个case都满足，那么执行的时候会随机执行一个。并不会同时执行。随机的目的就是避免饥饿。

select存在这么几种机制
1. 没有case
    
    这种情况下，select是永远的阻塞状态。
2. 只有一个case
    
    当case中的chan是nil时，陷入阻塞状态
3. 只有一个case和default

4. 含有多个case


## defer
defer执行的位置是在return之前进行执行，并且执行的顺序是栈的顺序。并且defer的后面只能跟函数。

通常来说我们使用defer都会为了关闭某个操作，比如response.body,就需要我们手动关闭，比如数据库的连接，这种操作下，使用defer再合适不过了，并且再在最新的go版本中，defer的性能损耗大大降低 6ns左右，可以说忽略不计了。

```go
func createPost(db *gorm.DB) error {
    tx := db.Begin()
    defer tx.Rollback()
    
    if err := tx.Create(&Post{Author: "Draveness"}).Error; err != nil {
        return err
    }
    
    return tx.Commit().Error
}
```
这里要注意一件事，tx.Commit().Error先执行，然后这个时候发现有defer，不会立即执行return，而是执行defer，执行完毕后，再把tx.Commit().Error的结果返回出来。也就是说return其实不是一次执行完毕的事情。

这里要注意两个非常重要的事情
- defer执行的顺序
- defer预计算的顺序

我们直到defer后面跟函数，然后后面执行的顺序是栈，也就是defer栈，按照先入后出（后执行）的顺序进行执行，但是有一点要注意，虽然defer是在整个大部队后面执行的，并且是栈顺序，然而，它的预操作是按照整个函数的顺序执行的,defer 关键字会立刻拷贝函数中引用的外部参数.

```go
func test(i int) int{
    // 这里打印的就是i 而不是i++后的结果，就是因为defer 把外部参数复制了。
    defer fmt.Println(i)
    i++
    return i
}
```
要改变这种方式也很简单，不要外部参数即可。

```go
func test(i int) int{
    
    defer func(){
        fmt.Println(i)
    }()
    i++
    return i
}
```

我们再看两个例子

输入1

```go
func test1(i int)(ii int){
	defer func() {
		ii++
	}()
	ii = 1+i
	return 
}
// output 3
-----
func test2(i int)int{
	defer func() {
		i++
	}()
	i++
	return i  
}
// output:2
```

这就是return的区别了，前面那种执行的过程是，return的结果在defer后面执行，后面那种是return的结果在defer前面就复制了，然后再执行defer，然后return结果，这个结果不受到defer的影响。

**实际执行的是等于**

```go
func test2(i int)int{
	defer func() {
		i++
	}()
	i++
    //
    x：= i
    //
	return x // x的值不受到i的影响
}

```

再看一下让人更震惊的：

```go
func test1(i int)(ii int){
	defer func() {
		ii++
	}()
	ii = 1+i
	return 1 // 这里等于 ii = 1
}
// 结果不是1，而是2 ，因为这里的return后面的值还是ii，ii被重置为1，然后再++ ，然后再返回。
```

```go
func test1(i int)int{
	defer func() {
		i++
		fmt.Println(i)
	}()
	i++
	return 5
}

// 返回5，打印3 ，证明return还是等到defer执行完再返回，不过return后面的变量是新变量了。
```


也就是说，return后面的结果就是ii，这种情况就是ii被1+i重置为了1，然后还是要再执行defer，得到的最终的值，才是返回值。

defer很容易出问题，所以这种技巧尽量不要采用。defer老老实实的使用的场景就是关闭某某。

我们来总结以下：

- 如果是带值返回的那种，返回的就是那个显式的变量
- 如果不带值的那种，返回的其实是return后面那个结果的复制值。所以这个之前的参数再怎么在defer里反应，即便return等到defer执行完再返回，也不会改变这个复制值。
## panic和recover

panic 就是恐慌，使得程序在这个goroutine中立刻停下，当然整个程序也会停止，并且其它的在这个goroutine后面的任何代码都不会执行了(不过Panic不会影响defer函数的执行，但是只是这个goroutine中的defer，其它goroutine中如果没有执行到就不会执行了)recover就是恢复或者忽略Panic，并且recover函数，必须在defer函数中使用。而且在一个goroutine中的潜逃的defer也不会收到影响

```go
func main() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
// defer 都可以正常执行
```

正确的recover的执行

```go
func main(){
    defer func(){
        if err := recover();err != nil {
            fmt.Println("panic:",err)
        }
    }()
    get()
}

func get(){
    panic("get is panic.")
}
```
注意一下

```go
func main() {
	defer recover()
	get()
}
```
虽然recover也是函数，但是这种方式，recover不起作用，提示的错误是 不应该在defer后面直接跟recover，应该使用一个匿名函数。
## make和new
- make 的作用是初始化内置的数据结构，也就是切片、哈希表和 Channel，值得注意的是切片可以给定len和cap，后面的哈希表和chan只有len，并且slice可以省略cap
- new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针
