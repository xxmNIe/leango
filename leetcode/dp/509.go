package main

import "fmt"

func fib(n int) int {
	if n <=1{
		return 0
	}
	dp := make([]int,n+1)
	dp[1] = 1
	for i :=2;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Print(fib(4))
}