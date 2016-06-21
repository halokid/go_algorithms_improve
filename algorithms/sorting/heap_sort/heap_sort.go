package heap

import (
  "github.com/arnauddri/algorithms/data-structures/heap"
)

/**
堆排序
先把乱序的数组进堆， 堆的 Insert 方法已经排序好了，然后再出堆，写进数组，用堆的 Extract 方法出堆，会逐渐把堆里面最大的节点出堆
**/

func sort(arr []int) []int {
  h := heap.NewMin()
  for i := 0; i < len(arr); i++ {
    h.Insert(heap.Int(arr[i]))
  }
  
  for i := 0; i < len(arr); i++ {
    arr[i] = int(h.Extract().(heap.Int))
  }
  
  return arr
}