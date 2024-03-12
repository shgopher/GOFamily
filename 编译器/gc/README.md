<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-11-17 11:59:13
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-03-12 15:26:14
 * @FilePath: /GOFamily/编译器/gc/README.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
# Go 语言官方编译器 gc

## 前端组件

### 1。Lexer (词法分析器)

Lexer 是编译器前端的第一个重要组件。它的主要职责是对源代码进行词法分析，将源代码分割成一个个词法单元 (tokens)。这些词法单元包括关键字、标识符、运算符、字面量等基本单元。

在 Go 语言中，Lexer 的实现在 `src/go/scanner` 包中。它使用了一种高效的 DFA (确定有限状态自动机) 算法进行词法分析。同时，它还支持 Unicode 编码，可以对包含多语种字符的源代码进行正确扫描。

### 2。Parser (语法分析器)

Parser 接收 Lexer 输出的词法单元，根据语言的语法规则构造抽象语法树 (AST)。AST 是源代码的树状表示形式，反映了代码的层次结构和语义信息。

Go 的语法分析器实现在 `src/go/parser` 包中。它采用 LL(1) 自顶向下分析算法，支持对 Go 语言所有语法结构的解析。Parser 会进行语法检查，如果发现语法错误，会报告相应的错误信息。

### 3。Type Checker (类型检查器)

Type Checker 对 AST 进一步进行语义分析，尤其是类型信息的收集和一致性检查。它会为所有节点附加类型信息，检查变量使用是否正确，函数调用参数是否匹配等。如果发现类型错误，会输出相应的错误信息。

Type Checker 的实现位于 `src/go/types` 包中。除了类型检查，它还可以解析导入路径、收集方法信息等。

## 中端组件

### 1。SSA (静态单赋值形式)

SSA 是一种将程序的控制流展开的中间表示形式。Go 编译器在 Type Checker 之后，会将 AST 转换为 SSA 形式，以方便后续的优化分析和转换。

SSA 的实现位于 `src/cmd/compile/internal/ssa` 包。在 SSA 中，每个变量只会被赋值一次，所有复杂的控制流都被展开成基本块的形式。这种表示形式方便了大量的编译器优化，如值传播、常量折叠等。

### 2。Optimizer (优化器)

Go 编译器内置了多种优化器模块，用于对 SSA 形式的程序进行各种形式的优化转换。主要优化器包括：

- 逃逸分析器 (`src/cmd/compile/internal/escape`)
- 内联器 (`src/cmd/compile/internal/inliner`)
- 死码消除 (`src/cmd/compile/internal/deadcode`)
- 常量传播和折叠 (`src/cmd/compile/internal/opt`)
- 控制流优化 (`src/cmd/compile/internal/opt`)

其中，逃逸分析是 Go 编译器最重要的优化环节。它可以分析变量是否会逃逸到堆上，进而决定是否可以优化为堆分配或栈分配。这对于 Go 这样的带有垃圾回收的语言而言至关重要，能极大减少内存分配和垃圾回收的压力。

## 后端组件

### 1。Code Generator (代码生成器)

代码生成器会将优化后的 SSA 表示形式转换为对应的机器码指令序列。

Go 编译器采用了自定义的高效代码生成器，其实现位于 `src/cmd/compile/internal/gc` 中。它不仅支持多种常用硬件架构 (x86、ARM、RISC-V 等)，还支持硬件特性加速。例如在 ARM64 架构上支持了利用 SVE 向量指令集进行矢量化优化。

### 2。Assembler (汇编器)

汇编器将前端生成的机器码经过进一步处理，生成目标平台的二进制代码。Go 编译器目前使用 GNU 汇编器进行汇编。

### 3。Linker (链接器)

链接器将全部目标文件 (object files) 以及需要的系统库文件链接合并，生成最终的可执行目标文件。可执行文件是完全独立无需外部依赖的自包含程序文件。

Go 语言自带小型高效的链接器，实现位于 `src/cmd/link` 包中。它支持静态链接和动态链接两种方式，默认采用静态链接生成全自包含的可执行文件。

### 4。Runtime

Runtime 库提供了垃圾回收、goroutine 调度、处理系统调用等运行时支持。Go 语言运行时高度优化，垃圾回收器采用了三色标记-压缩算法，并行和并发的处理提高了效率。

Goroutine 调度器使用了 M:N 调度模型，将 goroutine 和系统线程进行高效多路复用。这使得 Go 语言可以轻松创建大量并发任务，并拥有优秀的并发性能。

### 其他特性

除了高效完整的编译器支持，Go 编译器还具备以下特性：

1. 多核并行编译：gc 编译器利用并发支持，使用可用的所有 CPU 核心并行编译源文件，提高整体编译速度。
   
2. 增量编译：只编译被修改的源文件和依赖它们的源文件，避免对无关代码进行重新编译。
   
3. 类型检查缓存：将类型检查结果缓存，以加速后续编译过程。

4. Go 编写：gc 编译器本身就是使用 Go 编写的，这使得它可以自举并确保与 Go 语言保持高度契合。

Go 编译器与 Go 语言及其工具链源码一并开源发布，方便社区贡献者阅读理解和改进编译器实现。

### 总结

gc 是 Go 语言默认的官方编译器，它提供了完整的前端、优化器和后端支持，能够将 Go 语言源代码编译为高效的机器码。通过采用先进的编译器技术和算法，如 SSA 表示、逃逸分析等，gc 编译器可以生成高度优化的执行程序。

与此同时，gc 编译器也提供了强大的运行时支持，包括垃圾回收、goroutine 调度等核心功能。这使得 gc 不仅仅是一个简单的编译器，更是 Go 语言全栈的编译和执行解决方案。

