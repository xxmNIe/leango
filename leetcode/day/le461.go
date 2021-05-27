package main

import "fmt"

func hammingDistance(x int, y int) int {
	count :=0
	for i:=0;i<31;i++ {
		count +=  x>>i&1 ^ y>>i&1
	}
	return count
}

func main() {
	fmt.Println(hammingDistance(-1,4))
}