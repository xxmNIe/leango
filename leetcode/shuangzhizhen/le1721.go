package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	left, right := head, head
	var p *ListNode
	for k > 1 && right != nil {
		right = right.Next
		k--
	}
	if right == nil && k < 1 {
		return head
	}
	p = right
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	p.Val, left.Val = left.Val, p.Val
	return head

}
