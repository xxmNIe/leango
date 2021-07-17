package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	n :=len(nums)
	count :=0
	var walk  func(deep int,tmp int)
	walk = func(deep int,tmp int) {
		if deep == n{
			if tmp == target{
				count ++
			}
			return
		}
		walk(deep+1,tmp+nums[deep])
		walk(deep+1,tmp-nums[deep])

	}

	walk(0,0)
	return count
}

func main() {
	fmt.Println(findTargetSumWays([]int{1},1))
}
