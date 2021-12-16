package main

import "fmt"

// func canConstruct(ransomNote string, magazine string) bool {
// 	mp := make(map[byte]int, 0)
// 	for i := range magazine {
// 		mp[magazine[i]]++
// 	}
// 	for i := range ransomNote {
// 		val, ok := mp[ransomNote[i]]
// 		if !ok || val < 1 {
// 			return false
// 		}
// 		mp[ransomNote[i]] = val - 1
// 	}
// 	return true
// }

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := [26]int{}
	for _, v := range magazine {
		cnt[v-'a']++
	}
	for _, v := range ransomNote {
		if cnt[v-'a'] < 1 {
			return false
		} else {
			cnt[v-'a']--
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("ac", ""))
}
