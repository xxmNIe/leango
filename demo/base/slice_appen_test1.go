package main

import "fmt"

type Iter interface {
	Foo()
	Bar()
}

type Base struct {
	A, B string
	Q    int
}

type impl1 struct {
	Base
}

func (i1 *impl1) Foo() {
	fmt.Println("li Foo")
}
func (l2 *impl1) Bar() {
	fmt.Println("li Bar")
}

type impl2 struct {
	Base
}

func (l2 *impl2) Foo() {
	fmt.Println("l1 Foo")
}

func (l2 *impl2) Bar() {
	fmt.Println("l2 Bar")
}

type queue struct {
	arr []Iter
}

func main() {

	arr := make([]Iter, 0)
	b := impl1{
		Base{
			A: "A",
			B: "B",
			Q: 1,
		},
	}
	for i := 0; i < 5; i++ {
		t := b
		t.Q = i + 10
		fmt.Println(t == b)
		arr = append(arr, &t)
	}
	for _, c := range arr {
		fmt.Printf("%d\n", c.(*impl1).Q)
	}
}
