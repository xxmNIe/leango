package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		h := min(height[l], height[r])
		tmp := h * (r - l)
		res = max(res, tmp)
		if height[l] < height[r] {
			l++
		} else {
			r--
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
	fmt.Println(maxArea([]int{1, 1}))
}
