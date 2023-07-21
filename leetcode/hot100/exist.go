package main

import (
	"fmt"
)

/*
79. 单词搜索
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	// 避免走回头路
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	// board[i][j]与word[k]是否匹配
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		// 不匹配
		if board[i][j] != word[k] {
			return false
		}
		// 匹配
		if k == len(word)-1 {
			return true
		}
		used[i][j] = true
		defer func() {
			used[i][j] = false
		}()
		for _, direction := range directions {
			newI := i + direction[0]
			newJ := j + direction[1]
			// 越界检查以及是否未使用
			if newI >= 0 && newI < m && newJ >= 0 && newJ < n && !used[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
	word = "SEE"
	fmt.Println(exist(board, word))
	word = "ABCB"
	fmt.Println(exist(board, word))
	// board = [][]byte{
	// 	{'A', 'A', 'A', 'A', 'A', 'A'},
	// 	{'A', 'A', 'A', 'A', 'A', 'A'},
	// 	{'A', 'A', 'A', 'A', 'A', 'A'},
	// 	{'A', 'A', 'A', 'A', 'A', 'A'},
	// 	{'A', 'A', 'A', 'A', 'A', 'A'},
	// 	{'A', 'A', 'A', 'A', 'A', 'A'},
	// }
	// word = "AAAAAAAAAAAAAAB"
	// fmt.Println(exist(board, word))
}
