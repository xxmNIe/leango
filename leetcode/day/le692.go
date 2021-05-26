package main

import (
	"container/heap"
	"fmt"
)

//func topKFrequent(words []string, k int) []string {
//	sort.Strings(words)
//	maps := make(map[string]int)
//	for _,v:=range words {
//		maps[v]+=1
//	}
//	uni := make([]string,0,len(maps))
//
//	for w := range maps {
//		uni= append(uni, w)
//	}
//	sort.Slice(uni, func(i, j int) bool {
//		s,t := uni[i],uni[j]
//		return maps[s] > maps[t] || maps[s]==maps[t] && s<t
//	})
//
//	return uni[:k]
//}

type pair struct {
	w string
	c int
}

type hp []pair

func (h hp) Len() int {return len(h)}
func (h hp) Less(i,j int) bool {a,b :=h[i],h[j];return a.c <b.c || a.c == b.c && a.w > b.w}
func (h hp) Swap(i,j int) {h[i],h[j] = h[j],h[i]}
//????
func (h *hp) Push (v interface{}) { *h = append(*h,v.(pair))}
func (h *hp) Pop () interface{} {a := *h;v :=a[len(a)-1];*h = a[:len(a)-1];return v}
func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _,w :=range words {
		cnt[w]++
	}
	h := &hp{}
	for w,c :=range cnt {
		heap.Push(h,pair{w,c})
		if h.Len()>k {
			heap.Pop(h)
		}
	}
	ans := make([]string,k)
	for i:= k-1;i>=0;i--{
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}
//type pair struct {
//	w string
//	c int
//}
//type hp []pairfunc topKFrequent(words []string, k int) []string {
//	sort.Strings(words)
//	maps := make(map[string]int)
//	for _,v:=range words {
//		maps[v]+=1
//	}
//	uni := make([]string,0,len(maps))
//
//	for w := range maps {
//		uni= append(uni, w)
//	}
//	sort.Slice(uni, func(i,jint) bool {
//		s,t := uni[i],uni[j]
//		return maps[s] > maps[t] || maps[s]==maps[t] && s<t
//	})
//
//	return uni[:k]
//}
//func (h hp) Len() int            { return len(h) }
//func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.c < b.c || a.c == b.c && a.w > b.w }
//func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
//func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
//
//func topKFrequent(words []string, k int) []string {
//	cnt := map[string]int{}
//	for _, w := range words {
//		cnt[w]++
//	}
//	h := &hp{}
//	for w, c := range cnt {
//		heap.Push(h, pair{w, c})
//		if h.Len() > k {
//			heap.Pop(h)
//		}
//	}
//	ans := make([]string, k)
//	for i := k - 1; i >= 0; i-- {
//		ans[i] = heap.Pop(h).(pair).w
//	}
//	return ans
//}


func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"},2))
}