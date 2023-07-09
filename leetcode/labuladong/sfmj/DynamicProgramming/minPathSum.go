package main

/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/

func minPathSum(grid [][]int) int {
	// dp[i][j]表示从0，0到i，j的最小路径和
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
}
