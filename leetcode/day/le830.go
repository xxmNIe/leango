package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int)  (ans int) {
	left1,left2 := 0,0
	sum1,sum2 :=0,0
	for right,num := range nums {
		sum1+=num
		for left1 <= right && sum1 >goal {
			sum1-=nums[left1]
			left1++
		}
		sum2+=num
		for left2<=right && sum2 >=goal {
			sum2-=nums[left2]
			left2++
		}
		ans+=left2-left1
	}
	return
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1},2))
}