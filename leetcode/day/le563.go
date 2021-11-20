package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	res := 0
	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lv := f(node.Left)
		rv := f(node.Right)
		res += int(math.Abs(float64(lv) - float64(rv)))
		return node.Val + lv + rv
	}
	f(root)
	return res
}

func main() {
	t := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(findTilt(t))
}
