package main

import (
	"fmt"
	"math"
)

/*
931. 下降路径最小和
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。
在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者
(row + 1, col + 1) 。
*/

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = 66666
		}
	}
	var dp func(matrix [][]int, i, j int) int
	dp = func(matrix [][]int, i, j int) int {
		if i < 0 || j < 0 || i >= n || j >= n {
			return 66666
		}
		if i == 0 {
			return matrix[0][j]
		}
		if cache[i][j] != 66666 {
			return cache[i][j]
		}
		min := dp(matrix, i-1, j-1)
		mid := dp(matrix, i-1, j)
		right := dp(matrix, i-1, j+1)
		if mid < min && mid < right {
			min = mid
		}
		if right < min && right < mid {
			min = right
		}
		cache[i][j] = matrix[i][j] + min
		return cache[i][j]
	}
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		tmp := dp(matrix, n-1, i)
		if tmp < res {
			res = tmp
		}
	}
	return res
}

func main() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	matrix = [][]int{
		{-19, 57},
		{-40, -5},
	}
	fmt.Println(minFallingPathSum(matrix))
}
