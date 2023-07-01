package main

import "fmt"

/*
1020. 飞地的数量
给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。

一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。

返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。
*/

func numEnclaves(grid [][]int) int {
	res := 0
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	// 把上下边界的岛屿淹掉
	for i := 0; i < len(grid[0]); i++ {
		dfs(grid, 0, i)
		dfs(grid, len(grid)-1, i)
	}
	// 把左右边界的岛屿淹掉
	for i := 0; i < len(grid); i++ {
		dfs(grid, i, 0)
		dfs(grid, i, len(grid[0])-1)
	}
	// 计算飞地数量
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	grid = [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numEnclaves(grid))
}
