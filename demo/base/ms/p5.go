package main

import "fmt"

/**
slice 扩容
 */

func test_append(s []int) []int {
	for i:=0;i<1000;i++ {
		s = append(s,i)
	}
	return s
}

func main() {
	s := make([]int,0,4)
	c:= test_append(s)
	fmt.Printf("%p\n",s)
	fmt.Printf("%p\n",c)
	fmt.Println(s)
}
