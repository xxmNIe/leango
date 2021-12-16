package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	head := &ListNode{}
	p := head
	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		p.Next = v
		p = p.Next
		if v.Next != nil {
			heap.Push(h, v.Next)
		}
	}
	return head.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
