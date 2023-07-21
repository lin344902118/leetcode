package main

import "fmt"

/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

func subsets(nums []int) [][]int {
	m := len(nums)
	res := make([][]int, 0)
	res = append(res, []int{})
	tmp := make([]int, 0)
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == m {
			return
		}
		for i := index; i < m; i++ {
			tmp = append(tmp, nums[i])
			res = append(res, append([]int{}, tmp...))
			backtrack(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack(0)
	return res
}

func main() {
	nums := []int{0}
	fmt.Println(subsets(nums))
}
