package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	sb := []byte(s)
	left, right := 0, len(sb)-1
	for left <= right {
		for left < right && !isy(sb[left]) {
			left++
		}
		
		for left < right && !isy(sb[right]) {
			right--
		}
		sb[left], sb[right] = sb[right], sb[left]
		left++
		right--
	}
	return string(sb)
}

func isy(b byte) bool {
	if strings.Contains("aeiuoAEIOU", string(b)) {
		return true
	}
	return false
}

func main() {
	fmt.Println(reverseVowels(""))
}
