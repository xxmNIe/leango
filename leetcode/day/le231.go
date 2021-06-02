package main

import "fmt"

//func isPowerOfTwo(n int) bool {
//	if n<1{
//		return false
//	}
//
//	f := false
//	for i:=0;i<31;i++ {
//		if n>>i &1==1 {
//			if f {
//				return false
//			}else {
//				f = true
//			}
//		}
//	}
//	return true
//}
func isPowerOfTwo(n int) bool {
	if n <0 {
		return false
	}

	return n&-n ==n

}
func main() {
	fmt.Println(isPowerOfTwo(3))
}