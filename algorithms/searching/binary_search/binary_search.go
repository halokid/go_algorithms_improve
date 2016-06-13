package binarysearch

import ()

/**
 二分法查找的算法关键，就是先算出 middle 的位置， 然后比较要查找的内容
 
 持续的，通过对比 middle位置的值要和寻找的值， 不断地以 middle+1 去代替 init， 或者 middle -1 去代替end, 不断地以middle为分界线去移动要寻找的范围
**/
func search(sortedArray []int, el int) int {      //在searchArray里面查找el
  init, end := 0, len(sortedArray) - 1
  
  for init <= end {
    middle := ( (end - init) >> 1) + init
    
    if sortedArray[middle] == el {
      return middle
    }
    
    if sortedArray[middle] < el {
      init = middle + 1
    } else {
      end = middle - 1
    }
  }
  
  return -1
}