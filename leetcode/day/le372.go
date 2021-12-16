package main

var MOD = 1337

func superPow(a int, b []int) int {
	return dfs(a, b, len(b)-1)
}

func dfs(a int, b []int, u int) int {
	if u == -1 {
		return 1
	}
	return pow(dfs(a, b, u-1), 10) * pow(a, b[u]) % MOD
}

func pow(a, b int) int {
	ans := 1
	a %= MOD
	for b > 0 {
		ans = ans * a % MOD
		b--
	}
	return ans
}
