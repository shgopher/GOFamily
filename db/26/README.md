reis vs memchached

- reids有5中类型，memchached仅仅支持字符串
- redis支持持久化，后者不支持，主要是redis不近可以运行在内存中，它还可以RDB快照和AOF日志，它可以将数据储存在硬盘上，进行持久化的保存，
- redis可以支持分布式，memchached不支持分布式
