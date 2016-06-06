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
  
  nextNode, _ := l.Get(index)
  
}
















