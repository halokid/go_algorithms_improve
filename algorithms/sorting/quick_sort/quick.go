package quick

import (
  // "fmt"
  // "os"
)

func sort(arr []int) []int {
  // fmt.Println(len(arr))
  // os.Exit(0)
  var recurse func(left int, right int)            // 定义地鬼函数
  var partition func(left int, right int, pivot int) int      //定义分区间函数

  //实现 递归函数
  partition = func(left int, right int, pivot int) int {
    v := arr[pivot]     //取某个 点 作为比较对象
    right--             // ？？？
    arr[pivot], arr[right] = arr[right], arr[pivot]  // 交换数组里面比较对象 和 最右边的点的值
  
    // 以 最左边的点作为起始， 一直比较选择的点的值 v， 小于 v 的值 就不断把 left 的位置交换， left++ 
    for i := left; i < right; i++ {
      if arr[i] <= v {
        arr[i], arr[left] = arr[left], arr[i]
        left++
      }
    }
  
    // 此时 left 就是数组里面有多少个 元素是比 v 大的点的数值
    arr[left], arr[right] = arr[right], arr[left]
    return left
  }
  
  recurse = func(left int, right int) {
    /**
      我们走一次下面的循环吗， 梳理  left,  right 的数值改变:
      第一次, 假如数组的长度是100：
      left = 0,   right = 100
      那么
      pivot = (0 + 100) / 2 = 50
      则经过一次  partition 之后， 假如  pivot  和 left 之间有 30个 数值 是小于 v， arr[pivot] 的话
      则得到
      left = 30,   right = 99
      
      ------------------------------------------------------------
      这里解答几个疑问：
      1,  为什么 partition 要  right-- ？
           因为调用那里是  recurse(0, len(arr))
           所以 arr 的最大索引就是  len(arr) - 1, 一开始 right = len(arr), 所以要进行  right-- 处理
           
           
      2,  为什么 arr[pivot], arr[right] = arr[right], arr[pivot]
           因为以 中间作为基准  pivot, 在比较 arr[i] 和 v 的过程中，left++， 有可能left最后的结果会大于 pivot， 快速排序就是要得到 pivot 的值之后，然后把pivot放到最右边， 再进行递归对比
      
      3,   arr[left], arr[right] = arr[right], arr[left] 怎么解释？？
           当 left 不断地增大， 直到退出循环之后， 此时要把 arr[right] 也就是原来 v 的值作为基准放在 新的区间的最右边， 而新的区间就是 0 ---- left 值 之间的区间，所以执行这个逻辑
      ------------------------------------------------------------
    
    **/
    if left < right {
      pivot := (right + left) / 2       //选择的元素一开始一般是中间
      pivot = partition(left, right, pivot)
			recurse(left, pivot)
			recurse(pivot+1, right)
		}
	}

	recurse(0, len(arr))
	return arr
}



