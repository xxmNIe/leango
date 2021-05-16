package main

import "fmt"

func main() {
	n :=[]int{2,1,3}
	nextPermutation(n)
	fmt.Println(n)
}

func nextPermutation(nums []int)  {
	n := len(nums)
	i := n-2
	//从右向左找到第一个小于后边的数
	for i>=0 && nums[i]>=nums[i+1] {
		i--
	}
	//如果 1>=0   证明找到了上述
	if i >=0 {
		j := n - 1
		//找到这个数后边第一个大于这个数的数
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		//交换此时i+1及之后的数肯定是升序，所以之后要翻转取得最小的
		nums[i],nums[j] = nums[j],nums[i]
	}
	reverse1(nums[i+1:])
}
func reverse1(a []int) {
	for i, n :=0,len(a);i < n/2;i++ {
		a[i],a[n-1-i] = a[n-1-i],a[i]
	}
}


//
//func nextPermutation(nums []int)  {
//	lens := len(nums)
//	flag :=false
//	if lens==1{
//		return
//	}
//out:	for i :=lens-1;i>0;i-- {
//			if i ==0 {
//				break
//			}
//			j :=i-1
//			for j>=0 {
//				if nums[i] >nums[j] {
//					nums[i],nums[j] = nums[j],nums[i]
//					flag = true
//					break out
//				}
//				j--
//			}
//			i--
//		}
//	if flag == true{
//		return
//	}
//	for i :=0;i<lens-1;i++ {
//		flag = false
//		for j:=0;j<lens-1;j++ {
//			if nums[j]>nums[j+1] {
//				nums[j],nums[j+1] = nums[j+1],nums[j]
//				flag = true
//			}
//		}
//		if flag == false {
//			return
//		}
//	}
//}
