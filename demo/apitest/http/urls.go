package main

import "fmt"

func main() {

	m :=map[string]string {"a":"b","c":"d"}
	for d := range m {
		fmt.Println(d)
	}

	//maps := url.Values{
	//	"name":{"a","b"},
	//	"age":{"c"},
	//}
	//
	//a :=maps.Encode()
	//
	//fmt.Println(a)
}
