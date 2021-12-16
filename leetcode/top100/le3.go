package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int, 0)
	l, r := 0, 0
	res := 0
	for r < len(s) {
		if mp[s[r]] > 0 {
			mp[s[l]]--
			l++
			continue
		} else {
			if r-l+1 > res {
				res = r - l + 1
			}
			mp[s[r]]++
			r++

		}

	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
