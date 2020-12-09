# 储存设备
## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)
## 基本命令
- mount – 挂载一个文件系统

- umount – 卸载一个文件系统

- fsck – 检查和修复一个文件系统

- fdisk – 分区表控制器

- mkfs – 创建文件系统

- fdformat – 格式化一张软盘

- dd — 把面向块的数据直接写入设备
- genisoimage (mkisofs) – 创建一个 ISO 9660的映像文件

- wodim (cdrecord) – 把数据写入光存储媒介

- md5sum – 计算 MD5检验码

## 挂载和卸载储存设备。

用到的命令是
- mount
- umount

Linux 操作系统为例，你会注意到系统看似填充了多于它所需要的内存。 这不意味着 Linux 正在使用所有的内存，它意味着 Linux 正在利用所有可用的内存，来作为缓存区。
