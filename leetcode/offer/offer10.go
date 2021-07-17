package main

import "fmt"

func fib(n int) int {
	if n <2 {
		return n
	}
	f1,f2 :=0,1
	for i:=2;i<=n;i++ {
		f1,f2 = f2,(f2+f1)%(1e9+7)
	}
	return f2%(1e9+7)
}

func main() {
	fmt.Println(fib(5))
}