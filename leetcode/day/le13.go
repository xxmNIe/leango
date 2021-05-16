package main

import "fmt"

func main() {
	//keys := []int32{'M','D','C','L','X','V','I'}
	//values :=[]int{1000,500,100,50,10,5,1}
	fmt.Println(romanToInt("XXVII"))


}


func romanToInt(s string) int {
	res :=0
	for i :=range s {
		if i==0 {
			res +=get(s[i])
			continue
		}
		if get(s[i]) >get(s[i-1]){
			res -=2*get(s[i-1])
			res +=get(s[i])
		}else {
			res +=get(s[i])
		}
	}
	return res
}
func get(c uint8)int{
	switch c {
		case 'M':
			return 1000
		case 'D':
			return 500
		case 'C':
			return 100
		case 'L':
			return 50
		case 'X':
			return 10
		case 'V':
			return 5
		case 'I':
			return 1
	default:
		return 0
	}
}
