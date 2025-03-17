(window.webpackJsonp=window.webpackJsonp||[]).push([[37],{466:function(t,s,a){"use strict";a.r(s);var n=a(36),r=Object(n.a)({},(function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"变量声明"}},[t._v("变量声明")]),t._v(" "),a("h2",{attrs:{id:"变量声明符号-var-和"}},[t._v("变量声明符号 "),a("code",[t._v("var")]),t._v(" 和 "),a("code",[t._v(":=")])]),t._v(" "),a("p",[t._v("在 go 语言中，我们使用 "),a("code",[t._v("var")]),t._v(" 来表示声明一个变量")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"hello world"')]),t._v("\n")])])]),a("p",[t._v("这就是一个标准的变量声明方式，var 符号在最前面，接着就是变量 a，变量后面紧跟着是变量的类型，这里是 string 类型，也就是字符串类型，"),a("code",[t._v("=")]),t._v(" 后面是要赋值的具体值")]),t._v(" "),a("p",[t._v("我们也可以直接声明不赋予初始值，go 语言默认，声明即赋予初始值，那么这里的字符串类型初始值就是一个空字符串 "),a("code",[t._v('""')])]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n")])])]),a("p",[t._v("go 语言具有类型推断能力，所以我们可以省略类型，让 go 语言的编译器去推断类型")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a  "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"hello world"')]),t._v("\n")])])]),a("p",[t._v("编译器会自动推断 a 为 string 类型")]),t._v(" "),a("p",[t._v("在函数体内部 (这是一个先决条件) 我们"),a("strong",[t._v("也")]),t._v("可以使用省略的写法，就是使用一个符号 "),a("code",[t._v(":=")]),t._v(" 来充当 "),a("code",[t._v("var")]),t._v(" 的角色，也就是初始化的工作，比如说")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("main")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  a "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"hello world"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("可以看到，整个的用法是 "),a("code",[t._v("[变量] [:=] [初始值]")]),t._v(" 这三者缺一不可，而且还不能多，不能在 a 后面带有类型，不能省略初始值，且仅限于函数/方法内部使用")]),t._v(" "),a("p",[t._v("go 语言支持多变量同时赋值")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" b "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" c"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" d "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"1"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"2"')]),t._v("\n\na"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" b "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"hello world"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("12")]),t._v("\n")])])]),a("p",[t._v("其中，使用 var 进行声明的时候，如果是多个变量同时声明，必须是相同类型；使用 "),a("code",[t._v(":=")]),t._v(" 进行多变量赋值时，多个变量可以不同类型，因为全靠编译器推断")]),t._v(" "),a("h2",{attrs:{id:"常见的变量声明方式"}},[t._v("常见的变量声明方式")]),t._v(" "),a("p",[t._v("从广义上来说，go 语言只有两种变量，包一级的变量和函数一级的变量")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" DefaultValue "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("NewMethod")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("n "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("其中 "),a("code",[t._v("DefaultValue")]),t._v(" 就是包级变量，"),a("code",[t._v("n")]),t._v(" 就是函数级变量 (也是形式变量)")]),t._v(" "),a("p",[t._v("下面我列举一些常见的声明方式")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n a "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n b "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int32")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" c "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("map")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("//包级变量")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("hi")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  d "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("12")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("//仅限函数内部使用，变量后面不能有类型")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" c "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("map")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 函数级变量")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" e "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"hello world"')]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("//自动推断变量类型")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\n")])])]),a("p",[t._v("可以说声明的方式很多，不过呢，在一个项目中应该尽量保证声明方式的一致性，因为可以加强代码的统一性，减少理解代码的难度。")]),t._v(" "),a("h2",{attrs:{id:"go-语言可导出变量"}},[t._v("go 语言可导出变量")]),t._v(" "),a("p",[t._v("go 语言跟一般的语言不同，它使用变量"),a("strong",[t._v("首")]),t._v("字母的大小写来区分变量的可导出性质，大写 (如果使用中文作为变量名称，默认是可导出的) 代表可导出，小写代表仅限包内部使用 (包这一级，多个文件只要是同一个包就可以使用)")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("package")]),t._v(" Example\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" OutPutName "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("0o77")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" inName "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("0x99")]),t._v("\n")])])]),a("p",[t._v("其中 "),a("code",[t._v("OutPutName")]),t._v(" 是一个可导出的变量，"),a("code",[t._v("inName")]),t._v(" 是不可导出变量")]),t._v(" "),a("p",[t._v("go 语言还拥有比如函数，方法，结构体，接口，等等各类型的组件，这里你可以先不懂到底是什么，你先有一个印象，但凡首字母是大写的都是可以导出的，小写就是包内使用，没错，go 语言就是这么简单")]),t._v(" "),a("h2",{attrs:{id:"包级变量"}},[t._v("包级变量")]),t._v(" "),a("p",[a("code",[t._v("第一种声明形式")]),t._v("："),a("strong",[t._v("声明的同时，显式初始化")])]),t._v(" "),a("ul",[a("li",[a("code",[t._v('var a = method("t")')]),t._v(" go 编译器根据右侧的返回值自动确定左侧变量的类型")]),t._v(" "),a("li",[a("code",[t._v("var a = 3.14")]),t._v(" 在没有具体返回值，没有具体类型的情况下，go 赋予它默认类型，比如 float 的默认类型就是 float64，int 类型的默认类型就是 int")]),t._v(" "),a("li",[t._v("如果要显式的赋予类型，并且保证命名的一致性 "),a("code",[t._v("var a int = 12")]),t._v(" 的行为应该避免，应该写成 "),a("code",[t._v("var a = int(12)")]),t._v(" 用来保证一致性"),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n  a "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("3.14")]),t._v("\n  b "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("int")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("12")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n  e "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" errors"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("New")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"EOF"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])])])]),t._v(" "),a("p",[a("code",[t._v("第二种声明形式")]),t._v("："),a("strong",[t._v("声明但是延迟初始化")])]),t._v(" "),a("p",[t._v("只有 "),a("code",[t._v("var a int64")]),t._v(" 这一种方式，不过呢，go 语言的声明是直接赋予零值的，比如说这里的 a 默认就是 "),a("code",[t._v("0")])]),t._v(" "),a("p",[t._v("go 语言变量声明的聚集和就近原则：将同一类型的放在一个 var() 内部；或者另一种分类方法：将有初始值的放在一个 var 里，将延迟初始化的放在另一个 var 里。")]),t._v(" "),a("p",[t._v("分类一")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\na "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),t._v("\nb "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),t._v("\nc "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\nd "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\ne "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\nf "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("分类二")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n  a "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("12")]),t._v("\n  b "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"string"')]),t._v("\n  c "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"string"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("\n  d "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),t._v("\n  c "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n  f "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("bool")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n")])])]),a("p",[t._v("接下来谈一谈就近原则，变量的声明和变量的使用尽量的近，不要都声明在头部，如果一个变量被全局大量使用，那么可以放在头部，如果就是仅仅使用少量的次数，还是应该在使用的前面就进进行声明")]),t._v(" "),a("h2",{attrs:{id:"函数级变量"}},[t._v("函数级变量")]),t._v(" "),a("p",[a("code",[t._v("第一种声明形式")]),t._v("："),a("strong",[t._v("延迟初始化")])]),t._v(" "),a("p",[t._v("在函数体内使用 var")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Method")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("int")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("如果变量特别多，也可以使用 "),a("code",[t._v("var()")]),t._v(" 的方法在函数内部使用")]),t._v(" "),a("p",[a("code",[t._v("第二种声明形式")]),t._v("："),a("strong",[t._v("声明且显式初始化的局部变量")])]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("method1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  a "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("12")]),t._v("\n  b "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("int32")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("20")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// 改变默认类型")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("h2",{attrs:{id:"小心-shadow-的变量"}},[t._v("小心 shadow 的变量")]),t._v(" "),a("p",[t._v("我们知道，当有两个两个以上的变量在赋值时，如果其中有一个未被提前声明，那么就需要使用 "),a("code",[t._v(":=")]),t._v("，这个时候系统会自动判断有哪些未提前声明，然而有一种场景下系统会发生误判，准确的来说这是一种歧义，系统的判断会跟程序员的心理不一致，出现了变量 shadow 的行为，让我们看一下代码：")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("WithName")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// tracing 为 bool 类型")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" tracing"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    a"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" example"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Method")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("else")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    a"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" example"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Method1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" \n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("在外层，a 已经提前声明，但是在 if 这个作用域中，由于 err 并未提前声明，所以使用了 "),a("code",[t._v(":=")]),t._v("，由于系统"),a("strong",[t._v("无法获知")]),t._v("这里的 a 是否需要再次声明，所以 go 语言默认 a 是一个新的变量，这样外层的 a 就无法得到新的值，外层 a 也就被内层的 a 给 shadow 了。")]),t._v(" "),a("p",[t._v("如果想改变这种 bug，我们可以将 err 也提前声明：")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("WithName")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" err "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("error")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// tracing 为 bool 类型")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" tracing"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    a"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" example"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Method")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("else")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    a"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" example"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Method1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" \n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("p",[t._v("或者也可以改变内部的变量名称，来改变这种 shadow：")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("WithName")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("var")]),t._v(" a "),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// tracing 为 bool 类型")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" tracing"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    ai"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" example"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Method")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n    a "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" ai\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("else")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    ai"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" example"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Method1")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n    a "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("=")]),t._v(" ai\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" \n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("h2",{attrs:{id:"参考资料"}},[t._v("参考资料")]),t._v(" "),a("ul",[a("li",[t._v("https://juejin.cn/post/7241452578125824061")])])])}),[],!1,null,null,null);s.default=r.exports}}]);