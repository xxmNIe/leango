package main

import "fmt"

func smallestEqual(nums []int) int {
	n := len(nums)
	res := -1
	for i := 0; i < n; i++ {
		if nums[i] == (i % 10) {
			res = i
			break
		}
	}
	return res
}

func main() {
	fmt.Println(smallestEqual([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
