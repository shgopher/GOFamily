# go 语言官方编译器 gc

Go 语言默认的编译器全称是 gc，它是 Go 语言官方实现的编译器，主要功能包括：

- gc - Go 语言编译器前端，负责分析和转换 Go 语言源代码。
- ssa - 生成静态单赋值 (SSA) 中间表示。
- scheduler - 根据依赖关系安排指令执行顺序。
- linker - 链接程序依赖的库。
- runtime - Go 运行时，提供垃圾回收，并发等功能。
- assembler - 将汇编指令转换为机器码。

gc 编译器由 Go 语言的创造者 Robert Griesemer、Rob Pike 及 Ken Thompson 开发，首次在2007年与 Go 一起开源发布。

gc 实现了对 Go 语言的完整支持，可以编译包括复杂特性如 goroutine、channel 和接口在内的所有 Go 语言程序。它输出自包含的可执行文件，不需要外部依赖。

gc 编译器写在 Go 语言本身并采用了 Go 的并发特性，编译速度非常快。它与 Go 的发布周期同步，确保语言 feature 得到及时支持。

所以 Go 语言默认编译器的全称就是 gc，它是官方提供的 Go 语言编译器。