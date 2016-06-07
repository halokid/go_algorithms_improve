/**
* Author: JJyy<82049406@qq.com>
* 分离链实现的 哈希表 
* 用 Horner方法  实现哈希化字符串
**/
package hashtable

import (
	"errors"
	"github.com/arnauddri/algorithms/data-structures/linked-list"
	"math"
)

type HashTable struct {
	Table				map[int]*list.List				// list.List 这个用法是怎么回事呢？？是引用 github.com/arnauddri/algorithms/data-structures/linked-list  这个包， 这个包就是list包，然后包的源码里面有一个结构体是List 结构体， 所以这里就是定义一个 Table 为一个map ,  map的key是int类型， 值是 list.List 链表类型, 而且前面有 * ，所以这里是一个链表的指针
	Size				int
	Capacity		int
}

type item struct {
	key			string
	value		interface{}
}


func New(cap int) *HashTable {
	table := make(map[int]*list.List, cap)
	return &HashTable{ Table: table,	Size: 0,	Capacity:	cap }
}

func (ht *HashTable) Get(key string) (interface{}, error) {
	index := ht.position(key)									// position  func
	item, err := ht.find(index, key)					// find func 	
	
	if item == nil {
		return "", errors.New("Not found")
	}
	
	return item.value, err
}


func (ht *HashTable) Put(key, value string) {
	index := ht.position(key)
	
	if ht.Table[index] == nil {
		ht.Table[index] = list.NewList()
	}
	
	item := &item{ key:  key, value:  value }
	
	a, err := ht.find(index, key)
	if err != nil {
		ht.Table[index].Append(item)
		ht.Size++
	} else {
		a.value = value
	}
}



//返回某个 key应该在 HashTable 的什么位置上
func (ht *HashTable) position(s string) int {
	return hashCode(s) % ht.Capacity
}


//霍纳方法根据长度去计算 hash 字符串,  时间为  L (O(L))
func hashCode(s string) int {
	hash := int32(0)
	
	for i := 0; i < len(s); i++ {
		hash = int32(hash << 5 - hash) + int32(s[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}




















