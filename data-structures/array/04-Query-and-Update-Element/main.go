package main

import (
  "./array"
  "fmt"
)

func main() {
  arr := array.New(20)
  for i := 0; i < 10; i++ {
    arr.AddLast(i)
  }
  fmt.Println(arr)
  fmt.Println("-------------------------------------")

  arr.Add(1, 100)
  fmt.Println(arr)
  fmt.Println("-------------------------------------")

  arr.AddFirst(-1)
  fmt.Println(arr)

  //fmt.Println(arr.String())
}