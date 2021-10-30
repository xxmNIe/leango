package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Sb []byte

func (s Sb) Len() int           { return len(s) }
func (s Sb) Less(i, j int) bool { return s[i] < s[j] }
func (s Sb) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func reorderedPowerOf2(n int) bool {
	mp := make(map[string]int, 0)

	num := 0
	for i := 0; num < 10e9; i++ {
		num = 1 << i
		bs := []byte(strconv.Itoa(num))
		sort.Sort(Sb(bs))
		mp[string(bs)] = 1
	}
	nc := Sb(strconv.Itoa(n))
	sort.Sort(nc)

	if _, ok := mp[string(nc)]; ok {
		return true
	}
	return false
}

func main() {
	fmt.Println(reorderedPowerOf2(53454323))
	//fmt.Println(string([]byte(strconv.Itoa(24))))

}
