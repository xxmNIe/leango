package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	n :=len(arr)
	xors := make([]int,n+1)
	for i,v := range arr {
		xors[i+1] = xors[i] ^v
	}
	ans :=make([]int,len(queries))
	for i,arrs := range queries {
		ans[i] = xors[arrs[0]] ^ xors[arrs[1]+1]
	}
	return ans
}


func xorQueries2(arr []int, queries [][]int) []int {
	n :=len(queries)

	res := make([]int,n)

	for i,arrs := range queries {
		a :=arr[arrs[0]]
		for j:=arrs[0]+1;j<=arrs[1];j++{
			a ^= arr[j]
		}
		res[i] =a
	}
	return res
}

func main() {
	fmt.Println(xorQueries([]int{1,3,4,8},[][]int{{0,1},{1,2},{0,3},{3,3}}))
}