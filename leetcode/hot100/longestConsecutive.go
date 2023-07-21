package main

import (
	"fmt"
	"sort"
)

/*
128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	left := 0
	res := 1
	repeat := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			if i-left+1-repeat > res {
				res = i - left + 1 - repeat
			}
		} else if nums[i] == nums[i-1] {
			repeat++
		} else {
			left = i
			repeat = 0
		}
	}
	return res
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums = []int{1, 2, 0, 1, 1, 2, 0, 5}
	nums = []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	fmt.Println(longestConsecutive(nums))
}
