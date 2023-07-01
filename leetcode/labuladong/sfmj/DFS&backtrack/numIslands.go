package main

import "fmt"

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/

func numIslands(grid [][]byte) int {
	res := 0
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		// 已经被淹了，返回
		if grid[i][j] == '0' {
			return
		}
		// 将陆地淹掉，并淹掉其上下左右的陆地
		grid[i][j] = '0'
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))
}
