package main

import "fmt"

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 找到目标值索引
	index := halfSearch(nums, 0, len(nums)-1, target)
	if index == -1 {
		return []int{-1, -1}
	}
	leftIndex := index
	rightIndex := index
	// 往左右扩展
	for leftIndex > 0 {
		if nums[leftIndex-1] != target {
			break
		}
		leftIndex--
	}
	for rightIndex < len(nums)-1 {
		if nums[rightIndex+1] != target {
			break
		}
		rightIndex++
	}
	return []int{leftIndex, rightIndex}
}

func halfSearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	nums = []int{2, 2}
	target = 2
	fmt.Println(searchRange(nums, target))
}
