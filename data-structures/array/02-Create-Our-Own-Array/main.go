package main

import (
  "./array"
  "fmt"
)

func main() {
  a := array.New(7)

  fmt.Println(a)
  fmt.Println(a.GetCapacity(), a.GetSize(), a.IsEmpty())
}
