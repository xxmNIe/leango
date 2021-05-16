package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeKdigits2("1432219",3))
}

func removeKdigits2(num string, k int) string {
	stack :=[]byte{}

	for i:=range num{
		digit :=num[i]
		for k>0 && len(stack)>0 && digit <stack[len(stack)-1]{
			stack = stack[:len(stack)-1]
			k--
		}
		stack=append(stack,digit)
	}
	stack = stack[:len(stack)-k]

	ans := strings.TrimLeft(string(stack),"0")
	if ans ==""{
		return "0"
	}
	return ans



}

func removeKdigits(num string, k int) string {
	lens := len(num)
	res  := make([]byte,0)
	font :=0
	tail :=1
	for ;tail<lens;{
		for num[tail] >=num[font]&&k>0 {
			tail++
			k--
		}
		if k ==0 {
			res = append(res, num[font])
			break
		}
		if num[tail] < num[font]{
			k--
		}
		font = tail
		tail++
	}
	for tail < lens{
		res = append(res, num[tail])
		tail = tail+1
	}
	i:=0
	for ;i<len(res);i++{
		if res[i] != '0'{
			break
		}
	}
	ss := string(res[i:])
	if ss==""{
		ss="0"
	}
	return ss
}