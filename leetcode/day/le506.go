package main

import (
	"fmt"
	"sort"
	"strconv"
)

// func findRelativeRanks(score []int) []string {
// 	mp := make(map[int]int)
// 	for i, v := range score {
// 		mp[v] = i
// 	}
// 	sort.Slice(score, func(i, j int) bool { return score[i] > score[j] })
// 	res := make([]string, len(score))
// 	for i, v := range score {
// 		idx := mp[v]
// 		if i == 0 {
// 			res[idx] = "Gold Medal"
// 		} else if i == 1 {
// 			res[idx] = "Silver Medall"
// 		} else if i == 2 {
// 			res[idx] = "Bronze Medal"
// 		} else {
// 			res[idx] = strconv.Itoa(i + 1)
// 		}
// 	}
// 	return res
// }
var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
	n := len(score)
	type pair struct{ x, y int }
	arr := make([]pair, n)
	for i, v := range score {
		arr[i] = pair{i, v}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].y > arr[j].y })
	ans := make([]string, n)
	for i, v := range arr {
		if i < 3 {
			ans[v.x] = desc[i]
		} else {
			ans[v.x] = strconv.Itoa(i + 1)
		}
	}
	return ans
}

func main() {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
