package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}

	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

func main() {
	fmt.Println(([]int{3, 30, 34, 5, 9}))
}
