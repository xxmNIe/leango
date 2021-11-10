package main

import "fmt"

// func findDuplicate(nums []int) int {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		v := nums[i]

// 		for j := i + 1; j < n; j++ {

// 			if v == nums[j] {
// 				return v
// 			}
// 		}
// 	}
// 	return -1
// }

// func findDuplicate(nums []int) int {
// 	slow, fast := 0, 0
// 	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
// 	}
// 	slow = 0
// 	for slow != fast {
// 		slow = nums[slow]
// 		fast = nums[fast]
// 	}

// 	return slow
// }

func findDuplicate(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		cnt := 0
		mid := left + (right-left)/2
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}

		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// func findDuplicate(nums []int) int {
// 	fast, slow := nums[nums[0]], nums[0]
// 	for fast != slow {
// 		fast = nums[nums[fast]]
// 		slow = nums[slow]
// 	}
// 	fmt.Println("1")
// 	slow = 0

// 	for fast != slow {
// 		fast = nums[fast]
// 		slow = nums[slow]
// 	}
// 	return fast

// }
func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
