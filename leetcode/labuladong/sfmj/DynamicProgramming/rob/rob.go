package main

import "fmt"

/*
198. 打家劫舍
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，
一夜之内能够偷窃到的最高金额。
*/

func rob(nums []int) int {
	// // dp[i][0] 表示不打劫第i家
	// // dp[i][1] 表示打劫第i家
	// n := len(nums)
	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, 2)
	// }
	// for i := 0; i < n; i++ {
	// 	if i-1 == -1 {
	// 		dp[i][0] = 0
	// 		dp[i][1] = nums[i]
	// 		continue
	// 	}
	// 	if dp[i-1][0] > dp[i-1][1] {
	// 		dp[i][0] = dp[i-1][0]
	// 	} else {
	// 		dp[i][0] = dp[i-1][1]
	// 	}
	// 	if dp[i-1][1] > dp[i-1][0]+nums[i] {
	// 		dp[i][1] = dp[i-1][1]
	// 	} else {
	// 		dp[i][1] = dp[i-1][0] + nums[i]
	// 	}
	// }
	// if dp[n-1][0] > dp[n-1][1] {
	// 	return dp[n-1][0]
	// }
	// return dp[n-1][1]
	rb, urb := nums[0], 0
	for i := 1; i < len(nums); i++ {
		// urb+nums[i]上家不抢，抢这家, rb抢上架，不抢这家
		rb, urb = max(urb+nums[i], rb), max(urb, rb)
	}
	return max(rb, urb)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
