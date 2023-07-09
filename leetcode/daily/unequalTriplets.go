package main

import "fmt"

/*
2475. 数组中不等三元组的数目
给你一个下标从 0 开始的正整数数组 nums 。请你找出并统计满足下述条件的三元组 (i, j, k) 的数目：
0 <= i < j < k < nums.length
nums[i]、nums[j] 和 nums[k] 两两不同 。
换句话说：nums[i] != nums[j]、nums[i] != nums[k] 且 nums[j] != nums[k] 。
返回满足上述条件三元组的数目。
*/

func unequalTriplets(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i] == nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i] != nums[k] && nums[j] != nums[k] {
					res++
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{4, 4, 2, 4, 3}
	fmt.Println(unequalTriplets(nums)) // 3
	nums = []int{1, 1, 1, 1, 1}
	fmt.Println(unequalTriplets(nums)) // 0
}
