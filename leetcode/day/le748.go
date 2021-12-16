package main

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	cnt := [26]int{}
	for _, c := range licensePlate {
		if unicode.IsLetter(c) {
			cnt[(c|32)-'a']++
		}
	}
	ans := ""
out:
	for _, word := range words {
		tmp := [26]int{}
		for _, c := range word {
			tmp[c-'a']++
		}
		for i := 0; i < 26; i++ {
			if tmp[i] < cnt[i] {
				continue out
			}
		}
		if ans == "" || len(word) < len(ans) {
			ans = word
		}
	}
	return ans
}


