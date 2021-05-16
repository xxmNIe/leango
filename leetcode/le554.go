package main

import "fmt"

func main() {
	//fmt.Println(leastBricks([][]int{{1,2,2,1},{3,1,2},{1,3,2},{2,4},{3,1,2},{1,3,1,1}}))
	fmt.Println(leastBricks([][]int{{1},{1},{1}}))
	fmt.Println(leastBricks([][]int{{1,1,1,1,1},{1,1,1,1,1},{1,1,1,1,1}}))
	fmt.Println(leastBricks([][]int{{7,1,1,1},{1,4,2,3},{7,3}}))
}

func leastBricks(wall [][]int) int {
	maps := make(map[int][]int)
	for i,v:=range wall {
		tmp := 0
		for _,vv := range v {
			tmp += vv
			maps[i] = append(maps[i],tmp)
		}
	}
	nmp := make(map[int]int)
	for _,v:=range maps {
		for _,vv := range v[:len(v)-1] {
			nmp[vv]++
		}
	}
	res :=0
	for _,v :=range nmp{
		if   v>res {
			res = v
		}
	}
	if len(wall[0])>1 && res ==0 {
		res = len(wall)
	}
	return len(wall)-res

}