package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	sum := (1 + n) * (n / 2)
	if n%2 != 0 {
		sum += (n+1) / 2
	}

	for _, v := range nums {
		sum -= v
	}
	return sum
}
 
func main() {
	fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1}))
}
