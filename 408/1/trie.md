# 字典树
字典树的意思就是每个节点都布满了字符，然后不同的路径加和了不同的字符，然后就组成了单词，这就是字典树的意义。

举例说明
```go
          a        b
      c      d   e    f
```
那么可以组成的单词就是 a c  ad be bf 字典树可以把所有的字符都搞进去，变成了多叉树。所以其实字典树是一个多叉树。

![p](https://upload.wikimedia.org/wikipedia/commons/b/be/Trie_example.svg)
