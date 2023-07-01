package main

import (
	"fmt"
	"strings"
)

/*
⼒扣第 694 题「不同的岛屿数量」，题⽬还是输⼊⼀个⼆维矩阵，0 表示海⽔，1 表示陆地，这次让你计算
不同的 (distinct) 岛屿数量
*/

func numDistinctIslands(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var sb strings.Builder
	var dfs func(grid [][]int, i, j int, dir byte)
	dfs = func(grid [][]int, i, j int, dir byte) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		sb.WriteByte(dir)
		sb.WriteByte(',')
		dfs(grid, i-1, j, '1') // 上
		dfs(grid, i+1, j, '2') // 下
		dfs(grid, i, j-1, '3') // 左
		dfs(grid, i, j+1, '4') // 右
		sb.WriteByte('-')
		sb.WriteByte(dir)
		sb.WriteByte(',')
	}
	directions := make(map[string]struct{})
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				sb = strings.Builder{}
				dfs(grid, i, j, '#')
				fmt.Println(sb.String())
				directions[sb.String()] = struct{}{}
			}
		}
	}
	return len(directions)
}

func main() {
	grids := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}
	fmt.Println(numDistinctIslands(grids))
}
