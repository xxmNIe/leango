package main

import (
	"container/list"
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	stack := &list.List{}
	for i, b := range s {
		if (i == 0 || i == n-1) && b == ' ' {
			continue
		}
		pre := stack.Back()
		if pre == nil {
			stack.PushBack(b)
			continue
		}
		if pre.Value == ' ' && b == ' ' {
			continue
		}
		stack.PushBack(b)
	}
	res := make([]rune, 0)
	s2 := list.New()
	for stack.Len() > 0 {
		e := stack.Back()
		for e != nil && e.Value != ' ' {
			s2.PushBack(e.Value)
			stack.Remove(e)
			e = stack.Back()
		}
		e2 := s2.Back()
		for e2 != nil {
			res = append(res, e2.Value.(rune))
			s2.Remove(e2)
			e2 = s2.Back()

		}
		res = append(res, ' ')
		if e == nil {
			continue
		}
		stack.Remove(e)
	}
	return strings.TrimRight(strings.TrimLeft(string(res), " "), " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
