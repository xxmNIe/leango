package main

import "strconv"

func getMinSwaps(num string, k int) int {
	lens := len(num)
	bytes := []byte(num)
	for k>0 {
		for i := lens - 1; i > 0 && k>0; i-- {
			for j := i - 1; j > 0; j-- {
				if bytes[i] > bytes[j] && k>0 {
					bytes[i],bytes[j] = bytes[j],bytes[i]
					k--
					break
				}
			}
		}
	}
	s,_ := strconv.Atoi(string(bytes))
	return s
}

func main() {
	getMinSwaps("5489355142",4)
}