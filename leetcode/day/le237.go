package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	pre, cur := node, node.Next

	for cur.Next != nil {
		v := cur.Val
		pre.Val = v
		cur = cur.Next
		pre = pre.Next
	}
	pre.Val = cur.Val
	pre.Next = nil
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	deleteNode(l1.Next.Next.Next)
	display(l1)

}

func display(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
}
