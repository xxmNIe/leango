package main

import "fmt"

func removeDuplicates(nums []int) int {
	left, right := 0, 0
	n := len(nums)
	for right < n {
		for right < n-1 && nums[right] == nums[right+1] {
			right++
		}
		nums[left] = nums[right]
		left++
		right++
	}
	return left
}

func main() {
	a := []int{}
	fmt.Println(a[:removeDuplicates(a)])
	//fmt.Println(a[:3])
}
