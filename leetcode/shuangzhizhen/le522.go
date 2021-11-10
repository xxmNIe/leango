package main

import "fmt"

func findLUSlength(strs []string) int {
	mp := make(map[string]int, 0)

	for _, s := range strs {
		for i := 0; i < (1 << len(s)); i++ {
			t := ""
			for j := 0; j < len(s); j++ {
				if ((i >> j) & 1) != 0 {
					t = t + string(s[j])
				}
			}
			if _, ok := mp[t]; ok {
				mp[t] = mp[t] + 1
			} else {
				mp[t] = 1
			}
		}
	}
	res := -1
	for s := range mp {
		if mp[s] == 1 {
			res = max(len(s), res)
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
}
