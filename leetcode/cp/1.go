package main

import "fmt"

func sumBase(n int, k int) int {
	if n<k{
		return n
	}

	res :=0
	for n%k!=0 || n/k!=0{
		res = res *10+n%k
		n/=k
	}
	re :=0
	for res !=0{
		tmp := res%10
		re += tmp
		res /=10
	}
	return re
}

func main() {
	fmt.Println(sumBase(34,6))
}