package merge
/**
归并排序
算法时间复杂度 O(N * logN)
**/
func sort(arr []int) {
  var s = make([]int, len(arr)/2 + 1)       //定义 s slice
  if len(arr) < 2 {
    return
  }
  
  mid := len(arr) / 2
  
  sort(arr[:mid])     //分开两段进行归并排序
  sort(arr[mid:])
  
  if arr[mid-1] <= arr[mid] {     // 假如前一段的最大位少于 后一段的最小位， 那么排序已经是正确的了，函数返回
    return
  }

  //排序还没正确，则继续执行逻辑
  copy(s, arr[:mid])              // 赋予 前一段的数组给 s, 这个时候 s 和 arr  是相同的
  
  l, r := 0, mid                  // l = 0,   r = mid
  
  for i := 0; ; i++ {             // i 没有停下来的条件
    if s[l] <= arr[r] {           //刚开始循环的时候  s 是等于 arr 的， i 是从 0 开始的, 所以一开始是 如果 最左小于最右的话， 那么就重新定义  arr[i] 为最小的值
      arr[i] = s[l]
      l++
      
      if l == mid {       //如果归并的时候， 左段 的最小值 一直都 小于  右端段的 最小值，那么就是很好的结果， 最后会导致  l == mid , 排序完成， 跳出
        break;
      }
    } else {              //如果归并的时候 左段的最小值 还大于 有段的最小值，那么就用 右段的最小值去赋予 arr[i], 这个时候记住， 归并排序的一个特点就是， 归并分段去判断的时候， 段的排序已经排好了的， 所以这个时候就继续递增  r++ , 一直去判断  s[l] 和  arr[r] 的值， 假如一直都是大于 arr[r]， 那好办， 一直赋予  arr[r] 给 arr[i], 一直到 r == len(arr) 的时候， 把左段整段拷贝到 arr[i] 的右边去
      arr[i] = arr[r]
      r++
      if r == len(arr) {
        copy(arr[i+1:], s[l:mid])
        break
      }
    }
  } // END FOR
}