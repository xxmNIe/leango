package main

import "fmt"

func maxGeneticDifference(parents []int, queries [][]int) []int{
	res := make([]int,0)
	cache := make(map[int]map[int]int)
	for _,pair := range queries {
		node,val := pair[0],pair[1]
		nodecp := node
		xor :=0
		for node != -1 {
			if ca,ok := cache[node];ok {
				if ca == nil {
					ca = make(map[int]int)
				}else if tmp,ok :=ca[val];ok {
					 xor = max(tmp,xor)

					 break
				}
			}
				xor = max(node^val, xor)
				node = parents[node]
				//cache[node][val] = xor

		}
		if _,ok :=cache[nodecp];!ok{
			cache[nodecp] = make(map[int]int)
		}
		cache[nodecp][val] = xor
		res  = append(res, xor)
	}
	return res
}

func max(x,y int) int {
	if x>y{return x}
	return y
}

func main() {
	fmt.Println(maxGeneticDifference([]int{3,7,-1,2,0,7,0,2},[][]int{{4,6},{1,15},{0,5}}))
}