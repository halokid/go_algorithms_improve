/**
*  
* Author: JJyy<82049406@qq.com>
* 图的数据结构的代码 
**/

package graph

import (
  "fmt"
)

type VertexId unit      //Vertex 顶点， 定义顶点的值, 定义图的各个点的内省为 unit

type Vertices []VertexId        //把所有的顶点都放在一个数组里面

type Edge struct {      //定义 边 的结构
  From   VertexId       //来源的点
  To     VertexId       //要去的点 
}

type graph struct {       //定义一个图的结构
  edges         map[VertexId]map[VertexId]int       //图的边， 每条边包含从哪里来， 到哪里去 from & to
  edgesCount    int       //图的边的数量
  isDirected    bool      //是否是定向图
}


/**
  golang 的通道一定是 某一个已经定义了的类型
**/

type EdgesIterable interface {      //定义一个 interface接口类型
  //边的迭代函数
  EdgesIter() <-chan Edge           //以 EdgesIter函数的返回类型，然后写进 Edge 通道, Edge为一个结构体类型
}

type VerticesIterable interface {     //定义一个 interface接口类型
  //点的迭代函数
  VerticesIter() <-chan VertexId      //以 VerticesIter 函数的返回类型， 然后写进 VertexId 通道, VertexId 为一个类型,  实际上是 uint 的别称
}

//这个函数是把  边 写进一个通道去
func (g *graph) EdgesIter() <-chan Edge {
  ch := make(chan Edge)
  
  go func() {
    for from, connectedVertices := range g.edges {    // A-FOR
      for to, _ := range connectedVertices {    // B-FOR 
        if g.isDirected {       //如果图是某个点直达某点的话， 也就是定向的， 那么就不用计较每点的大小了
          ch <- Edge{ from, to }  //把定义好的结构体写入通道， 你可以把他想象成一个等待处理的通道，而且这个通道是可以并行处理的， 你可以想象成放在一个篮子里面，但是不是放进丢列里面
        }  else {               // 如果不是定向查找的话，要判断 点的大小
          if from < to {
            ch <- Edge{ from, to }    //判断 from 小于 to 的时候，写进通道
          }
        }
      }   // END  A-FOR
    }   // END B-FOR
    close(ch)   //关闭通道
  } ()
  return ch     //返回通道
}


//这个函数是把 点 写进一个通道里面去
func (g *graph) VerticesIter() <-chan VertexId {
  ch := make(chan VertexId)     //首先初始化 ch 通道
  
  go func() {
    for vertex, _ := range g.edges {    //循环把每一个 点 都放进通道 
      ch <- vertex
    }
    close(ch)
  } ()
  
  return ch             //返回通道
}


//检查 点 是否存在
func (g *graph) CheckVertex(vertex VertexId) bool {
  _, exists := g.edges[vertex]
  
  return exists
}

//定义一个点, 相当于linux 的touch
func (g *graph) TouchVertex(vertex VertexId) {
  if _, ok := g.edges[vertex]; !ok {      // !ok 表示点不存在，则可以定义创建
    g.edges[vertex] = make(map[VertexId]int)
  }
}


//添加一个点
func (g *graph) AddVertex(vertex VertexId) error {
  i, _ := g.edges[vertex]
  if i != nil {
    return errors.New("Vertex already exists")
  }
  
  g.edges[vertex] = make(map[VertexId]int)    // g.edges 是 二维 map
  
  return nil
}


//删除一个点
func (g *graph) RemoveVertex(vertex VertexId) error {
  if !g.IsVertex(vertex) {
    return errors.New("Unknow vertex")
  }
  
  delete(g.edges, vertex)     //删除某个点
  
  for _, connectedVertices := range g.edges {     //删除 
    delete(connectedVertices, vertex)
  }
  
  return nil
}


//判断某个点 是否是 图里面的点
func (g *graph) IsVertex(vertex VertexId) (exist bool) {
  _, exist = g.edges[vertex]
  
  return
}

//获取 图 的点的数量
func (g *graph) VertexCount() int {
  return len(g.edges)
}


// 给图加一条边
func (g *graph) AddEdge(from, to VertexId, weight int) error {
  if from == to {
    return errors.New("Cannot add self lopp")
  }
  
  if !g.CheckVertex(from) || !g.CheckVertex(to) {
    return errors.New("Vertices donnot exist")
  }
  
  i, _ := g.edges[from][to]     //得到 从 from 到 to 的点
  j, _ := g.edges[to][from]     //得到 从 to 到 from 的点

  //如果这两个值其中一个存在， 则证明边已经存在, 有可能是 A -> B, 或者 B -> A， 都是一样的
  if i > 0 || j > 0  {
    return errors.New("Edge already defined")
  }
  
  g.TouchVertex(from)
  g.TouchVertex(to)
  
  g.edges[from][to] = weight      //二维map的值就是这条边的权重
  
  if !g.isDirected {      //如果不是定向图， 则反过来权重也是一样， 定向图是没有把线反过来这个说法的
    g.edges[to][from] = weight
  }

  //边的数量增加一
  g.edgeCount++

  return nil
}


//删除一条边
func (g *graph) RemoveEdge(from, to VertexId) error {
  i, _ := g.edges[from][to]
  j, _ := g.edges[to][from]
  
  if i == -1 || j == -1 {
    return errors.New("Edge doesn't exist")
  } 
  
  g.edges[from][to] = -1
  
  if !g.isDirected {
    g.edges[to][from] = -1
  }
  
  g.edgeCount--
  
  return nil
}


//判断是否为 边
func (g *graph) IsEdge(from, to VertexId) bool {
  // 获取边的起点， 定义为 connected, 得到一维 map 的位置
  connected, ok := g.edges[from]
  
  if !ok {
    return false
  }
  
  // to 就是一维 map 的二维的 key
  weight := connected[to]     //根据二维的key 得出 边的权重
  return weight > 0 
}

func (g *graph) Order() int {
  return len(g.edges)
}


func (g *graph) GetEdge(from, to VertexId) int {
  return g.edges[from][to]
}


func (g *graph) GetNeighbour (vertex VertexId) VerticesIterable {
  iterator := func() <-chan VertexId {
    ch := make(chan VertexId)
    
    go func() {
      if connected, ok := g.edges[vertex]; ok {
        for VertexId, _ := range connected {
          ch <- VertexId
        }
      }
      close(ch)
    }() 
    return ch
  }
  
  return VerticesIterable( &vertexIterableHelper{ iterFunc: iterator} )
}


type vertexIterableHelper struct {
  iterFunc func() <-chan VertexId
}

func (helper *vertexIterableHelper) VerticesIter() <-chan VertexId {
  return helper.iterFunc()
}















