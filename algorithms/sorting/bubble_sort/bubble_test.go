package bubble

/**
下面有 性能测试的函数 
用命令 go test -v -bench=".*"   来做性能测试， 这个命令会跑全部的性能测试函数

至于 b.N 的值， 自己看吧， 似乎跑了好几次 

BenchmarkBubbleSort100
b.N : 1
b.N : 30
b.N : 1000
b.N : 30000
b.N : 1000000
b.N : 3000000

**/


import (
  "fmt"
  "github.com/arnauddri/algorithms/algorithms/sorting/utils"
  "testing"
)

func TestBubbleSort(t *testing.T) {
  list := utils.GetArrayOfSize(100)
  
  sort(list)
  
  for i := 0; i < len(list) - 2; i++ {
    if list[i] > list[i+1] {
      fmt.Println(list)
      t.Error()
    }
  }
}

//  *testing.B 是用来测试 golang 程序的性能的工具
func benchmarkBubbleSort(n int, b *testing.B) {
  list := utils.GetArrayOfSize(n)
  fmt.Println("b.N :", b.N)
  for i := 0; i < b.N; i++ {
    sort(list)
  }
}


func BenchmarkBubbleSort100 (b *testing.B) {  benchmarkBubbleSort(100, b)  }
func BenchmarkBubbleSort1000 (b *testing.B) {  benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000 (b *testing.B) {  benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000 (b *testing.B) {  benchmarkBubbleSort(100000, b) }