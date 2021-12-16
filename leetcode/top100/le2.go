package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l, s *ListNode
	p, r := l1, l2
	for p != nil && r != nil {
		p, r = p.Next, r.Next
	}
	if p == nil {
		l = l2
		s = l1
	} else {
		l = l1
		s = l2
	}
	p, r = l, s
	carry := 0
	var pre *ListNode
	for l != nil && s != nil {
		pre = l
		tv := l.Val
		l.Val = (tv + s.Val + carry) % 10
		carry = (tv + s.Val + carry) / 10
		l, s = l.Next, s.Next
	}

	for l != nil {
		tv := l.Val
		pre = l
		l.Val = (tv + carry) % 10
		carry = (tv + carry) / 10
		l = l.Next
	}
	if carry == 1 {
		for pre.Next != nil {
			pre = pre.Next
		}
		pre.Next = &ListNode{Val: 1}
	}
	return p
}

func main() {
	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l3 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	display(addTwoNumbers(l2, l3))
}

func display(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
}
