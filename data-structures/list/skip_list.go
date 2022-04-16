/*
Package ConcurrentSkipList provide an implementation of skip list. It's thread-safe in concurrency and high performance.
*/
package  list

import (
  "errors"
  "math"
  "sync/atomic"

  "github.com/OneOfOne/xxhash"
)

// Comes from redis's implementation.
// Also you can see more detail in William Pugh's paper <Skip Lists: A Probabilistic Alternative to Balanced Trees>.
// The paper is in ftp://ftp.cs.umd.edu/pub/skipLists/skiplists.pdf
const (
  MAX_LEVEL   = 32
  PROBABILITY = 0.25
  SHARDS      = 32
)

// shardIndex is used to indicate which shard a given index belong to.
var shardIndexes = make([]uint64, SHARDS)

// init will initialize the shardIndexes.
func init() {
  var step uint64 = 1 << 59 // 2^64/SHARDS
  var t uint64 = math.MaxUint64

  for i := SHARDS - 1; i >= 0; i-- {
    shardIndexes[i] = t
    t -= step
  }
}

// ConcurrentSkipList is a struct contains a slice of concurrent skip list.
type ConcurrentSkipList struct {
  skipLists []*skipList
  level     int
}

// NewConcurrentSkipList will create a new concurrent skip list with given level.
// Level must between 1 to 32. If not, will return an error.
// To determine the level, you can see the paper ftp://ftp.cs.umd.edu/pub/skipLists/skiplists.pdf.
// A simple way to determine the level is L(N) = log(1/PROBABILITY)(N).
// N is the count of the skip list which you can estimate. PROBABILITY is 0.25 in this case.
// For example, if you expect the skip list contains 10000000 elements, then N = 10000000, L(N) ≈ 12.
// After initialization, the head field's level equal to level parameter and point to tail field.
func NewConcurrentSkipList(level int) (*ConcurrentSkipList, error) {
  if level <= 0 || level > MAX_LEVEL {
    return nil, errors.New("invalid level, level must between 1 to 32")
  }

  skipLists := make([]*skipList, SHARDS, SHARDS)
  for i := 0; i < SHARDS; i++ {
    skipLists[i] = newSkipList(level)
  }

  return &ConcurrentSkipList{
    skipLists: skipLists,
    level:     level,
  }, nil
}

// Level will return the level of skip list.
func (s *ConcurrentSkipList) Level() int {
  return s.level
}

// Length will return the length of skip list.
func (s *ConcurrentSkipList) Length() int32 {
  var length int32
  for _, sl := range s.skipLists {
    length += sl.getLength()
  }

  return length
}

// Search will search the skip list with the given index.
// If the index exists, return the value and true, otherwise return nil and false.
func (s *ConcurrentSkipList) Search(index uint64) (*Node, bool) {
  sl := s.skipLists[getShardIndex(index)]
  if atomic.LoadInt32(&sl.length) == 0 {
    return nil, false
  }

  result := sl.searchWithoutPreviousNodes(index)
  return result, result != nil
}

// Insert will insert a value into skip list. If skip has these this index, overwrite the value, otherwise add it.
func (s *ConcurrentSkipList) Insert(index uint64, value interface{}) {
  // Ignore nil value.
  if value == nil {
    return
  }

  sl := s.skipLists[getShardIndex(index)]
  sl.insert(index, value)
}

// Delete the node with the given index.
func (s *ConcurrentSkipList) Delete(index uint64) {
  sl := s.skipLists[getShardIndex(index)]
  if atomic.LoadInt32(&sl.length) == 0 {
    return
  }

  sl.delete(index)
}

// ForEach will create a snapshot first shard by shard. Then iterate each node in snapshot and do the function f().
// If f() return false, stop iterating and return.
// If skip list is inserted or deleted while iterating, the node in snapshot will not change.
// The performance is not very high and the snapshot with be stored in memory.
func (s *ConcurrentSkipList) ForEach(f func(node *Node) bool) {
  for _, sl := range s.skipLists {
    if sl.getLength() == 0 {
      continue
    }

    nodes := sl.snapshot()
    stop := false
    for _, node := range nodes {
      if !f(node) {
        stop = true
        break
      }
    }

    if stop {
      break
    }
  }
}

// Sub will return a slice the skip list who starts with startNumber.
// The startNumber start with 0 as same as slice and maximum length is skip list's length.
func (s *ConcurrentSkipList) Sub(startNumber int32, length int32) []*Node {
  // Ignore invalid parameter.
  if startNumber > s.Length() || startNumber < 0 || length <= 0 {
    return nil
  }

  var result []*Node
  var position, count int32 = 0, 0
  for _, sl := range s.skipLists {
    if l := sl.getLength(); l == 0 || position+l <= startNumber {
      position += l
      continue
    }

    nodes := sl.snapshot()
    for _, node := range nodes {
      if position < startNumber {
        position++
        continue
      }

      if count == length {
        break
      }

      result = append(result, node)
      count++
    }

    if count == length {
      break
    }
  }

  return result
}

// Locate which shard the given index belong to.
func getShardIndex(index uint64) int {
  result := -1
  for i, t := range shardIndexes {
    if index <= t {
      result = i
      break
    }
  }

  return result
}

// Hash will calculate the input's hash value using xxHash algorithm.
// It can be used to calculate the index of skip list.
// See more detail in https://cyan4973.github.io/xxHash/
func Hash(input []byte) uint64 {
  h := xxhash.New64()
  h.Write(input)
  return h.Sum64()
}


