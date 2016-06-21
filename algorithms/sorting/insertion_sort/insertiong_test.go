package insertion

//  go test -v -bench=".*"

import (
  "fmt"
  "github.com/arnauddri/algorithms/algorithms/sorting/utils"
  "testing"
)

func TestInsertionSort(t *testing.T) {
  list := utils.GetArrayOfSize(100)
  
  sort(list)
  
  for i := 0; i < len(list) - 2; i++ {
    if list[i] > list[i+1] {
      fmt.Println(list)
      t.Error()
    }
  }
}

func benchmarkInsertionSort(n int, b *testing.B) {
  list := utils.GetArrayOfSize(n)
  for i := 0; i < b.N; i++ {
    sort(list)
  }
}

func BenchmarkInsertionSort100 (b *testing.B) {
  benchmarkInsertionSort(100, b)
}



