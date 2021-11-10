package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	rs := []rune(s)
	l, r := 0, len(rs)-1
	for l <= r {
		if !unicode.IsLetter(rs[l]) && !unicode.IsNumber(rs[l]) {
			l++
			continue
		}
		if !unicode.IsLetter(rs[r]) && !unicode.IsNumber(rs[r]) {
			r--
			continue
		}
		if unicode.ToLower(rs[l]) != unicode.ToLower(rs[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))

}
