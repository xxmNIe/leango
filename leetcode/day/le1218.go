package main

import "fmt"

func longestSubsequence(arr []int, difference int) int {
	mp := make(map[int]int, len(arr))
	res := 1
	for _, n := range arr {
		if v, ok := mp[n-difference]; ok {
			mp[n] = v + 1
			res = max(v+1, res)
		} else {
			mp[n] = 1
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}
