package main

import "math"

func getMinDistance(nums []int, target int, start int) int {
	lens := len(nums)

	var  res float64 = math.MaxFloat64
	for i:=0;i<lens;i++ {
		if nums[i] == target {
			res = math.Min(math.Abs(float64(i-start)),res)
		}
	}
	return int(res)
}
