package main

import "fmt"

/*
121. 买卖股票的最佳时机
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/

func maxProfit(prices []int) int {
	// // 暴力
	// m := len(prices)
	// max := 0
	// for i := 0; i < m-1; i++ {
	// 	for j := i + 1; j < m; j++ {
	// 		if prices[j]-prices[i] > max {
	// 			max = prices[j] - prices[i]
	// 		}
	// 	}
	// }
	// return max
	// 动态规划 dp[i][0]表示第i天不持有股票，dp[i][1]表示第i天持有股票
	m := len(prices)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = dp[i-1][0]
		if dp[i-1][1]+prices[i] > dp[i-1][0] {
			dp[i][0] = dp[i-1][1] + prices[i]
		}
		dp[i][1] = dp[i-1][1]
		if -prices[i] > dp[i-1][1] {
			dp[i][1] = -prices[i]
		}
	}
	return dp[m-1][0]
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	// prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
