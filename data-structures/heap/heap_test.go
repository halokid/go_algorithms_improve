package heap

import (
  "fmt"
  "testing"
)

/**
[0 1 3 2 5 6 7 4 9 8]
[1 2 3 4 5 6 7 8 9]
[2 4 3 8 5 6 7 9]
[3 4 6 8 5 9 7]
[4 5 6 8 7 9]
[5 7 6 8 9]
[6 7 9 8]
[7 8 9]
[8 9]
[9]
sorted:     [0 1 2 3 4 5 6 7 8 9]
**/
func TestMinHeap (t *testing.T) {
  h := NewMin()

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
  
  sorted := make([]Int, 0)
  i := 0
  for h.Len() > 0 {
    // fmt.Println(h.data[0])
    // fmt.Println(len(h.data))
    fmt.Println(h.data)
    // fmt.Println(i)
    i++
    
    /**
    [0 1 3 2 5 6 7 4 9 8]
    Extract:    0
    [1 2 3 4 5 6 7 8 9]
    Extract:    1
    [2 4 3 8 5 6 7 9]
    Extract:    2
    [3 4 6 8 5 9 7]
    Extract:    3
    [4 5 6 8 7 9]
    Extract:    4
    [5 7 6 8 9]
    Extract:    5
    [6 7 9 8]
    Extract:    6
    [7 8 9]
    Extract:    7
    [8 9]
    Extract:    8
    [9]
    Extract:    9
    sorted:     []
    **/
      
    // sorted = append(sorted, h.Extract().(Int))
    fmt.Println("Extract:   ", h.Extract().(Int))
  }
  fmt.Println("sorted:    ", sorted)
  
  for i := 0; i < len(sorted)-2; i++ {
    if sorted[i] > sorted[i+1] {
      fmt.Println(sorted)
      t.Error()
    }
  }
  
}