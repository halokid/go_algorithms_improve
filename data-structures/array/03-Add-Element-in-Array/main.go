package main

import (
  "./array"
  "fmt"
)

func main() {
  a := array.New(5)

  a.AddLast(2)
  a.AddLast(3)
  a.AddFirst(1)

  fmt.Println(a)
}
