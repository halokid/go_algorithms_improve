package bubble

import ()

/**
最差时间复杂度  O(n^2)
最优时间复杂度  O(n)
平均时间复杂度  O(n^2)
最差空间复杂度 总共 O(n), 需要辅助空间 O(1)
**/

/**
冒泡算法理解：
假如是从小到大排列，分别两层缓存， 外一层为从右到左循环，以右为边界， 然后内嵌的循环从左到右，  分别两两对比， 把大的移到右边， 算法循环理解为：

第一次   把最大的移到最右边
第二次   把剩下最大的（倒数第二大）移到 (最右边的位置 - 1), 内嵌循环最右边边界 
第三次   把剩下最大的（倒数第三大）移到 (最右边的位置 - 2), 内嵌循环最右边边界
**/
func sort(arr []int) {
  for itemCount := len(arr); ; itemCount-- {
    swap := false     //标记内嵌的循环有没有交换
    
    for i := 1; i < itemCount; i++ {
      if arr[i-1] > arr[i] {
        arr[i-1], arr[i] = arr[i], arr[i-1]
        swap = true
      }
    }   //END for, 里面这个for循环肯定是要循环到最后
    
    if swap == false {    //假如剩下的最大位已经是在 右边边界上， 则没有swap， break外层的循环
      break;
    }
  }
}
