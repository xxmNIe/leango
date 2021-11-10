package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil {
		return []int{}
	}
	var pre, cur = head, head.Next

	cp := make([]int, 0)
	cnt := 1
	for cur.Next != nil {
		if (cur.Val > pre.Val && cur.Val > cur.Next.Val) || (cur.Val < pre.Val && cur.Val < cur.Next.Val) {
			cp = append(cp, cnt)
		}
		pre = pre.Next
		cur = cur.Next
		cnt++
	}

	n := len(cp)
	if n < 2 {
		return []int{-1, -1}
	}
	res := []int{10e5, cp[n-1] - cp[0]}
	for i := 1; i <= n-1; i++ {
		if cp[i]-cp[i-1] < res[0] {
			res[0] = cp[i] - cp[i-1]
		}
	}
	return res

}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 2,
									Next: &ListNode{
										Val: 7,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(nodesBetweenCriticalPoints(l1))
}
