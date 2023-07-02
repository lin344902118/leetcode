package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
773. 滑动谜题
在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示。
一次 移动 定义为选择 0 与一个相邻的数字（上下左右）进行交换.
最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
给出一个谜板的初始状态 board ，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。
*/

func slidingPuzzle(board [][]int) int {
	m, n := 2, 3
	sb := strings.Builder{}
	target := "123450"
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sb.WriteString(strconv.Itoa(board[i][j]))
		}
	}
	start := sb.String()
	// 每个位置相邻元素的索引
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}
	queue := make([]string, 0)
	visited := make(map[string]struct{}, 0)
	queue = append(queue, start)
	visited[start] = struct{}{}

	step := 0
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			idx := strings.Index(cur, "0")
			for _, n := range neighbor[idx] {
				// 交换0与相邻元素位置
				tmp := []byte(cur)
				tmp[idx], tmp[n] = tmp[n], tmp[idx]
				newBoard := string(tmp)
				_, exist := visited[newBoard]
				if !exist {
					queue = append(queue, newBoard)
					visited[newBoard] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

func main() {
	board := [][]int{
		{1, 2, 3},
		{5, 4, 0},
	}
	fmt.Println(slidingPuzzle(board))
}
