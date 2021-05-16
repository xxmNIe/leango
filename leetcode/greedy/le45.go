package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(jump2([]int{2,3,1,1,4}))
}
func jump(nums []int) int {
	lens := len(nums)
	if lens <=1{
		return 0
	}
	dp := make([]int,lens)
	for i:=1;i<lens;i++{
		dp[i]=math.MaxInt64
	}
	for i:=0;i<lens;i++{
		jump := nums[i]
		if i!=0 {
			dp[i] = min(dp[i-1]+1, dp[i])
		}
		for j:= 0;j<=jump &&j+i<lens;j++{
			dp[j+i] = min(dp[j+i],dp[i]+1)
		}
	}
	return dp[lens-1]
}
//func min(x,y int)int{
//	if x <y {
//		return x
//	}
//	return y
//}

func jump2(nums []int) int {
	lens:=len(nums)
	end :=0
	maxpos:=0
	step :=0

	for i:=0;i<lens-1;i++{
		maxpos = max(maxpos,i+nums[i])
		if i==end{
			end = maxpos
			step++
		}
	}
	return step
}

func max(x,y int)int{
	if x >y {
		return x
	}
	return y
}