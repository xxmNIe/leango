package main

import "fmt"

//type TreeNode struct {
//	 Val int
//	 Left *TreeNode
//	 Right *TreeNode
// }

//func isBalanced(root *TreeNode) bool {
//	if root ==nil {
//		return true
//	}
//
//
//	return abs(getdeep(root.Left),getdeep(root.Right)) <=1 && isBalanced(root.Left) && isBalanced(root.Right)
//
//}

//func getdeep(node *TreeNode) int {
//	if node == nil {
//		return 0
//	}
//	left := getdeep(node.Left)
//	right := getdeep(node.Right)
//	return max(left,right)+1
//}
//
//func max(x,y int) int {
//	if x >y {
//		return x
//	}
//	return y
//}
//
//func abs(x,y int) int {
//	if x-y > 0{
//		return x-y
//	}
//	return y-x
//}

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