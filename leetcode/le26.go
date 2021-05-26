package main

import "fmt"

func removeDuplicates(nums []int) int {
	i,j:=1,1
	n := len(nums)
	if n <2{
		return n
	}
	for ;i<n;i++ {
		for i<n&& nums[i]==nums[i-1]  {
			i++
		}
		if i>=n {
			continue
		}
		nums[j]=nums[i]
		j++
	}
	return j
}

func main() {
	fmt.Println(removeDuplicates([]int{1,1}))
}
