package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	var f func(node *Node) int
	f = func(node *Node) int {
		if node == nil {
			return 0
		}
		tdp := 0
		for _, n := range node.Children {
			tmp := f(n)
			if tmp > tdp {
				tdp = tmp
			}
		}
		return tdp + 1
	}
	return f(root)
}

func main() {
	cd2 := []*Node{&Node{}, &Node{Children: []*Node{&Node{}}}}
	cd := []*Node{&Node{}, &Node{}, &Node{Children: cd2}}
	t := &Node{Children: cd}
	fmt.Println(maxDepth(t))
}
