package main

func minInterval(intervals [][]int, queries []int) []int {
 mapsl := make(map[int]int)
 mapsh := make(map[int][]int)

 for _,v :=range intervals {
 	mapsl[v[0]]=v[1]
 	if mapsh[v[1]] == nil {
		mapsh[v[1]] = make([]int,0)
	}
	mapsh[v[1]] = v
 }
 return nil
}