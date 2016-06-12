package pq

import (
  "fmt"
  "testing"
)

//测试最大堆
func TestMaxPriorityQueue(t *testing.T) {
  h := NewMax()       // h 是一个 heap

  /**
    注意，下面是把  *NewItem(8, 10)类似的 结构体的的 指针  加入到堆里面去，记住是 指针
    按照这些  指针 的大小组成一个堆 ,  每个类似  *NewItem(8, 10)这样的都是一个整数
    
    
    这个做法似乎有点问题啊，下面的权重都是 按照 NewItem 的次序来排列的， 先new 的指针 就整数会小一些， 而且权重设置也小一些， 所以当然是这样的算法拉， 当我实际应用的时候，可不一定会按照这个顺序
  **/
  h.Insert(*NewItem(8, 10))
  h.Insert(*NewItem(7, 11))
  h.Insert(*NewItem(6, 12))
  h.Insert(*NewItem(3, 13))
  h.Insert(*NewItem(1, 14))
  h.Insert(*NewItem(0, 15))
  h.Insert(*NewItem(2, 16))
  h.Insert(*NewItem(4, 17))
  h.Insert(*NewItem(9, 18))
  h.Insert(*NewItem(5, 19))
  
  sorted := make([]Item, 0)     // Item{ Value: xx, Priority:  yy}
  for h.Len() > 0 {
    sorted = append(sorted, h.Extract())      //从堆的第一个节点开始，把堆的节点一个一个提取出来，append到 sorted 去
  }
  
  /**
    heap 代码的 insert 方法已经 进行了排序  
  **/
  
  for i := 0; i < len(sorted) - 2; i++ {
    if sorted[i].Priority < sorted[i+1].Priority {
      fmt.Println(sorted)
      t.Error()
    }
  }
  
}


//测试最小堆
func TestMinPriorityQueue(t *testing.T) {
  h := NewMin()
  
  h.Insert(*NewItem(8, 10))
  h.Insert(*NewItem(7, 11))
  h.Insert(*NewItem(6, 12))
  h.Insert(*NewItem(3, 13))
  h.Insert(*NewItem(1, 14))
  h.Insert(*NewItem(0, 15))
  h.Insert(*NewItem(2, 16))
  h.Insert(*NewItem(4, 17))
  h.Insert(*NewItem(9, 18))
  h.Insert(*NewItem(5, 19))
  
  sorted := make([]Item, 0)
  for h.Len() > 0 {
    sorted = append(sorted, h.Extract())
  }
  
  for i := 0; i < len(sorted) - 2; i++ {      // 为什么  i < len(sorted) - 2, 因为最后两个元素不用比较了 ??? 
    if sorted[i].Priority > sorted[i+1].Priority {
      fmt.Println(sorted)
      t.Error()
    }
  }
}


func TestChangePriority(t *testing.T) {
  h := NewMax()

  h.Insert(*NewItem(8, 10))
  h.Insert(*NewItem(7, 11))
  h.Insert(*NewItem(6, 12))
  h.Insert(*NewItem(3, 13))
  h.Insert(*NewItem(1, 14))
  h.Insert(*NewItem(0, 15))
  h.Insert(*NewItem(2, 16))
  h.Insert(*NewItem(4, 17))
  h.Insert(*NewItem(9, 18))
  h.Insert(*NewItem(5, 19))
  

/**
//这句对应方法 ChangePriority 的两句 
  popped.Priority = priority
  pq.data.Insert(popped)
  
  这个是最后 NewItem  的, 而且测试代码写了 权重是 66, 因为上面是  NewMax() , 而且这个是最后加入到， 所以 insert 之后这个就应该是最大堆的顶点
  
  所以  popped := h.Extract()  自然就是 弹出 8
**/
  h.ChangePriority(8, 66)         
	popped := h.Extract()

	if popped.Value != 8 {
		t.Error()
	}
}

















