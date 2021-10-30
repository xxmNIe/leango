package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countValidWords(sentence string) int {
	count := 0

	strs := strings.Split(sentence, " ")
out:
	for _, str := range strs {
		str = strings.Trim(str, " ")
		if str == "" {
			continue
		}
		l, b := 0, 0
		for i, c := range str {
			n := len(str)
			if unicode.IsDigit(c) {
				continue out
			}
			if isP(c) {
				if i != n-1 {
					continue out
				}
				if b > 0 {
					continue out
				}
				b++
			}
			if c == '-' {
				if i == 0 || i == n-1 {
					continue out
				}
				if l > 0 {
					continue out
				}
				if !unicode.IsLower(rune(str[i-1])) || !unicode.IsLower(rune(str[i+1])) {
					continue out
				}
				l++
			}
		}
		count++
	}
	return count
}

func isP(c rune) bool {
	if c == '!' || c == '.' || c == ',' {
		return true
	}
	return false
}

func main() {
	fmt.Println(countValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."))
}
