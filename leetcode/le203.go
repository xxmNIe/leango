package main

import "fmt"

type ListNode struct {
  Val int
  Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	root :=&ListNode{Next: head}
	pre := root
	cur := root.Next
	for cur!=nil {
		if cur.Val == val{
			pre.Next = cur.Next
		}else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return root.Next
}

func main() {
	a := &ListNode{Val: 1,Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: nil,
						},
					},
				},
			},
		},
	}}
	b:=removeElements(a,6)

	for b !=nil {
		fmt.Println(b.Val)
		b = b.Next
	}
}
