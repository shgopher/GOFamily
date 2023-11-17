<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-11-17 12:28:57
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-11-17 12:34:29
 * @FilePath: /GOFamily/编译器/llvm/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# llvm
LLVM 可以通过以下两种主要方式为 Go 语言代码提供编译支持：

## gollvm
gollvm 项目实现了将 Go 代码编译为 LLVM IR 的 frontend。它可以直接生成 LLVM IR，然后通过 LLVM 进行后端代码生成。

主要步骤是：

- gollvm 解析 Go 代码，生成对应的 LLVM IR
- 执行 LLVM 优化流水线进行优化
- LLVM 后端生成目标平台的机器码
- 嵌入 LLVM pass
## 可以通过修改 Go Compiler 工具链，在编译过程中调用 LLVM Pass 执行优化。

主要步骤是：

- Go Compiler 前端生成 initial IR
- 将 IR 传递给 LLVM Pass 执行优化
- LLVM Pass 输出优化后的 IR
- Go Compiler 后端根据 IR 生成机器码

这种方式需要 invasive 修改 Go 编译器，较为复杂。

总结：

gollvm 通过完全 External 的方式引入 LLVM，而嵌入 LLVM Pass 需要修改 Go Compiler。

gollvm 方式集成更简单，但是需要保证 IR 的转换正确性；嵌入 Pass 可以利用 Go Compiler 的 Context 信息进行更准确的优化。

根据需求和成本进行抉择。
## llvm -golang
llvm-golang 是另一个用于 Go 语言编译的 LLVM 集成项目。

它与 gollvm 的主要区别有：

实现方式不同
- gollvm 是独立的 Go frontend，将 Go 代码编译成 LLVM IR
- llvm-golang 直接使用 LLVM 对 Go 代码进行编译
编译入口不同
- gollvm 通过调用 gollvm 命令，传入 Go 源文件
- llvm-golang 编译时调用 clang 命令，并使用-femulate-llvm-golang 参数
编译过程不同
- gollvm 编译生成完整的 LLVM IR 然后优化
- llvm-golang 是逐函数生成 LLVM IR 并编译
项目状态不同
- gollvm 活跃维护，可正常使用
- llvm-golang 最后更新在5年前，未维护
总之，llvm-golang 是直接使用 LLVM 编译 Go 的早期尝试，但维护情况不佳。gollvm 作为独立 frontend 集成 LLVM，是更可靠的解决方案。

但 llvm-golang 的直接编译方式也具有借鉴意义，合理结合两种方式可以获得更好的 LLVM 集成效果。
