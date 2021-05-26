package main

import (
	"fmt"
	"math"
)

/*
	100(mid⋅t+dist[n−1])≤mid⋅hr.
 */
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	var hr int = int(math.Round(hour*100))

	if hr <(100*(n-1)) {
		return  -1
	}
	l :=1
	h :=200000000
	for l<h {
		mid :=(l+h)/2
		t :=0
		for i:=0;i<n-1;i++ {
			t+= (dist[i]-1)/mid+1
		}
		t*=mid
		t +=dist[n-1]
		if t *100 <= hr*mid {
			h = mid
		} else {
			l = mid +1
		}
	}
	return l
}

func main() {
	fmt.Println(minSpeedOnTime([]int{1,1,100000},2.01))
	//fmt.Println(minSpeedOnTime([]int{1,3,2},1.9))


}