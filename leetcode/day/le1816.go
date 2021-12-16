package main

import "fmt"

func truncateSentence(s string, k int) string {
	pos, i := -1, 0
	for i = 0; i < len(s); i++ {

		if !isBlack(s[i]) {
			continue
		}
		k--
		if k == 0 {
			pos = i
			break
		}
	}
	if i == len(s) {
		pos = len(s)
	}
	return s[:pos]
}

func isBlack(b byte) bool {
	if b == ' ' {
		return true
	}
	return false
}

func main() {
	fmt.Println(truncateSentence("Hello aaa", 2))
}
