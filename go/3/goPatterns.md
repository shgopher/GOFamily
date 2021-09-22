# go常见设计模式

常见的设计模式有创建型，结构型，行为型，这里只讲go项目中常见的设计模式。

创建型：
- 单例模式
- 工厂三种模式

结构型：
- 策略模式
- 模版模式

行为型：
- 代理模式
- 选项模式

本文档在这里只会涉及到这些内容。因为在go的项目中，这几个才是最常用的，更多关于设计模式的内容，在下面的408篇中会更加详细的阐述。

## 单例模式
单例模式的意义就是，只执行一次，在go语言中，我们使用懒汉模式，或者饿汉模式，都是不优雅的，使用`sync.Once` 才是正确的方式。

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

## 简单工厂

## 抽象工厂

## 工厂方法

## 策略

## 模版

## 代理

## 选项

