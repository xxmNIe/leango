package main

import (
	"fmt"
	"net/url"
)

func main() {
	maps := url.Values{
		"name":{"a","b"},
		"age":{"c"},
	}

	a :=maps.Encode()

	fmt.Println(a)
}
