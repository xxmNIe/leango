package main

import "fmt"

/*
矩阵 前缀
 */
type NumMatrix struct {
	matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	xn := len(matrix)
	yn :=0
	if matrix[0] !=nil {
		yn= len(matrix[0])
	}
	sum := NumMatrix{make([][]int,xn+1)}
	for i:=0;i<xn+1;i++ {
		sum.matrix[i] = make([]int,yn+1)
	}
	for i:=1;i<xn+1;i++{
		for j :=1;j<yn+1;j++ {
			sum.matrix[i][j] = sum.matrix[i-1][j]+sum.matrix[i][j-1]-sum.matrix[i-1][j-1]+matrix[i-1][j-1]
		}
	}
	return sum
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrix[row2+1][col2+1]-this.matrix[row1][col2+1]-this.matrix[row2+1][col1]+this.matrix[row1][col1]
}

func main() {
	m :=[][]int{

	{3, 0, 1, 4, 2},
	{5, 6, 3, 2, 1},
	{1, 2, 0, 1, 5},
	{4, 1, 0, 1, 7},
	{1, 0, 3, 0, 5},
	}
	//m2 :=[][]int{
	//	{-4,-5},
	//}

	obj := Constructor(m)
	fmt.Println(obj.SumRegion(2,1,4,3))
	//fmt.Println(obj.SumRegion(0,1,0,1))
}