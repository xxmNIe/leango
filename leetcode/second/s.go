package main

import (
	"fmt"
	"strconv"
)

func nextBeautifulNumber(n int) int {
	for !check(n) {
		n++
	}
	return n
}

func check(n int) bool {
	tags := [10]int{}
	str := strconv.Itoa(n)

	for _, c := range str {
		tags[c-'0'] += 1
	}
	for i := range tags {
		if tags[i] != 0 && i != tags[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(nextBeautifulNumber(1234))
}
