package main

import "fmt"

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	ans := make([]int, n)
	g := make([][]int, n)
	for _, rich := range richer {
		g[rich[1]] = append(g[rich[1]], rich[0])
	}
	for i := range ans {
		ans[i] = -1
	}
	var dfs func(int)
	dfs = func(x int) {
		if ans[x] != -1 {
			return
		}
		ans[x] = x
		for _, y := range g[x] {
			dfs(y)
			if quiet[ans[y]] < quiet[ans[x]] {
				ans[x] = ans[y]
			}
		}
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	return ans

}
func main() {
	fmt.Println(loudAndRich([][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
}
