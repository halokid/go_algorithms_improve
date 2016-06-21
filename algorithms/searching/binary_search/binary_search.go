package binarysearch

import ()

/**
 二分法查找的算法关键，就是先算出 middle 的位置， 然后比较要查找的内容
 
 持续的，通过对比 middle位置的值要和寻找的值， 不断地以 middle+1 去代替 init， 或者 middle -1 去代替end, 不断地以middle为分界线去移动要寻找的范围
**/
func search(sortedArray []int, el int) int {      //在searchArray里面查找el
  init, end := 0, len(sortedArray) - 1
  
  for init <= end {
    middle := ( (end - init) >> 1) + init    //求出中间位置的节点的逻辑是这个算法的关键
    
    if sortedArray[middle] == el {
      return middle
    }
    
    if sortedArray[middle] < el {
      init = middle + 1       //假如中间节点小于 要寻找的 el， 则 init（头节点）加一之后，再用14行的规则求中间节点
    } else {
      end = middle - 1        //逻辑同理
    }
  }   // END FOR
  
  return -1
}