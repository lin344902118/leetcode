package main

import (
	"fmt"
)

/*
55. 跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。
*/

func canJump(nums []int) bool {
	// 判断每个位置能达到的最大长度
	n := len(nums)
	max := 0
	for i := 0; i < n-1; i++ {
		if max < i+nums[i] {
			max = i + nums[i]
		}
		if max <= i {
			return false
		}
	}
	return max >= n-1
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))  // true
	nums = []int{3, 2, 1, 0, 4} // false
	fmt.Println(canJump(nums))
	nums = []int{2, 0, 0} // true
	fmt.Println(canJump(nums))
	nums = []int{0} // true
	fmt.Println(canJump(nums))
	nums = []int{0, 2, 3} // false
	fmt.Println(canJump(nums))
	nums = []int{2, 5, 0, 0} // true
	fmt.Println(canJump(nums))
	nums = []int{1, 0, 4, 4}
	fmt.Println(canJump(nums)) // false
}
