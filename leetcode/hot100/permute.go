package main

import "fmt"

/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	path := make([]int, 0, n)
	used := make(map[int]struct{}, 0)
	var backtrack func()
	backtrack = func() {
		if len(path) == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			_, e := used[nums[i]]
			if e {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]] = struct{}{}
			backtrack()
			path = path[:len(path)-1]
			delete(used, nums[i])
		}
	}
	backtrack()
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
