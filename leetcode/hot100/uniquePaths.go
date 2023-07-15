package main

import "fmt"

/*
62. 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
*/

func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		}
		if i < 0 || j < 0 {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		res := dp(i, j-1) + dp(i-1, j)
		cache[i][j] = res
		return res
	}
	return dp(m-1, n-1)
}

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}
