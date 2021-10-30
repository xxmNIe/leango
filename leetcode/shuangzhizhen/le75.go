package main

import "fmt"

func sortColors(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}

	for i := 0; i < n; i++ {
		cur := 0
		for cur < n-i-1 {
			if nums[cur] > nums[cur+1] {
				nums[cur], nums[cur+1] = nums[cur+1], nums[cur]
			}
			cur++
		}
	}

}

func main() {
	a := []int{1, 0, 2, 1, 1, 0}
	sortColors(a)
	fmt.Println(a)
}
