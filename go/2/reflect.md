# go语言的反射

反射的意义是可以在go代码执行的过程中，动态操作对象，其中最重要的两个函数就是
- reflect.TypeOf

    ```go
    func main() {
	var a interface{}
	a = 1
	fmt.Println(reflect.TypeOf(a))
	a = "1"
	fmt.Println(reflect.TypeOf(a))
	a = true
	fmt.Println(reflect.TypeOf(a))
	a= map[int]int{1:1}
	fmt.Println(reflect.TypeOf(a))
    }
    ```
- reflect.ValueOf

    获得类型的value值，进而对值在运行时进行操作。

前者是为了动态的获取目标的类型，后者是为了获得目标的value值。

- reflect.DeepEqual()

	使用这个函数，可以获取**深度的比较**，比如两个结构体的比较，map的比较。