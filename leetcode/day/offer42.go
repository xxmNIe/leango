package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	pre,now :=nums[0],nums[0]
	res :=nums[0]
	for i:=1;i<n;i++ {
		pre,now = now,nums[i]
		if nums[i]+pre < nums[i] {
			now = nums[i]
		}else {
			now =pre+nums[i]
		}
		res =max(now,res)
	}
	return res
}


func max34(x,y int) int {
	if x>y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
