#  唯一id生成算法
> [uuid参考](https://github.com/imgoogege/go.uuid)

常见的满足 分布式的 唯一id生成算法有

- uuid
- 数据库自增id auto_increment
- 数据库多主模式
- 号段模式
- redis incr
- twitter 雪花算法 [snowflake](https://github.com/twitter-archive/snowflake) [go版本](https://github.com/bwmarrin/snowflake)
- 滴滴打车 [tinyid](https://github.com/didi/tinyid)
- 百度 [uidgenerator](https://github.com/baidu/uid-generator)
- 美团 [leaf](https://github.com/Meituan-Dianping/Leaf)

didi和meituan都是使用的数据库多主模式的算法 百度是使用的 twitter的算法 当然它是Java版本，上面有go的版本。生成唯一id的时候 uuid并不合适。因为uuid本身没有意义，而且太长，这样的话不利于储存。

[**gotools**](https://github.com/googege/gotools/tree/master/id) 这里有实现了的源代码

## uuid
uuid 是一个16个字节 128bit的字符串，基本上你拿它当mysql的键性能就死定了，还有就是它没有任何的规则性，
当然如果你把它当值储存，以及本来就要求没有什么含义，uuid当然适合不过了，而且如果是满足并发分布式的使用
version1 可满足这个需求。
## 数据库自增
数据库自增id auto_increment mysql，压不住分布式的压力。
## 基于数据库集群模式
使用主从模式的集群，使用双主集群模式，两个mysql都生成id
```bash
set @@auto_increment_offset = 1;     -- 初始值
set @@auto_increment_increment = 2;  -- 步长
```
```bash
set @@auto_increment_offset = 1;     --  初始值
set @@auto_increment_increment = 2;  -- 步长
```
这样的话就是 1 3 5 7 9  vs 2 4 6 8 10 完全可以满足了。

问题来了，假如说每次都扩容数据库，那么该如何搞定这些ide的步长以及值呢，为了不冲突，就不得不人工去更改初始值以及步长
不太利于自动扩容。当然如果有个算法可以自动的计算初始值和步长也OK。
## 基于数据库的号段模式
此方法较为常用，

|id|	b_type|	max_id|	step	|version|
|:---:|:--:|:---:|:--:|:---:|
|1|	101	|1000|	2000|	0|

- id:本次获取的id
- b_type: 是标记来访的用途
- max_id: 是最大的id值
- step 是从第一个到max_id之间的长度
- version 就是版本号，代表获取的是第几段，这是为了避免id重复搞的一种乐观锁

 每次请求直接将这一段输入到内存中，比如使用一个数组这种，然后再享用，如果是并发来获取数据
 因为用乐观锁在，所以也不会出现重复数据的行为。

 每次取完  max_id = max_id + step  verison++
 我们可以将version储存在一个地方，然后每次取完数据肩擦version是否已经存在了，如果存在就放弃这次的行为
 重新取值。

 另外数据库会有锁的。完全不用怕数据重复的问题。

## twitter 的 snow flake
基本原理就是基于一个int64的长类型，然后这个类型中有64个bits，这就是它的核心思想，说白了 位运算是它的核心思想。
## 基于redis的原子操作 incr
```go
set id 1
incr id
```
由于redis是原子操作的，所以线程安全。

但是reids是运行在内存中的，它有两种持久化的方法1 rdb 2 aof

rdb 可能会造成id重复但是它恢复的适合速度快，aof不会id重复，但是它恢复的慢。这个需要取舍。
