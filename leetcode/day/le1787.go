package main

func minChanges(nums []int, k int) int {
	count :=0
	a:=0
	for i:=0;i<k-1;i++ {
		a  ^= nums[i]
	}
	if nums[k]!= a {
		nums[k] = a
		count++
	}
	return 0
}
