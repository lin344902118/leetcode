package main

import (
	"fmt"
	"sort"
)

/*
31. 下一个排列
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。
*/

func nextPermutation(nums []int) {
	// 由后往前后判断，是否存在一个数大于其前一个数，如果存在。该前一个数安排合适的位置并返回。如果一直找不到，
	// 说明不存在字典序更大的排列，交换整个数组。
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			// 找到i之后第一个比i-1大的数
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					nums[j], nums[i-1] = nums[i-1], nums[j]
					break
				}
			}
			sort.Ints(nums[i:])
			return
		}
	}
	// 未找到
	sort.Ints(nums)
	return
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{1, 3, 4, 2}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{1, 5, 1}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = make([]int, 0, 100)
	for i := 100; i > 0; i-- {
		nums = append(nums, i)
	}
	nextPermutation(nums)
	fmt.Println(nums)
}
