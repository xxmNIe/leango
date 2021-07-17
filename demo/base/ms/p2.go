package main

import "fmt"
/*
 继承 实现
 */
type A struct{}

func (a A) f() {
	fmt.Println("A.f()")
}
func (a A) g() {
	fmt.Println("A.g()")
}

type B struct {
	A
}

func (b B) f() {
	fmt.Println("B.f()")
}

type I interface {
	f()
}

func printType(i I) {
	fmt.Printf("%T\n", i)
	if a, ok := i.(A); ok {
		a.f()
		a.g()
	}
	if b, ok := i.(B); ok {
		b.f()
		b.g()
	}
}

func main() {
	printType(A{})
	printType(B{})
}
