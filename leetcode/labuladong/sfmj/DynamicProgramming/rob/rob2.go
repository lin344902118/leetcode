package main

import "fmt"

/*
213. 打家劫舍 II
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
*/

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	memo1 := make([]int, n)
	memo2 := make([]int, n)
	for i := 0; i < n; i++ {
		memo1[i] = -1
		memo2[i] = -1
	}
	var dp func(nums, memo []int, start, end int) int
	dp = func(nums, memo []int, start, end int) int {
		if start > end {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		rb := dp(nums, memo, start+2, end) + nums[start]
		urb := dp(nums, memo, start+1, end)
		if rb > urb {
			memo[start] = rb
			return rb
		} else {
			memo[start] = urb
			return urb
		}
	}
	// 抢劫第0家，不能抢劫第n家
	rb := dp(nums, memo1, 0, n-2)
	// 不抢劫第0家，抢劫第n家
	urb := dp(nums, memo2, 1, n-1)
	if rb > urb {
		return rb
	} else {
		return urb
	}
}

func main() {
	nums := []int{2, 3, 2}
	nums = []int{1, 2, 3, 1}
	nums = []int{1, 2, 3}
	fmt.Println(rob(nums))
}
