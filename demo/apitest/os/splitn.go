package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/geeo/errr/errrt/"
	base := "/geeo/"
	ss := strings.SplitN(s[len(base):], "/", 2)
	fmt.Println(ss)
}
