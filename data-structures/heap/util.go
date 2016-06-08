package heap

type Int int

func (x Int) Less(than Item) bool {
  return x < than.(Int)       // than.(Int)的用法跟之前 hashtable 的那个用法是一样的， 把 item int化
}