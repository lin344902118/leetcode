package main

import "fmt"

/*
  爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

var cache = map[int]int{1: 1, 2: 2}

func climbStairs(n int) int {
	v, exist := cache[n]
	if exist {
		return v
	} else {
		res := climbStairs(n-1) + climbStairs(n-2)
		cache[n] = res
		return res
	}
}

func main() {
	n := 2
	fmt.Println(climbStairs(n)) // 2
	n = 3
	fmt.Println(climbStairs(n)) // 3
}
