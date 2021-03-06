package main

import "fmt"

func countPairs(deliciousness []int) (ans int) {
	maxVal :=deliciousness[0]
	for _,val := range deliciousness {
		if val >maxVal {
			maxVal = val
		}
	}
	maxSum := maxVal *2
	cnt := map[int]int{}
	for _,val := range deliciousness {
		for sum :=1;sum <=maxSum; sum <<=1 {
			ans += cnt[sum - val]
		}
		cnt[val]+=1
	}
	return ans%(1e9 + 7)
}

func main() {
	fmt.Println(countPairs([]int{1,1,1,3,3,3,7}))
}