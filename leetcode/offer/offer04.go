package main

import "fmt"

//func findNumberIn2DArray(matrix [][]int, target int) bool {
//	n,m:=len(matrix),0
//	if n>0 {
//		m = len(matrix[0])
//	}else {
//		return false
//	}
//	if m< 1 {
//		return false
//	}
//	left,right :=0,m-1
//	f,mid := matrix[0],0
//	for left <= right {
//		mid = (left+right)/2
//		if f[mid] == target {
//			return true
//		}else if f[mid] > target {
//			right = mid-1
//		}else {
//			left = mid+1
//		}
//	}
//	col := 0
//	if right >=0 {
//		col = right
//	}
//	left,right = 0,n-1
//	for left <=right {
//		mid = (left+right)/2
//		if matrix[mid][col] == target {
//			return true
//		}else if matrix[mid][col] > target {
//			right = mid-1
//		}else {
//			left = mid+1
//		}
//	}
//	return false
//}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix)==0 || len(matrix[0])==0 {
		return false
	}
	n,m :=len(matrix),len(matrix[0])
	for i,j :=0,m-1;i<n&& j>=0;{
		if matrix[i][j] == target{
			return true
		}else if matrix[i][j] > target {
			j-=1
		}else {
			i+=1
		}
	}
	return  false
}



func main() {
	//m := [][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}
	m :=[][]int{{-5}}
	//m := [][]int{
	//	{1,2,3,4,5},
	//	{6,7,8,9,10},
	//	{11,12,13,14,15},
	//	{16,17,18,19,20},
	//	{21,22,23,24},
	//}
	fmt.Println(findNumberIn2DArray(m,5))
}