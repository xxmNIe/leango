package main

import "fmt"

func numWays(n int) int {
	f1,f2 :=1,1
	for n>1 {
		f1,f2 = f2,(f1+f2)%(1e+7)
		n-=1
	}
	return f2
}

func main() {
	fmt.Println(numWays(2))
}