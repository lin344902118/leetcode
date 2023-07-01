package main

import "fmt"

/*
51. N 皇后
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

func solveNQueens(n int) [][]string {
	// 构造棋盘
	board := make([][]byte, n)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = '.'
	}
	for i := 0; i < n; i++ {
		tmp := make([]byte, n)
		copy(tmp, buf)
		board[i] = tmp
	}
	res := make([][]string, 0)
	var backtrack func(board [][]byte, row int)
	// 回溯
	backtrack = func(board [][]byte, row int) {
		// 到达最后一行
		if len(board) == row {
			res = append(res, convert(board))
			return
		}
		for col := 0; col < len(board[row]); col++ {
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = 'Q'
			backtrack(board, row+1)
			board[row][col] = '.'
		}
	}
	backtrack(board, 0)
	return res
}

func convert(s [][]byte) []string {
	res := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, string(s[i]))
	}
	return res
}

func isValid(board [][]byte, row, col int) bool {
	n := len(board)
	// 判断列是否有冲突
	for i := 0; i <= row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 判断右上方是否有冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 判断左上方是否有冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func main() {
	n := 8
	fmt.Println(solveNQueens(n))
}
