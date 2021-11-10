package main

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	A, B := 0, 0
	mp := make(map[byte]int, 0)
	for i := range secret {
		if secret[i] == guess[i] {
			A++
			continue
		}
		v1, v2 := mp[secret[i]], mp[guess[i]]
		mp[secret[i]]--
		mp[guess[i]]++
		if v1 > 0 {
			B++
		}
		if v2 < 0 {
			B++
		}

	}

	return fmt.Sprintf("%dA%dB", A, B)
}
func main() {
	fmt.Println(getHint("1807", "7810"))
}
