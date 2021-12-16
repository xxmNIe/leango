package main

import "fmt"

// func dailyTemperatures(temperatures []int) []int {
// 	n := len(temperatures)
// 	if n == 1 {
// 		return []int{0}
// 	}
// 	ans := make([]int, n)
// 	for i := range temperatures {
// 		j := i + 1
// 		for j < n && temperatures[j] <= temperatures[i] {
// 			j++
// 		}
// 		if j == n {
// 			ans[i] = 0
// 		} else {
// 			ans[i] = j - i
// 		}

// 	}
// 	return ans
// }

//err
// func dailyTemperatures(temperatures []int) []int {
// 	n := len(temperatures)
// 	if n == 1 {
// 		return []int{0}
// 	}
// 	stack := make([]int, n)
// 	ans := make([]int, n)
// 	idx := 0
// 	for i := 0; i < n; i++ {
// 		if i == 0 || idx == 0 {
// 			stack[idx] = i
// 			idx++
// 			continue
// 		}
// 		top := stack[idx-1]
// 		if temperatures[i] <= temperatures[top] {
// 			stack[idx] = i
// 			idx++
// 		} else {
// 			for idx > 0 && temperatures[top] < temperatures[i] {

// 				top = stack[idx-1]
// 				ans[top] = i - top
// 				idx--
// 			}
// 			stack[idx] = i
// 			idx++
// 		}

// 	}
// 	return ans
// }

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, 0)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
		}
		temper := temperatures[i]
		for len(stack) > 0 && temper > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
