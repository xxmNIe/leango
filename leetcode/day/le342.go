package main

import "fmt"

func isPowerOfFour(n int) bool {
	if n<=0{
		return false
	}
 	match := 0x2aaaaaaa
 	if n & match >0{
 		return false
	}

 	return n & -n ==n
}

func main() {
	fmt.Println(isPowerOfFour(32))
}