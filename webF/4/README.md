link 是 HTML 方式， @import 是 CSS 方式

link 最大限度支持并行下载，@import 过多嵌套导致串行下载，出现FOUC

link 可以通过 rel="alternate stylesheet"指定候选样式

浏览器对 link 支持早于 @import，可以使用 @import 对老浏览器隐藏样式

@import 必须在样式规则之前，可以在 css 文件中引用其他文件

总体来说：link 优于 @import
