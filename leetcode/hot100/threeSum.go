package main

import (
	"fmt"
	"sort"
)

/*
15. 三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		// 跳过相邻重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 排序后，第一个数一定是负数
		if nums[i] > 0 {
			break
		}
		// 双指针
		left := i + 1
		right := len(nums) - 1
		target := -nums[i]
		// 最大的两个数之和都无法满足
		if nums[right]+nums[right-1] < target {
			continue
		}
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 跳过左边重复元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 跳过右边重复元素
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				// 和小于期望值，左指针往后移动
				left++
			} else {
				// 和大于期望值，右指针往前移动
				right--
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, -2, -1}
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
