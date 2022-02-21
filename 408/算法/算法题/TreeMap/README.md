# TreeMap

基础知识：基于红黑树（平衡二叉搜索树）的一种树状 hashmap，增删查改、找求大最小均为logN复杂度，Python当中可以使用SortedDict替代；SortedDict继承了普通的dict全部的方法，除此之外还可以peekitem(k)来找key里面第k大的元素，popitem(k)来删除掉第k大的元素，弥补了Python自带的heapq没法logN时间复杂度内删除某个元素的缺陷；最近又在刷一些hard题目时候突然发现TreeMap简直是个神技，很多用别的数据结构写起来非常麻烦的题目，TreeMap解决起来易如反掌。

常见题目：
- Leetcode 729 My Calendar I
- Leetcode 981 Time Based Key-Value Store
- Leetcode 846 Hand of Straights
- Leetcode 218 The Skyline Problem
- Leetcode 480. Sliding Window Median (这个题用TreeMap超级方便)
- Leetcode 318 Count of Smaller Numbers After Self (这个题线段树、二分索引树、TreeMap都可以)