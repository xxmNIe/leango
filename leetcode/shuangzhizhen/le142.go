package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {
	twoNode := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	nodeEnd := &ListNode{
		Val:  -4,
		Next: twoNode,
	}
	lc := &ListNode{
		Val:  3,
		Next: twoNode,
	}
	twoNode.Next.Next = nodeEnd
	fmt.Printf("%v", detectCycle(lc))
}

// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return nil
// 	}
// 	fast, low := head.Next, head
// 	for fast != nil && low != nil {
// 		if fast == low {
// 			break
// 		}
// 		if fast.Next == nil {
// 			return nil
// 		}
// 		fast = fast.Next.Next
// 		low = low.Next
// 	}
// 	cntCycle := 1
// 	for fast.Next != low {
// 		cntCycle++
// 		fast = fast.Next
// 	}
// 	p := head
// 	twoFlag := 0
// 	cntFull := 0
// 	for twoFlag < 2 {
// 		if twoFlag == 2 {
// 			break
// 		}
// 		if p == low {
// 			twoFlag++
// 		}

// 		p = p.Next
// 		cntFull++
// 	}
// 	diff := cntFull - 2*cntCycle
// 	r := head
// 	for diff > 1 {
// 		r = r.Next
// 		diff--
// 	}
// 	return r

// }

// func main() {
// 	twoNode := &ListNode{
// 		Val: 2,
// 		Next: &ListNode{
// 			Val:  0,
// 			Next: nil,
// 		},
// 	}
// 	nodeEnd := &ListNode{
// 		Val:  -4,
// 		Next: twoNode,
// 	}
// 	lc := &ListNode{
// 		Val:  3,
// 		Next: twoNode,
// 	}
// 	twoNode.Next.Next = nodeEnd
// 	// p := lc
// 	// cnt := 10
// 	// for cnt > 0 {
// 	// 	fmt.Println(p.Val)
// 	// 	p = p.Next
// 	// 	cnt--
// 	// }
// 	fmt.Printf("%v", detectCycle(lc))
// }
