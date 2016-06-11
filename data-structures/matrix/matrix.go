//package matrix provides some methods to use with matrix
package matrix

import (
  "errors"
)

type Matrix struct {
  rows      int           //number of rows   行数
  cols      int           //number of columns     列数
  Elements  []float64     //matrix stored as a flat array:  Aij = Elements[i*step + j], i 是第几行， j是第几列
  step      int           //offset between  rows, 第几列的转移量,也就是说从第一行 移动 第几列 
}


func MakeMatrix(Elements []float64, rows, cols int) *Matrix {
  A := new(Matrix)      //新建一个结构体的写法
  A.rows = rows
  A.cols = cols
  A.step = cols
  A.Elements = Elements
  
  return A
}

func (A *Matrix) CountRows() int {
  return A.rows
}

func (A *Matrix) CountCols() int {
  return A.cols
}

func (A *Matrix) GetElm(i int, j int) float64 {
  return A.Elements[i*A.step + j]
}

func (A *Matrix) SetElm(i int, j int, v float64) {      // 这里set 元素的时候就已经是这样定的了  i*A.step + j, 这里的 i, j 跟 行 和 列 都没关系???有， i就是行， j就是列
  A.Elements[i*A.step + j] = v
}

//取得对角线的元素形成slice 并返回
func (A *Matrix) diagonalCopy() []float64 {
  diag := make([]float64, A.cols)
  
  for i := 0; i < len(diag); i++ {
    diag[i] = A.GetElm(i, i)      //对角线的时候 i, j 的值都一样，都是i
  }
  return diag
}


//把 A 矩阵拷贝到 B矩阵
func (A *Matrix) copy() *Matrix {
  B := new(Matrix)
  B.rows = A.rows
  B.cols = A.cols
  B.step = B.step

  //真个矩阵 的元素的数量就是 矩阵的 长乘与宽， 但是矩阵下标的设计并不是一个 二维 数组， 而是一个通过某个规则设置的key访问的数组， 这个规矩就是  i*step + j, i代表行数， j代表列数
  B.Elements = make([]float64, A.cols*A.rows)
  
  for i := 0; i < A.rows; i++ {
    for j := 0; j < A.cols; j++ {
      B.Elements[i*A.step + j] = A.GetElm(i, j)
    }
  }
  return B
}

//追踪矩阵就是一条斜线的元素的值相加？？ 是的， 具体看 test 文件
func (A *Matrix) trace() float64 {
  var tr float64 = 0

  for i := 0; i < A.cols; i++ {
    tr += A.GetElm(i, i)
  }
  return tr
}


/**
下面的函数都很好理解，基本上都是对 矩阵的元素进行 加减乘除 的操作的
**/

func (A *Matrix) add(B *Matrix) error {
  if A.cols != B.cols || A.rows != B.rows {
    return errors.New("Wrong input sizes")
  }
  
  for i := 0; i < A.rows; i++ {
    for j := 0; j < A.cols; j++ {
      A.SetElm(i, j, (A.GetElm(i, j) + B.GetElm(i, j)) )
    }
  }
  
  return nil
}


func (A *Matrix) subtract(B *Matrix) error {
  if A.cols != B.cols || A.rows != B.rows {
    return errors.New("Wrong input sizes")
  }
  
  for i := 0; i < A.rows; i++ {
    for j := 0; j < A.cols; j++ {
      A.SetElm(i, j, (A.GetElm(i, j,) - B.GetElm(i, j)) )
    }
  }
  
  return nil
}

//让矩阵的元素的数值 增大 a 倍的函数
func (A *Matrix) scale(a float64) {
  for i := 0; i < A.rows; i++ {
    for j := 0; j < A.cols; j++ {
      A.SetElm(i, j, a * A.GetElm(i, j))
    }
  }
}



func Add(A *Matrix, B *Matrix) *Matrix {
  result := MakeMatrix(make([]float64, A.cols * A.rows), A.cols, A.rows)
  
  for i := 0; i < A.rows; i++ {
    for j := 0; j < A.cols; j++ {
      result.SetElm(i, j, (A.GetElm(i, j) + B.GetElm(i, j)) )
    }
  }
  
  return result
}


func Substract(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			result.SetElm(i, j, A.GetElm(i, j)-B.GetElm(i, j))
		}
	}

	return result
}



func Multiply(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			sum := float64(0)
			for k := 0; k < A.cols; k++ {
				sum += A.GetElm(i, k) * B.GetElm(k, j)
			}
			result.SetElm(i, j, sum)
		}
	}

	return result
}







