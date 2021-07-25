package main

import "fmt"

func restoreArray1(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	mp := make(map[int]int,n)
	list := make(map[int]int,n)
	for _,v := range adjacentPairs {
		head := v[0]
		tall := v[1]
		mp[head] = tall
		list[head] +=1
		list[tall] -=1
	}
	leader :=-1
	for k,v := range list {
		if v == 1 {
			leader = k
			break
		}
	}
	res := []int{leader}
	for {
		if v,ok := mp[leader] ;ok {
			res= append(res, v)
			leader = v
		}else {
			break
		}
	}
	return res

}
func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	mp := make(map[int][]int,n)
	for _,p :=range adjacentPairs {
		v,w := p[0],p[1]
		mp[v] = append(mp[v], w)
		mp[w] = append(mp[w], v)
	}
	ans := make([]int,n+1)
	for e,adj := range mp {
		if len(adj) == 1 {
			ans[0] = e
			break
		}
	}

	ans[1] = mp[ans[0]][0]
	for i:=2;i<n+1;i++ {
		adj := mp[ans[i-1]]
		if ans[i-2] == adj[0] {
			ans[i]  = adj[1]

		}else {
			ans[i] = adj[0]
		}
	}

	return ans


}
func main() {
	fmt.Println(restoreArray([][]int{{2,1},{3,4},{3,2}}))
}