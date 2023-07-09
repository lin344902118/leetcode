package main

/*
509. 斐波那契数
斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，
后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给定 n ，请计算 F(n) 。
*/

// var cache = make([]int, 31)

// func fib(n int) int {
// 	cache[n] = 0
// 	cache[1] = 1
// 	if n == 0 {
// 		return 0
// 	}
// 	if cache[n] != 0 {
// 		return cache[n]
// 	}
// 	res := cache[n-1] + cache[n-2]
// 	cache[n] = res
// 	return res
// }

// func fib(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	dp := make([]int, n+1)
// 	dp[0] = 0
// 	dp[1] = 1
// 	for i := 2; i <= n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	return dp[n]
// }

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp_i_1 := 1
	dp_i_2 := 0
	for i := 2; i <= n; i++ {
		dp_i := dp_i_1 + dp_i_2
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i_1
}

func main() {

}
