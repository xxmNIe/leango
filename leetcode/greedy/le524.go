package main

import "fmt"

// func findLongestWord(s string, dictionary []string) string {
// 	ans := ""
// 	for _, dic := range dictionary {
// 		l, r := 0, 0
// 		for l < len(s) && r < len(dic) {
// 			for l < len(s) && s[l] != dic[r] {
// 				l++
// 			}
// 			if l < len(s) {
// 				l++
// 				r++
// 			}
// 		}
// 		if r == len(dic) {
// 			if len(dic) > len(ans) {
// 				ans = dic
// 			} else if len(dic) == len(ans) && dic < ans {
// 				ans = dic
// 			}
// 		}
// 	}
// 	return ans
// }
func findLongestWord(s string, dictionary []string) string {
	ans := ""
	for _,t := range dictionary {
		i := 0
		for j := 0; j < len(s); j++ {
			if s[j] == t[i] {
				i++
			}
			if i == len(t){
				if len(t) > len(ans) ||(len(t) == len(ans) && t < ans){
					ans = t
				}
				break
			}
		}

	}
	return ans
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}
