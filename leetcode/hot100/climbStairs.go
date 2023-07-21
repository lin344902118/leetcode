package main

import "fmt"

/*
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

var cache = map[int]int{1: 1, 2: 2}

func climbStairs(n int) int {
	// if n == 1 || n == 2 {
	// 	return n
	// }
	// dp := make([]int, n)
	// dp[0] = 1
	// dp[1] = 2
	// for i := 2; i < n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n-1]
	v, e := cache[n]
	if e {
		return v
	} else {
		res := climbStairs(n-1) + climbStairs(n-2)
		cache[n] = res
		return res
	}
}

func main() {
	n := 45
	fmt.Println(climbStairs(n))
}
