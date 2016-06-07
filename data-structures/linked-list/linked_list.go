/**
* Author: JJyy<82049406@qq.com>
* 链表的实现 
**/
package list

import (
  "errors"
)


type List struct {      //定义 链表结构
  Length    int         //链表的长度
  Head      *Node       //链表头
  Tail      *Node       //链表尾
}


//新建一个链表
func NewList() *List {
  l := new(List)      //新建一个结构体存在内存里面，用 new 关键字
  l.Length = 0
  return l 
}

type Node struct {          //定义 node 的结构体
  Value   interface{}       // Node 的值
  Prev    *Node             // 前一个 node
  Next    *Node             // 后一个 node
}


//新建一个 node
func NewNode(value interface{}) *Node {
  return &Node{ Value:   value}
}


func (l *List) Len() int {      //返回链表的长度
  return l.Length
}


func (l *List) IsEmpty() bool {     //判断链表是否为空
  return l.Length == 0
}


//链表的 头点 插入法，即从第一个节点的位置 插入链表
func (l *List) Prepend(value interface{}) {
  node := NewNode(value)
  
  if l.Len() == 0 {             //当链表为空的时候,初始化head， 然后收尾相连
    l.Head = node
    l.Tail = l.Head
    
  //当链表不为空的时候， 很明显这个处理是一个后进链表就成为 链表的 head 的一个逻辑，难道这个就是 分离式链表的算法过程？？？
  //当有一个新的 node要进来链表的时候， 那么原来的 head node 就成为新的node 的next node， 新的node 就代替他成为 head
  } else {                      
    formerHead := l.Head
    formerHead.Prev = node
    
    node.Next = formerHead
    l.Head = node
  }
  
  l.Length++
}


//链表的 尾点 插入法  即从最后一个节点的位置 插入链表
func (l *List) Append(value interface{}) {
  node := NewNode(value)
  
  if l.Len() == 0 {
    l.Head = node
    l.Tail = l.Head
  } else {
    formerTail := l.Tail
    formerTail.Next = node
    
    node.Prev = formerTail
    l.Tail = node
    
  }
  l.Length++
}


//在 指定的 index 的位置加上某个节点
func (l *List) Add(value interface{}, index int) error {
  if index > l.Len() {
    return errors.New("Index out of range")
  }
  
  node := NewNode(value)

  /**
    当 链表为空， 即是要add 的位置就是在链表开头，则使用 Prepend 效率最高
  **/
  if l.Len() == 0 || index == 0 {
    l.Prepend(value)
    return nil
  }

  /**
    当 链表不为空， 但是要add 的位置就是在链表的末尾，则使用 Append 效率最高
  **/
  if l.Len() - 1 == index {
    l.Append(value)
    return nil
  }
  
  nextNode, _ := l.Get(index + 1)
  prevNode := nextNode.Prev
  
  prevNode.Next = node
  node.Prev = prevNode
  
  nextNode.Prev = node
  node.Next = nextNode
  
  l.Length++
  
  return nil
}


//注意返回的是 error 类型
//移除链表里面的某个元素
func (l *List) Remove(value interface{}) error {
  if l.Len() == 0 {             //如果链表为空
    return errors.New("Empty list")
  }
  
  if l.Head.Value == value {      //如果刚好是要删除链表的第一个节点
    l.Head = l.Head.Next
    l.Length--
    return nil
  }
  
  found := 0
  for n := l.Head; n != nil; n = n.Next {     //循环链表
    
    if *n.Value.(*Node) == value && found == 0  {       // *n.Value.(*Node) 这是什么用法啊？？？？
      n.Next.Prev, n.Prev.Next = n.Prev, n.Next
      l.Length--
      found++
    }
  }
  
  if found == 0 {
    return errors.New("Node not found")
  }
  
  return nil
}


//获取链表上的某个节点
//返回两种类型的数据, node 和  error
//在链表上获得某个节点都是从第一个节点开始获取的， 时间为 O(n)
func (l *List) Get(index int) (*Node, error) {
  if index > l.Len() {
    return nil, errors.New("index out of range")
  }
  
  node := l.Head        //从第一个点开始获取
  for i := 0; i < index; i++ {        //一直获取的指定的 index 位置上
    node = node.Next
  }
  
  return node, nil
}


//在链表中寻找某个节点， 返回 index位置
func (l *List) Find(node *Node) (int, error) {
  if l.Len() == 0 {
    return 0, errors.New("Empty list")
  }
  
  index := 0
  found := -1
  l.Map( func(n *Node)  {       //函数作为参数~~好用法, n 就是从 map函数里面得到的 n节点
    index++
    if n.Value == node.Value && found == -1 {
      found = index
    }
  })    //END MAP

  if found == -1 {
    return 0, errors.New("Item not found")
  }
  
  return found, nil
}


//合并两个链表
//逻辑操作好两个链表的 头 尾 属性就可以了
func (l *List) Concat(k *List) {
  l.Tail.Next, k.Tail.Prev = k.Head, l.Tail
  l.Tail = k.Tail
  l.Length += k.Length
}  


// Map 仅仅是循环链表？？？
/**
解释一下这种函数作为参数引用的流程， 我们看下调用代码是
=========================================================================

l.Map( func(n *Node)  { 
    index++
    if n.Value == node.Value && found == -1 {
      found = index
    }
  })
  
上面的代码逻辑解释的流程为  list类型直接调用 Map函数， Map函数的参数是一个函数 f

执行流程为:

首先在  Map 函数里面 执行下面的代码
for node := list.Head; node != nil; node = node.Next  {
  n := node.Value.(*Node)
  
  上面这段逻辑都是执行 map 函数里面的代码， 然后得到 n 之后，干什么呢？？？
  
  得到 n 之后就开始执行   l.Map 指定的函数里面的逻辑，因为 l.Map 里面执行的函数就是
  f  func(node *Node)
  
  所以      f(n)
  就等于是执行了
  ***********************
  index++
  if n.Value == node.Value && found == -1 {
      found = index
    }
  ************************
  上面这段代码。。。。。
}

这样分开理解比较好理解， 可以很好的解释了 golang 用函数作为参数的执行过程

==========================================================================

**/
func (list *List) Map( f func(node *Node) ) {
  for node := list.Head; node != nil; node = node.Next  {
    n := node.Value.(*Node)       //这种 用法是既表示  n是属于 node.value 类型， 又是属于 node 类型 ？？？
    f(n)
  } 
}

func (list *List) Each(f func(node Node)) {
  for node := list.Head; node != nil; node = node.Next {
    f(*node)
  }
}










