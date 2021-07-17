package main

import "fmt"

func numWays1(n int, relation [][]int, k int) int {
	edges := make([][]int,n)
	ans :=0
	for _,r := range relation {
		edges[r[0]] = append(edges[r[0]],r[1] )
	}
	//fmt.Println(edges)
	var dfs func(int,int)
	dfs = func(x,step int){
		if step == k {
			if n-1 == x {
				ans++
			}
			return

		}
		for _,v :=range edges[x] {
			dfs(v,step+1)
		}
	}
	dfs(0,0)
	return ans
}

func numWays(n int, relation [][]int, k int) int {
	dp := make([][]int,k+1)
	for i :=range dp {
		dp[i] = make([]int,n)
	}
	dp[0][0]=1
	for i:=0;i<k;i++ {
		for _,r :=range relation {
			src,dst := r[0],r[1]
			dp[i+1][dst] += dp[i][src]
		}
	}
	return dp[k][n-1]
}

func main() {
	fmt.Println(numWays(5,[][]int{[]int{0,2},[]int{2,1},[]int{3,4},[]int{2,3},[]int{1,4},[]int{2,0},[]int{0,4}},3))
}
