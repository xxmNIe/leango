package main

import "fmt"

func sortColors(nums []int) {
	count0 := swap(nums, 0)
	swap(nums[count0:], 1)
}

func swap(nums []int, target int) (count int) {
	for i, c := range nums {
		if c == target {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
	return count
}
func main() {
	a := []int{1, 0, 2, 1, 1, 0}
	sortColors(a)
	fmt.Println(a)
}
