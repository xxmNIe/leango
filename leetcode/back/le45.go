package main

func jump(nums []int) int {
	l := len(nums)
	end := 0
	maxpos := 0
	step := 0
	for i := 0; i < l-1; i++ {
		maxpos = max(maxpos, nums[i]+i)
		if i == end {
			step++
			end = maxpos
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
