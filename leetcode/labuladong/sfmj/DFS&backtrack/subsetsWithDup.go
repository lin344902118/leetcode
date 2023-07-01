package main

import (
	"fmt"
	"sort"
)

/*
90. 子集 II
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
*/

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtrack func(nums []int, index int)
	backtrack = func(nums []int, index int) {
		res = append(res, append([]int{}, path...))
		for i := index; i < len(nums); i++ {
			// 相同的值不处理
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	// 先排序
	sort.Ints(nums)
	backtrack(nums, 0)
	return res
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
