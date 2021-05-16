package main

import "fmt"

func main() {
	fmt.Println(xorOperation(4,3))
}

func xorOperation(n int, start int) int {
	res :=0
	for i:=0;i<n;i++ {
		res = res ^ (start+2*i)
	}
	return res
}
