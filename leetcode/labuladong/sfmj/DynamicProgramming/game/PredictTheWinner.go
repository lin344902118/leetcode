package main

import "fmt"

/*
486. 预测赢家
给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏。

玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手。开始时，两个玩家的初始分值都是 0 。
每一回合，玩家从数组的任意一端取一个数字（即，nums[0] 或 nums[nums.length - 1]），
取到的数字将会从数组中移除（数组长度减 1 ）。
玩家选中的数字将会加到他的得分上。当数组中没有剩余数字可取时，游戏结束。

如果玩家 1 能成为赢家，返回 true 。如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，也返回 true 。
你可以假设每个玩家的玩法都会使他的分数最大化。
*/

func PredictTheWinner(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	return stoneGame(nums) >= 0
}

func stoneGame(nums []int) int {
	// dp[i][j][0]表示从i到j个数字中先手获得的最大分数，dp[i][j][1]表示从i到j个数字中后手获得最大分数
	n := len(nums)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
		}
		dp[i][i][0] = nums[i]
		dp[i][i][1] = 0
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			left := nums[i] + dp[i+1][j][1]
			right := nums[j] + dp[i][j-1][1]
			if left > right {
				dp[i][j][0] = left
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}
		}
	}
	return dp[0][n-1][0] - dp[0][n-1][1]
}

func main() {
	nums := []int{1, 5, 233, 7}
	fmt.Println(PredictTheWinner(nums))
}
