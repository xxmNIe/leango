package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abfc","aedbc"))
}


func longestCommonSubsequence(text1 string, text2 string) int {

	len1,len2 := len(text1),len(text2)
	dp :=make([][]int,len1+1)

	for i:=0;i<len1+1;i++{
		dp[i] = make([]int,len2+1)
	}

	for i:=0;i<len1;i++{
		for j := 0;j<len2;j++{
		//	fmt.Println(i+1,":",j+1)
			if text1[i]==text2[j]{
				dp[i+1][j+1] = dp[i][j]+1
			}else{

				dp[i+1][j+1] = max(dp[i+1][j],dp[i][j+1])
			}
		}
	}
	return dp[len1][len2]
}

func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}
