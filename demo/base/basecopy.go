package main

import "fmt"

type A interface {
	Len() int
}

type B struct {
	a string
}

func (b *B) Len() int {
	return len(b.a)
}

var mp = make(map[string]A, 0)

func main() {
	b := &B{"abc"}
	mp["a"] = b
	d := mp["a"]
	e := d.(A)
	b.a = "sdasdss"
	fmt.Println(b == e)
	fmt.Println(b, "  ", e)

}
