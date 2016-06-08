package heap

import (
  "sync"
)

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


func (h *Heap) Insert(n Item) {
  h.Lock()
  defer h.Unlock()
  
  h.data = append(h.data, n)        // append 是标识 在 slice 的末尾处插入节点， 则这时节点的位置为  data[len() - 1],刚好是下面  siftUp()  的 i， parent 的初始值
  h.siftUp()
  
  return
}



func (h *Heap) Extract() (el Item) {
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
  
  h.data = append([]Item{last}, h.data[1:h.Len()-1]...)
  h.siftDown()
  
  return
}


//这个是最小堆的操作
// 堆的 索引 index， 也就是这里的 i，是按顺序来的，从 顶点开始 一直向下，从左到右 , 索引逐渐增加
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


//返回两个 item 的整型类型的大小对比的结果
func (h *Heap) Less(a, b Item) bool {
  if h.min {          // 如果 h 的 min 属性是 true 的话， 那么就返回  a < b  
    return a.Less(b)
  } else {            // 如果不是， 返回  b < a
    return b.Less(a)
  }
}




















