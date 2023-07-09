package main

import (
	"fmt"
	"sort"
)

/*
  三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j],
nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		// 用双指针
		left := i + 1
		right := len(nums) - 1
		target := -nums[i]
		if nums[right]+nums[right-1] < target {
			continue
		}
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				fmt.Println(left, right, nums)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(threeSum(nums))
	// nums = []int{0, 1, 1}
	// fmt.Println(threeSum(nums))
	// nums = []int{0, 0, 0}
	// fmt.Println(threeSum(nums))
	// nums = []int{3, 0, -2, -1, 1, 2}
	// fmt.Println(threeSum(nums))
	// nums = []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	// fmt.Println(threeSum(nums))
	nums := []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(nums))
}
