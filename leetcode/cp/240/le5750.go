package main

import "fmt"

func main() {
	fmt.Println(maximumPopulation([][]int{{1950,1961},{1960,1971},{1970,1981}}))
}

func maximumPopulation(logs [][]int) int {
 	maps := make(map[int]int)
 	for _,li :=range logs {
 		for i:=li[0];i<li[1];i++ {
 			maps[i]++
		}
	}
	myear := 2050
	res :=0
	for year,num := range maps {
		if num >=res {
			if num == res && myear> year {
				myear = year
			}else if  num !=res{
				myear = year
			}
			res = num
		}
	}
	return myear
}
