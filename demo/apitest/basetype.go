package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	m
	p = iota
)

func main() {
	fmt.Println(x, y, z, k,m, p)
}

