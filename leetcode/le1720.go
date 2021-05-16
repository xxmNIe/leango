package main

import "fmt"

func decode(encoded []int, first int) []int {
	last :=first
	res :=make([]int,len(encoded)+1)
	res[0]=first
	for i:=0;i<len(encoded);i++ {
		res[i+1] = encoded[i]^last
		last = res[i+1]
	}
	return res
}
func main() {
	fmt.Print(decode([]int{6,2,7,3},4))
}