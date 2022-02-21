# BFS
基础知识：

常见的BFS用来解决什么问题？(1) 简单图（有向无向皆可）的最短路径长度，注意是长度而不是具体的路径（2）拓扑排序 （3） 遍历一个图（或者树）

BFS基本模板（需要记录层数或者不需要记录层数）

多数情况下时间复杂度空间复杂度都是O（N+M），N为节点个数，M

基于树的BFS：不需要专门一个set来记录访问过的节点
- Leetcode 102 Binary Tree Level Order Traversal
- Leetcode 103 Binary Tree Zigzag Level Order Traversal
- Leetcode 297 Serialize and Deserialize Binary Tree （很好的BFS和双指针结合的题）
- Leetcode 314 Binary Tree Vertical Order Traversal

基于图的BFS：（一般需要一个set来记录访问过的节点）
- Leetcode 200. Number of Islands
- Leetcode 133. Clone Graph
- Leetcode 127. Word Ladder
- Leetcode 490. The Maze
- Leetcode 323. Connected Component in Undirected Graph
- Leetcode 130. Surrounded Regions
- Leetcode 752. Open the Lock
- Leetcode 815. Bus Routes
- Leetcode 1091. Shortest Path in Binary Matrix
- Leetcode 542. 01 Matrix
- Leetcode 1293. Shortest Path in a Grid with Obstacles Elimination

拓扑排序：（https://zh.wikipedia.org/wiki/%E6%8B%93%E6%92%B2%E6%8E%92%E5%BA%8F）
- Leetcode 207 Course Schedule （I, II）
- Leetcode 444 Sequence Reconstruction
- Leetcode 269 Alien Dictionary
- Leetcode 310 Minimum Height Trees
- Leetcode 366 Find Leaves of Binary Tree