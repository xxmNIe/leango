package main

import (
	"fmt"
	"unicode"
)

func findWords(words []string) []string {
	const rowIdx = "12210111011122000010020202"
	ans := make([]string, 0)
out:
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		for _, ch := range word[1:] {
			if rowIdx[unicode.ToLower(rune(ch))-'a'] != idx {
				continue out
			}
		}
		ans = append(ans, word)
	}
	return ans
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
