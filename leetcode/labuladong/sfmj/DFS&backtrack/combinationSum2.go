package main

import (
	"fmt"
	"sort"
)

/*
40. 组合总和 II
给定一个候选人编号的集合 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。
*/

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sum := 0
	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			if sum+nums[i] > target {
				return
			}
			path = append(path, nums[i])
			sum += nums[i]
			backtrack(nums, i+1)
			path = path[:len(path)-1]
			sum -= nums[i]
		}
	}
	sort.Ints(candidates)
	backtrack(candidates, 0)
	return res
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
