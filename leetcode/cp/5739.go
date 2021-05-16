package main

import "fmt"

func longestBeautifulSubstring(word string) int {
	lens :=len(word)
	if lens <5{
		return 0
	}

	maxd :=0

	for i:=0;i<lens;{
		tmpl:=0
		var head uint8 ='a'
		step :=1
		if word[i]!='a'{
			i++
			continue
		}
		for ;i<lens && word[i]>=head;{
			if word[i]>head{
				step++
			}
			head = word[i]
			i++
			tmpl++
		}
		if step== 5 {
			maxd = max(maxd, tmpl)
		}
	}
	return maxd
}

func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou"))
}
