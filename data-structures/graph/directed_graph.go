package graph

type DirGraph struct {      //定义一个  定向图 的结构体
  graph
}


func NewDirected() *DirGraph {
  return &DirGraph {
    graph {
      edgeCount:      0,
      edges:          make( map[VertexId]map[VertexId]int ),
      isDirected:     true,
    }
  }
}



func (g *graph) GetPredecessors(vertex VertexId) VerticesIterable {
  iterator := func() <-chan VertexId {
    ch := make(chan VertexId)
    
    go func() {
      if connected, ok := g.edges[vertex]; ok {
        for VertexId, _ := range connected {
          if g.IsEdge(VertexId, vertex) {
            ch <- VertexId
          }
        }
      }
      close(ch)
    } ()
    return ch
  }
  
  return VerticesIterable( &vertexIterableHelper{ iterFunc: iterator} )
}


func (g *graph) GetSuccessors(vertex VertexId) VerticesIterable {
  
}


