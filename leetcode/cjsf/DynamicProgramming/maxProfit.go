package main

import "fmt"

/*
  买卖股票的最佳时机
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxPro := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > maxPro {
			maxPro = prices[i] - min
		}
	}
	return maxPro
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // 5
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices)) // 0
}
