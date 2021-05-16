package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deleteAndEarn([]int{2,3,4}))
}

func deleteAndEarn(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	maps := make(map[int]int)
	dp := make([]int,1)
	dp = append(dp, nums[0])
	for i:=0;i<n;i++ {
		maps[nums[i]]++
		if nums[i]!=dp[len(dp)-1] {
			dp = append(dp, nums[i])
		}
	}
	last := dp[1]
	dp[1] = dp[1] * maps[dp[1]]
	for i:=2;i<len(dp);i++ {

		if dp[i] -last == 1{
			last = dp[i]
			dp[i] = max(dp[i-1],dp[i-2]+dp[i]*maps[dp[i]])
		} else {
			last = dp[i]
			dp[i] = dp[i]*maps[dp[i]]+dp[i-1]
		}
	}

	return dp[len(dp)-1]
}
func max(x,y int) int{
	if x>y {
		return x
	}
	return y
}
