package main

import "fmt"

/*
  买卖股票的最佳时机 II
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。
你也可以先购买，然后在 同一天 出售。
返回 你能获得的 最大 利润 。
*/

func maxProfit(prices []int) int {
	// 使用前缀差计算
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4} // 7
	fmt.Println(maxProfit(prices))
	prices = []int{1, 2, 3, 4, 5} // 4
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1} // 0
	fmt.Println(maxProfit(prices))
}
