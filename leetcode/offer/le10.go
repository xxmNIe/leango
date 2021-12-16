package main

import "fmt"

// func fib(n int) int {
// 	mp := make(map[int]int, 0)
// 	var f func(x int) int
// 	f = func(x int) int {
// 		if x < 2 {
// 			return x
// 		}
// 		var r1, r2 int
// 		r1, ok := mp[x-1]
// 		if !ok {
// 			r1 = f(x - 1)
// 			mp[x-1] = r1
// 		}
// 		r2, ok = mp[x-2]
// 		if !ok {
// 			r2 = f(x - 2)
// 			mp[x-2] = r2
// 		}

// 		res := (r1 + r2) % (1e9 + 7)
// 		mp[x] = r1 + r2
// 		return res
// 	}
// 	return f(n)
// }

func fib(n int) int {
	if n < 2 {
		return n
	}
	mp := make([]int, n+1)
	mp[0], mp[1] = 0, 1
	for i := 2; i < n+1; i++ {
		mp[i] = (mp[i-1] + mp[i-2]) % (1e9 + 7)
	}
	return mp[n]
}

func main() {
	fmt.Println(fib(45))
}
