package main

type Trie struct {
	trie  [][]int
	count []int
	index int
}

// func (t *Trie) Constructor() Trie {

// 	tmp := make([][]int, 10000)
// 	for i := 0; i < 10000; i++ {
// 		tmp[i] = make([]int, 26)
// 	}
// 	return Trie{
// 		trie:  tmp,
// 		count: make([]int, 10000),
// 		index: 0,
// 	}
// }

// func (t *Trie) Insert(s string) {
// 	p := 0
// 	for i := range s {
// 		u := s[i] - 'a'
// 		if t.trie[p][u] == 0 {
// 			t.index++
// 			t.trie[p][u] = t.index
// 		}
// 		p = t.trie[p][u]
// 	}
// 	t.count[p]++
// }

// func (t *Trie) Search(s string) bool {
// 	p := 0
// 	for i := range s {
// 		u := s[i] - 'a'
// 		if t.trie[p][u] == 0 {
// 			return false
// 		}
// 		p = t.trie[p][u]
// 	}
// 	return t.count[p] != 0
// }

// func (t *Trie) StartsWith(s string) bool {
// 	p := 0
// 	for i := range s {
// 		u := s[i] - 'a'
// 		if t.trie[p][u] == 0 {
// 			return false
// 		}
// 		p = t.trie[p][u]

// 	}
// 	return true
// }


func (t *Trie) Constructor() Trie {

	tmp := make([][]int, 10000)
	for i := 0; i < 10000; i++ {
		tmp[i] = make([]int, 26)
	}
	return Trie{
		trie:  tmp,
		count: make([]int, 10000),
		index: 0,
	}
}

func (t *Trie) Insert(s string) {
	p:=0
	for i :=range s{
		u := s[i]-'a'
		if t.trie[p][u] == 0 {
			t.index++
			t.trie[p][u] = t.index
		}
		p = t.trie[p][u]
	}
	t.count[p]++
}

func (t *Trie) Search(s string) bool {
	p :=0 
	for i := range s {
		 u := s[i]-'a'
		 if t.trie[p][u] ==0 {
			 return false
		 }
		 p = t.trie[p][u]
	}
	return t.count[p] !=0
}

func (t *Trie) StartsWith(s string) bool {
	p :=0 
	for i := range s {
		 u := s[i]-'a'
		 if t.trie[p][u] ==0 {
			 return false
		 }
		p = t.trie[p][u]
	}
	return true
}
