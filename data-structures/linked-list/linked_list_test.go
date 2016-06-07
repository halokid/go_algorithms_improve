package list

import (
  "fmt"
  "testing"
)

func TestLinkList(t *testing.T) {
  //test Prepend/Get
  l := NewList()
  
  l.Prepend(NewNode(1))
  l.Prepend(NewNode(2))
  l.Prepend(NewNode(3))
  
  zero := *slice(l.Get(0))[0].(*Node).Value.(*Node)
  one := *slice(l.Get(1))[0].(*Node).Value.(*Node)
  two := *slice(l.Get(2))[0].(*Node).Value.(*Node)
  
  if zero != *NewNode(3) || one != *NewNode(2) || two != *NewNode(1) {
    fmt.Println(*one.Value.(*Node), *NewNode(2))
    fmt.Println(zero.Value)
    fmt.Println(one.Value)
    fmt.Println(two.Value)
    t.Error()
  }
  
  // fmt.Println(*one.Value.(*Node), *NewNode(2))
  fmt.Println(zero.Value)
  fmt.Println(one.Value)
  fmt.Println(two.Value)
}



func slice(args ...interface{}) []interface{} {
  return args
}
