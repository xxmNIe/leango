package main

import "fmt"

func reverseParentheses(s string) string {
	stack := [][]byte{}
	str := []byte{}

	for i :=range s {
		if s[i] == '(' {
			stack = append(stack, str)
			str = []byte{}
		}else if s[i] == ')'{
			for i,j :=0,len(str);i<j/2;j-- {
				str[i],str[j-i-1] = str[j-1-i],str[i]
			}
			str = append(stack[len(stack)-1],str...)
			stack = stack[:len(stack)-1]
		}else {
			str = append(str, s[i])
		}
	}
	return string(str)
}



//func reverseParentheses(s string) string {
//	n := len(s)
//	stack := make([]int,n)
//	cur :=-1
//	for i,v:= range s{
//		if v == ')' {
//			cur++
//			stack[cur] = i
//
//		}
//	}
//	bytes := []byte(s)
//	var rev func(left,right int)
//	rev = func(left, right int) {
//		i,j :=left+1,right -1
//		for i < j {
//			bytes[i],bytes[j] = bytes[j],bytes[i]
//			i++
//			j--
//		}
//	}
//	cur=0
//	for i:=stack[0]-1;i>=0;i-- {
//		if cur >=len(stack){
//			break
//		}
//		for i>=0 && s[i]!='('{
//			i--
//		}
//		rev(i,stack[cur])
//		cur++
//	}
//	str :=""
//	for i:=0 ;i<n;i++ {
//		if bytes[i]=='(' || bytes[i]== ')'{
//			continue
//		}
//		str +=string(bytes[i])
//	}
//	return str
//}
func main() {
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}

