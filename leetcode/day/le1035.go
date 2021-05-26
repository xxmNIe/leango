package main

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1,n2 := len(nums1),len(nums2)
	dp := make([][]int,n1+1)
	for i:=0;i<n1+1;i++ {
		dp[i] = make([]int,n2+1)
	}

	for i:=0;i<n1;i++ {
		for j :=0;j<n2;j++ {
			if nums1[i] == nums2[j]{
				dp[i+1][j+1] = dp[i][j]+1
			}else {
				dp[i+1][j+1] =max(dp[i][j+1],dp[i+1][j])
			}
		}
	}
	return dp[n1][n2]
}

func max(x,y int) int {
	if x>y {
		return x
	} else {
		return y
	}
}
func main() {
	fmt.Println(maxUncrossedLines([]int{1,4,2},[]int{1,2,4}))
}