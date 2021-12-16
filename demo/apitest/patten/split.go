package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "fesf"
	for _, v := range strings.Split(str, "/") {
		fmt.Println(v)
	}

}
