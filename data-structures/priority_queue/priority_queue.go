package pq

import (
	"github.com/arnauddri/algorithms/data-structures/heap"         //引入heap 堆文件
	"github.com/arnauddri/algorithms/data-structures/queue"        //引入 queue  队列文件
)

/**
结论是：
似乎这个算法挺扯蛋的， 因为都是安排好了权重顺序， 但是也可以学习学习算法逻辑

怎么定义有优先级的队列， 不是在原来的队列上改进下， 而是要结合 heap 和 queue 的程序文件一起呢？？
这里有一些 golang 的引用文件的用法，可以参考一下golang 语言的特性
调用 引用文件 的结构体或者函数的时候， 如果返回指针类型，用 *heap 
调用 引用文件 的结构体或者函数的时候， 如果返回普通类型，用 heap 
**/


type Item struct {
  Value         interface{}         // item的 value 是 interface 类型
  Priority      int                 //优先级的属性
}

//添加 item， 属性是 值 和 权重
func NewItem(value interface{}, priority int) (i *Item) {
  return &Item{
    Value:      value,
    Priority:   priority,
  }
}


func (x Item) Less(than heap.Item) bool {
  return x.Priority < than.(Item).Priority
}

// PQ实际上是一个堆来的。。。。
type PQ struct {
  data heap.Heap      //heap 是引用的文件， Heap 是这个文件里面的一个结构体, PQ只有一个属性， data ， 就是 heap 文件里面的  Heap 结构体
}


func NewMax() (q *PQ) {
  return &PQ{
    data:   *heap.NewMax(),       // 这里的  *heap.NewMax() 是什么写法？？ 本来在文件 heap 里面，也没有这种写法，文件里面只有  func NewMax() *Heap {}  的定义， 所以这里这个用法的意思就是引用了 文件 heap 里面的 NewMax 方法 , 写法是  *heap.NewMax()  
  }
}


func NewMin() (q *PQ) {
  return &PQ{
    data:     *heap.NewMin(),       //跟上面一样
  }
}

func (pq *PQ) Len() int {
  return pq.data.Len()
}

func (pq *PQ) Insert(el Item) {
  pq.data.Insert(heap.Item(el))     //看好这里的定义，   heap.Item(el) 是什么写法来的？？ heap.Item(el) 是表示这里的  el 是属于 heap.Item类型的， 然后  这里 Item 类型是在 引用文件  heap 里面定义了的， 所以写法就是  heap.Item(el)
}

/**
看下面的这种golang的写法，貌似甚屌哦 
在 pa.data.Extract()  是在引用文件 heap 里面的一个方法， 该方法要输入参数  el Item， 但是在这里写的时候，本来自己文件里面的方法是有定义了   func (pa *PQ) Extract() (el Item) 然后下面  return pa.data.Extract()这里直接就不用再输入  el item 类型了，直接可以传承本身函数的参数？？？？
**/
func (pa *PQ) Extract() (el Item) {
  return pa.data.Extract().(Item)
}


func (pq *PQ) ChangePriority(val interface{}, priority int) {
  var storage  = queue.New()
  
  popped := pq.Extract()      //这里调用的 Extract() 是本文件的 Extract() , popped 返回本文件定义的  item
  
  for val != popped.Value {
    if pq.Len() == 0 {
      panic("Item not found")
    }
    
    storage.Push(popped)
    popped = pq.Extract()
  } 
  
  popped.Priority = priority
  pq.data.Insert(popped)
  
  for storage.Len() > 0 {
    pq.data.Insert(storage.Shift().(heap.Item))
  }
}




