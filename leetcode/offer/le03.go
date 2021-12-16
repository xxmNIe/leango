package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}
