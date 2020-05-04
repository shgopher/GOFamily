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
[gotools](https://github.com/googege/gotools/tree/master/id) 这里有实现了的源代码
