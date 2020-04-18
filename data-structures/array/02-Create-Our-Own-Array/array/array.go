package array

type Array struct {
  data  []int       // 数组的容量
  size  int         // 数组的元素个数
}

func New(capacity int) *Array {
  return &Array{
    data: make([]int, capacity),
  }
}

func (a *Array) GetCapacity() int {
  return len(a.data)
}

func (a *Array) GetSize() int {
  return a.size
}

func (a *Array) IsEmpty() bool {
  // 返回数组是否为空
  return a.size == 0
}



