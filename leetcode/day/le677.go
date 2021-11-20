package main

import (
	"fmt"
	"strings"
)

type MapSum struct {
	mp map[string]int
}

func Constructor() MapSum {
	return MapSum{mp: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.mp[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	for k := range this.mp {
		tmp, ok := this.mp[k]
		if !ok {
			continue
		}
		if strings.Index(k, prefix) == 0 {
			res += tmp
		}
	}
	return res
}

func main() {
s :="abcabccccd"
p :="bcb"

fmt.Println(strings.Index(s,p))

}
