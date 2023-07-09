package main

import "fmt"

/*
188. 买卖股票的最佳时机 IV
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格，和一个整型 k 。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// dp[n][k][0] 表示第n天可以完成k笔交易并且未持有股票
	// dp[n][k][1] 表示第n天可以完成k笔交易并且持有股票
	for i := 0; i < n; i++ {
		for j := k; j > 0; j-- {
			if i-1 == -1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
				continue
			}
			if dp[i-1][j][0] > dp[i-1][j][1]+prices[i] {
				dp[i][j][0] = dp[i-1][j][0]
			} else {
				dp[i][j][0] = dp[i-1][j][1] + prices[i]
			}

			if dp[i-1][j][1] > dp[i-1][j-1][0]-prices[i] {
				dp[i][j][1] = dp[i-1][j][1]
			} else {
				dp[i][j][1] = dp[i-1][j-1][0] - prices[i]
			}
		}
	}
	return dp[n-1][k][0]
}

func main() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(k, prices))
}
