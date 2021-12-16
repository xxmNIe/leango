package main

func findAnagrams(s string, p string) []int {
	slen, plen := len(s), len(p)
	if slen < plen {
		return []int{}
	}
	var sCount, pCount [26]int
	ans := make([]int, 0)
	for i, ch := range p {
		pCount[ch-'a'] += 1
		sCount[s[i]-'a'] += 1
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	for i, ch := range s[:slen-plen] {
		sCount[ch-'a']--
		sCount[s[i+plen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func findAnagrams(s string, p string) []int {
	slen, plen := len(s), len(p)
	if slen < plen {
		return []int{}
	}
	ans := make([]int, 0)
	diff := [26]int{}
	for i, ch := range p {
		diff[ch-'a'] -= 1
		diff[s[i]-'a'] += 1
	}
	cnt := 0
	for _, n := range diff {
		if n != 0 {
			cnt++
		}
	}
	if cnt == 0 {
		ans = append(ans, 0)
	}
	for i, ch := range s[:slen-plen] {
		if diff[ch-'a'] == 1 {
			cnt--
		} else if diff[ch-'a'] == 0 {
			cnt++
		}
		diff[ch-'a']--

		if diff[s[i+plen]-'a'] == -1 {
			cnt--
		} else if diff[s[i+plen]-'a'] == 0 {
			cnt++
		}
		diff[s[i+plen]-'a']++
		if cnt == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
