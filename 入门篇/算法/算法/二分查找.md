# 二分查找
必须符合以下三点才可以使用二分查找
- 数据单调（单调递增或者单调递减）
- 存在上下的分界点（可以得到mid中间值）
- 可以通过索引获取到值（数组）
### 解题模版
```go
func Search(target,value){
  left,right := 0,len(value)-1 // 设置左右的双指针索引
  for left <= right { // 只要左比右边小于等于就OK
    mid = (left + right) / 2 // 获取到中间值
    if value[mid]== target { // 判断中间值是否等于目标
      return value[mid]
    }else if value[mid] > target{ // 看目标在哪个区域就把下次寻找的数据往哪个方向找。
      right = mid -1 // 注意如果是那种不能+1 就可以跨界的情况，right = mid即可
    }else {
      left = mid+1
    }
  }
}
```
**注意如果是那种不能+1 就可以跨界的情况[比如浮点数这种的等等]，right = mid 即可**
