package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, low := head.Next, head
	for fast != nil && low != nil {
		if fast == low {
			return true
		}
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		low = low.Next
	}
	return false
}
