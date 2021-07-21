package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetterss string) int {
	res :=0
	strs := strings.Split(text," ")
	bc := []byte(brokenLetterss)
out:	for _,str :=range strs {
		for _,c :=range bc {
			if strings.Contains(str,string(c)) {
				continue out
			}
		}
		res+=1
	}
	return res
}
func main() {
	fmt.Println(canBeTypedWords("leet code","lt"))
}
