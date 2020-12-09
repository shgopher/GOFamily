# 只能在Linux中使用的命令

## [目录](./summary.md)

> 换言之在macOS中无法使用。

- free:主要是看空闲内存使用内存的命令。

- 注意 Linux中`control + c control + v`并不能像macOS一样复制个粘贴。

## tab 升级版
- `Alt-?`	显示可能的自动补全列表。在大多数系统中，你也可以完成这个通过按 两次 tab 键，这会更容易些。
- `Alt-*`	插入所有可能的自动补全。当你想要使用多个可能的匹配项时，这个很有帮助。

## 脚本启动时的home配置文件Linux是：


## 登录 shell 会话的启动文件
文件	内容
/etc/profile	应用于所有用户的全局配置脚本。

~/.bash_profile	用户私人的启动文件。可以用来扩展或重写全局配置脚本中的设置。

~/.bash_login	如果文件 ~/.bash_profile 没有找到，bash 会尝试读取这个脚本。

~/.profile	如果文件 ~/.bash_profile 或文件 

~/.bash_login 都没有找到，bash 会试图读取这个文件。 这是基于 Debian 发行版的默认设置，比方说 Ubuntu。

## 非登录 shell 会话的启动文件
文件	内容
/etc/bash.bashrc	应用于所有用户的全局配置文件。
~/.bashrc	用户私有的启动文件。可以用来扩展或重写全局配置脚本中的设置。


然而 macOS中并不是这样。
## [目录](./summary.md)
