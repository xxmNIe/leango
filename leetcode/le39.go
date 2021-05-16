package main

import "fmt"

var ress [][]int
func combinationSum(candidates []int, target int) [][]int {
	ress = make([][]int,0)
	n :=len(candidates)
	ans :=make([]int,0)
	combin(candidates,n,0,target,ans)
	return ress
}

func combin(candidates []int,length int, start int, own int, ans []int) {
	if own==0{
		tmp := make([]int,len(ans))
		copy(tmp,ans)
		ress=append(ress,tmp)
	}
	for i:=start;i<length;i++{
		if own-candidates[i]>=0{
			ll:=len(ans)
			ans = append(ans,candidates[i])
			combin(candidates,length,i,own-candidates[i],ans)
			ans=ans[:ll]
		}
	}
}

func main() {

	fmt.Println(combinationSum([]int{2,3,7},7))
}