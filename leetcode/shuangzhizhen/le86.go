package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func partition(head *ListNode, x int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	empty := &ListNode{
// 		Next: head,
// 	}
// 	var l, r *ListNode = &ListNode{}, nil
// 	r = l
// 	p := l
// 	for empty.Next != nil {
// 		v := empty.Next
// 		empty.Next = empty.Next.Next
// 		if v.Val < x {
// 			v.Next = p.Next
// 			p.Next = v
// 			p = v
// 			if r != l && r.Val >= x {

// 			} else {
// 				r = p
// 			}
// 		} else {
// 			r.Next = v
// 			r = r.Next

// 		}

// 	}
// 	r.Next = nil
// 	return l.Next
// }

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallhead := small
	large := &ListNode{}
	largehead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largehead.Next
	return smallhead.Next
}

func main() {
	l := &ListNode{
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
	display(partition(l, 3))
}

func display(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
}
