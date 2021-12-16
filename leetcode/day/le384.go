package main

import "math/rand"

type Solution struct {
	num, original []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	copy(this.num, this.original)
	return this.num
}

func (this *Solution) Shuffle() []int {
	// shuffled := make([]int, len(this.num))

	// for i := range shuffled {
	// 	j := rand.Intn(len(this.num))
	// 	shuffled[i] = this.num[j]
	// 	this.num = append(this.num[:j], this.num[j+1:]...)
	// }
	// this.num = shuffled
	// return this.num

	n := len(this.num)
	for i := range this.num {
		j := i + rand.Intn(n-i)
		this.num[i], this.num[j] = this.num[j], this.num[i]
	}
	return this.num
}
