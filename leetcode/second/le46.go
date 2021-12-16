package main

import "fmt"

// func permute(nums []int) [][]int {
// 	n := len(nums)
// 	ans := make([][]int, 0)
// 	visited := make([]bool, n)
// 	var walk func(an []int)
// 	walk = func(an []int) {
// 		if len(an) == n {
// 			tmp := make([]int, n)
// 			copy(tmp, an)
// 			ans = append(ans, tmp)
// 			return
// 		}
// 		for i := 0; i < n; i++ {
// 			if visited[i] {
// 				continue
// 			}
// 			visited[i] = true
// 			an = append(an, nums[i])
// 			walk(an)
// 			visited[i] = false
// 			an = an[:len(an)-1]
// 		}
// 	}
// 	walk([]int{})
// 	return ans
// }

func permute(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	if n < 2 {
		return [][]int{nums}
	}
	var walk func(now int)
	walk = func(now int) {
		if now == n {
			tmp := make([]int, n)
			copy(tmp, nums)
			ans = append(ans, tmp)
			return
		}

		for i := now; i < n; i++ {
			nums[i], nums[now] = nums[now], nums[i]
			walk(now + 1)
			nums[i], nums[now] = nums[now], nums[i]
		}
	}
	walk(0)
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
