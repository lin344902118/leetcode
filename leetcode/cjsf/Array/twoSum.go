package main

import "fmt"

/*
  两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。
*/

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		v, exist := cache[target-nums[i]]
		if exist {
			return []int{i, v}
		}
		cache[nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // 0 1
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target)) // 1 2
	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target)) // 0 1
}
