package main

import "fmt"

// func buddyStrings(s string, goal string) bool {
// 	if s == goal {
// 		t_mp := [26]int{}
// 		for i := range s {
// 			if t_mp[s[i]-'a'] > 0 {
// 				return true
// 			}
// 			t_mp[s[i]-'a']++
// 		}
// 		return false
// 	}
// 	ns, ng := len(s), len(goal)
// 	if ns != ng {
// 		return false
// 	}
// 	mp := make(map[byte]int, 0)
// 	for i := range s {
// 		mp[s[i]] += 1
// 	}
// 	for i := range goal {
// 		if cnt, ok := mp[goal[i]]; ok {
// 			if cnt <= 0 {
// 				return false
// 			}
// 			mp[goal[i]]--
// 			continue
// 		}
// 		return false
// 	}
// 	return true
// }

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		seen := [26]bool{}
		for _, ch := range s {
			if seen[ch-'a'] == true {
				return true
			}
			seen[ch-'a'] = true
		}
	}
	f, se := -1, -1
	for i := range s {
		if s[i] != goal[i] {
			if f == -1 {
				f = i
			} else if se == -1 {
				se = i
			} else {
				return false
			}
		}
	}
	return se != -1 && s[f] == goal[se] && s[se] == goal[f]
}

func main() {
	fmt.Println(buddyStrings("ab", "ab"))
}
