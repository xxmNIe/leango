package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	mp := make(map[int]int, n)
	s := make([]int, n)
	idx := 0
	for i := len(nums2) - 1; i >= 0; i-- {
		for idx != 0 && s[idx-1] <= nums2[i] {
			idx -= 1
		}
		if idx == 0 {
			mp[nums2[i]] = -1
		} else {
			mp[nums2[i]] = s[idx-1]
		}
		s[idx] = nums2[i]
		idx++
	}
	ans := []int{}

	for _, c := range nums1 {
		if num, ok := mp[c]; ok {
			ans = append(ans, num)
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}

func main() {
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
