package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] <nums[j]})
	left,right,res,totle :=0,1,0,0
	for right < len(nums) {
		totle += (nums[right] -nums[right-1])*(right-left)
		for totle > k {
			totle -=(nums[right]-nums[left])
			left++
		}
		res = maxx(res,right-left+1)
		right +=1
	}
	return res
}

func maxx(x,y int) int {
	if x>y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxFrequency([]int{1,2,4},5))
}