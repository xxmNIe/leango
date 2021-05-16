package main

import (
	"fmt"
	"sort"
)

var resd [][]int
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	resd = make([][]int,0)
	len := len(candidates)
	if len ==0{
		return resd
	}
	ll := make([]int,0)
	dfs(candidates,0,target,len,ll)
	return resd
}

func dfs(candidates []int,start int,need int,lens int,ll []int){
		pt := len(ll)
		if need ==0{
			tmp := make([]int,pt)
			copy(tmp,ll)
			resd = append(resd,tmp)
			return

		}
		for i:=start;i<lens;i++{
			if i>0 && candidates[i]==candidates[i-1]&& i-1>=start{
				continue
			}
			if need-candidates[i]>=0{
				ll = append(ll, candidates[i])
				dfs(candidates,i+1,need-candidates[i],lens,ll)
				ll = ll[:pt]
			}
		}

}

func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5},8))
}