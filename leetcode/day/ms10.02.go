package main

import (
  "fmt"
  "sort"
)

func groupAnagrams(strs []string) [][]string {
  mp :=map[string][]string{}
  for _,str := range strs {
    s := []byte(str)
    sort.Slice(s, func(i, j int) bool {
      return s[i] < s[j]
    })
    sstr := string(s)
    mp[sstr] = append(mp[sstr], str)
  }
  ans :=make([][]string,0,len(mp))
  for _,s :=range mp {
    ans = append(ans, s)
  }
  return ans
}

func main() {
  fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}