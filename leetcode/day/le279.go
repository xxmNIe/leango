package main

import (
	"fmt"
	"math"
)
/*
f[i] = 1+min(i~/n)f[i-j*j]

 */
func numSquares(n int) int {
	dp := make([]int,n+1)
	for i:=1;i<=n;i++ {
		minn :=math.MaxInt32
		for j :=1;j*j<=i;j++ {
			minn = min(dp[i-j*j],minn)
		}
		dp[i] = 1+minn
	}
	return dp[n]
}

func min(x,y int) int{
	if x<y {
		return x
	}
	return y
}

func main() {
	fmt.Println(numSquares(12))
}
