package main

import "fmt"

func hammingWeight2(num uint32) int {
	count :=0
	for i:=0;i<32;i++ {
		if (num >>i &1) ==1{
			count++
		}
	}
	return count
}
func hammingWeight(num uint32) int {
	count :=0
	for ;num >0;num &=num -1 {
		count++
	}
	return count
}

func main() {
	var a uint32 =00000000000000000000000010000000
	fmt.Println(hammingWeight(a))
}
