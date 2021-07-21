package main

import "fmt"

func addRungs(rungs []int, dist int) int {
	step :=0
	add:=0
	for _,num :=range rungs {
		//for num-step >dist{
		//	step += dist
		//	add +=1
		//}
		distence :=num-step
		if  distence > dist {
			if distence%dist ==0{
				add+= distence/dist-1
			}else {
				add+= distence/dist
			}

		}
		step = num
	}
	return add
}

func main() {
	//fmt.Println(addRungs([]int{3,6,8,10},3))
	fmt.Println(addRungs([]int{3},1))
	fmt.Println(addRungs([]int{1,3,5,10},2))
}