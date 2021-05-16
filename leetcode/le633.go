package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	if c<=2{
		return true
	}

	sqrtnum:= int(math.Sqrt(float64(c)))
	i,j :=0,sqrtnum

	for i<=j{
		i2 :=i*i
		j2:=j*j
		if i2+j2 == c{
			return true
		}else if i2+j2 < c{
			i++
		}else{
			j--
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(6))
}