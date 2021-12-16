package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {

	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	for x, y := 0, n-1; x < m && y >= 0; {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func main() {
	// fmt.Println(findNumberIn2DArray([][]int{{1, 4, 7, 11, 15},
	// 	{2, 5, 8, 12, 19},
	// 	{3, 6, 9, 16, 22},
	// 	{10, 13, 14, 17, 24},
	// 	{18, 21, 23, 26, 30}}, 18))
	fmt.Println(findNumberIn2DArray([][]int{{}}, 18))
}
