package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(jump([]int{2,3,1,1,4}))
//}
//
//func jump(nums []int) int {
//	end :=0
//	maxpos :=0
//	step :=0
//	lens :=len(nums)
//	for i:=0;i<lens-1;i++{
//		maxpos = max(maxpos,i+nums[i])
//		if i == end {
//			end = maxpos
//			step++
//		}
//	}
//	return step
//}
//
//func max(x,y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}