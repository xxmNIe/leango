package main

import "fmt"

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
	}
	return dp[n]
}

func main() {
	fmt.Println(numWays(44))
}
