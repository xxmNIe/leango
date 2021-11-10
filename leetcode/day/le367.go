package main

import "fmt"

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)/2
		squ := mid * mid
		if squ == num {
			return true
		}
		if squ < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
}
