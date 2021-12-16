package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	tops := make([][]int, 2)
	res := 0
	for i := 0; i < 2; i++ {
		tops[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		topx, topy := 0, 0
		for j := 0; j < n; j++ {
			topx = max(topx, grid[i][j])
			topy = max(topy, grid[j][i])
		}
		tops[0][i] = topx
		tops[1][i] = topy
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += min(tops[0][i], tops[1][j]) - grid[i][j]
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	a := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(a))
}
