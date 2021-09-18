# 并查集
并查集是一种森林，它叫做有根的森林。
## 并查集的伪代码

 ```go
 func MakeUnion (x){
   x.root = x
 }
 func Search(i) int {
   for array[i] != i {
     i = array[i]
   }
   return i
 }
 func unin(x,y) {
   array [x.Search()] = y.Search()
 }
 ```
## 并查集的含义
并查集存在的意义是为了寻找是否是属于一个集体，举个例子 你是个炮兵 你想看你的好友，他也当兵他是不是跟你一个部队，那么就可以使用并查集。
