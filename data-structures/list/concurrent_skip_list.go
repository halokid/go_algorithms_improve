package list

import (
  "math/rand"
  "sync"
  "sync/atomic"
)

type skipList struct {
  level int
  length  int32
  head  *Node
  tail  *Node
  mutex  sync.RWMutex
}

func newSkipList(level int) *skipList {
  head := newNode(0, nil, level)
  var tail *Node
  for i := 0; i < len(head.nextNodes); i++ {
    head.nextNodes[i] = tail
  }

  return &skipList{
    level:  level,
    length: 0,
    head:   head,
    tail:   tail,
  }
}

func (s *skipList) searchWithPreviousNodes(index uint64) ([]*Node, *Node) {
  previousNodes := make([]*Node, s.level)

  currentNode := s.head

  for l := s.level -1; l >= 0; l-- {
    for currentNode.nextNodes[l] != s.tail && currentNode.nextNodes[l].index < index {
      currentNode = currentNode.nextNodes[l]
    }
    previousNodes[l] = currentNode
  }

  if currentNode.nextNodes[0] != s.tail {
    currentNode = currentNode.nextNodes[0]
  }

  return previousNodes, currentNode
}

func (s *skipList) searchWithoutPreviousNodes(index uint64) *Node {
  currentNode := s.head

  // read lock and unlock
  s.mutex.RLock()
  defer s.mutex.RUnlock()

  // iterate from top level to bottom level
  for l := s.level - 1; l >= 0; l-- {
    for currentNode.nextNodes[l] != s.tail && currentNode.nextNodes[l].index < index {
      currentNode = currentNode.nextNodes[l]
    }
  }

  currentNode = currentNode.nextNodes[0]
  if currentNode == s.tail || currentNode.index > index {
    return nil
  } else if currentNode.index == index {
    return currentNode
  } else {
    return nil
  }
}

func (s *skipList) insert(index uint64, value interface{}) {
  s.mutex.Lock()
  defer s.mutex.Unlock()

  previousNodes, currentNode := s.searchWithPreviousNodes(index)

  if currentNode != s.head && currentNode.index == index {
    currentNode.value = value
    return
  }

  // make a new value
  newNode := newNode(index, value, s.randomLevel())

  // adjust pointer, similar to update linked list
  for i := len(newNode.nextNodes) - 1; i >= 0; i-- {
    // Firstlt, new value point to next value
    newNode.nextNodes[i] = previousNodes[i].nextNodes[i]

    // Secondly, previous nodes porint to new value
    previousNodes[i].nextNodes[i] = newNode

    // Finally, in order to release the slice, point to nil
    previousNodes[i] = nil
  }

  // todo: atomic operate
  atomic.AddInt32(&s.length, 1)

  for i := len(newNode.nextNodes); i < len(previousNodes); i++ {
    previousNodes[i] = nil
  }
}

// delete will find the index is existed or not firstly.
// If existed, delete it and update length, otherwise do nothing.
func (s *skipList) delete(index uint64) {
  // Write lock and unlock.
  s.mutex.Lock()
  defer s.mutex.Unlock()

  previousNodes, currentNode := s.searchWithPreviousNodes(index)

  // If skip list length is 0 or could not find value with the given index.
  if currentNode != s.head && currentNode.index == index {
    // Adjust pointer. Similar to update linked list.
    for i := 0; i < len(currentNode.nextNodes); i++ {
      previousNodes[i].nextNodes[i] = currentNode.nextNodes[i]
      currentNode.nextNodes[i] = nil
      previousNodes[i] = nil
    }

    atomic.AddInt32(&s.length, -1)
  }

  for i := len(currentNode.nextNodes); i < len(previousNodes); i++ {
    previousNodes[i] = nil
  }
}

// snapshot will create a snapshot of the skip list and return a slice of the nodes.
func (s *skipList) snapshot() []*Node {
  s.mutex.RLock()
  defer s.mutex.RUnlock()

  result := make([]*Node, s.length)
  i := 0

  currentNode := s.head.nextNodes[0]
  for currentNode != s.tail {
    node := &Node{
      index:     currentNode.index,
      value:     currentNode.value,
      nextNodes: nil,
    }

    result[i] = node
    currentNode = currentNode.nextNodes[0]
    i++
  }

  return result
}

// getLength will return the length of skip list.
func (s *skipList) getLength() int32 {
  return atomic.LoadInt32(&s.length)
}

func (s *skipList) randomLevel() int {
  level := 1
  for rand.Float64() < PROBABILITY && level < s.level {
    level++
  }
  return level
}







