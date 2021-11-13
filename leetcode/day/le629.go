package main

func kInversePairs(n int, k int) int {
	const mod int = 1e9 + 7
	f := [2][]int{make([]int, k+1), make([]int, k+1)}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			cur := i & 1
			pre := cur ^ 1
			f[cur][j] = 0
			if j > 0 {
				f[cur][j] = f[cur][j-1]
			}
			if j >= i {
				f[cur][j] -= f[pre][j-i]
			}
			f[cur][j] += f[pre][j]
			if f[cur][j] >= mod {
				f[cur][j] -= mod
			} else if f[cur][j] < 0 {
				f[cur][j] += mod
			}
		}
	}
	return f[n&1][k]
}
