package main

// func maxProduct(words []string) int {
// 	ans :=0
// 	masks := make([]int,len(words))
// 	for i,word := range words {
// 		for _,ch := range word{
// 			masks[i] |= 1<<(ch-'a')
// 		}
// 	}
// 	for i,x := range masks{
// 		for j,y :=range masks[:i]{
// 			if x&y ==0 && len(words[i])* len(words[j]) >ans {
// 				ans =  len(words[i])* len(words[j])
// 			}
// 		}
// 	}
// 	return ans
// }

func maxProduct(words []string) int {
	masks := make(map[int]int)
	for _,word := range words{
		mask :=0
		for _,ch := range word{
			mask |= 1<<(ch-'a')
			if len(word) > masks[mask]{
				masks[mask]  = len(word)
			}
		}
	}
	ans := 0 
	for x,lenX :=range masks{
		for y,lenY := range masks{
			if x&y == 0 && lenX*lenY > ans{
				ans = lenX*lenY
			}
		}
	}
	return ans
}
