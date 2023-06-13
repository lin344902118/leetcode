package main

import "fmt"

/*
304. 二维区域和检索 - 矩阵不可变
给定一个二维矩阵 matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，
右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角
 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
*/

type NumMatrix struct {
	matrix [][]int
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{
			matrix: matrix,
		}
	}
	// 题目保证m、n>0,所以不会越界
	m := len(matrix)
	n := len(matrix[0])
	preSum := make([][]int, m+1)
	for i := 0; i < len(preSum); i++ {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSum[i][j] = matrix[i-1][j-1] + preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{
		matrix: matrix,
		preSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	m := Constructor(matrix)
	fmt.Println(m.SumRegion(2, 1, 4, 3))
	fmt.Println(m.SumRegion(1, 1, 2, 2))
	fmt.Println(m.SumRegion(1, 2, 2, 4))
}
