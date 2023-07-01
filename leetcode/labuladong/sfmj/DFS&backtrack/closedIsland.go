package main

/*
1254. 统计封闭岛屿的数目
二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，
封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
请返回 封闭岛屿 的数目。
*/

func closedIsland(grid [][]int) int {
	res := 0
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
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
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func main() {
}
