package main

import "fmt"

/*
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

func largestRectangleArea(heights []int) int {
	// 暴力
	// m := len(heights)
	// if m == 0 {
	// 	return 0
	// }
	// res := 0
	// for i := 0; i < m; i++ {
	// 	left := i
	// 	for left > 0 && heights[left-1] >= heights[left] {
	// 		left--
	// 	}
	// 	right := i
	// 	for right < m-1 && heights[right+1] >= heights[right] {
	// 		right++
	// 	}
	// 	if res < (right-left+1)*heights[i] {
	// 		res = (right - left + 1) * heights[i]
	// 	}
	// }
	// return res
	// 单调栈
	m := len(heights)
	// 记录第i个柱子最左边的柱子的编号
	left := make([]int, m)
	// 记录第i个柱子最右边的柱子的编号
	right := make([]int, m)
	for i := 0; i < m; i++ {
		right[i] = m
	}
	monoStack := make([]int, 0)
	// 计算左边柱子的编号
	for i := 0; i < m; i++ {
		// 将大于当前元素的栈顶元素删除，保证栈的单调性,删除的时候就确定了被删除元素的右边界
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		// i左边的柱子编号为栈顶元素,没有栈顶元素的为-1
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		// 将当前元素索引入栈
		monoStack = append(monoStack, i)
	}
	res := 0
	// 计算最大矩形面积
	for i := 0; i < m; i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if area > res {
			res = area
		}
	}
	return res
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
	heights = []int{2, 4}
	fmt.Println(largestRectangleArea(heights))
	heights = []int{2, 1, 5, 6, 2, 3, 7, 8}
	fmt.Println(largestRectangleArea(heights))
	heights = []int{2, 1, 2}
	fmt.Println(largestRectangleArea(heights))
	heights = []int{5, 4, 4, 6, 3, 2, 9, 5, 4, 8, 1, 0, 0, 4, 7, 2}
	fmt.Println(largestRectangleArea(heights))
	heights = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(largestRectangleArea(heights))
}
