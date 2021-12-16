package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	sb := strings.Builder{}
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			sb.WriteByte(s[i] | 32)
		} else {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(toLowerCase("Hello"))
}
