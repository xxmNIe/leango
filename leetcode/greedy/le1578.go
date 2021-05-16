package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCost("aabaa",[]int{1,2,3,4,1}))
}


func minCost(s string, cost []int) int {
	res :=0
	lens := len(s)
	if lens ==1{
		return 0
	}
	i:=0
	for i<lens{
		maxvalue :=cost[i]
		sum := 0
		tmp := s[i]
		for i<lens&& s[i]==tmp{
			maxvalue = max1(cost[i],maxvalue)
			sum+=cost[i]
			i++
		}
		res += sum-maxvalue


	}
	return res
}
func min1(x,y int)int{
		if x>y{
			return y
		}
		return x
}
func max1(x,y int)int{
	if x<y{
		return y
	}
	return x
}

//
//func minCost(s string, cost []int) int {
//	lens := len(s)
//	if lens ==1{
//		return 0
//	}
//	dp :=make([]int,lens)
//	for i:=0;i<lens-1;i++{
//		if s[i]==s[i+1]{
//			dp[i+1] = min1(cost[i],cost[i+1])
//		}
//	}
//
//}
//func min1(x,y int)int{
//		if x>y{
//			return y
//		}
//		return x
//}