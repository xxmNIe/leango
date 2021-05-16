package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange2([]int{1, 2, 5},11))
}

func coinChange(coins []int, amount int) int {
	if amount<1{
		return 0
	}
	count :=make([]int,amount)
	var conchange func([]int,int)int
	conchange = func(coins []int,rem int)int {
		if rem <0{
			return -1
		}
		if rem ==0{
			return 0
		}
		if count[rem-1]!=0{
			return count[rem-1]
		}
		min :=math.MaxInt64
		for _,coin :=range coins{
			res := conchange(coins,rem-coin)
			if res >=0 &&res <min{
				min = 1+res
			}
		}
		if min ==math.MaxInt64{
			min = -1
		}
		count[rem-1]=min
		return count[rem-1]
	}

	return conchange(coins,amount)
}


func coinChange2(coins []int, amount int) int {
	max :=amount+1
	count :=make([]int,max)
	for i:=1;i<max;i++{
		count[i] = max
	}
	count[0]=0

	for i:=1;i<=amount;i++{
		for _,j :=range coins{
			if j <=i{
				count[i] = min(count[i],count[i-j]+1)
			}
		}
	}
	if count[amount] ==max{
		return -1
	}
	return count[amount]
}
func min(x,y int) int{
	if x<y{
		return x
	}
	return y
}