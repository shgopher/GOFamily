# 单例模式
单例模式的意义就是，只执行一次，在go语言中，我们使用懒汉模式，或者饿汉模式，都是不优雅的，使用`sync.Once` 才是正确的方式。
## 懒汉模式
此模式下当包被加载的时候就要执行这个动作了
```go
type singleton struct {}
// 当包加载的时候，就开始执行了初始化这个动作
var app *singleton = &singleton{}

func CreateApp()*singleton {
	return app
}
```
## 饿汉模式
此模式下，只有方法被使用的时候才开始执行
```go
type singleton struct{}

var app *singleton

func CreateApp() *singleton {
	if app != nil {
		app = &singleton{}
	}

	return app
}
```
## sync.Once

```go
type singleton struct {}
var app *singleton
var once *sync.Once
// 使用CreateApp once.Do只能执行一次
func CreateApp()*singleton{
	once.Do(func() {
		app = &singleton{}
	})
	return app
}
```