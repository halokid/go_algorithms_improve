package selection


/**
下面的 sort 我们开展一下逻辑思考：
第一遍:
当 i == 0， 以arr[i]为比较基准， j == (i+1) == 1 的时候， min = i = 0
第一个 for 的逻辑是
当 arr[j] < arr[min] 的时候， 把 min 设置为 j(也就是1) , 这个的逻辑就是 在 从 1 开始 到数组的最右边的元素 之间把最小的元素找出来， 并且把 这个元素的索引值 赋给 min
交换 arr[i]  和  arr[min]， 也就是交换  arr[1] 和  arr[0], 把最小的值 arr[min]， 交换到 FOR2 的循环元素 arr[0]去， 所以最小值就确定在了  arr[0]

--------------------------------------------------------------

第二遍:
当 i == 1， 以arr[i]为比较基准， j == (i+1) == 2 的时候， min = i = 1
第一个 for 的逻辑是
当 arr[j] < arr[min] 的时候， 把 min 设置为 j(也就是2) 



----------- 逻辑总结 -------------
从最左边的点开始， 依次为基准， 然后选择 基准到最右边的点之间最小的元素， 然后交换到基准的点的索引位置

**/

func sort(arr []int) []int {
  for i := 0; i < len(arr); i++ {
    min := i
    for j := i + 1; j < len(arr); j++ {
      if arr[j] < arr[min] {
        min = j
      }
    } // END FOR 1
  
   if min != i {        //加多一步 min 的检查， 确定 不为 i ，表明在 FOR1 里面发生了对于min的新的设定
    arr[i], arr[min] = arr[min], arr[i]
   }  
  } // END FOR 2
  
  return arr
}