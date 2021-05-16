package main

import "fmt"

func main() {
fmt.Println(strStr("hello","ll"))
}


func strStr(haystack string, needle string) int {
	if len(needle)==0{
		return 0
	}

	next :=getNext(needle)

	i,j:=0,0

	for;i<len(haystack) && j<len(needle);{
		if j==-1|| haystack[i]==needle[j]{
			i++
			j++
		}else{
			j=next[j]
		}
	}
	if j==len(needle){
		return i-j
	}
	return -1
}


func getNext(needle string)[]int{
	next := make([]int,len(needle))

	i,k:=0,-1
	next[0]=-1
	for ;i<len(needle)-1;{
		if k==-1 || needle[i]==needle[k]{
			i++
			k++
			next[i]=k
		}else{
			k=next[k]
		}
	}

	return next
}