package main

import (
	"fmt"
	"sort"
)

// func findLHS(nums []int) int {
// 	if len(nums) == 1 {
// 		return 0
// 	}
// 	sort.Ints(nums)
// 	res, l, r := 0, 0, 0
// 	for r < len(nums) {
// 		tmp := 0
// 		for r < len(nums) && nums[r]-nums[l] < 2 {
// 			r++
// 			tmp++
// 		}
// 		if nums[r-1] == nums[l] {
// 			continue
// 		}
// 		res = max(res, tmp)
// 		l++
// 		r = l
// 	}
// 	return res
// }

func findLHS(nums []int) int {

	sort.Ints(nums)
	begin := 0
	res := 0
	for end, num := range nums {
		for num-nums[begin] > 1 {
			begin++
		}
		if num-nums[begin] == 1 && end-begin+1 > res {
			res = end - begin + 1
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

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
