package main

import "fmt"

func main() {
	a := 0x7fffffffffffffff
	fmt.Printf("%x\n",a)
	//out :7fffffffffffffff
	b:=-0x7fffffffffffffff
	fmt.Printf("%x\n",b)
	//out:-7fffffffffffffff
	//c := 0xffffffffffffffff
	//fmt.Printf("%x",c)
	//err :超出范围

	//a =a<<1
	//fmt.Printf("%x\n",a)
	//out:-2  证明位移包括符号位     7fffffffffffffff -> fffffffffffffffe
	a = a>>1
	fmt.Printf("%x\n",a)
	//out 3fffffffffffffff   正数左右位移都是补0
	a= a>>100
	//out -1


}
