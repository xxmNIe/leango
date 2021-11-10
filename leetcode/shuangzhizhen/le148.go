package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	fast, slow := head, head
	for fast != tail {
		fast = fast.Next
		slow = slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head,mid), sort(mid,tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	empty := &ListNode{}
	empty_p, p1, p2 := empty, head1, head2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			empty_p.Next = p1
			p1 = p1.Next
		} else {
			empty_p.Next = p2
			p2 = p2.Next
		}
		empty_p = empty_p.Next
	}

	//链表。。。。不需要for
	if p1 != nil {
		empty_p.Next = p1
		p1 = p1.Next
	} else if p2 != nil {
		empty_p.Next = p2
		p2 = p2.Next
	}

	return empty.Next
}
