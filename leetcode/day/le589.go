package main

import "fmt"

// func maxCount(m int, n int, ops [][]int) int {

// 	if len(ops) == 0 {
// 		return m * n
// 	}
// 	maxd := 0
// 	matrix := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		matrix[i] = make([]int, n)
// 	}

// 	for _, pair := range ops {
// 		for i := 0; i < pair[0]; i++ {
// 			for j := 0; j < pair[1]; j++ {
// 				matrix[i][j] += 1
// 				maxd = max(matrix[i][j], maxd)
// 			}
// 		}
// 	}
// 	res := 0
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == maxd {
// 				res++
// 			}
// 		}
// 	}
// 	return res
// }

func maxCount(m int, n int, ops [][]int) int {
	//op_count := len(ops)
	if len(ops) == 0 {
		return m * n
	}
	min_m, min_n := 40001, 40001

	for _, pair := range ops {
		min_m = min(pair[0], min_m)
		min_n = min(pair[1], min_n)
	}
	return min_m * min_n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxCount(40000, 4000, [][]int{}))
}
