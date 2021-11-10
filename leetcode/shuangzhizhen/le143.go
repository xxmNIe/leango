package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	arr := make([]*ListNode, 0)
	p := head
	for p != nil {
		v := p.Next
		p.Next = nil
		arr = append(arr, p)
		p = v
	}
	l, r := 0, len(arr)-1
	for l < r {
		arr[l].Next = arr[r]
		if l+1 != r {
			arr[r].Next = arr[l+1]
		}
		l++
		r--
	}

}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(l)
	display(l)
}
func display(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
}
