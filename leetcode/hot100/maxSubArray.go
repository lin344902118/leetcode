package main

import "fmt"

/*
53. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。
*/

func maxSubArray(nums []int) int {
	n := len(nums)
	// dp[i]表示0到第j个数连续子数组和
	dp := make([]int, n)
	res := nums[0]
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}
		dp[i] = nums[i]
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
	nums = []int{1}
	fmt.Println(maxSubArray(nums))
	nums = []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
}
