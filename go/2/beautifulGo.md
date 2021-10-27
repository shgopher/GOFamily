# 优雅的go项目

- 符合go 编码规范 和最佳实践
- 易阅读，易理解，易维护
- 易测试，易扩展
- 代码质量高

我们使用以下方法来遵守这些原则：

- 代码结构
    - 目录结构
    - 按照功能拆分模块
- 代码规范
    - 编码规范
    - 最佳实践
- 代码质量
    - 编写可测试的代码
    - 高单元测试覆盖率
    - code review
- 编程哲学
    - 面向接口编程
    - 面向对象编程
- 软件设计方法
    - 设计模式
    - 遵循SOLID原则

## 代码结构
核心内容就是标准的代码目录，以及代码的拆分方式。

目前最流行的就是按照功能拆分。比如拆分为 用户，订单，计费等三个模块。每一个模块都可以独立运行，诶，这就可以实现高内聚低耦合的设计理念，因为互不影响嘛。
## 代码规范
可以参考一下[uber的规范](https://github.com/imgoogege/uber_go_guide_cn)说明
## 代码质量

- code review
- 单元测试

将代码从不可mock的状态更改为可mock的代码演示：
```go
// 不可mock

type Student struct {

}
// SomeMethod 这里的a.Some() 是属于some的特有的方法，mock的时候是无法访问
// 的，所以这里，应该把这个sturct和这个函数解耦
func SomeMethod(a some)Student{
    return a.Some()
}

// SomeMehod1 这里将struck改成一个接口，这样就解耦了
type Server interface{
    Some()Student
}
func SomeMethod1(s Server)Student{
    s.Some()
}
```

## 编程哲学

- 面向接口
    - 只要可以抽象的地方一律使用接口进行抽象，将结构体和函数分离
- 面向对象
    - 接口实现 “多态”
    - 嵌入实现 “继承”
    - 结构体实现 “抽象”


## 软件设计方法

- [设计模式](./goPatterns.md)
- solid理论

|SOLID理论|中文描述|具体介绍|
|:---:|:---:|:---:|
|SRP|单一功能原则|一个类或者模块只负责完成一个职责或者功能|
|OCP|开闭原则|软件实体应该对扩展开放，对修改关闭|
|LSP|里氏替换原则|如果s是t的子类型，那么t的对象可以替换为s的对象并且不会破坏程序|
|DIP|依赖倒置原则|依赖于抽象而不是实例，本质是面向接口编程而不是面向现实编程|
|ISP|接口分离原则|客户端程序不应该依赖它不需要的方法|

> 可以[参考](https://github.com/googege/geekbang-go/blob/master/SOLID原则介绍.md)这里
