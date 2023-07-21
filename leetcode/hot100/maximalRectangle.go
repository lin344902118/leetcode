package main

import "fmt"

/*
85. 最大矩形
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*/

func maximalRectangle(matrix [][]byte) int {
	ans := 0
	if len(matrix) == 0 {
		return ans
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	// 转化为柱状图，left[i][j]表示在第i行和j列之前的最大宽度
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			// 第0列，最大宽度是1，其他列的最大宽度是前面的最大宽度加一
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		for i, l := range left {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
				stk = stk[:len(stk)-1]
			}
			up[i] = -1
			if len(stk) > 0 {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = nil
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			down[i] = m
			if len(stk) > 0 {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			if area > ans {
				ans = area
			}
		}
	}
	return ans
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
