package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	left, right := 0, 0
	for right < n {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
