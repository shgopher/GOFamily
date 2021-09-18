# 解题模版
## 遇到比较内容的时候

**比较 s和t（都是字符串）的区别**
```go
ma := make(map[rune]int)
for _,v := range s {
  ma[v]++
}
for _,v := range t {
  ma[v]-- // 这里 -- 是重点
}
for _,v := range ma {
  if v != 0 { // 这里不等于0是重点
    return false
  }
}
```
## 递归解题模版

```go
func recursion(level int,parm1,parm2...) parm{
  // terminator 终止条件
  if xxx {
    return XXX
  }

  // process 本层要处理的东西
  xxxxxxx

  // trill down 往下一层输入的信息
  recursion(level+1,parmNew1,parmNew2)

  // clear states  可选，递归返回的时候要处理的东西
  xxxxxxx
}
```
终止条件 + 本层处理 + 往下层走 + 递归返回的时候的处理[可选]

## 分治的解题模版

```go
func DivideConquer(problem,parm) parm  {
  // 判断问题是否结束
  if problem == nil {
    return result
  }
  // 开始准备数据，拆分数据
  data := prepareProblem(problem,parm)
  sonProblem := splitProblem(problem,data)
  // 开始执行往下一层的递归任务
  p1  := DivideConquer(sonProblem[0])
  p2  := DivideConquer(sonProblem[1])
  p3  := DivideConquer(sonProblem[2])
  p4  := DivideConquer(sonProblem[3])
}
// 开始合并这几个子问题
return merge(p1,p2,p3,p4)
```
我们可以看到其实分治使用的就是递归，只不过分治的思想加入来进去。可以看出来递归是及其重要的一个算法思想，很多其它的算法都要使用递归来解决问题。
## 广度优先搜索解题模版
```go
func BFS(n *Node,value interface{}) *Node {
  ma := make(map[*node]int)// 这是去重的辅助map
  queue := make([]*Node,0)
  queue = append(queue,n.Value) // 将数据加入到queue中
  ma[n]++ // 记录这个节点

  for len(queue) > 0 {

    //这一部分是为了寻找数据
    ma[t]++
    t = queue[0]
    queue = queue[1:] // 这是pop了数据
    if t.Value == value {
      return t
    }
    // 这一部分是为了将数据加入到queue中，然后为了找到他们的孩子数据。
    for _,v  := range t.Child {
        if _,ok := ma[v];!ok { // 树的话不用判重。如果是图这一步不能少。
          queue = append(queue,v)
        }
      }
    }
}
```
## 深度优先搜索解题模版
```go
 a = map[int]int{}
func DFS(n *Node,value ) *Node{
  a[node]++
  if n.Child == nil && n.Value != value {
      return nil
    }
    if n.Value == value {
      return n
    }
    for _,v := range n.Child {
      if _,ok := ma[v];ok {
          if DFS(v,value) != nil {
          return t
          }
        }
    }
    return nil
}
```
## 二分查找的代码
```go
func Search(target,value){
  left,right := 0,len(value)-1
  for left <= right {
    mid = left + (right-left)>> 1
    if value[mid]== target {
      return value[mid]
    }else if value[mid] > target{
      right = mid -1
    }else {
      left = mid+1
    }
  }
}
```
## 树的遍历模版代码

**前中后**
```go
func P(){
  // xxxxx 在这里处理 本函数的本次的东西 就是前序遍历
  p(left)
  // xxxxx  这里就是中序遍历 其中比如说判断是否是搜索二叉树的时候就可以中序遍历然后看是否是前面小于后面就OK了。
  p(right)
  // xxxx 放到这里处理 那肯定是 后序遍历了。
}
```

**层序遍历** 类似于广度优先搜索，解决的问题不过是每层的节点如何划分层次

```go
func p(n *node){
  // 定义数据结构
  queue := NewQueue
  ma :=  map[*node]int
  result := make([][]int)
  // 开始初始化
  queue.Add(n)
  ma[n]++
  for len(queue) != 0 {
    n := len(queue)
    re := make([]int,0)
    for i := 0;i < n;i++ {
      d := queue.Pop()
      re = append(re,d)
      for _,v := range d.Children {
        queue.push(v)
      }
    }
    result = append(result ,re)
  }
}
```
## 动态规划的模版 ！！！！！ 超级重要 望周知

```go
func DP(){
  // 状态的定义 一般是设置一个二维或者一维或者三维的数组
  dp = [m][n]int{}
  // 初始化这个数组 例如把第一层都手动加上好数据
  dp[1] = xxxx
  // 开始动态推演
  for i := 0;i <= m;i++ {
    for j := 0;j <= n ;j++ {
        dp [i][j] = min{dp[i-1][j],dp[i][j-1]} // 这里仅仅作为事例。
       }
    }
    return dp[m][n] // 一般都是这个矩阵的最右边的那个是最优解或者是最左上角的那个是也就是dp[0][0] 这个就是最优解了。
}
```
## 链表题解题模版
说明：日常生活中双链表是最常见的，毕竟好操作嘛，但是吧，面试的时候往往是单链表，可能更考验出人的能力？

```go
func LinkedList(root){
  pre := nil //  维护一个前节点，极其重要，望周知。
  cur := root
  for cur != nil {
    if pre == nil {
       ////
    }else {
       /////
    }
    pre = cur
    cur = cur.next
  }
}
```
