package main

/*
877. 石子游戏
Alice 和 Bob 用几堆石子在做游戏。一共有偶数堆石子，排成一行；每堆都有 正 整数颗石子，数目为 piles[i] 。

游戏以谁手中的石子最多来决出胜负。石子的 总数 是 奇数 ，所以没有平局。

Alice 和 Bob 轮流进行，Alice 先开始 。 每回合，玩家从行的 开始 或 结束 处取走整堆石头。
这种情况一直持续到没有更多的石子堆为止，此时手中 石子最多 的玩家 获胜 。

假设 Alice 和 Bob 都发挥出最佳水平，当 Alice 赢得比赛时返回 true ，当 Bob 赢得比赛时返回 false 。
*/

func stoneGame(piles []int) bool {
	if len(piles) == 1 {
		return true
	}
	n := len(piles)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
		}
		dp[i][i][0] = piles[i]
		dp[i][i][1] = 0
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			left := dp[i+1][j][1] + piles[i]
			right := dp[i][j-1][1] + piles[j]
			if left > right {
				dp[i][j][0] = left
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}
		}
	}
	return dp[0][n-1][0]-dp[0][n-1][1] > 0
}

func main() {
}
