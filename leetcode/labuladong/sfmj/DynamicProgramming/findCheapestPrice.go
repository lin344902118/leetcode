package main

import (
	"fmt"
	"math"
)

/*
787. K 站中转内最便宜的航班
有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，
表示该航班都从城市 fromi 开始，以价格 pricei 抵达 toi。

现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，
使得从 src 到 dst 的 价格最便宜 ，并返回该价格。
 如果不存在这样的路线，则输出 -1。
*/

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	indegree := make(map[int][][]int, 0)
	for _, f := range flights {
		from := f[0]
		to := f[1]
		price := f[2]
		_, exist := indegree[to]
		if !exist {
			indegree[to] = make([][]int, 0)
		}
		indegree[to] = append(indegree[to], []int{from, price})
	}
	k++
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			cache[i][j] = -888
		}
	}

	var dp func(s int, k int) int
	dp = func(s, k int) int {
		if s == src {
			return 0
		}
		if k == 0 {
			return -1
		}
		if cache[s][k] != -888 {
			return cache[s][k]
		}
		res := math.MaxInt32
		value, exist := indegree[s]
		if exist {
			for _, v := range value {
				from := v[0]
				price := v[1]
				subProblem := dp(from, k-1)
				if subProblem != -1 {
					if subProblem+price < res {
						res = subProblem + price
					}
				}
			}
		}
		if res == math.MaxInt32 {
			cache[s][k] = -1
			return -1
		}
		cache[s][k] = res
		return res
	}
	return dp(dst, k)
}

func main() {
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src := 0
	dst := 3
	k := 1
	fmt.Println(flights)
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}
