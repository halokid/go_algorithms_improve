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


// Table 是一 链表的数组结构，下标是 int
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


func (ht *HashTable) Del(key string) error {
	index := ht.position(key)
	l := ht.Table[index]
	var val *item				// item 是一个结构,  struct {key,  valuie}
	
	l.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})
	
	if val == nil {
		return nil
	}
	
	ht.Size--
	return l.Remove(val)
}


func (ht *HashTable) ForEach(f func(*item)) {
	for k := range ht.Table {			//遍历 链表的数组
		if ht.Table[k] != nil {			//如果数组的元素不为空
			ht.Table[k].Each( func(node list.Node) {
				f(node.Value.(*item))
			})
		}
	}
}



//返回某个 key应该在 HashTable 的什么位置上
func (ht *HashTable) position(s string) int {
	return hashCode(s) % ht.Capacity
}


func (ht *HashTable) find(i int, key string) (*item, error) {
	l := ht.Table[i]
	var val *item

	/**
	为什么会有  node.Value.(*item) 这种写法呢？？？
	首先我们看看 本身  item  就包含  key: value
	
	我们看看链表的结构是这样的 
	
	type Node struct {          //定义 node 的结构体
  Value   interface{}       // Node 的值
  Prev    *Node             // 前一个 node
  Next    *Node             // 后一个 node
	} node

	========================================================================
	这个 node.Value  就是 item 的 key 的值
	
	而 item 得到这个key之后， 根据某一个算法规则算出 item.value ， 这个就是所谓的键值对照
	
	
	另外 关于  node.Value.(*item) 这个是什么写法呢？？？
	
	我们看  实际上  node.Value 就是  *item.key  
	node.Value.(*item).key  就是要声明一次 这个  node.Value 和  *item.key 是同一块内存里面的东西	
	========================================================================
	上面这段逻辑完全是错误的
	
	实际的情况是这样的
	node.Value  实际上就是储存  item 的,  这个本身就是 item 的值,  node.Value 是 interface{} 类型， 可以储存任何东西， 可以储存结构体
	
	node.Value.(*item) 的意思是， 把 node.Value 作为一个  *item 类型 的意思
	
	**/
	l.Each( func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)				//这里是返回一个 item
		}
	})
	
	if val == nil {
		return nil, errors.New("Not found")
	}
	
	return val, nil
}


//霍纳方法根据长度去计算 hash 字符串,  时间为  L (O(L))
func hashCode(s string) int {
	hash := int32(0)			//int32类型 初始化为0
	
	for i := 0; i < len(s); i++ {
		hash = int32(hash << 5 - hash) + int32(s[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))				//abs 取绝对值
}




















