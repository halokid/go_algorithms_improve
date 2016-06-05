/**
* copy for  https://github.com/arnauddri/algorithms
* add chinese explain
* JJyy<82049406@qq.com>
* 代码只是一个结构化的过程， 要完全理解还得要有抽象意识，多空间意识
**/

package bst

import ()

//定义节点的类型
type Node struct {
	Value   int 						//节点的值				
	Parent  *Node						//节点的父节点指针地址	
	Left	  *Node						//节点的左节点的指针地址
	Right   *Node 					//节点的右节点的指针地址
}

//返回一个新的节点， 并且定义了节点的值
func NewNode(i int) *Node {
	return &Node{Value: i}
}


//比较两个节点的值的大小， 大于返回 1， 小于返回 -1， 等于返回0
func (n *Node) Compare(m *Node) int {
	if n.Value < m.Value {
		return -1
	} else if n.Value > m.Value {
		return 1
	} else {
		return 0
	}
}

//定义树的结构
type Tree struct {
	// head 表示的是树的顶点，  或者在寻找线路上的前一个点
	Head   *Node				//树的起点，起点的指针值， 假如在寻找树的时候，这个值就是寻找的前一个节点
	Size   int					//树的体积的大小（ 有多少个节点 ） 
}

//返回一棵树， 而且定义了树的顶点位置， 定义了树的 size
func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{ Head: n, Size: 1 }			// Head 这里肯定是一个 Node 类型
}


/**
 这个 Insert  的过程是这样的
 首先 假如 树为空的话， 就把节点n设置为顶点
 然后假如继续加入其他点的话， 那么就判断跟顶点的大小，然后按照 左节点小右节点大的原则一直往下加下去，形成一颗树， 所以此函数可以递归处理, 无论加入的点是左节点还是右节点，函数里面自会判断
**/

func (t *Tree) Insert(i int) {
	n := &Node{ Value: i }			//定义一个 Node类型的 n 变量, 把参数 i 赋予这个节点(node)的 value 属性
	if t.Head == nil {			//如果这个树没有顶点的话，那么就定义 n 为这个树的顶点
		t.Head = n
		t.Size++
		return
	}
	
	h := t.Head			//如果这个树本来就有顶点，那么取出来 赋予 h 变量
	
	for {														//无限for， 直到 break 才跳出for
		if  n.Compare(h) == -1 {			//如果 n节点的值小于 h节点的值
			if h.Left == nil {					//如果节点 h 没有左节点
				h.Left = n								//因为 n节点的值小于 h节点的值，那么把 n节点作为 h节点的左节点
				n.Parent = h							//顺理把 h节点定义为 n节点的父节点
				break
			} else {										//如果 h 已经有了左节点, 把 h的左节点指向自己，证明它没有左节点 ??
				h = h.Left
			}
		} else {											//如果 n节点的值大于 h节点的值
			if h.Right == nil {					//如果 h节点没有右节点
				h.Right = n								//那么就定义 h的右节点为 n节点
				n.Parent = h							//定义 n节点的父节点就是 h节点
				break
		  } else {
				h = h.Right							//如果 h已经有了右节点，那么就继续寻找下去， 移动到 h的右节点继续寻找合适的位置 insert 进去， 所以这里把 h.Right 的 点的指针赋予 h，以便整个函数能递归
			}
		}
	}
	t.Size++				//树的体积增加
}


//寻找树的某个节点， 返回一个 节点的指针
func (t *Tree) Search(i int) *Node {
	h := t.Head						//寻找路线上的前一个节点
	n := &Node{ Value: i }			//要寻找的节点n， 这个节点n 的值就是参数里面的 i
	
	for h != nil {						//如果树有顶点
		switch h.Compare(n) {			//线比较顶点h 和要寻找的节点 n的大小
		case -1:				//如果 顶点h 小于 要寻找的节点n
			h = h.Right		// 则 寻找的路线应该向右走， 那么h就向右走， 把 h的右节点继续赋予 h
		case 1:					//如果 点h 大于 要寻找的节点 n
			h = h.Left		//则 寻找的路线应该向左走， 那么h 就向左走， 把 h的左节点继续赋予 h
		case 0:
			return h			//如果 h == n ， 那么就是寻找到了 节点n， 这个时候直接返回h
		default:
			panic("	Node not found")
			
		}
	}
	panic("Node not found")
}


//删除某个节点
func (t *Tree) Delete(i int) bool {
	var parent *Node			//定义一个父节点
	
	h := t.Head				//树的顶点,寻找是从顶点开始的, 当继续向下寻找的时候，h就是寻找线路上的某个点，寻找状态点
	n := &Node{ Value: i }		//要寻找的节点 n
	
	for h != nil {			//如果顶点存在
		switch n.Compare(h) {
		case -1:			//如果 节点n的值小于 节点h的值
			parent = h		//继续向下寻找， 父节点就是 节点h
			h = h.Left		//因为 n 小于 h， 所以向 左下 寻找， 把 h移动到 h的左节点
		case 1:
			parent = h		//继续向下寻找， 父节点就是 节点h
			h = h.Right		//因为 n 大于 h， 所以向 右下 寻找， 把 h移动到 h的右节点
		case 0:					//如果 节点n 等于 节点 h, 那么就是要删除 节点h了
		
			/****  下面的逻辑都是删除 节点h 的  ***************/
			
			//下面的逻辑是移动节点 h 的左右 子节点 及下面的子树的
			if h.Left !=	nil {		//	如果 h节点的左节点不为空
			
				//下面三句是用 h节点的左节点 去代替h节点
				h.Value = h.Left.Value
				h.Left = h.Left.Left
				h.Right = h.Left.Right
			
				/**
				* 如果是左右节点都有的情况下,就用子左节点去代替父节点， 然后以子右节点为顶点，重新再建一颗新的树， 就是subTree
				**/
				
				right := h.Right			//先把 h节点的右节点赋予变量 right
				if right != nil {			//如果 h节点的右节点也不为空，这种情况是 h节点的左右两个节点都是不为空的情况
					subTree := &Tree{Head: h}				//以右节点为顶点建一颗子树 ???
					IterOnTree(right, func(n *Node) {			//查看 IterOnTree 函数的逻辑
							subTree.Insert(n.Value)
					})
				}
				t.Size--
				return true
			}
			
			if h.Right != nil {				//如果 h节点的右节点不为空，排除了上面左节点不为空的情况了，因为上面已经返回了, 这种情况就是这个 h节点只有右节点的情况， 比较好处理
				h.Value =  h.Right.Value
				h.Left = h.Right.Left
				h.Right = h.Right.Right
				
				t.Size--
				return true
			}
			
			if parent == nil {			//一开始就定义了 parent， 直到找到要寻找的节点， parent就是 h节点，也就是寻找路线上的上一个节点, 这种情况就是树为空 或者 树只有一个节点, 而且这个节点就是要寻找的节点， 那么去就掉这个唯一的节点， 则树为空
				t.Head = nil
				t.Size--
				return true
			}
			
			if parent.Left == n {			//下面这种情况都是树原来有三个节点的情况，左右子节点哪个是要寻找的节点，分别删除哪个
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			t.Size--
			return true
		}
	}			//END FOR
	return false				//找不到要寻找的节点	
}


func IterOnTree(n *Node, f func(*Node)) bool {
	if n == nil {		//？？？
		return true
	}
	if !IterOnTree(n.Left, f) {
		return false
	}
	
	f(n)
	
	return IterOnTree(n.Right, f)
}




