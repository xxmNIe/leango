package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTimeRequired([]int{46,13,54,51,38,49,54,67,26,78,33},10))
}

func minimumTimeRequired(jobs []int, k int) int {
	n :=len(jobs)
	sum := make([]int,k)
	ans :=math.MaxInt64
	walk := func(int,int,int){}
	walk = func(u int, maxd int,used int) {
		if maxd >= ans {
			return
		}
		if u==n {
			ans = maxd
			return
		}
		if used <k {
			sum[used] =jobs[u]
			walk(u+1,max2(sum[used],maxd),used+1)
			sum[used] =0
		}
		for i:=0;i<used;i++ {
			sum[i]+= jobs[u]
			walk(u+1,max2(sum[i],maxd),used)
			sum[i]-=jobs[u]
		}
	}
	walk(0,0,0)
	return ans
}
func max2(x,y int ) int{
	if x >y {
		return x
	}
	return y
}