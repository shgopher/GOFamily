# GO的坑
### go中的for循环要想退出一定要使用这种方式

```go
L1:
for{
  L2:
  for{
    L3:
    for{
      break L1// 看你是L几就是调到哪里
    }
  }
}
```
```go
for{
  for{
    for{
      break  // 并不能跳出全部，只能跳出最里面的那层，
    }
  }
}
```

### go所有的数据都是数据的复制，并不是指针类型，如果想使用指针类型请明示

```go
func A(i int)int{
  i++
  return i
}
func main(){
  t := 2
  A(t)
  fmt.Println(t)//结果还是2，因为你在A中只是复制而已
}
```
### go中所有的数据只要声明，就可以初始化，

引用类型的初始化是nil，引用类型的初始化需要make或者明示 比如

`a=[]int{1,2,3,4,5}`，还有nil本身不含任何的类型，所以先声明才有类型。

### defer的执行顺序，和它初始化的顺序完全是两码事，
defer的执行是按照stack来执行的，不过它的初始化是按照正常的顺序来执行的
```go
func Test(i int){
  defer func (i int)  {
    fmt.Println(i)
  }(i)
  i++
}
// output:i 而不是i+=1
```
### 对于接口的时候非常严格，指针和非指针对象完全是不同的。
也就是说是整形实现来接口还是指针类型实现来接口完全不同。
### rang中的数据是对原有数据的复制，其实所有的地方使用的数据都是复制。所以要谨记这一条。

```go
func main(){
    test([]int{1,2,3})
    test1([]int{1,2,3})
}

func test(value []int){
    for _,v := range value {
        v++
    }
    fmt.Println(value)
}
// 很明显想改变就使用这种方法。所以为什么说让你slice后的数据尽可能是指针类型，因为range会复制。
func test1(value []int){
    for i:= 0;i<len(value);i++ {
        value[i]++
    }
    fmt.Println(value)
}
```
### copy 对于数据的复制是值的复制不是引用的复制，所以对新的slice进行更改不会对老的底层数组造成任何的影响
### go中 `^` 这个 符号不表示次方，只表示异 ，go中的次方只能用 math.Pow
### go里 前套了匿名字段的struct，如果重写了匿名字段的方法，在接口实现的时候，匿名字段不会调用重写的方法还是调用自己的方法。

```go

func main() {

	s := Shape{name: "Shape"}
	c := Circle{Shape: Shape{name: "Circle"}, r: 10}
	r := Rectangle{Shape: Shape{name: "Rectangle"}, w: 5, h: 4}

	listshape := []ShapeInterface{&s, &c, &r}
	for _, si := range listshape {
		si.PrintArea()
	}

}


type ShapeInterface interface {
	Area() float64
	GetName() string
	PrintArea()
}

// 标准形状，它的面积为0.0
type Shape struct {
	name string
}

func (s *Shape) Area() float64 {
	return 0.0
}

func (s *Shape) GetName() string {
	return s.name
}

func (s *Shape) PrintArea() {
	fmt.Printf("%s : Area %v\r\n", s.name, s.Area()) // 这里明确指出调用s.Area()
}

// 矩形 : 重新定义了Area方法
type Rectangle struct {
	Shape // Rectangle 嵌套了 shap然后等于它也实现了这个接口，然后当调用s.PrintArea()的时候，里面的s.Area() 不是这个地方重新的这个而是shap中定义的那个
	w, h float64
}

func (r *Rectangle) Area() float64 {
	return r.w * r.h+88
}

// 圆形  : 重新定义 Area 和PrintArea 方法
type Circle struct {
	Shape
	r float64
}

func (c *Circle) Area() float64 {
	return c.r * c.r * math.Pi
}

func (c *Circle) PrintArea() {
	fmt.Printf("%s : Area %v\r\n", c.GetName(), c.Area()) // 这里由于 circle重新实现了printarea 所以c.Area()就是使用的它自己重写的那个。
}

```
### go slice 切片的时候容易从前面的开始index开始算一直到底层数组的最后一个

```go
a := [10]int{}
b := a[3:4]
//len(b)==1
//cap(b)= 7
```
所以记住容量并不是10，它从它的初始值开始算了。
