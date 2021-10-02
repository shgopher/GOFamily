# 工厂模式
## 简单工厂
简单工厂的目的也是为了创建，我们可以看一下一段代码
```go
type People struct {
	name string
	year int
}

func (p *People) Post() {
	fmt.Println(p.name, p.year)
}

func NewPeople(name string, year int) *People {
	return &People{
		name: name,
		year: year,
	}
}
```
使用这种方式，newpeople可以接受参数，然后返回一个结构体。
## 抽象工厂
抽象工厂和简单工厂的区别是它不返回结构体，返回一个接口。
```go
type People interface {
	Do(req *http.Request) (*http.Response, error)
}
type people struct {
	name string
	year int
}

func (p *people) Do(req *http.Request)(*http.Response, error) {
	rec := httptest.NewRecorder()
	return rec.Result(),nil
}

func Newpeople(name string, year int) People {
	return &people{
		name: name,
		year: year,
	}
}
```
可以看到 new方法返回的是一个接口类型，当然了return的是一个实现了接口的structure，只不过这里有一个隐含的类型转换。

这样做的好处是什么呢，比如现在有另一个new，newStudent ,它当然也定义了一个Do

## 工厂方法
工厂方法的宗旨就是解耦，使用子函数来代替自身，这样就可以解除耦合，举个简单的例子：

```go
func main() {
	// 16岁的分成一组
	student16 := NewStudent(16)
	student16("red")
	student16("green")
	student16("blue")
	// 17岁的分成一组
	student17 := NewStudent(17)
	student17("red")
	student17("green")
	student17("blue")
	// 18岁的分成一组
	student18 := NewStudent(18)
	student18("red")
	student18("green")
	student18("blue")
}

type Student struct {
	name string
	year int
}

func NewStudent(year int) func(name string) Student {
	return func(name string) Student {
		return Student{
			name: name,
			year:year,
		}
	}
}


```
我们可以清晰的看出，我们将年龄作为一个root，然后name是child ，那么就是将 year和name解耦了，我们使用子函数去创建name这child选项。

所以可以看到上文中，我们创建了16岁的红蓝绿，17岁的红蓝绿，和18岁的红蓝绿 我们将year和name解耦了。