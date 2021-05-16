package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1,2},[]int{}))
}

func findContentChildren(g []int, s []int) int {
	lens :=len(s)
	if lens ==0{
		return 0
	}
	leng :=len(g)
	sort.Ints(g)
	sort.Ints(s)
	i,j,res:=0,0,0
	for ;i<leng;i++{
		for ;j<lens;{
			if g[i]<=s[j]{
				res++
				j++
				break
			}
			j++
		}
	}
	return res
}
