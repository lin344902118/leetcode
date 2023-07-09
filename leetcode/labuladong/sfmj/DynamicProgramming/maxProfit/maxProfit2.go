package main

import "math"

/*
122. 买卖股票的最佳时机 II
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。
*/

func maxProfit(prices []int) int {
	// n := len(prices)
	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, 2)
	// }
	// dp[0][0] = 0
	// dp[0][1] = -prices[0]
	// for i := 1; i < n; i++ {
	// 	dp[i][0] = dp[i-1][0]
	// 	if dp[i-1][1]+prices[i] > dp[i-1][0] {
	// 		dp[i][0] = dp[i-1][1] + prices[i]
	// 	}
	// 	dp[i][1] = dp[i-1][1]
	// 	if dp[i-1][0]-prices[i] > dp[i-1][1] {
	// 		dp[i][1] = dp[i-1][0] - prices[i]
	// 	}
	// }
	// return dp[n-1][0]
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
	for i := 0; i < n; i++ {
		tmp := dp_i_0
		if dp_i_1+prices[i] > dp_i_0 {
			dp_i_0 = dp_i_1 + prices[i]
		}
		if tmp-prices[i] > dp_i_1 {
			dp_i_1 = tmp - prices[i]
		}
	}
	return dp_i_0
}

func main() {
}
