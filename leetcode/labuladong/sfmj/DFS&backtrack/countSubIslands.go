package main

/*
1905. 统计子岛屿
给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。
一个 岛屿 是由 四个方向 （水平或者竖直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。

如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，
也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，
那么我们称 grid2 中的这个岛屿为 子岛屿 。

请你返回 grid2 中 子岛屿 的 数目 。
*/

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	m := len(grid1)
	n := len(grid1[0])
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
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
	// 把grid2中不是子岛屿的淹掉
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// grid2是岛屿，grid1是海水，必然不是子岛屿，淹掉
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(grid2, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(grid2, i, j)
			}
		}
	}
	return res
}

func main() {
}
