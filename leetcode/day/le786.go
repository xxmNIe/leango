package main

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	type pair struct{ x, y int }
	tmpArr := make([]pair, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			tmpArr = append(tmpArr, pair{arr[i], arr[j]})
		}
	}
	sort.Slice(tmpArr, func(i, j int) bool {
		return tmpArr[i].x*tmpArr[j].y < tmpArr[i].x*tmpArr[i].y
	})

	return []int{tmpArr[k-1].x, tmpArr[k-1].y}
}
