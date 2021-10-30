package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	mp := make(map[int]int, 0)
// 	var pre, cur *ListNode
// 	pre, cur = nil, head
// 	for cur != nil {
// 		v := cur.Val
// 		if cnt := mp[v]; cnt == 2 {
// 			pre.Next = cur.Next
// 			cur.Next = nil
// 			cur = pre.Next
// 		} else {
// 			if pre == nil {
// 				pre = cur
// 			} else {
// 				pre = pre.Next
// 			}
// 			cur = cur.Next
// 		}
// 		mp[v] = +1
// 	}

// 	return head
// }

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	mp := make(map[int]int, 0)
// 	var pre, cur *ListNode
// 	pre, cur = nil, head
// 	for cur != nil {
// 		v := cur.Val
// 		mp[v] += 1
// 		cur = cur.Next
// 	}
// 	pre = nil
// 	cur = head
// 	head = nil
// 	for cur != nil {
// 		v := cur.Val
// 		if cnt := mp[v]; cnt < 2 {
// 			if pre == nil {
// 				pre = cur
// 				head = pre
// 				cur = cur.Next
// 			} else {
// 				pre.Next = cur
// 				cur = cur.Next
// 				pre = pre.Next
// 			}
// 			pre.Next = nil
// 		} else {
// 			cur = cur.Next
// 		}

// 	}
// 	return head
// }
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	empty := &ListNode{
		Next: head,
	}
	cur := empty
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return empty.Next
}
func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
	}
	display(deleteDuplicates(l))
}

func display(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
}
