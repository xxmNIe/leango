package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteListNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head
	for ; p != nil && p.Next != nil; p = p.Next {
		p.Next = p.Next.Next
	}
	return head
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

	display(deleteListNode(l1))
}

func display(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
}
