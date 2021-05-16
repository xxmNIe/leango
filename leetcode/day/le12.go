package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
}

func intToRoman(num int) string {
	keys :=[]int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	values := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}

}




func intToRoman2(num int) string {
	keys :=[]int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	values := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}

	res :=""
	i := 0
	for num >0 &&i<len(keys) {
		if num-keys[i] >=0 {
			res +=values[i]
			num -=keys[i]
		}else {
			i++
		}
	}
	return res
}
