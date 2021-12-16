package main

func maxSumOfOneSubarray(nums []int, k int) (ans []int) {
	var sum1, maxSum1 int
	for i, num := range nums {
		sum1 += num
		if i >= k-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				ans = []int{i - k + 1}
			}
		}
		sum1 -= nums[i-k+1]
	}
	return ans
}
func maxSumOfTwoSubarrays(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1index int
	var sum2, maxSum12 int
	for i := k; i < len(nums); i++ {
		sum1 += nums[i-k]
		sum2 += nums[k]
		if i >= 2*k-1 {
			if sum1 >= maxSum1 {
				maxSum1 = sum1
				maxSum1index = i - 2*k + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum12 + sum2
				ans = []int{maxSum1index, i - k + 1}
			}
			sum1 -= nums[i-2*k+1]
			sum2 -= nums[i-k+1]
		}
	}
	return
}

func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1Index int
	var sum2, maxSum2, maxSum12Idnex1, maxSum12Index2 int
	var sum3, maxTotal int
	for i := k * 2; i < len(nums); i++ {
		sum1 += nums[i-2*k]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= 3*k-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Index = i - 3*k + 1
			}
			if maxSum1+sum2 > maxSum2 {
				maxSum2 = maxSum1 + sum2
				maxSum12Idnex1, maxSum12Index2 = maxSum1Index, i-2*k+1
			}
			if maxSum2+sum3 > maxTotal {
				maxTotal = maxSum2 + sum3
				ans = []int{maxSum12Idnex1, maxSum12Index2, i - k + 1}
			}
			sum1 -= nums[i-3*k+1]
			sum2 -= nums[i-2*k+1]
			sum3 -= nums[i-k+1]	
		}

	}
	return
}
