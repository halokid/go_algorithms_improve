package matrix

import (
  "testing"
)

func TestTrace(t *testing.T) {
  a := []float64{1, 2, 3, 4}      //定义一个 slice
  A := MakeMatrix(a, 2, 2)
  
  if A.trace() != 5 {
    t.Error()
  }
}