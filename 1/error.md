# go错误处理

我们知道传统的编程语言喜欢使用try模式，将很多错误一并处理，go语言不允许这么做，go的哲学就是，只要是错误，你就得处理，处理的越细致越好，看一下传统的go错误处理形如下文：

```go
func fast()error {
    return fmt.Errorf()
}

func main(){
    if err := fast();err != nil {
        //xxx
    }
}
```
这就是一个非常常见，以及非常传统的go生成错误，以及处理错误的方式，使用`fmt.Errorf`生成标准错误，然后处理错误的时候，使用一个if语句，如果错误不等于nil（可以回忆以下这里为啥是不等于nil，我上文说过）那么就对错误进行处理，其中这里的`fmt.Errorf`可以更换以下,`errors.New("")`也是生成了一个新的error对象。

接下来我们谈一下，最新的go版本里新添加的几个函数。

-  `errors.Unwrap`
    ```go
    func Unwrap(err error) error {
        // 先断言，
    u, ok := err.(interface {
    Unwrap() error
    })
    // 没有就返回nil，有就返回对象实例上面自己的Uwrap()方法
    if !ok {
    return nil
    }
    return u.Unwrap()
    }
    ```
    断言的时候，看看这个实现了error接口的对象实例，究竟有没有实现了Unwrap()error方法，这里的括号里面就是说要断言的具体的类型。这里是简略了。

    这里的省略，如果不省略就是下面这个样子

    ```go
    type A interface {
        Unwrap()error
    }

    func Unwrap(err error) error {
    u, ok := err.(A)
    if !ok {
    return nil
    }
    return u.Unwrap()
    }
    ```
    这是个小的知识点，可以留意以下。

- `errors.Is(err,target error)`
    这个语句的意思是说，err是否等于后面的结果，如果等于就是返回的true，如果不等于就返回false
    ```go
    // 本代码来自go标准库源码

    func Is(err, target error) bool {
	if target == nil {
		return err == target
	}
    // 这句话的意思就是说可以比较吗
	isComparable := reflectlite.TypeOf(target).Comparable()
	for {
        // 如果可以比较，那么以及err等于target，就返回true
		if isComparable && err == target {
			return true
		}
        // 下文就是不匹配了，如果不匹配的话，就调用err的IS()方法。意思就是如果这个实现了这个error接口的对象也实现了IS方法的话，那么就调用这个方法。如果匹配这个target的话也是返回true。
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}
		// TODO: consider supporting target.Is(err). This would allow
		// user-definable predicates, but also may allow for coping with sloppy
		// APIs, thereby making it easier to get away with them.
		if err = Unwrap(err); err == nil {
			return false
		    }
	    }
    }
    ```

- `errors.As(err error,target interface{})`
    As在err链中找到与目标实例匹配（意思是实现了共同的方法）的第一个错误，如果匹配，则将target设置为该错误值并返回true。否则，它返回false。

    通常来说，target本身是一个指针类型，然后传入的是这个指针类型的地址。下文中的代码里有相应的显示。

    这里稍微解释一下，err是一个调用的链条，只要其中一个跟target类型一样了，然后就把target赋值成err的值。而且这个target必须是指针类型。

    我们来看一下官方的案例

    ```go
        func main() {
            if _, err := os.Open("non-existing"); err != nil {
                var pathError *fs.PathError
                // 
                if errors.As(err, &pathError) {
                    fmt.Println("Failed at path:", pathError.Path)
                } else {
                    fmt.Println(err)
                }
            }

        }
    ```