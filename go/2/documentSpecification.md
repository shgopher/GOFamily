# 文档规范
## README 规范
1. 项目名称
2. 功能特征
3. 软件的架构图【可选】
4. 快速开始例子
5. 项目依赖包
6. 构建项目的方法
7. 运行项目的方法
8. 使用项目的方法
9. 如何贡献
10. 社区
11. 关于作者
12. 谁在用
13. 许可证
## 项目文档规范
文档通常放置在 docs/ 里面，下面是规范文档的内容

![](https://gitee.com/shgopher/img/raw/master/%E6%96%87%E6%A1%A3%E8%A7%84%E8%8C%83.png)
## API接口文档规范
API 作为一个文件夹进行存放内容，其中大概是包含了以下内容：

- README.md 基本介绍
- CHANGELOG.md 变更历史说明log
- generic.md 保存通用的基本参数，返回参数，等通用数据
- struct.md 列出文档中使用的数据结构
- user.md secret.md policy.md restful资料
- err_code.md 错误码的描述文件

其中 user.md secret.md 存放的就是最重要的具体接口描述，它包含了以下[内容](https://github.com/googege/iam/blob/master/docs/guide/zh-CN/api/user.md)：

- 接口的描述
- 请求方法，比如 `GET` or  `POST`等，以及path
- 输入参数
- 输出参数
- 请求的例子
