package list

type Node struct {
  index     uint64
  value     interface{}
  nextNodes   []*Node
}

func newNode(index uint64, value interface{}, level int) *Node {
  return &Node{
    index:     index,
    value:     value,
    nextNodes: make([]*Node, level, level),
  }
}

func (n *Node) Index() uint64 {
  return n.index
}

func (n *Node) Value() interface{} {
  return n.value
}





