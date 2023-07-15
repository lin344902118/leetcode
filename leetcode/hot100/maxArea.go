package main

import "fmt"

/*
11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。
*/

func maxArea(height []int) int {
	// // 思路一,超出时间限制
	// max := 0
	// for i := 0; i < len(height)-1; i++ {
	// 	for j := i + 1; j < len(height); j++ {
	// 		min := height[i]
	// 		if height[j] < min {
	// 			min = height[j]
	// 		}
	// 		cap := min * (j - i)
	// 		if cap > max {
	// 			max = cap
	// 		}
	// 	}
	// }
	// return max
	// 思路二
	leftIndex := 0
	rightIndex := len(height) - 1
	max := 0
	for {
		if leftIndex >= rightIndex {
			break
		}
		min := height[leftIndex]
		if height[rightIndex] < min {
			min = height[rightIndex]
		}
		area := min * (rightIndex - leftIndex)
		if area > max {
			max = area
		}
		if height[rightIndex] < height[leftIndex] {
			rightIndex--
		} else {
			leftIndex++
		}
	}
	return max
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
