package insertion

import ()
/**
插入排序

算法复杂度为 O(n^2)
**/
func sort(arr []int) {
  for i := 1; i < len(arr); i++ {
    value := arr[i]   //从数组的第二位开始循环, value 一开始是第二位
    j := i - 1        //从数组的 第一位 开始
    
    //下面的逻辑是先比较 第一 和 第二 位的数值
    for j >= 0 && arr[j] > value {
      arr[j+1] = arr[j]     //  大的值放右边
      j = j - 1
    } // END FOR 1 ,  循环之后， j 点的值为 (j-1)
  
    /**
     因为循环， 实际上 上面的 for循环, 一直循环到 符合条件的 j, 因为附加的循环条件 j = j -1
     本来应该把 作为对比标准的 value 赋予当时符合条件的j的， 但是由于 j经过了 j-1 之后，下面就
     成为   arr[j+1] = value 
    **/
    arr[j+1] = value   //把下一位的数值
  } //END FOR 2
}