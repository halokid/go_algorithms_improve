package hashtable

import (
  "fmt"
  "testing"
)

func TestHt(t *testing.T) {
  ht := New(1000)
  ht.Put("foo", "bar")
  ht.Put("fiz", "buzz")
  ht.Put("bruce", "wayne")
  ht.Put("peter", "parker")
  ht.Put("clark", "kent")
  
  //test sample Get
  val , err := ht.Get("foo")
  if err != nil || val != "bar" || ht.Size != 5 {
    fmt.Println(val, err)
    t.Error()
  }
  
  ht.Put("peter", "bob")
  //test "peter" has been updated

  index := ht.position("foo")
  fmt.Println("foo index: ", index)
    
  var it *item
  it, _ = ht.find(index, "foo")
  // fmt.Println(it)
  fmt.Printf("it: %v\n", it)
  
  isPoint := fmt.Sprintln("%p", it)
  fmt.Println(isPoint)
} 