package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {return nums[i] <nums[j]})
	n := len(nums)
	res :=0
	for i:=0;i<n/2;i++ {
		res = maxd(nums[i]+nums[n-i-1],res)
	}
	return res
}

func maxd(x,y int) int {
	if x>y {
		return x
	}
	return y
}
func main() {
	fmt.Println(minPairSum([]int{3,5,4,2,4,6}))
}