package main

import "fmt"

func main() {
	//fmt.Println(maxDistance([]int{55,30,5,4,2},[]int{100,20,10,10,5}))
	fmt.Println(maxDistance([]int{2,2,2},[]int{10,10,1}))
}
func maxDistance(nums1 []int, nums2 []int) int {
	n1 :=len(nums1)
	n2 := len(nums2)
	i :=0
	res := 0
	for j:=0;j<n2;j++ {
		for i< n1 && nums1[i]>nums2[j]{
			i++
		}
		if i<n1 {
			res = max(res,j-i)
		}
	}
	return res
}
func max(x,y int )int {
	if x>y{
		return x
	}
	return y
}

//func maxDistance(nums1 []int, nums2 []int) int {
//	n1 :=len(nums1)
//	n2 := len(nums2)
//	res :=0
//	last :=0
//	for j:=0;j<n2;j++ {
//		if j<n2-1 && nums2[j]==nums2[j+1]{
//			continue
//		}
//		outer := nums2[j]
//		for i:=j;i>=last;i-- {
//			if i >=n1{
//				continue
//			}
//			if nums1[i]>outer  {
//				last=i
//				break
//			}
//			if i>last && nums1[i]==nums1[i-1]{
//				continue
//			}
//				res = max(res,j-i)
//
//		}
//	}
//	return res
//}
