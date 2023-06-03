package main

import "fmt"

/*
  旋转图像
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
*/

func rotateImage(matrix [][]int) {
	length := len(matrix)
	//先上下交换
	for i := 0; i < length/2; i++ {
		temp := matrix[i]
		matrix[i] = matrix[length-i-1]
		matrix[length-i-1] = temp
	}
	//在按照对角线交换
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateImage(matrix)
	fmt.Println(matrix) // [[7,4,1],[8,5,2],[9,6,3]]
	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotateImage(matrix)
	fmt.Println(matrix) // [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
}
