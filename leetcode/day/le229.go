package main

import "fmt"

func majorityElement(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	maj1, maj2 := 0, 0
	cnt1, cnt2 := 0, 0
	for _, c := range nums {
		if cnt1 > 0 && c == maj1 {
			cnt1++
		} else if cnt2 > 0 && c == maj2 {
			cnt2++
		} else if cnt1 == 0 {
			maj1 = c
			cnt1 = 1
		} else if cnt2 == 0 {
			maj2 = c
			cnt2 = 1
		} else {
			cnt2--
			cnt1--
		}
	}
	res := []int{}
	c1, c2 := 0, 0
	for _, c := range nums {
		if cnt1 > 0 && c == maj1 {
			c1 += 1
		}
		if cnt2 > 0 && c == maj2 {
			c2++
		}
	}

	if cnt1 > len(nums)/3 {
		res = append(res, maj1)
	}
	if cnt2 > len(nums)/3 {
		res = append(res, maj2)
	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
