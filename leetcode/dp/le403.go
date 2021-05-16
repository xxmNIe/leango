package main

import (
	"fmt"
)

func main() {
	fmt.Println(canCross([]int{0,1,3,5,6,8,12,17}))
}




func canCross(stones []int) bool {
	n := len(stones)
	dp := make([][]bool,n)
	for i := 0;i < n;i++{
		dp[i] = make([]bool,n)
	}

	dp[0][0] = true

	for i:=1;i<n;i++{
		if stones[i] - stones[i-1] >i{
			return false
		}
	}
	for i:=1;i<n;i++{
		for j :=i-1;j>=0;j-- {
			k :=stones[i] - stones[j]
			if k >j+1{
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i==n-1 &&dp[n-1][k]{
				return true
			}
		}
	}
	return false
}

//func canCross(stones []int) bool {
//	end :=0
//	maxpos :=0
//	minpos :=0
//	step :=1
//	for i:=0;i<len(stones);i++ {
//		maxpos = stones[i]+step+1
//		minpos = stones[i]+step-1
//		if i == end{
//			end = maxpos
//			s
//		}
//
//	}
//
//	return true
//}