package main

import (
	"fmt"
	"time"
)

var a map[int]int

// func main() {
// 	var m0 map[int]int
// 	var m1 = map[int]int{}
// 	m2 := map[int]int{}
// 	m3 := make(map[int]int)
// 	m4 := new(map[int]int)
// 	fmt.Println(m0 == nil, m1 == nil, m2 == nil, m3 == nil, m4 == nil) //m0 true  other :all false
// 	m1[1] = 2
// 	m2[1] = 2
// 	m3[1] = 2
// 	//m4[1]=2   err
// 	m4 = &map[int]int{}
// 	(*m4)[1] = 1
// 	fmt.Println(m1, m2, m3, m4) //map[1:2] map[1:2] map[1:2] &map[1:1])
// }
type S struct {
	A, B string
}

var mp map[S]int

func main() {
	mp = make(map[S]int)
	a := S{"A", "b"}
	mp[a] = 1
	go func(s S) {
		a.A = "a"
		fmt.Println(s)
	}(a)
	time.Sleep(2 * time.Second)
	fmt.Println(a)
	fmt.Println(mp[a])
}
