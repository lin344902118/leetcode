package main

import (
	"fmt"
)

/*
698. 划分为k个相等的子集
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
*/

func canPartitionKSubsets(nums []int, k int) bool {
	// 数组长度小于k，返回false
	if len(nums) < k {
		return false
	}
	// 数组数字总和不是k的整数倍，返回false
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	used := 0
	memo := make(map[int]bool)
	var backtrack func(k, bucket, start, target, used int, nums []int) bool
	backtrack = func(k, bucket, start, target, used int, nums []int) bool {
		if k == 0 {
			return true
		}
		if bucket == target {
			res := backtrack(k-1, 0, 0, target, used, nums)
			memo[used] = res
			return res
		}
		_, exist := memo[used]
		if exist {
			return memo[used]
		}
		for i := start; i < len(nums); i++ {
			if ((used >> i) & 1) == 1 {
				continue
			}
			if nums[i]+bucket > target {
				continue
			}
			bucket += nums[i]
			used |= 1 << i
			if backtrack(k, bucket, i+1, target, used, nums) {
				return true
			}
			used ^= 1 << i
			bucket -= nums[i]
		}
		return false
	}
	return backtrack(k, 0, 0, target, used, nums)
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	fmt.Println(canPartitionKSubsets(nums, 4))
}
