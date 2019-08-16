# OOP设计原则

- SRP	The Single Responsibility Principle	单一责任原则
- OCP	The Open Closed Principle	开放封闭原则
- LSP	The Liskov Substitution Principle	里氏替换原则
- ISP	The Interface Segregation Principle	接口分离原则
- DIP	The Dependency Inversion Principle	依赖倒置原则

SRP：一个类只负责一件事 比如一个接口中的method只有一个功能
OCP：类应该可以对外扩展 go中就是实现这个接口
LSP：子类对象必须能够替换掉所有父类对象。
LSP:使用多个专门的接口比使用单一的总接口要好
DIP:高层模块不应该依赖于低层模块，二者都应该依赖于抽象；
抽象不应该依赖于细节，细节应该依赖于抽象。
