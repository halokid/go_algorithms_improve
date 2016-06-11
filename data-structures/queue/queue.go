package queue

import "sync"

type Queue struct {
  queue []interface{}     //队列的实际值
  len     int             //队列的长度
  lock    *sync.Mutex     //队列的锁
}


func New() *Queue {
  queue := &Queue{}
  queue.queue = make([]interface{}, 0)
  queue.len = 0
  queue.lock = new(sync.Mutex)
  
  return queue
}


func (q *Queue) Len() int {
  // q.lock.Lock()
  // defer q.lock.Unlock()
  
  return q.len
}

func (q *Queue) isEmpty() bool {
  q.lock.Lock()
  defer q.lock.Unlock()
  
  return q.len == 0
}

//从队列中弹出首位元素, 把队列长度减一， 并发挥弹出的元素
func (q *Queue) Shift() (el interface{}) {    //注意这里返回的是   interface{} 类型
  el, q.queue = q.queue[0], q.queue[1:]
  q.len--
  return
}

//把元素推进队列, append 从 索引0  开始组成 slice
func (q *Queue) Push(el interface{}) {
  q.queue = append(q.queue, el)
  q.len++
  
  return
}


//返回队列的首位元素
func (q *Queue) Peek() interface{} {
  q.lock.Lock()
  defer q.lock.Unlock()
  
  return q.queue[0]
}  





