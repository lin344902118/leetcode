package main

import (
	"fmt"
)

/*
322. 零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。
*/

func coinChange(coins []int, amount int) int {
	// cache := make([]int, amount+1)
	// for i := 0; i < len(cache); i++ {
	// 	cache[i] = -666
	// }
	// var dp func(coins []int, amount int) int
	// dp = func(coins []int, amount int) int {
	// 	if amount == 0 {
	// 		return 0
	// 	}
	// 	if amount < 0 {
	// 		return -1
	// 	}
	// 	if cache[amount] != -666 {
	// 		return cache[amount]
	// 	}
	// 	res := math.MaxInt32
	// 	for _, coin := range coins {
	// 		tmp := dp(coins, amount-coin)
	// 		if tmp == -1 {
	// 			continue
	// 		}
	// 		if tmp+1 < res {
	// 			res = tmp + 1
	// 		}
	// 	}
	// 	if res == math.MaxInt32 {
	// 		cache[amount] = -1
	// 		return -1
	// 	}
	// 	cache[amount] = res
	// 	return res
	// }
	// return dp(coins, amount)
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
