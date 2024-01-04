<!--
 * @Author: shgopher shgopher@gmail.com
 * @Date: 2023-05-14 23:08:19
 * @LastEditors: shgopher shgopher@gmail.com
 * @LastEditTime: 2024-01-05 00:35:48
 * @FilePath: /GOFamily/并发/atomic/README.md
 * @Description: 
 * 
 * Copyright (c) 2024 by shgopher, All Rights Reserved. 
-->
# atomic
原子包是比互斥锁更底层的包，如果在简单的场景下，使用 sync.Mutex 可能会比较复杂，并且耗费资源，那么使用更加底层的 atomic 就更加划算了

所谓原子操作，就是当某 goroutine 去执行原子操作时，其它 goroutine 只能看着，这个操作要么成功，要么失败，不会有第三个状态


## issues
###
