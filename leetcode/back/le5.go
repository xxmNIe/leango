package main

// 给你一个字符e串 s，找到 s 中最长的回文子串。
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool,n)
	start,end :=0,0
	for i :=0;i<n;i++ {
		dp[i] = make([]bool,n)
		dp[i][i] = true
	}
	for L := 2;L<=n;L++ {
		for i :=0;i<n;i++ {
			j := i + L - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				}else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == true && j-i > start-end {
				start,end = i,j
			}
 		}
	}
	return s[start:end+1]
}
func main() {
	println(longestPalindrome("cbbd"))
}