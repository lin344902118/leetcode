package main

/*
输⼊股票价格数组 prices，你最多进⾏ max_k 次交易，每次交易需要额外消耗 fee 的⼿续费，⽽且每次交
易之后需要经过 cooldown 天的冷冻期才能进⾏下⼀次交易，请你计算并返回可以获得的最⼤利润。
*/

func maxProfit(k, fee, cooldown int, prices []int) int {
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
				dp[i][j][1] = -prices[0] - fee
				continue
			}
			if i-cooldown-1 < 0 {
				if dp[i-1][j][0] > dp[i-1][j][1]+prices[i] {
					dp[i][j][0] = dp[i-1][j][0]
				} else {
					dp[i][j][0] = dp[i-1][j][1] + prices[i]
				}

				if dp[i-1][j][1] > -prices[i]-fee {
					dp[i][j][1] = dp[i-1][j][1]
				} else {
					dp[i][j][1] = -prices[i] - fee
				}
				continue
			}
			if dp[i-1][j][0] > dp[i-1][j][1]+prices[i] {
				dp[i][j][0] = dp[i-1][j][0]
			} else {
				dp[i][j][0] = dp[i-1][j][1] + prices[i]
			}

			if dp[i-1][j][1] > dp[i-cooldown][j-1][0]-prices[i]-fee {
				dp[i][j][1] = dp[i-1][j][1]
			} else {
				dp[i][j][1] = dp[i-cooldown][j-1][0] - prices[i] - fee
			}
		}
	}
	return dp[n-1][k][0]
}

func main() {
}
