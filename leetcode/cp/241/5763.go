package main

import "fmt"

func checkZeroOnes(s string) bool {
	max1,max0:=0,0
	for i:=0;i<len(s); {
		count :=0
		for i<len(s)&& s[i] =='1' {
			count++
			i++
		}
		max1 = max(count,max1)
		count=0
		for i<len(s) &&s[i]=='0'  {
			count++
			i++
		}
		max0 = max(count,max0)
	}
	return max1 >max0

}
func max (x,y int) int {
	if x>y {
		return x
	}
	return y
}

func main() {
	fmt.Println(checkZeroOnes("11010001011111"))
}