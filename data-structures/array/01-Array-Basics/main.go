package main

import "fmt"

func main() {
  var arr [10]int
  for i := 0; i < len(arr); i++ {
    arr[i] = i
  }

  scores := [...]int{100, 99, 66}
  for i := 0; i < len(scores); i++ {
    fmt.Println(scores[i])
  }
  fmt.Println("-----------------------")
  for _, score := range scores {
    fmt.Println(score)
  }

  fmt.Println("-----------------------")
  scores[0] = 96
  for i := 0; i < len(scores); i++ {
    fmt.Println(scores[i])
  }
}

