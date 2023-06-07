package main

import "fmt"

/*
  杨辉三角
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}
	// 等于1直接返回
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := make([][]int, 0, numRows)
	// 将1加入
	res = append(res, []int{1})
	// 生成第2到第numRows行
	for i := 2; i <= numRows; i++ {
		// 每行的个数同行数
		tmp := make([]int, i)
		// 第一个为1
		tmp[0] = 1
		// 后面的为上一行的相邻的值
		for j := 1; j < i-1; j++ {
			tmp[j] = res[len(res)-1][j-1] + res[len(res)-1][j]
		}
		// 最后一个为1
		tmp[i-1] = 1
		res = append(res, tmp)
	}
	return res
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(generate(i))
	}
}
