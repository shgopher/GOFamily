<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-11-17 11:59:13
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-12-06 14:25:27
 * @FilePath: /GOFamily/编译器/gc/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# go 语言官方编译器 gc
## 前端组件
- Lexer：进行词法分析，将代码分割成词法单元
- Parser：进行语法分析，检查代码正确性，生成 AST
- Type checker：进行类型检查，确保类型使用正确
## 中端组件
- SSA：构建程序的静态单赋值形式中间表示
- Optimizer：进行各种形式的优化，如内联、逃逸分析等
## 后端组件
- Code generator：将优化后的代码生成对应的汇编代码
- Assembler：汇编器，将汇编代码转换为机器代码
- Linker：链接器，将对象代码链接成最终的可执行文件
- 运行时，Runtime：提供垃圾回收、协程调度等运行时功能

gc 不仅包含 frontend，还包含了中间 representation、优化器、代码生成器等组件，可以完成 Go 代码的全部编译工作。它既是一个编译器，也提供了部分运行时支持。所以说 gc 是 Go 语言的整体编译与工具链解决方案。

gc 编译器由 Go 语言的创造者 Robert Griesemer、Rob Pike 及 Ken Thompson 开发，首次在2007年与 Go 一起开源发布。

gc 实现了对 Go 语言的完整支持，可以编译包括复杂特性如 goroutine、channel 和接口在内的所有 Go 语言程序。它输出自包含的可执行文件，不需要外部依赖。

gc 编译器写在 Go 语言本身并采用了 Go 的并发特性，编译速度非常快。它与 Go 的发布周期同步，确保语言 feature 得到及时支持。

所以 Go 语言默认编译器的全称就是 gc，它是官方提供的 Go 语言编译器。
## 参考资料
- https://go.dev/src/cmd/compile/README