package main

import "math"

/*
714. 买卖股票的最佳时机含手续费
给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
*/

func maxProfit(prices []int, fee int) int {
	// n := len(prices)
	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, 2)
	// }
	// // 买入的时候支付手续费
	// for i := 0; i < n; i++ {
	// 	if i-1 == -1 {
	// 		dp[i][0] = 0
	// 		dp[i][1] = -prices[i] - fee
	// 		continue
	// 	}
	// 	if dp[i-1][0] > dp[i-1][1]+prices[i] {
	// 		dp[i][0] = dp[i-1][0]
	// 	} else {
	// 		dp[i][0] = dp[i-1][1] + prices[i]
	// 	}
	// 	if dp[i-1][1] > dp[i-1][0]-prices[i]-fee {
	// 		dp[i][1] = dp[i-1][1]
	// 	} else {
	// 		dp[i][1] = dp[i-1][0] - prices[i] - fee
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
		if tmp-prices[i]-fee > dp_i_1 {
			dp_i_1 = tmp - prices[i] - fee
		}
	}
	return dp_i_0
}

func main() {
}
