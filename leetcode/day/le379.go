package main

import "fmt"

// func integerReplacement(n int) int {
// 	if n == 1 {
// 		return 0
// 	}
// 	if n%2 == 0 {
// 		return 1 + integerReplacement(n/2)
// 	} else {
// 		return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
// 	}

// }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func integerReplacement(n int) int {
// 	mome := map[int]int{}
// 	var f func(int) int
// 	f = func(k int) (v int) {
// 		if k == 1 {
// 			return 0
// 		}

// 		if v, ok := mome[k]; ok {
// 			return v
// 		}
// 		if k%2 == 0 {
// 			v = 1 + f(k/2)
// 		} else {
// 			v = 1 + min(f(k+1), f(k-1))
// 		}
// 		mome[k] = v
// 		return v
// 	}
// 	return f(n)
// }

func integerReplacement(n int) int {
	step := 0
	for n != 1 {
		if n%2 == 0 {
			n >>= 1
		} else {
			if n != 3 && (n>>1)&1 == 1 {
				n++
			} else {
				n--
			}
		}
		step++
	}
	return step
}

func main() {
	fmt.Println(integerReplacement(65535))
}
