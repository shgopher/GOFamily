<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-10-12 11:34:58
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2023-10-12 11:48:28
 * @FilePath: /GOFamily/工程/go面试题/中高级go面试题一.md
 * @Description: 
 * 
 * Copyright (c) 2023 by shgopher, All Rights Reserved. 
-->
1. 自我介绍
2. 代码效率分析，考察局部性原理
3. 多核 CPU 场景下，cache 如何保持一致、不冲突？
4. uint 类型溢出
5. 介绍 rune 类型
6. 编程题：3 个函数分别打印 cat、dog、fish，要求每个函数都要起一个 goroutine，按照 cat、dog、fish 顺序打印在屏幕上 100 次。
7. 介绍一下 channel，无缓冲和有缓冲区别
8. 是否了解 channel 底层实现，比如实现 channel 的数据结构是什么？
9. channel 是否线程安全？
10. Mutex 是悲观锁还是乐观锁？悲观锁、乐观锁是什么？
11. Mutex 几种模式？
12. Mutex 可以做自旋锁吗？
13. 介绍一下 RWMutex
14. 项目中用过的锁？
15. 介绍一下线程安全的共享内存方式
16. 介绍一下 goroutine
17. goroutine 自旋占用 cpu 如何解决 (go 调用、gmp)
18. 介绍 linux 系统信号
19. goroutine 抢占时机 (gc 栈扫描)
20. Gc 触发时机
21. 是否了解其他 gc 机制
22. Go 内存管理方式
23. Channel 分配在栈上还是堆上？哪些对象分配在堆上，哪些对象分配在栈上？
24. 介绍一下大对象小对象，为什么小对象多了会造成 gc 压力？
25. 项目中遇到的 oom 情况？
26. 项目中使用 go 遇到的坑？
27. 工作遇到的难题、有挑战的事情，如何解决？
28. 如何指定指令执行顺序？
