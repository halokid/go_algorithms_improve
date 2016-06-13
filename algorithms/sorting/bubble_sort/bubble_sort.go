package bubble

import ()

func sort(arr []int) {
  for itemCount := len(arr) - 1; ; itemCount-- {
    swap := false
    
    for i := 1; i < itemCount; i++ {
      if arr[i-1] > arr[i] {
        arr[i-1], arr[i] = arr[i], arr[i-1]
        swap = true
      }
    }   //END for, 里面这个for循环肯定是要循环到最后
    
    if swap == false {
      break;
    }
  }
}