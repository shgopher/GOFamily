# 排序算法
![p](https://raw.githubusercontent.com/imgoogege/Sorting-Algorithm/master/res/sort.png)
接下来我们按照重要程度来排序的说一下各种排序算法
## 经典的十种排序算法
**快速排序**

毫无疑问 ，快排作为用到最多的，使用最广泛的排序算法，以及它还有衍生的“快速选择”，不过要注意的是快排不是稳定的排序，下面的同样拥有良好性能的归并排序是稳定的排序。如果要求稳定那么只能选归并排序。

```go
func QuickSort(arr []int)[]int{
    quickSort(0,len(arr)-1,arr)
    return arr
}

func quickSort(left,right int,arr  []int){ // 严格按照递归的模版来写。
    // stop
    if left > right {
        return
    }
    //process
    ge := f(left,right,arr)
    // trill down
    quickSort(left,ge-1,arr)
    quickSort(ge+1,right,arr)
    // clear states
}

func f(left,right int,arr []int)int{ // 辅助函数的意思是寻找坐标的被排列的位置，以及将切片排好位置。
    pri := left
    index := pri+1
    for i := index;i <= right;i++ {
        if arr[i] < arr[pri] {
           swap(i,index,arr)
            index++
        }
    }
    swap(pri,index-1,arr)
    return index -1
}
func swap(i,j int,arr []int){
    arr[i],arr[j]  = arr[j],arr[i]
}

```
时间复杂度分析：每次的排列是`n` 然后将这些数据分治是`log n` 乘积就是 `nlog n`

`快速选择的代码:`

```go

```
这种算法主要是用在topk的问题上。

**归并排序**

```go

```
完美的使用了分治的思想，先分 再和 。

**堆排序**

```go

```
堆排序是不断堆化的过程。所以跟topk问题也是很大的源源。

**冒泡排序**

```go
func Mao(value []int)[]int{
  for i := 0;i < len(value);i ++ {
    for j := 0;j < len(value)-1-i;j++ {
      if value [j] > value[j+1]{
        value[j],value[j+1] = value[j+1],value[j]
      }
    }
  }
  return value
}

```
`冒泡法的优化方法:`

当某次的循环中所有的都不再交换顺序，证明已经排好了。
```go
func Mao(value []int)[]int{
  L:
  for i := 0;i < len(value);i ++ {
    ma := 0
    for j := 0;j < len(value)-1-i;j++ {
      if value [j]> value[j+1]{
        ma ++
        value[j],value[j+1] = value[j+1],value[j]
      }
    }
    if ma == 0 {
      break L
    }
  }
  return value
}

```

**选择排序**

```go
func Sort(arr []int)[]int{
  for i := 0;i<len(arr)-1;i++ {
    min := i
    for j := i+1;j<len(arr);j++ {
      if arr[min] > arr[j] {
        min = j
      }
    }
    arr[min],arr[i] = arr[i],arr[min]
  }
}
```
**插入排序**

```go
```

**希尔排序**

```go
```

**筒排序**

```go
```
**计数排序**

```go
```
**基数排序**

```go
```

## go语言中的排序算法
先参考[这里](https://segmentfault.com/a/1190000016514382)

go的写法 如果 数据< 12 就是希尔排序 然后大于12的时候 当排序的次数小于n的时候就是 快排 然后大于n了 就开始堆排序。嗯 还挺复杂。。
