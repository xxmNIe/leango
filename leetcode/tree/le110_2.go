package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(node * TreeNode) int {
	if node == nil{
		return 0
	}
	left := height(node.Left)
	right := height(node.Right)
	if left == -1 || right == -1 || abs(left,right) >1 {
		return  -1
	}
	return max(left,right)+1
}


func max(x,y int) int {
	if x >y {
		return x
	}
	return y
}

func abs(x,y int) int {
	if x-y > 0{
		return x-y
	}
	return y-x
}

func main() {
	m := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isBalanced(m))
}