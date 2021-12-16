package main

import "fmt"

func maxPower(s string) int {
	n := len(s)
	res := 1
	for i := 1; i < n; i++ {
		tmp := 1
		for i < n && s[i] == s[i-1] {
			i++
			tmp++
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func main() {
	fmt.Println(maxPower("ccbccbb"))
}
