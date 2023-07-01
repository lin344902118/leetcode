package main

/*
695. 岛屿的最大面积
给你一个大小为 m x n 的二进制矩阵 grid 。
岛屿 是由一些相邻的 1 (代表土地) 构成的组合，
这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。
你可以假设 grid 的四个边缘都被 0（代表水）包围着。
岛屿的面积是岛上值为 1 的单元格的数目。
计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
*/

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	m := len(grid)
	n := len(grid[0])
	var dfs func(grid [][]int, i, j int) int
	// 计算岛屿面积
	dfs = func(grid [][]int, i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1) + 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				tmp := dfs(grid, i, j)
				if tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}

func main() {
}
