package main

import "sort"

type TopVotedCandidate struct {
	tops, times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tops := []int{}
	top := -1
	voteCount := map[int]int{-1: -1}
	for _, p := range persons {
		voteCount[p]++
		if voteCount[p] >= voteCount[top] {
			top = p
		}
		tops = append(tops, top)
	}

	return TopVotedCandidate{tops: tops, times: times}
}

func (this *TopVotedCandidate) Q(t int) int {
	return this.tops[sort.SearchInts(this.times, t+1)-1]
}
