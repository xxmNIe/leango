package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	if m == 1 && n == 1 {
		if matrix[0][0] == target {
			return true
		} else {
			return false
		}
	}

	for i, j := 0, n-1; i < m && j >= 0; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 18))
}
