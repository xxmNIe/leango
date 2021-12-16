package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	for k > 0 {
		if nums[0] < 1 {
			for i := 0; i < len(nums) && nums[i] < 1 && k > 0; i++ {
				nums[i] = -nums[i]
				k--
			}
		} else {
			nums[0] = -nums[0]
			k--
		}
		sort.Ints(nums)
	}
	ans := 0
	for _, v := range nums {
		ans += v
	}
	return ans
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}
