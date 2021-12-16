package main

// func movingCount(m int, n int, k int) int {
// 	if k == 0 {
// 		return 1
// 	}
// 	cnt := 0
// 	mp := make([][]bool, m)
// 	for i := range mp {
// 		mp[i] = make([]bool, n)
// 	}

// 	var f func(x, y int)
// 	f = func(x, y int) {
// 		if x < 0 || x >= m || y < 0 || y >= n {
// 			return
// 		}
// 		if mp[x][y] {
// 			return
// 		}
// 		if (x%10)+(x/10)+(y/10)+(y%10) <= k {
// 			fmt.Println(x, ":", y)
// 			cnt++
// 		}
// 		mp[x][y] = true
// 		f(x+1, y)
// 		f(x-1, y)
// 		f(x, y+1)
// 		f(x, y-1)
// 	}
// 	f(0, 0)
// 	return cnt
// }

func sum(x, y int) int {
	return (x % 10) + (x / 10) + (y / 10) + (y % 10)

}

func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	cnt := 0
	mp := make([][]bool, m)
	for i := range mp {
		mp[i] = make([]bool, n)
	}
	mp[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 || sum(i, j) > k {
				continue
			}

			if i >= 1 {
				mp[i][j] = mp[i][j] || mp[i-1][j]
			}
			if j >= 1 {
				mp[i][j] = mp[i][j] || mp[i][j-1]
			}
			if mp[i][j] {
				cnt += 1
			}
		}
	}
	return cnt + 1
}

var mp map[string]int

func main() {
	//fmt.Println(movingCount(16, 8, 4))
	mp["a"] = 1
}
