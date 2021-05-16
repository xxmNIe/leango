package main

import "fmt"

type ListNode struct {
      Val int
      Next *ListNode
}

func main() {
      l :=&ListNode{Val: 1,Next: &ListNode{Val: 2,Next: &ListNode{Val: 3,Next: &ListNode{4,nil}}}}

      r := swapPairs(l)
      c :=0
      for r!=nil {
            fmt.Println(r.Val,"\t")
            c++
            if c>10{
                  break
            }
            r = r.Next
      }
}


func swapPairs(head *ListNode) *ListNode {
      p := head
      pre :=head
      for p != nil {
            if p.Next !=nil {
                if pre == head {
                      tmp := p.Next
                      p.Next = tmp.Next
                      tmp.Next = p
                      head = tmp
                      pre = p
                }else{
                      tmp := p.Next
                      pre.Next = tmp
                      p.Next = tmp.Next
                      tmp.Next = p
                      pre = p
                }
            }
            p = p.Next
      }
      return head
}


