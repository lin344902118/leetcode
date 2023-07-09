package main

import (
	"fmt"
	"math"
)

/*
121. 买卖股票的最佳时机
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/

func maxProfit(prices []int) int {
	// dp[i][0] 表示第i天未持有股票最大利润
	// dp[i][1] 表示第i天持有股票最大利润
	// 求dp[i][0]
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
	// 	if -prices[i] > dp[i-1][1] {
	// 		dp[i][1] = - prices[i]
	// 	}
	// }
	// return dp[n-1][0]
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
	for i := 0; i < n; i++ {
		if dp_i_1+prices[i] > dp_i_0 {
			dp_i_0 = dp_i_1 + prices[i]
		}
		if -prices[i] > dp_i_1 {
			dp_i_1 = -prices[i]
		}
	}
	return dp_i_0
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
