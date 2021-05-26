package main

import "fmt"

//func isMatch(s string, p string) bool {
//	ns,np := len(s),len(p)
//	dp := make([][]bool,np+1)
//	for i:=0;i<np+1;i++{
//		dp[i] = make([]bool,ns+1)
//	}
//	for i:=1;i<np+1;i++ {
//		for j :=1;j<ns+1;j++ {
//			if s[j-1] == p[i-1]  {
//				dp[i][j] = true
//			}else{
//				//不相等
//				if   p[i] == '.' && dp[i-1][j-1]{
//					dp[i][j] = true
//				}else if p[i] == '*' && dp[i-1][j-1] == true{
//					dp[i][j] = true
//				}else {
//					dp[i][j] = false
//				}
//			}
//		}
//	}
//	return dp[np][ns]
//}
func isMatch(s string, p string) bool {
	m,n := len(s),len(p)
	 matchs :=func(i,j int) bool {
	 	if i==0 {
	 		return false
		}
		if p[j-1] == '.'{
			return true
		}
		return s[i-1] == p[j-1]
	 }

	 f:=make([][]bool,m+1)
	 for i:=0;i<len(f);i++ {
	 	f[i] = make([]bool,n+1)
	 }
	 f[0][0] = true
	 for i:=0;i<=m;i++ {
	 	for j:=1;j<=n;j++ {
	 		if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matchs(i,j-1) {
					f[i][j] = f[i][j]|| f[i-1][j]
				}
			}else if matchs(i,j){
				f[i][j] = f[i][j]|| f[i-1][j-1]
			}
		}
	 }
	 return f[m][n]
}

func main() {
	fmt.Println(isMatch("aa","a"))
}