package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int, 0)
	n := len(s)
	r, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			mp[s[i-1]]--
		}
		for r+1 < n && mp[s[r+1]] == 0 {
			mp[s[r+1]]++
			r++
		}
		ans = max(ans, r-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
