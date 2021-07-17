package main

import "fmt"
/*
  string   a 地址不改变
 */
func main() {
	a := "012"
	pa := &a
	b := []byte(*pa)
	pb := &b

	fmt.Printf("%p\n",&a)
	a = a+"3"
	fmt.Printf("%p\n",&a)
	*pa += "4"
	b[1] = '5'

	println(*pa)
	println(string(*pb))
}



