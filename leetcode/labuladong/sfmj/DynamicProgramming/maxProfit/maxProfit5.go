package main

import "math"

/*
309. 最佳买卖股票时机含冷冻期
给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit(prices []int) int {
	// n := len(prices)
	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, 2)
	// }
	// for i := 0; i < n; i++ {
	// 	if i-1 == -1 {
	// 		dp[i][0] = 0
	// 		dp[i][1] = -prices[i]
	// 		continue
	// 	}
	// 	if i-2 == -1 {
	// 		if dp[i-1][0] > dp[i-1][1]+prices[i] {
	// 			dp[i][0] = dp[i-1][0]
	// 		} else {
	// 			dp[i][0] = dp[i-1][1] + prices[i]
	// 		}
	// 		if dp[i-1][1] > -prices[i] {
	// 			dp[i][1] = dp[i-1][1]
	// 		} else {
	// 			dp[i][1] = -prices[i]
	// 		}
	// 		continue
	// 	}
	// 	if dp[i-1][0] > dp[i-1][1]+prices[i] {
	// 		dp[i][0] = dp[i-1][0]
	// 	} else {
	// 		dp[i][0] = dp[i-1][1] + prices[i]
	// 	}
	// 	if dp[i-1][1] > dp[i-2][0]-prices[i] {
	// 		dp[i][1] = dp[i-1][1]
	// 	} else {
	// 		dp[i][1] = dp[i-2][0] - prices[i]
	// 	}
	// }
	// return dp[n-1][0]
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
	dp_pre_0 := 0 // dp[i-2][0]
	for i := 0; i < n; i++ {
		tmp := dp_i_0
		if dp_i_1+prices[i] > dp_i_0 {
			dp_i_0 = dp_i_1 + prices[i]
		}
		if dp_pre_0-prices[i] > dp_i_1 {
			dp_i_1 = dp_pre_0 - prices[i]
		}
		dp_pre_0 = tmp
	}
	return dp_i_0
}

func main() {
}
