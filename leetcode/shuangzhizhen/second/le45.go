package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	// if n < 3 {
	// 	return 1
	// }
	step := 0
	end := 0
	maxLen := -1
	for i := 0; i < n-1; i++ {
		if i+nums[i] > maxLen {
			maxLen = i + nums[i]
		}
		if i == end {
			end = maxLen
			step++
		}

	}
	return step
}

func main() {
	fmt.Println(jump([]int{0}))
}
