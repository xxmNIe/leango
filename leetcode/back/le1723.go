package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTimeRequired([]int{46,13,54,51,38,49,54,67,26,78,33},10))
}

func minimumTimeRequired1(jobs []int, k int) int {
	n :=len(jobs)
	sum := make([]int,k)
	ans :=math.MaxInt64
	walk := func(int,int){}
	walk = func(u int, maxd int) {
		if maxd >= ans {
			return
		}
		if u==n {
			ans = maxd
			return
		}
		for i:=0;i<k;i++ {
			sum[i]+= jobs[u]
			walk(u+1,max(sum[i],maxd))
			sum[i]-=jobs[u]
		}
	}
	walk(0,0)
	return ans
}
func max1(x,y int ) int{
	if x >y {
		return x
	}
	return y
}