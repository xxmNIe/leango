package main

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

	}
	reverse(nums[i+1:])
}

func reverse(ss []int) {
	for i, n := 0, len(ss); i < n/2; i++ {
		ss[i], ss[n-i-1] = ss[n-i-1], ss[i]
	}
}
