package main

func decode(encoded []int) []int {
	n :=len(encoded)
	res := make([]int,n+1)
	a:=0
	for i:=0;i<n+1;i++{
		a ^= i
	}
	for i:=n-1;i>=0;i-=2 {
		a ^= encoded[i]
	}
	res[0]=a
	for i:=1;i<n+1;i++ {
		res[i] = res[i-1]^encoded[i-1]
	}
	return res
}