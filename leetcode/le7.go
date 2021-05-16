package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(-123))
}


func reverse(x int) int {
	var res int32 =0
	y := int32(x)
	for y!=0 {
		tmp := y%10
		value := res*10+tmp
		if (value -tmp)/10 !=res {
			return 0
		}
		res = value
		y = y/10
	}
	return int(res)
}