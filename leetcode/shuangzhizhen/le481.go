package main

import "fmt"

func magicalString(n int) int {
	if n < 4 {
		return 1
	}
	arr := make([]int, n, n)
	arr[0], arr[1], arr[2] = 1, 2, 2
	left, right := 2, 3
	res := 1
	num := 1
	for right < n {
		cnt := arr[left]
		for i := 0; i < cnt && right < n; i++ {
			arr[right] = num
			if num == 1 {
				res++
			}
			right++
		}
		left++
		num = num ^ 2 ^ 1
	}
	fmt.Println(arr)
	return res
}

func main() {
	fmt.Println(magicalString(7))
}