Go 开发团队一直在持续改进完善 gc，以支持更多最新的语言特性和平台，满足日益增长的性能和兼容性要求。gc 编译器的持续进化有力地支撑了 Go 语言工程级应用的快速发展。

## 真实案例的分析
好的，我们来补充一些实际的例子，帮助理解编译器的工作原理。

编译器前端示例：

### Lexer 示例

假设有以下 Go 源代码：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Lexer 会将该源代码分割为以下词法单元：

```
package
main
import
"fmt"
func
main
(
)
{
fmt
.
Println
(
"Hello, World!"
)
}
```

### Parser 示例

经过词法分析后，Parser 会将这些词法单元构造成抽象语法树 (AST)：

```go
// FileNode
File{
    Package: main
    Imports: []*ImportSpec{
        &ImportSpec{
            Path: &BasicLit{
                Value: "fmt"
            }
        }
    }
    Decls: []Decl{
        &FuncDecl{
            Name: &Ident{Name: "main"}
            Body: &BlockStmt{
                List: []Stmt{
                    &ExprStmt{
                        X: &CallExpr{
                            Fun: &SelectorExpr{
                                X: &Ident{Name: "fmt"}
                                Sel: &Ident{Name: "Println"}
                            }
                            Args: []Expr{
                                &BasicLit{
                                    Value: "Hello, World!"
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

### Type Checker 示例

Type Checker 会对 AST 进行类型分析，并附加类型信息：

```go
// File Node with Type Information
File{
    Package: main
    Imports: []*ImportSpec{
        &ImportSpec{
            Path: &BasicLit{
                Value: "fmt"
            }
        }
    }
    Decls: []Decl{
        &FuncDecl{
            Name: &Ident{Name: "main", Type: func()}
            Body: &BlockStmt{
                List: []Stmt{
                    &ExprStmt{
                        X: &CallExpr{
                            Fun: &SelectorExpr{
                                X: &Ident{Name: "fmt", Type: *fmt}
                                Sel: &Ident{Name: "Println", Type: func(a ...interface{}) (n int, err error)}
                            }
                            Args: []Expr{
                                &BasicLit{
                                    Value: "Hello, World!",
                                    Type: string
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

编译器中端示例：

### SSA 示例

在 SSA 中，代码被转换为基本块的形式，控制流完全展开：

```
# Example SSA representation
func main():
bb0: 
    t0 = new [1]string { /* str */ }
    t1 = &t0[0]
    *t1 = "Hello, World!"
    t2 = staticbytes(&<string  Value>)
    t3 = &t2[0]
    t4 = slice t2[:]
    t5 = make([]interface{}, 1)
    t6 = &t5[0]
    t7 = &t4
    t8 = *t7
    *t6 = t8
    t9 = fmt.Println(&t5...)
    return

# Optimized SSA representation
func main():
bb0:
    t0 = staticbytes(&<string  Value="Hello, World!">)
    t1 = &t0[0]  
    t2 = slice t0[:]
    t3 = make([]interface{}, 1)
    t4 = &t3[0]
    t5 = &t2
    t6 = *t5
    *t4 = t6
    t7 = fmt.Println(&t3...)
    return
```

可以看到，在优化后的 SSA 表示中，一些冗余的内存分配和字符串构造操作都被优化掉了。

编译器后端示例：

### Code Generator 示例

在 `amd64` 平台上，优化后的 `fmt.Println("Hello, World!")` 会被生成如下汇编代码：

```
0x001b  TEXT    "".main(SB), ABIInternal, $24-0
        MOVQ    $go.string."Hello, World!"(SB), AX      // 加载字符串常量
        MOVQ    AX, (SP)                                // 将字符串常量传入参数区
        MOVQ    $1, 8(SP)                               // 设置 slice 长度
        MOVQ    $1, 16(SP)                              // 设置 slice 容量
        CALL    runtime.convTstring(SB)                 // 调用 convTstring
        MOVQ    8(SP), AX                               // 加载转换后的 slice
        MOVQ    AX, (SP)                                // 将 slice 作为参数
        CALL    fmt.Println(SB)                         // 调用 fmt.Println
        MOVQ    24(SP), BP                              // 恢复 BP
        ADDQ    $24, SP                                 // 调整栈指针
        RET                                             // 返回
```

可以看到，Go 编译器后端生成了紧凑高效的汇编指令，包括字符串常量加载、参数传递、函数调用等操作。

### Linker 示例

假设我们编译了以下两个 Go 源文件：

```go
// file1.go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
    sayHi()
}

// file2.go  
package main

func sayHi() {
    fmt.Println("Hi")
}
```

Go 编译器会先将它们分别编译为目标文件 `file1.o` 和 `file2.o`。链接器 (`cmd/link`) 会将这两个目标文件以及需要的运行时库链接合并，生成最终的可执行文件 `main`。

该过程类似于：

```
$ go tool compile -o file1.o file1.go
$ go tool compile -o file2.o file2.go  
$ go tool link -o main file1.o file2.o
```

最终输出的 `main` 可执行文件中，包含了 `main` 包定义、`fmt.Println` 导入符号以及 `sayHi` 函数实现等所有需要的代码和元数据。

通过这些实例，我们可以更好地理解 Go 编译器各组件的具体工作方式，以及它们是如何高效协作将 Go 源代码转换为机器码的。这种透明度和可阅读性也是 Go 编译器的一大优势。

## 参考资料

- https://go.dev/src/cmd/compile/README
- https://golang.org/doc/ssa
- https://go101.org/article/compiler.html
- https://dmitri.shuralyov.com/idiomatic-go#compiler