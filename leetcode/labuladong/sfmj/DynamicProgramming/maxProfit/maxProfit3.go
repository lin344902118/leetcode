package main

import (
	"fmt"
	"math"
)

/*
123. 买卖股票的最佳时机 III
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit(prices []int) int {
	// n := len(prices)
	// if n == 0 {
	// 	return 0
	// }
	// k := 2
	// dp := make([][][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([][]int, k+1)
	// 	for j := 0; j < k+1; j++ {
	// 		dp[i][j] = make([]int, 2)
	// 	}
	// }
	// // dp[n][k][0] 表示第n天可以完成k笔交易并且未持有股票
	// // dp[n][k][1] 表示第n天可以完成k笔交易并且持有股票
	// for i := 0; i < n; i++ {
	// 	for j := k; j > 0; j-- {
	// 		if i-1 == -1 {
	// 			dp[i][j][0] = 0
	// 			dp[i][j][1] = -prices[0]
	// 			continue
	// 		}
	// 		if dp[i-1][j][0] > dp[i-1][j][1]+prices[i] {
	// 			dp[i][j][0] = dp[i-1][j][0]
	// 		} else {
	// 			dp[i][j][0] = dp[i-1][j][1] + prices[i]
	// 		}

	// 		if dp[i-1][j][1] > dp[i-1][j-1][0]-prices[i] {
	// 			dp[i][j][1] = dp[i-1][j][1]
	// 		} else {
	// 			dp[i][j][1] = dp[i-1][j-1][0] - prices[i]
	// 		}
	// 	}
	// }
	// return dp[n-1][k][0]
	dp_i10 := 0
	dp_i11 := math.MinInt32
	dp_i20 := 0
	dp_i21 := math.MinInt32
	for i := 0; i < len(prices); i++ {
		if dp_i21+prices[i] > dp_i20 {
			dp_i20 = dp_i21 + prices[i]
		}
		if dp_i10-prices[i] > dp_i21 {
			dp_i21 = dp_i10 - prices[i]
		}
		if dp_i11+prices[i] > dp_i10 {
			dp_i10 = dp_i11 + prices[i]
		}
		if -prices[i] > dp_i11 {
			dp_i11 = -prices[i]
		}
	}
	return dp_i20
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}
