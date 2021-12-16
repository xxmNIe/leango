package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeKLists(lists []*ListNode) *ListNode {
// 	var list *ListNode
// 	for _, l := range lists {
// 		list = mergeTwo(list, l)
// 	}
// 	return list
// }

// func mergeTwo(a, b *ListNode) *ListNode {
// 	if a == nil || b == nil {
// 		if a == nil {
// 			return b
// 		}
// 		return a
// 	}

// 	head := &ListNode{}
// 	tail, aPtr, bPtr := head, a, b
// 	for aPtr != nil && bPtr != nil {
// 		if aPtr.Val < bPtr.Val {
// 			tail.Next = aPtr
// 			aPtr = aPtr.Next
// 		} else {
// 			tail.Next = bPtr
// 			bPtr = bPtr.Next
// 		}
// 		tail = tail.Next
// 	}
// 	if aPtr != nil {
// 		tail.Next = aPtr
// 	}
// 	if bPtr != nil {
// 		tail.Next = bPtr
// 	}
// 	return head.Next
// }

func mergeKLists(lists []*ListNode) *ListNode {

	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	return mergeTwo(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwo(a, b *ListNode) *ListNode {
	if a == nil || b == nil {
		if a == nil {
			return b
		}
		return a
	}

	head := &ListNode{}
	tail, aPtr, bPtr := head, a, b
	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}
	if aPtr != nil {
		tail.Next = aPtr
	}
	if bPtr != nil {
		tail.Next = bPtr
	}
	return head.Next
}
