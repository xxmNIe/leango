package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	p := root
	for p != nil {
		if p.Val == val {
			return p
		} else if p.Val > val {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return p
}

func main() {
	bst := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	fmt.Println(searchBST(bst, 4))
}
