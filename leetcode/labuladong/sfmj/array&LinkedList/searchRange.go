package main

import "fmt"

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left > len(nums)-1 || nums[left] != target {
		return []int{-1, -1}
	}
	res := make([]int, 0, 2)
	res = append(res, left)
	for i := left; i < len(nums); i++ {
		if nums[i] != target {
			res = append(res, i-1)
			break
		}
	}
	if nums[len(nums)-1] == target {
		res = append(res, len(nums)-1)
	}
	return res
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target)) // [3,4]
	target = 6
	fmt.Println(searchRange(nums, target)) // [-1,-1]
	nums = []int{}
	target = 0
	fmt.Println(searchRange(nums, target)) // [-1,-1]
}
