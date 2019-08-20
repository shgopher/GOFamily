# 数据淘汰策略

当数据占内存超过一个值的时候，就要施行数据淘汰策略了。

描述

- volatile-lru	从已设置过期时间的数据集中挑选最近最少使用的数据淘汰
- volatile-ttl	从已设置过期时间的数据集中挑选将要过期的数据淘汰
- volatile-random	从已设置过期时间的数据集中任意选择数据淘汰
- allkeys-lru	从所有数据集中挑选最近最少使用的数据淘汰
- gallkeys-random	从所有数据集中任意选择数据进行淘汰
- noeviction	禁止驱逐数据
