package main

import "fmt"

func distributeCandies(candyType []int) int {
	n := len(candyType)
	if n < 4 {
		return 1
	}
	typemap := make(map[int]struct{}, 0)
	count := 0
	for _, c := range candyType {
		typemap[c] = struct{}{}
		count = len(typemap)
		if count > n/2 {
			return n / 2
		}
	}
	return count

}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}
