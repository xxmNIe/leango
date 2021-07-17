package main

import "fmt"

func majorityElement(nums []int) int {
	n := len(nums)
	if n <1 {
		return -1
	}
	major := nums[0]
	ctn :=1
	for i:=1;i<n;i++ {
		if ctn == 0 {
			major = nums[i]
			ctn=1
		}else if nums[i] == major {
			ctn++
		}else {
			ctn--
		}
	}
	ctn =0
	for i:=0;i<n;i++ {
		if nums[i] ==major{
			ctn++
		}
	}
	if ctn <=n/2 {
		return -1
	}
	return major
}

func main() {
	fmt.Println(majorityElement([]int{10,9,9,9,10}))
}