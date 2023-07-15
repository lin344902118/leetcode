package main

import "fmt"

/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap(height []int) int {
	// 柱子升序和降序都要考虑
	n := len(height)
	if n < 2 {
		return 0
	}
	res := 0
	left := -1
	// 升序需要扣掉的面积
	loss := 0
	// 找到升序所有面积
	for i := 0; i < n; i++ {
		if height[i] != 0 {
			if left == -1 {
				left = i
				continue
			}
			if height[i] < height[left] {
				loss += height[i]
			} else {
				area := (i - left - 1) * height[left]
				res += area - loss
				left = i
				loss = 0
			}
		}
	}
	// 降序等于倒过来升序
	right := -1
	loss = 0
	for i := n - 1; i >= left; i-- {
		if height[i] != 0 {
			if right == -1 {
				right = i
				continue
			} else {
				if height[i] < height[right] {
					loss += height[i]
				} else {
					area := (right - i - 1) * height[right]
					res += area - loss
					right = i
					loss = 0
				}
			}
		}
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height)) // 6
	height = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height)) // 9
	height = []int{4, 2, 3}
	fmt.Println(trap(height)) // 1
}
