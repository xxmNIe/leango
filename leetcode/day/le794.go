package main

import (
	"strings"
)

func validTicTacToe(board []string) bool {
	xCount, oCount := 0, 0
	for _, s := range board {
		xCount += strings.Count(s, "X")
		oCount += strings.Count(s, "O")
	}
	return !(xCount != oCount && oCount != xCount-1 || xCount != oCount && win(board, 'O') || oCount != xCount-1 && win(board, 'X'))
}

func win(board []string, c byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == c && board[i][1] == c && c == board[i][2] ||
			board[0][i] == c && board[1][i] == c && board[2][i] == c {
			return true
		}
	}

	if board[0][0] == c && board[1][1] == c && board[2][2] == c || board[0][2] == c && board[1][1] == c && board[2][0] == c {
		return true
	}
	return false
}
