package main

import (
	"fmt"
	"unicode"
)

func detectCapitalUse(word string) bool {
	runes := []rune(word)

	up, le := 0, 0
	for _, r := range runes {
		if unicode.IsUpper(r) {
			if le > 0 {
				return false
			}
			up++
		} else {
			le++
		}
	}
	if up == len(runes) || le == len(runes) || le == len(runes)-1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(detectCapitalUse("UsA"))
}
