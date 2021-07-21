package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber(nums []int) int {
	mp :=make(map[int]struct{},len(nums))
	for _,num :=range nums {
		if _,ok:=mp[num];ok{
			return num
		}
		mp[num] = struct{}{}
	}
	return -1
}


func findRepeatNumber2(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <nums[j]
	})
	for i:=1;i<len(nums);i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}