package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	words := []byte(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == words[0] {
				if dfs(i, j, board, words) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(i, j int, board [][]byte, words []byte) bool {
	m, n := len(board), len(board[0])

	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != words[0] {
		return false
	}
	if len(words) == 1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = '1'

	if dfs(i+1, j, board, words[1:]) || dfs(i, j+1, board, words[1:]) || dfs(i-1, j, board, words[1:]) || dfs(i, j-1, board, words[1:]) {
		return true
	} else {
		board[i][j] = tmp
		return false
	}
}
