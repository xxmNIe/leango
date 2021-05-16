package main

import (
	"fmt"
	"sort"
)

/*
   count[i] = count[i-1]+1 or
							for   + count[i-coins[i]]+1
 */
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	if amount == 0{
		return 0
	}
	max := amount+1
	count :=make([]int,max)
	for i:=1;i<max;i++{
		count[i] = max+1
	}
	for i:=1;i<max;i++{
		tmp :=max+1
		for _,coin := range coins{
			if coin > i{
				break
			}
			tmp = min(tmp,count[i-coin]+1)
		}
		count[i] = tmp
	}
	if count[amount] >=max+1{
		return -1
	}
	return count[amount]
}

func min(x,y int)int {
	if x<y{
		return x
	}
	return y
}

func main() {
	fmt.Println(coinChange([]int{474,83,404,3},264))
}