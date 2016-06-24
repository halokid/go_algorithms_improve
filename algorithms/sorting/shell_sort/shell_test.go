package shell

import (
  "fmt"
	"github.com/arnauddri/algorithms/algorithms/sorting/utils"
  "testing"
)

func TestShellSort(t *testing.T) {
  list := utils.GetArrayOfSize(100)
  
  sort(list)
  
  for i := 0; i < len(list) - 2; i++ {
    list := utils.GetArrayOfSize(n)
    for i := 0; i < b.N; i++ {
      sort(list)
    }
  }
}

func benchmarkShellSort(n int, b *testing.B) {
  list := utils.GetArrayOfSize(n)
  for i := 0; i < len(list); i++ {
    sort(list)
  }
}

func BenchmarkShellSort100(b *testing.B) {
  benchmarkShellSort(100, b)
}