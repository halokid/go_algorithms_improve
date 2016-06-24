package shell

/**
算法逻辑:
假如 len(arr) = 10
increment = 10 / 2 = 5
然后进入 FOR 2， 循环条件是 increment > 0
初始化值为
i = 5, i < 10, i++ ，  进入 FOR 1
j = 5
temp 为 索引5 的值

FOR 3 的逻辑应该是   5 >= 5  && arr[5-5] > arr[5] 所以第一个区间就是 arr[0] 到 arr[5]
假如这个条件里面  arr[5-5] <= arr[5] 的话，则忽略 FOR 3

假如符合 arr[5-5] > arr[5], 进入循环
arr[5] = arr[5-5]
5 = 5 - 5 = 0 (  j = j - increment)
运算之后   j  的值为 0 


其实就是循环区间， 然后对比区间的 最左边的点 和 最右边的点， 假如 最右边比最左边的点小的话，则交换两个点

疑问：
j = j - increment 这个逻辑的理解 ？？？
第一次循环的时候 j=5,  increment=5,  j = j - increment 刚好就是 区间的最左边的点的索引值
那么当 i 慢慢 循环改变的时候， 我们就可以根据  j = j - increment 一直取得区间的最左边的点的索引值

------------------------------------------------------

increment = int(increment * 5 / 11)  的理解 ???
至于上面这个寻找 increment 的值的做法， 这是一种取得 increment 比较好的位置的一种值方法


**/


func sort(arr []int) {
  increment := len(arr) / 2
  
  for increment > 0 {
    for i := increment; i < len(arr); i++ {
      j := i
      temp := arr[i]
      
      for j >= increment &&  arr[j-increment] > temp {
        arr[j] = arr[j - increment]
        j = j - increment
      } // END FOR3
      arr[j] = temp
      
    } // END FOR1
    
    if increment == 2 {
      increment = 1
    } else {
      increment = int(increment * 5 / 11) 
    }
  } // END FOR2
}


