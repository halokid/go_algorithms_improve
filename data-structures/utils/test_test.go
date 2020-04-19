package utils

import "testing"

func TestComm(t *testing.T) {
  var a interface{}
  var b interface{}

  a = "a"
  b = "7"
  c := a.(string) > b.(string)

  t.Log(a)
  t.Log(b)
  t.Log(c)

}

