package heap

import (
  "sync"
)

/**
1, golang 的特性，   Heap 有属性  sync.Mutex 之后,  h* Heap 就可以使用  h.Lock(), h.Unlock() 等功能

2,  siftUp 是在 insert 的过程中，梳理一次最小堆的过程， 从第一位开始，使整个数组尽量接近按照最小堆的顺序来排列，最终肯定会把最小的点排在  h.data  的首位
    而  siftDown 是在  Extract 堆的过程中要做的， 从最后一位开始， 使整个数组尽量接近按照最小堆的顺序来排列， 最终肯定会把最小的点排在  h.data 的首位
    
其实任何数据结构会有很多表达方式，这里的方式是 insert 右排序一次，  Extract又排序一次， 只是这里的方式而已    

**/


type Item interface {
  Less(than Item) bool      //看见没  Less()的参数是自身，没有什么不可能
}

type  Heap struct {
  sync.Mutex
  data  []Item
  min   bool
}


func New() *Heap {
  return &Heap{
    data: make([]Item, 0),
  }
}


func NewMin() *Heap {
  return &Heap{
    data:   make([]Item, 0),
    min:    true,               // min 的true 标识， 表示这个 heap 是最小的
  } 
}


func NewMax() *Heap {
  return &Heap{
    data:   make([]Item, 0),
    min:    false,
  }
}


func (h *Heap) isEmpty() bool {
  return (len(h.data) == 0)
}


func (h *Heap) Len() int {
  return len(h.data)
}

func (h *Heap) Get(n int) Item {
  return h.data[n]
}

/**
按照顺序添加 
================================

  h.Insert(Int(8))
  h.Insert(Int(7))
  h.Insert(Int(6))
  h.Insert(Int(3))
  h.Insert(Int(1))
  h.Insert(Int(0))
  h.Insert(Int(2))
  h.Insert(Int(4))
  h.Insert(Int(9))
  h.Insert(Int(5))
  
===============================
[0 1 3 2 5 6 7 4 9 8]         //这个是 Insert 处理之后的第一次排序
[1 2 3 4 5 6 7 8 9]
[2 4 3 8 5 6 7 9]
[3 4 6 8 5 9 7]
[4 5 6 8 7 9]
[5 7 6 8 9]
[6 7 9 8]
[7 8 9]
[8 9]
[9]
**/


func (h *Heap) Insert(n Item) {
  h.Lock()
  defer h.Unlock()
  
  h.data = append(h.data, n)        // append 是标识 在 slice 的末尾处插入节点， 则这时节点的位置为  data[len() - 1],刚好是下面  siftUp()  的 i， parent 的初始值
  h.siftUp()
  
  return
}


// extract 提取的意思
//这个提取是把 堆的第一个节点提出去， 然后把剩下的节点重新组成堆
func (h *Heap) Extract() (el Item) {      //指定了返回 e1, 函数最后return 的时候不用 return e1， golang自己会返回
  h.Lock()
  defer h.Unlock()
  
  if h.Len() == 0 {
    return
  }
  
  el = h.data[0]
  last := h.data[h.Len() - 1]
  
  if h.Len() == 1 {
    h.data = nil
    return
  }

  //下面这两段的逻辑解释
  /**
    首先   last := h.data[h.Len() - 1] ， 取得 last 的值， 这个时候赋予  []Item{last}, 这里是一个slice, 一开始   h.data[1:h.Len()-1]   这里的 h.data 是一个数组， 这个数组的排序不一定是对的，有可能大小不一的
    
    append([]Item{last}, h.data[1:h.Len()-1]...)
    这句就是把 整个 h.data 的所有元素都一个一个 加入到  Item  里面去
  **/
  h.data = append([]Item{last}, h.data[1:h.Len()-1]...)
  h.siftDown()        // 梳理 h.data 为最小堆的函数
  
  return
}


//这个是最小堆的操作
// 堆的 索引 index， 也就是这里的 i，是按顺序来的，从 顶点开始 一直向下，从左到右 , 索引逐渐增加
// 这个操作的顺序是 最堆的最底层开始 ， 一直比较底层节点与它的父节点的大小，然后再决定是否交换的操作
// 如果想还原逻辑， 1 >> 1 等于 0

/**
siftUp 这个就是一个冒泡算法，去计算， 不断得把最小的点放到  h.data 数组的前面，但是不断insert进来之后，最后还不是排好的顺序，因为每个 insert 进来的话，只要一碰到 数组 左边的某个比它小的点，就会停止再去比较了 ，最后排出来的是一个大概排序好的列表 

**/
func ( h *Heap) siftUp() {    
  for i, parent := h.Len() - 1, h.Len() - 1; i > 0; i = parent {
     /** 
     假如当 h.Len() 是 7 的时候， 则按照堆来的应该是这样的 
     
             1
           /  \
          2    3
         /\    /\
        4  5  6  7
        
        i       初始为   7 - 1 = 6
        parent  初始为   7 - 1 = 6
    
  
    siftUp 就是从在底层向上寻找
     
     **/
    parent = i >> 1       // 找到最底层的 i 的父节点, 右移一位（就是等于除以2）来求出父节点的索引位置, 此时 parent = 3

    if h.Less(h.Get(i), h.Get(parent)) {      //如果 i 小于  h
      h.data[parent], h.data[i] = h.data[i], h.data[parent]       //比较大小之后互换
    } else {
      break
    }
  }
}

/**
 i << 1+1,  就是 i 向左移动两位， 也就是 作为索引的 i 来做就是等于 i 编程了它的子节点的索引值， 是 i 自身改变
 下一个循环就是从  i = child  开始的， 因为是持续 从 i 的子节点的索引开始计算的
 
 下面这个逻辑是循环匹配， 貌似性能一般般, 这个一个最小堆的比较逻辑， 顶点的值是真个堆最大的

**/
func (h *Heap) siftDown() {
  for i, child := 0, 1; i < h.Len() && (i << 1+1) < h.Len(); i = child {      //FUCK, 这么多条件
    child = i << 1+1
    
    
    // 在刚开始的循环得到 i 之后， 然后一直用 索引 i节点的值 去对比 索引值为 child节点的值， 假如 下一个节点的值一直比索引child的值小的话(也就是比上一个节点的值)， 那么child就一直增大，向堆下去移动
    
    if child + 1 <= h.Len() - 1 && h.Less(h.Get(child+1), h.Get(child)) {
      child++
    }
    
    if h.Less(h.Get(i), h.Get(child)) {       //如果 i节点  小于  child节点，其实这个时候 i 就是在 child 节点隔壁
      break
    }
    
    h.data[i], h.data[child] = h.data[child], h.data[i]
  } 
}



//返回两个 item 的整型类型的大小对比的结果
func (h *Heap) Less(a, b Item) bool {
  if h.min {          // 如果 h 的 min 属性是 true 的话， 那么就返回  a < b  
    return a.Less(b)
  } else {            // 如果不是， 返回  b < a
    return b.Less(a)
  }
}




















