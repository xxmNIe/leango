package main

import (
	"bytes"
	"fmt"
)

// func originalDigits(s string) string {
// 	mp := make(map[byte]int, 0)
// 	for i := range s {
// 		mp[s[i]]++
// 	}
// 	var res []int
// 	sli := []string{"zero", "one", "two", "there", "four", "five", "six", "seven", "eight", "nine"}
// 	for {
// 	out:
// 		for i, sl := range sli {
// 			tmp := []byte{}
// 			for j := range sl {
// 				if mp[sl[j]] <= 0 {
// 					for _, ts := range tmp {
// 						mp[ts]++
// 					}
// 					continue out
// 				}
// 				tmp = append(tmp, sl[j])
// 				mp[sl[j]]--
// 				if mp[sl[j]] == 0 {
// 					delete(mp, sl[j])
// 				}
// 			}
// 			res = append(res, i)
// 		}
// 		if len(mp) == 0 {
// 			break
// 		}
// 	}
// 	sort.Ints(res)
// 	str := ""
// 	for _, num := range res {
// 		str += strconv.Itoa(num)
// 	}
// 	return str
// }
func originalDigits(s string) string {
	c := map[rune]int{}
	for _, ch := range s {
		c[ch]++
	}

	cnt := [10]int{}
	cnt[0] = c['z']
	cnt[2] = c['w']
	cnt[4] = c['u']
	cnt[6] = c['x']
	cnt[8] = c['g']

	cnt[3] = c['h'] - cnt[8]
	cnt[5] = c['f'] - cnt[4]
	cnt[7] = c['s'] - cnt[6]

	cnt[1] = c['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = c['i'] - cnt[5] - cnt[6] - cnt[8]

	ans := []byte{}

	for i, c := range cnt {
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, c)...)
	}
	return string(ans)
}

func main() {

	//fmt.Println(originalDigits("zeroonetwothreefourfivesixseveneightnine"))
	fmt.Println(bytes.Repeat([]byte{byte('0')}, 2))

}
