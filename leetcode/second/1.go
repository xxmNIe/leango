package main

import "fmt"

func twoSum(nums []int, target int) []int {
	maps :=map[int]int{}
	for i,v := range nums {
		tmp := target - v
		if idx,ok:= maps[tmp];ok{
			return []int{idx,i}
		}else {
			maps[v] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3,3},6))
}