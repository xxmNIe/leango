package main

import "fmt"

func minArray1(numbers []int) int {
	n := len(numbers)
	if n <2{
		return numbers[0]
	}
	key := numbers[0]
	left,right := 0,n-1
	mid := (left+right)/2
	for left <right {
		if numbers[mid] < key {
			if  numbers[mid]>numbers[right]  {
				left = 0
				right = mid-1
			}else if numbers[mid] > numbers[left]{
				left = mid+1
				right = n-1
			}
			key = numbers[mid]
		}else {
			right = mid-1
		}
	}
	return  key
}

func minArray(numbers []int) int {
	n :=len(numbers)
	left,right := 0,n-1
	mid := left+(right-left)/2
	for left <right {
		mid = left+(right-left)/2
		if numbers[mid] <numbers[right] {
			right = mid
		}else if numbers[mid] > numbers[right] {
			left = mid+1
		}else {
			right -=1
		}
	}
	return numbers[left]
}

func main() {
	fmt.Println(minArray([]int{2,2,2,0,1}))
}