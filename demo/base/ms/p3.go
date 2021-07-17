package main

import "fmt"

func main() {
	var a []int
	var m map[string]int
	// all = nil
	m["a"]=5 //err
	a = append(a, 9)  //ok
	fmt.Println(a)
	fmt.Println(m)
}
