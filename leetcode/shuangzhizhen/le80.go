package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	l, r := 0, 0
	
	for r < n {
		cnt, cur := 0, nums[r]
		for r < n && cnt < 2 && nums[r] == cur {
			nums[l] = nums[r]
			l++
			cnt++
			r++
		}
		if cnt == 2 {
			for r < n && nums[r] == cur {
				r++
			}
			if r == n {
				break
			} else {
				cur = nums[r]
				cnt = 0
			}
		}

	}
	return l
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
