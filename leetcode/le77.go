package main

import "fmt"

var res [][]int
func combine(n int, k int) [][]int {
	if n==0 ||k==0||n<k{
		return [][]int{}
	}
	res= make([][]int,0)
	//visited :=make([]bool,n)
	li := make([]int,0)
	generate(1,n,li,k)
	return res
}

func generate(begin,n int,li []int,deep int,){
	if len(li)==deep{
		des := make([]int,deep)
		copy(des,li)
		res = append(res,des)
		return
	}
	for i:=begin;i<=n;i++{
		l := len(li)
		li = append(li,i)
		generate(i+1,n,li,deep)

		li = li[:l]

	}
}
func main() {
	fmt.Println(combine(4,3))

}