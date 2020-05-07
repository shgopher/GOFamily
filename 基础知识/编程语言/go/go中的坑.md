# GO的坑
1. go中的for循环要想退出一定要使用这种方式

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

2. go所有的数据都是数据的复制，并不是指针类型，如果想使用指针类型请明示

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
3. go中所有的数据只要声明，就可以初始化，但是引用类型的初始化是nil，引用类型的初始化需要make或者明示 比如 `a=[]int{1,2,3,4,5}`

4. defer的执行顺序，和它初始化的顺序完全是两码事，具体你可以把它想成栈就ok了。defer虽然是FILO但是初始化还是要顺序初始化的。这也正常
5. 对于接口的时候非常严格，指针和非指针对象完全是不同的。
6. rang中的数据是对原有数据的复制，其实所有的地方使用的数据都是复制。所以要谨记这一条。
7. copy 对于数据的复制是值的复制不是引用的复制，所以对新的slice进行更改不会对老的底层数组造成任何的影响
8. go中 `^` 这个 符号不表示次方，只表示异 ，go中的次方只能用 math.Pow
9. go里 前套了匿名字段的struct，如果重写了匿名字段的方法，在接口实现的时候，匿名字段不会调用重写的方法还是调用自己的方法。

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
