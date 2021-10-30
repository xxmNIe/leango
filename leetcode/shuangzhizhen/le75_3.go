package main

import "fmt"

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func main() {
	a := []int{1, 0, 2, 1, 1, 0}
	sortColors(a)
	fmt.Println(a)
}
