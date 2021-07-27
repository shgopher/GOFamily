# go语言项目线上事故排查
使用go的pprof包：
- runtime/pprof：采集程序（非Server）的运行数据进行分析
- net/http/pprof：采集 HTTP Server 的运行时数据进行分析

可以做到检测：
- CPU Profiling：`CPU 分析`，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置
- Memory Profiling：`内存分析`，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏
- Block Profiling：`goroutine运行分析`，记录 goroutine 阻塞等待同步（包括定时器通道）的位置
- Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况
## 内存泄漏
## 协程泄漏