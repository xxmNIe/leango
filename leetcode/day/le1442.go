package main

import "fmt"

func countTriplets(arr []int) int {
	n := len(arr)
	res :=0
	pre := make([]int,len(arr)+1)
	//前缀异或
	for i:=0;i<n;i++ {
		pre[i+1] = pre[i] ^ arr[i]
	}
	for i:=0;i<n-1;i++ {
		for j:=i+1;j<n;j++ {
			for k:=j; k<n;k++ {
				a := pre[i] ^ pre[j]
				b := pre[j] ^ pre[k+1]
				if a==b {
					res++
				}
			}
		}
	}
	return res
}







//func countTriplets(arr []int) int {
//	n := len(arr)
//	res :=0
//	pre := make([]int,len(arr)+1)
//	//前缀异或
//	for i:=0;i<n;i++ {
//		pre[i+1] = pre[i] ^ arr[i]
//	}
//	for j:=1;j<n;j++ {
//		i:=j-1;k:=j
//		for i>=0 || k<n {
//			a := pre[i] ^ pre[j]
//			b := pre[j] ^ pre[k+1]
//			if a == b{
//				res++
//				fmt.Println(i,j,k)
//			}
//			if i >0 {
//				i--
//			}else {
//				if i==0{
//					i=j-1
//				}
//				if k ==n-1 {
//					break
//				}else {
//					k++
//				}
//			}
//		}
//	}
//	return res
//}

func main() {
	fmt.Println(countTriplets([]int{2,3,1,6,7}))
}